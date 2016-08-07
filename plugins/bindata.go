// Code generated by go-bindata.
// sources:
// plugins/plugins.toml
// DO NOT EDIT!

package plugins

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _pluginsPluginsToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x96\x5b\x6f\xea\x38\x10\xc7\xdf\xf9\x14\xa3\xf4\x65\x77\x75\x28\x7b\xce\xea\x9c\xb7\x7d\xc8\x52\x7a\x5b\x5a\x50\x42\x91\xaa\xaa\x6a\x87\xc4\x09\xde\x3a\x76\x64\x3b\xb4\x7c\xfb\xb5\xcd\xad\x40\x10\x0e\x6a\x5f\xd2\x8c\xc7\xf3\xff\xcd\xdf\x17\x72\x06\x5d\x51\xce\x25\xcd\xa7\x1a\x7e\x4b\x7e\x87\x1f\x7f\x7e\xff\x0b\xda\xf6\xf1\x0b\x26\x0c\x93\x37\x2d\x4a\xb8\x15\x6a\x5a\x21\xdc\x21\xe5\xe4\x1b\x84\x8c\x41\x64\x27\x28\x88\x88\x22\x72\x46\xd2\xf3\xd6\x19\xc4\x84\x40\xff\xa6\xdb\xbb\x8f\x7b\x90\x09\x09\x8c\x26\x84\x2b\x02\x94\x9b\xb7\x02\x35\x15\xfc\xbc\xd5\x3a\xfb\x9a\x3f\xa3\x37\xec\x3f\x5c\xdd\xdc\x1b\x7c\x9e\xd1\xbc\x92\x4e\x00\x9a\xd7\xf9\x22\x9e\x96\xa6\x9a\x11\xf8\x1b\x82\x3b\xb4\x9d\xc3\x90\x55\x39\xe5\xdb\x78\x2a\x68\xb5\x9e\x9e\x4a\x37\xf2\xfc\xdc\x02\xe0\x58\xb8\x39\x5c\x49\x16\x98\x77\xc2\x71\xc2\x48\x6a\x42\x19\x32\x45\x4c\x24\x25\x2a\x91\xb4\x74\xcd\x99\xc4\xfb\x38\xea\xc3\x05\x6a\x9c\xa0\xb1\xf6\x1a\xd5\xd4\xf8\x8e\x32\x99\xda\xd9\x09\x6a\x92\x0b\x39\xb7\x89\x94\x6b\xe2\x4a\xd2\x02\x73\xa7\x51\x38\xae\xce\x4a\xaa\xa0\x0b\xe9\xa9\xa9\x61\xdf\xed\x53\xcf\x4b\xa2\x4c\xf0\xc9\x64\xa7\x3f\x03\x78\x3e\x80\x3b\xa3\xb2\x52\x5a\x68\xdc\x81\xd6\xb2\xaa\x61\x1e\xdb\xec\x91\xcd\x36\x1b\x2b\xa3\xcc\x48\xa8\x04\x39\x20\x4f\x9d\x2a\x30\x21\xde\xaa\xd2\xbb\x83\x6d\xf5\xaa\x64\x02\xd7\x8e\x81\xdd\x87\x7a\x15\xd4\x02\x5e\x2d\xd3\x2b\xd0\x0c\xe6\xa2\x82\x77\xe4\xda\x46\x97\xe3\x0a\x8b\xd2\x2c\x9a\x09\x6c\x6a\x9e\x27\xa2\x30\x65\xb1\xa4\x6f\xc4\x71\x38\xae\xc2\x2a\x04\x1b\x50\x0f\xf7\xbe\x41\xa0\xa6\xf8\x7d\xf9\xfc\xf1\xf3\x97\xf5\xd3\xba\x35\xb3\x39\xc1\x5d\x68\x0f\xcb\xcb\x78\xf4\x12\x0e\x6f\x82\x43\x4e\x9b\x99\xa9\x78\x6f\xbb\x53\x26\x7d\xcc\x8e\xdd\x84\xd8\xe5\x1b\xbb\x4f\xf1\x77\x4f\xb3\x41\xb3\x07\xb7\x8c\x26\x58\xb4\x93\x79\x21\x2b\xaf\x7d\x3e\x32\xe9\x5d\x9b\x7d\x62\x0f\xdb\x72\xcd\x1a\xd8\x5b\xa3\xd1\xf5\xcb\x43\xdc\x8b\x6c\xce\x26\xf2\x6f\xef\xf1\xe0\xaa\xd9\x4d\x6e\xaf\x3d\x9f\x05\x53\x2a\x25\xa4\xec\x8c\x22\x7a\xd1\x21\x1f\x34\xd3\x42\xb0\xdd\x36\x0b\xa2\x31\x35\x87\xbe\xae\xd3\xcf\x5a\xab\x3e\xff\x38\x74\xcf\xcc\x51\xa2\x0f\xd4\x63\x18\x85\x10\x9b\x43\xba\x4b\x82\xb3\x3a\x86\x55\xd9\xa3\xfa\x38\x43\xa5\x7d\x00\x42\x9b\x08\x21\xd7\xd4\xdd\x1e\x7e\x18\xeb\xea\x1e\x1c\xb9\x17\xc5\xf8\xaa\x39\x43\xee\x45\x30\xa1\x3a\x25\x19\xe1\xa9\xdf\xb1\xfe\x67\x93\xde\x94\x68\x47\xe9\x28\x59\xc2\xb0\x58\x54\x3a\x06\xd5\x35\x99\xe1\xd8\x8f\x62\x53\xf5\x38\x80\x28\x44\xea\x75\x76\xba\x2e\xb3\xa9\x21\x9b\xfa\x47\x51\xb2\x52\x0a\xaf\xfd\x7a\xd9\x1e\x46\x83\x51\x53\x92\x75\xf9\xa3\x20\x25\xf1\xba\x36\x87\x3d\x73\x5f\xda\x1b\xc4\xfe\xa6\x69\x49\xad\x5c\x29\xa4\xb6\xd3\x80\x7c\x90\xa4\x72\xff\xee\xe1\x99\xa1\x3a\xbe\x85\xea\x0a\x0e\xcb\xd2\x44\xdd\x87\x4c\xe7\xa3\x9d\x0a\x65\x0b\x1e\x74\x8e\x09\xa5\xbc\x9c\xa3\x92\xf4\xe6\xe6\xd3\x11\x27\x0a\x06\x93\xac\x52\x16\x2c\x85\xd8\xe0\xf3\x1c\x62\xc1\x96\xbf\x42\x1e\xbc\x6b\xd1\x93\x90\x45\x96\x99\x22\x5e\x3e\x0f\x5c\xea\xbe\xd7\x83\x7e\xaf\x13\x8d\x2e\x21\x15\x49\x55\x10\xae\xf7\x8c\x5e\x0d\xd4\xd1\x6f\xf4\x8f\x6f\x87\x34\xf3\xdb\x0f\x17\x97\xfb\x90\x36\x78\x12\xe0\x52\xb5\xce\x5c\x37\x54\xcf\xfa\x9f\xbd\x96\x1d\x94\x17\xf2\xed\x3a\x7d\x9f\xfc\x36\x86\xc5\x50\x23\xec\x6d\x80\x3a\xfa\xcf\x19\xff\x07\x00\x00\xff\xff\xa0\x2b\x4a\x17\x11\x0d\x00\x00")

func pluginsPluginsTomlBytes() ([]byte, error) {
	return bindataRead(
		_pluginsPluginsToml,
		"plugins/plugins.toml",
	)
}

func pluginsPluginsToml() (*asset, error) {
	bytes, err := pluginsPluginsTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/plugins.toml", size: 3345, mode: os.FileMode(420), modTime: time.Unix(1470525713, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
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
	"plugins/plugins.toml": pluginsPluginsToml,
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
	"plugins": &bintree{nil, map[string]*bintree{
		"plugins.toml": &bintree{pluginsPluginsToml, map[string]*bintree{}},
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

