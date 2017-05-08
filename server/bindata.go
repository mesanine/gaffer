// Code generated by go-bindata.
// sources:
// www/.index.html.swp
// www/index.html
// www/static/css/bootstrap.min.css
// www/static/css/bootstrap.min.css.map
// www/static/css/dashboard.css
// www/static/js/bootstrap.min.js
// www/static/js/jquery.min.js
// DO NOT EDIT!

package server

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// bindataRead reads the given file from disk. It returns an error on failure.
func bindataRead(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// wwwIndexHtmlSwp reads file data from disk. It returns an error on failure.
func wwwIndexHtmlSwp() (*asset, error) {
	path := "/home/kevin/repos/go/src/github.com/vektorlab/gaffer/www/.index.html.swp"
	name := "www/.index.html.swp"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// wwwIndexHtml reads file data from disk. It returns an error on failure.
func wwwIndexHtml() (*asset, error) {
	path := "/home/kevin/repos/go/src/github.com/vektorlab/gaffer/www/index.html"
	name := "www/index.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// wwwStaticCssBootstrapMinCss reads file data from disk. It returns an error on failure.
func wwwStaticCssBootstrapMinCss() (*asset, error) {
	path := "/home/kevin/repos/go/src/github.com/vektorlab/gaffer/www/static/css/bootstrap.min.css"
	name := "www/static/css/bootstrap.min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// wwwStaticCssBootstrapMinCssMap reads file data from disk. It returns an error on failure.
func wwwStaticCssBootstrapMinCssMap() (*asset, error) {
	path := "/home/kevin/repos/go/src/github.com/vektorlab/gaffer/www/static/css/bootstrap.min.css.map"
	name := "www/static/css/bootstrap.min.css.map"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// wwwStaticCssDashboardCss reads file data from disk. It returns an error on failure.
func wwwStaticCssDashboardCss() (*asset, error) {
	path := "/home/kevin/repos/go/src/github.com/vektorlab/gaffer/www/static/css/dashboard.css"
	name := "www/static/css/dashboard.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// wwwStaticJsBootstrapMinJs reads file data from disk. It returns an error on failure.
func wwwStaticJsBootstrapMinJs() (*asset, error) {
	path := "/home/kevin/repos/go/src/github.com/vektorlab/gaffer/www/static/js/bootstrap.min.js"
	name := "www/static/js/bootstrap.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// wwwStaticJsJqueryMinJs reads file data from disk. It returns an error on failure.
func wwwStaticJsJqueryMinJs() (*asset, error) {
	path := "/home/kevin/repos/go/src/github.com/vektorlab/gaffer/www/static/js/jquery.min.js"
	name := "www/static/js/jquery.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"www/.index.html.swp": wwwIndexHtmlSwp,
	"www/index.html": wwwIndexHtml,
	"www/static/css/bootstrap.min.css": wwwStaticCssBootstrapMinCss,
	"www/static/css/bootstrap.min.css.map": wwwStaticCssBootstrapMinCssMap,
	"www/static/css/dashboard.css": wwwStaticCssDashboardCss,
	"www/static/js/bootstrap.min.js": wwwStaticJsBootstrapMinJs,
	"www/static/js/jquery.min.js": wwwStaticJsJqueryMinJs,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"www": &bintree{nil, map[string]*bintree{
		".index.html.swp": &bintree{wwwIndexHtmlSwp, map[string]*bintree{}},
		"index.html": &bintree{wwwIndexHtml, map[string]*bintree{}},
		"static": &bintree{nil, map[string]*bintree{
			"css": &bintree{nil, map[string]*bintree{
				"bootstrap.min.css": &bintree{wwwStaticCssBootstrapMinCss, map[string]*bintree{}},
				"bootstrap.min.css.map": &bintree{wwwStaticCssBootstrapMinCssMap, map[string]*bintree{}},
				"dashboard.css": &bintree{wwwStaticCssDashboardCss, map[string]*bintree{}},
			}},
			"js": &bintree{nil, map[string]*bintree{
				"bootstrap.min.js": &bintree{wwwStaticJsBootstrapMinJs, map[string]*bintree{}},
				"jquery.min.js": &bintree{wwwStaticJsJqueryMinJs, map[string]*bintree{}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

