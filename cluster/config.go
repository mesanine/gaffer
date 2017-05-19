package cluster

import "github.com/vektorlab/gaffer/cluster/service"

// A Config emits a service configuration based on
// pre-defined options of a given cluster.
// Gaffer only supports configuring Mesos clusters for now
// but could be updated to support other systems in the future.
type Config interface {
	// Update returns the desired service configuration
	// based on the cluster input. If no change is required
	// Update returns nil.
	Update(Cluster) map[string][]*service.Service
}

var (
	_ Config = Mesos{}
	_ Config = Mock{}
)