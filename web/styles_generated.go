// Code generated by go-bindata.
// sources:
// dist/hydrocarbon.min.css
// DO NOT EDIT!

package web

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

var _distHydrocarbonMinCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x54\xdd\x6e\xdb\x3a\x0c\x7e\x95\x00\xc3\xb0\x1b\xa6\x88\xb3\xb4\xe8\xe4\x57\x38\x77\x7b\x02\xca\xa2\x6c\x9e\x52\x3f\x95\x68\x2f\x99\xe1\x77\x3f\xb0\x93\xae\xe9\xd9\x76\x25\x9a\xfa\xfb\xfe\x64\x04\xb4\xb6\x00\x76\x25\xc5\x4b\x00\x74\xae\x50\xad\x80\x39\x0b\x29\x60\x51\xee\x84\x00\x2b\x3b\x02\x1c\x1d\x27\xb0\x60\xb9\x07\x2b\xa9\x7b\x79\x1d\x93\x12\xd8\xe4\x2e\xd0\x61\x9c\xb0\x42\x87\x59\x39\x45\xe8\x28\x2a\x15\xe8\x58\x09\xba\xe4\x08\x9c\x03\x47\x02\x8e\x14\x59\x2a\x38\x1f\xc1\xf1\x04\x4e\xc0\x29\x50\x00\x0a\x96\x1c\x78\x26\x71\x95\x14\x3c\xf7\x6f\x47\x79\xee\xc7\x42\xe0\x53\x5a\x4f\xf4\xa9\x04\x18\x1a\x18\x8e\x30\x7c\x85\xe1\x04\xc3\x23\x0c\x4f\x30\x10\x3a\x2a\x30\xf4\x25\x8d\x19\x06\x0d\x02\x0c\xec\x0b\x06\x02\x0e\x3d\x70\xac\xf0\x62\x1d\x08\x5a\x12\x10\xea\x29\x3a\x10\x86\x80\xe5\x05\x02\xc5\x11\x22\x4e\x90\xec\xbf\xd4\x29\x24\x81\x34\x6a\x1e\x15\x32\xe4\x42\xf0\x0a\x65\xb4\x17\xa8\x50\x31\x64\xa8\xd4\x6d\xb8\x6a\x40\x11\xa8\x19\x23\x54\x2d\xfc\x42\xeb\x90\x62\x0f\x75\xb4\x50\xc7\x10\xb0\x5c\xa0\x8e\x19\x14\xad\x10\xe8\x26\x93\x3a\xd0\x95\x09\xe8\x00\xba\x82\x06\xe5\x40\xa0\x05\x54\x61\x84\x51\x60\xc2\x02\x13\x3b\x4a\x73\xc0\xd2\x73\x34\x87\x36\xa3\x73\x1c\x7b\x73\x68\x6d\x2a\x8e\x8a\x39\xb4\x13\xad\xce\xa0\xec\x51\xb8\x8f\xc6\x62\x25\xe1\x48\xad\x4f\x51\x0d\xc7\x81\x0a\xeb\xf6\xb1\xaf\xfc\x93\x4c\x73\x38\x7c\x5e\x3e\x9a\xf9\xe6\xc4\x5f\x95\xfe\x28\xe9\x2f\x8d\x6e\xf4\x67\xc7\x35\x0b\x5e\xcc\x16\x84\x65\x25\x37\xaf\x08\xf6\x03\x71\x3f\xa8\x69\x96\x24\x30\xca\x2c\x5c\x75\x5f\xf5\x22\x64\x62\x8a\xb4\xdc\xe5\xe6\x75\xde\xc6\xfa\xff\x09\x83\x7e\xbd\xff\xae\x61\xc9\xa7\xd5\x87\xdb\xcc\xeb\xad\x31\x77\x29\x2a\x45\x35\x5f\xbe\xb4\x6f\xe5\x76\xd6\xa6\xf8\x7c\x15\x6b\x5f\x33\x76\xf7\xea\xed\xbb\x24\x82\xb9\x92\x79\x2b\xae\xe8\x2d\x76\x2f\x2b\xd7\xe8\xd6\x15\xa9\x98\x4f\xde\xfb\xf6\x56\x36\xa7\xe6\xf9\xf8\xf5\xaa\xa8\xc7\xc0\x72\x31\xdf\x31\xd6\xfd\x77\x2a\xec\x5b\xa5\xb3\xee\x0b\x45\x47\x65\xbd\x29\x65\xe5\xc0\x3f\xe9\x1f\xea\xd9\xb2\xb0\x5e\x16\xc7\xd3\xc3\x0d\xe2\x9b\xaf\x4d\xa1\xb0\xc3\x51\x53\xfb\x83\x9d\x0e\xe6\xf9\xf1\xf3\xbd\x63\x4f\xf9\xdc\x7e\x10\xf4\xe1\xb4\xf0\x7c\x5d\xb0\xc9\xc9\x8a\xc2\xdd\x32\x34\xf3\xfb\xae\xe3\x29\x9f\x97\xe1\x78\xdf\x39\xe4\xf3\xb2\x3e\x9a\x39\xa7\xca\xab\x73\xa6\x90\xa0\xf2\x44\x0b\xc7\x3c\xea\xdf\xfa\x86\xe3\x84\xc2\xee\x26\xa3\x69\xf2\x79\x57\x93\xb0\xdb\x7d\xea\xe8\xe8\x4e\x87\xe5\x9a\x93\xf7\xfd\x68\x6b\x92\x51\xa9\x2d\x1b\xe0\x55\x6f\xd5\x14\xcc\xa1\x15\xf2\x7a\x97\xe3\x95\xf9\x55\xb3\x6b\x7a\xaf\x3f\x8b\x7b\xf2\xc7\x7c\x5e\x46\x99\xd3\x44\xc5\x4b\xfa\x61\x06\x76\x8e\x62\xfb\xa7\x17\xf1\x9b\x6b\xc7\xd3\xe3\xb7\x6f\x4f\xed\x7b\xf0\xf6\x7a\xc9\xb7\xf4\x09\xff\xca\x2d\xc7\x4d\xdd\x2d\x65\xad\x97\x84\x6a\x56\x98\x8b\xf0\xc3\x86\x7f\xbe\xf6\xb6\x7a\x11\xde\xe1\xc7\xc4\xbf\x93\x39\xe5\xf3\x6e\x73\xeb\x2e\x35\xbf\x93\xdb\x3a\x8e\xba\x54\x70\x13\xeb\x06\x67\x87\x66\x58\x59\xfe\x21\x7d\x8d\x3f\xb9\xe7\x66\xf9\x2f\x00\x00\xff\xff\xd2\xd3\xb7\x89\xa1\x05\x00\x00")

func distHydrocarbonMinCssBytes() ([]byte, error) {
	return bindataRead(
		_distHydrocarbonMinCss,
		"dist/hydrocarbon.min.css",
	)
}

func distHydrocarbonMinCss() (*asset, error) {
	bytes, err := distHydrocarbonMinCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dist/hydrocarbon.min.css", size: 1441, mode: os.FileMode(420), modTime: time.Unix(1482511063, 0)}
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
	"dist/hydrocarbon.min.css": distHydrocarbonMinCss,
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
	"dist": &bintree{nil, map[string]*bintree{
		"hydrocarbon.min.css": &bintree{distHydrocarbonMinCss, map[string]*bintree{}},
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
