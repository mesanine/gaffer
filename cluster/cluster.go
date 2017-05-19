package cluster

import (
	"github.com/vektorlab/gaffer/cluster/host"
	"github.com/vektorlab/gaffer/cluster/service"
	"os"
)

type ProcessList map[string]map[string]*os.Process

func (p ProcessList) Len() int {
	var l int
	for _, m := range p {
		l += len(m)
	}
	return l
}

func (p ProcessList) Started(hostID string, serviceID string) bool {
	if _, ok := p[hostID]; ok {
		if _, ok := p[hostID][serviceID]; ok {
			return true
		}
	}
	return false
}

type Stats struct {
	Started int
	Stopped int
	Hosts   map[string]struct {
		Started int
		Stopped int
	}
}

// State represents the state of a given cluster
type State int

func (s State) String() string {
	switch s {
	case CONVERGING:
		return "CONVERGING"
	case STARTING:
		return "STARTING"
	case STARTED:
		return "STARTED"
	case DEGRADED:
		return "DEGRADED"
	}
	return ""
}

const (
	_ State = iota
	// Management hosts are still coming online
	CONVERGING
	// One or more host services are still starting
	STARTING
	// All essential host services have started
	STARTED
	// One or more services have not reported back
	DEGRADED
)

// Cluster represents the overall configuration of a Mesos cluster
type Cluster struct {
	ID       string                        `json:"id"`
	Hosts    []*host.Host                  `json:"hosts"`
	Services map[string][]*service.Service `json:"services"`
}

func New(id string, size int) *Cluster {
	cluster := &Cluster{
		ID:       id,
		Hosts:    []*host.Host{},
		Services: map[string][]*service.Service{},
	}
	for i := 0; i < size; i++ {
		cluster.Hosts = append(cluster.Hosts, host.NewHost())
		cluster.Services[cluster.Hosts[i].ID] = []*service.Service{}
	}
	return cluster
}

func (c Cluster) State(p ProcessList) State {
	state := CONVERGING
	if c.Hosts == nil {
		// No hosts registered
		// CONVERGING
		return state
	}
	for _, host := range c.Hosts {
		if !host.Registered {
			// Not all hosts registered
			// CONVERGING
			return state
		}
	}
	// All hosts registered
	// STARTING
	state++
	var s int
	for _, services := range c.Services {
		s += len(services)
	}
	if len(c.ServicesFlat()) != p.Len() {
		// STARTING
		return state
	}
	// All services are running
	// STARTED
	state++
	return state
}

func (c Cluster) Stats(p ProcessList) *Stats {
	stats := &Stats{
		Started: 0,
		Stopped: 0,
		Hosts: map[string]struct {
			Started int
			Stopped int
		}{},
	}
	for hostID, services := range c.Services {
		stats.Started += len(services)
		hostStats := struct {
			Started int
			Stopped int
		}{}
		if running, ok := p[hostID]; ok {
			hostStats.Started = len(running)
			hostStats.Stopped = len(services) - len(running)
		} else {
			hostStats.Stopped = len(services)
		}
		stats.Hosts[hostID] = hostStats
	}
	return stats
}

func (c Cluster) Host(id string) *host.Host {
	for _, host := range c.Hosts {
		if host.ID == id {
			return host
		}
	}
	return nil
}

func (c Cluster) Service(host, id string) *service.Service {
	if services, ok := c.Services[host]; ok {
		for _, svc := range services {
			if svc.ID == id {
				return svc
			}
		}
	}
	return nil
}

func (c Cluster) ServicesFlat() []*service.Service {
	flat := []*service.Service{}
	for _, services := range c.Services {
		for _, service := range services {
			flat = append(flat, service)
		}
	}
	return flat
}
