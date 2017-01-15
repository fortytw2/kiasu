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

var _distHydrocarbonMinCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x54\x5d\x8e\xdb\x38\x0c\xbe\x4a\x80\xa2\xe8\x0b\x3d\x88\xd3\x74\xd0\x91\xaf\xb0\x6f\x3d\x01\x65\xd1\x36\x37\xb4\xa4\x91\x68\x4f\x52\xc3\x77\x5f\xc8\x49\x3a\x99\x6d\xfb\x64\x8a\xfa\xfb\xfe\x64\x04\xb4\x36\x01\xb6\x29\xf8\xcb\x08\xe8\x5c\xa2\x9c\x01\x63\x14\x52\xc0\xa4\xdc\x0a\x01\x66\x76\x04\x38\x39\x0e\x60\xc1\x72\x0f\x56\x42\x7b\x7a\x9d\x82\x12\xd8\xe0\x2e\xd0\xa2\x9f\x31\x43\x8b\x51\x39\x78\x68\xc9\x2b\x25\x68\x59\x09\xda\xe0\x08\x9c\x03\x47\x02\x8e\x14\x59\x32\xb8\xce\x83\xe3\x19\x9c\x80\x53\xa0\x11\x68\xb4\xe4\xa0\x63\x12\x97\x49\xa1\xe3\xfe\x7e\x54\xc7\xfd\x94\x08\xba\x10\xca\x89\x5d\x48\x23\x0c\x35\x0c\x07\x18\xbe\xc2\x70\x84\xe1\x1b\x0c\xcf\x30\x10\x3a\x4a\x30\xf4\x29\x4c\x11\x06\x1d\x05\x18\xb8\x4b\x38\x12\xf0\xd8\x03\xfb\x0c\x27\xeb\x40\xd0\x92\x80\x50\x4f\xde\x81\x30\x8c\x98\x4e\x30\x92\x9f\xc0\xe3\x0c\xc1\xfe\x4b\xad\x42\x10\x08\x93\xc6\x49\x21\x42\x4c\x04\xaf\x90\x26\x7b\x81\x0c\x19\xc7\x08\x99\xda\x0d\x57\x1e\x51\x04\x72\x44\x0f\x59\x13\x9f\xa8\x7c\x82\xef\x21\x4f\x16\xf2\x34\x8e\x98\x2e\x90\xa7\x08\x8a\x56\x08\x74\x93\x49\x1d\x68\x61\x02\x3a\x80\x16\xd0\xa0\x3c\x12\x68\x02\x55\x98\x60\x12\x98\x31\xc1\xcc\x8e\xc2\x32\x62\xea\xd9\x9b\x7d\x13\xd1\x39\xf6\xbd\xd9\x37\x36\x24\x47\xc9\xec\x9b\x99\x8a\x33\x28\x15\x0a\xf7\xde\x58\xcc\x24\xec\xa9\xe9\x82\x57\xc3\x7e\xa0\xc4\xba\x0d\xaa\xcc\x3f\xc9\xd4\xfb\xfd\xe7\xf5\xa3\x99\x77\x27\xfe\xaa\xf4\x47\x49\x7f\x69\x74\xa3\xbf\x38\xce\x51\xf0\x62\xb6\x20\xac\x85\xdc\x52\x10\x54\x03\x71\x3f\xa8\xa9\xd7\x20\x30\xc9\x22\x9c\xb5\xca\x7a\x11\x32\x3e\x78\x5a\x1f\x72\xf3\xba\x6c\xdf\xfc\xff\x09\x83\x5d\xb9\xff\xa1\x61\xa9\x0b\xc5\x87\xdb\xcc\xeb\xad\xb1\xb4\xc1\x2b\x79\x35\x5f\xbe\x34\xf7\x72\x3b\x6b\x53\x7c\xb9\x8a\x55\xe5\x88\xed\xa3\x7a\x55\x1b\x44\x30\x66\x32\xf7\xe2\x8a\xde\x62\x7b\x2a\x5c\xbd\x2b\x2b\x42\x32\x9f\xba\xae\x6b\x6e\x65\x7d\xac\xbf\x1f\xbe\x5e\x15\xed\x70\x64\xb9\x98\x1f\xe8\x73\xf5\x83\x12\x77\x8d\xd2\x59\xab\x44\xde\x51\x2a\x37\x85\xa8\x3c\xf2\x4f\xfa\x87\x7a\xb6\x2c\xac\x97\xd5\xf1\xfc\x74\x83\x78\xf7\xb5\x4e\x34\xee\x70\xd2\xd0\xbc\x85\xe4\xaa\xb7\x84\xd1\xd8\x44\x78\xaa\xca\xb8\x79\x63\xa7\x83\x79\xd9\x7f\x7e\xb4\xf1\x39\x9e\x9b\x0f\x2a\x3f\x1d\x57\x5e\xae\x0b\x36\x8d\x59\x51\xb8\x5d\x87\x7a\x79\xdf\x75\x38\xc6\xf3\x3a\x1c\x1e\x3b\xfb\x78\x5e\xcb\x4b\x5a\x62\xc8\x5c\xec\x34\x89\x04\x95\x67\x5a\xd9\xc7\x49\xff\xd6\x37\xec\x67\x14\x76\x37\x6d\x4d\x1d\xcf\xbb\x1c\x84\xdd\xee\x53\x4b\x07\x77\xdc\xaf\xd7\xf0\xbc\xef\x47\x9b\x83\x4c\x4a\x4d\xda\x00\x17\x13\x54\xc3\x68\xf6\x8d\x50\xa7\x0f\xe1\x2e\x72\x5c\x85\xbc\x46\xfa\xfa\x07\x79\x24\x7f\x88\xe7\x75\x92\x25\xcc\x94\x3a\x09\x6f\x66\x60\xe7\xc8\x37\x7f\x7a\x26\xbf\x59\x79\x38\x7e\x7b\x79\x79\x6e\xde\xd3\x58\xe9\x25\xde\x22\x29\xfc\x2b\xcc\xec\x37\x75\xb7\xe8\x35\x9d\x04\x54\x53\x60\xae\xc2\x4f\x1b\xfe\xe5\xda\xdb\xea\x55\x78\x87\x1f\x9f\xc1\x3b\x99\x63\x3c\xef\x36\xb7\x1e\xa2\xf4\x3b\xb9\xad\xe3\xa8\x0d\x09\x37\xb1\x6e\x70\x76\x68\x86\xc2\xf2\x0f\x91\xac\xbb\xa3\xfb\x5e\xaf\xff\x05\x00\x00\xff\xff\x77\x10\xda\x58\xb6\x05\x00\x00")

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

	info := bindataFileInfo{name: "dist/hydrocarbon.min.css", size: 1462, mode: os.FileMode(420), modTime: time.Unix(1484503282, 0)}
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

