// Code generated by go-bindata.
// sources:
// schema/1_init.sql
// schema/2_auto_update_column.sql
// schema/3_favorites.sql
// DO NOT EDIT!

package pg

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

var _schema1_initSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x56\xc1\x6e\xdb\x30\x0c\xbd\xfb\x2b\x78\x4b\x8a\xf5\x30\xec\x34\x60\xd8\xc1\x4d\x15\xcc\x68\xe2\x74\x8e\x0c\xa4\x1d\x06\x41\xb0\x98\x56\x98\x23\x19\x92\x9c\xb6\x7f\x3f\x48\x4d\xb3\xa4\x76\x32\x27\x68\x8a\x5e\x4d\x52\x24\xdf\x23\x1f\x3d\xc8\x48\x4c\x09\x90\x19\x25\xe9\x34\x99\xa4\x90\x0c\x21\x9d\x50\x20\xb3\x64\x4a\xa7\x50\xdd\x15\xe6\xa9\x72\xfa\x5b\x14\xad\x3c\x69\x7c\x31\x22\x30\x47\x14\x16\xfa\x11\x80\x14\x90\xe7\xc9\x25\x5c\x67\xc9\x38\xce\x6e\xe0\x8a\xdc\xc0\x25\x19\xc6\xf9\x88\xc2\x1d\x2a\x66\xb8\x12\x7a\xc1\xea\x5a\x8a\xfe\xd9\x79\x14\x01\x54\x65\x7d\x27\x15\x50\x32\xa3\x21\x55\x9a\x8f\x46\xe7\xfe\x25\x25\x9d\xe4\x25\xab\x4d\xf9\xda\x18\x01\x14\x06\xb9\x43\xc1\xb8\x03\x9a\x8c\xc9\x94\xc6\xe3\x6b\x7a\xbb\xf6\x59\xe7\x54\xfa\xc1\xe7\x01\xa8\x2b\x71\x58\x80\xc1\xb9\x41\x7b\xdf\x08\xf1\x36\xc5\x17\xb8\x5d\x14\xe4\x69\xf2\x33\x27\xde\x28\xd0\x16\x46\x56\x4e\xea\x46\x57\x11\xc0\x3d\x3e\xb2\x42\x97\xda\xb4\x74\x5c\x68\xb5\xa3\xdd\xc1\x0f\x32\xb8\xea\x93\x19\xcd\xe2\x01\xed\xfb\x6a\x6e\x27\x29\x81\x61\x36\x19\x6f\x40\x71\x06\xdf\xa1\xf7\xb9\x17\xca\xdf\x17\xf1\x0f\x8b\xae\x11\x9b\x60\xbc\xc4\x44\x67\xaf\xc7\xa0\xd2\xd6\x1d\x35\x06\x10\x26\x88\xbd\x44\x65\x64\x48\x32\x92\x0e\xc8\x74\x35\x59\x1f\x88\x7a\xdf\x63\x8b\x21\x02\x70\xd2\x95\xd8\x64\xb5\x41\xe8\xc6\xa8\x14\x5a\x39\x54\xee\x03\x12\xbe\x6e\xf3\x6d\x27\xa4\xb6\x68\x8e\x15\x8a\x53\xf3\x1e\x01\x70\xc5\xcb\x27\x27\x0b\x0b\x17\x93\xc9\x88\xc4\x69\xd3\x79\xce\x4b\x8b\xfe\x75\x5c\x70\xb9\x9b\x59\x54\x41\x29\x51\xb0\x8a\x5b\xfb\xa0\x8d\x68\x21\x99\x17\x4e\x2e\x71\x77\xae\x5e\x48\xd6\x5b\x4d\xca\x5c\x9a\x05\x8a\x43\xbc\xb9\xd7\x20\xe6\xf4\x1f\x6c\xd1\xd7\xf0\x99\x1d\x08\x6a\x04\x60\x9d\x91\x15\xb2\xa2\xb6\x4e\x2f\xd0\xf8\xa5\xdd\x06\x61\x5d\x4f\x28\xa5\xe2\x52\xb0\x5a\x39\x8f\xd5\x7f\x13\xc0\x27\x90\xca\xa1\x59\xf2\x12\x7a\x5f\xbe\x82\xe0\x4f\xb6\x17\xb2\xce\x75\x29\x42\x32\x0b\x0e\x1f\xdd\xaf\xdf\xdd\xd7\xe5\x75\xa3\x9d\x57\x60\x5d\x79\xd7\x88\xb7\x58\xcc\xe6\xce\x3c\xb7\x7e\xa4\xae\xbe\xc3\xd2\x84\x63\xe8\x49\xd9\x1a\xaf\x95\xa0\x1f\x4e\xd7\x69\x40\x34\xc8\x05\xb3\x8e\xbb\xda\xe2\x91\x50\x7a\xed\x6a\x3b\x51\xcf\x9a\xb6\xd9\xbb\x17\x4f\xef\xd9\x70\x7d\x3e\x90\x5b\x50\x84\xc2\x3a\xb3\x23\x70\x29\x0b\x6c\xec\x9c\x37\x95\xba\xe0\x3b\xfe\x39\xf6\xeb\x36\xdf\x8f\x9c\x45\x6b\xa5\x56\x27\x06\xed\x5d\xee\xba\x54\x4b\x5e\xca\xb6\xa0\x50\x00\x3e\x56\xd2\xa0\xed\xf6\xde\x0e\xa9\x5a\xe9\xea\xc9\x2f\xfa\x76\x2b\x9b\xec\xfd\x0d\x00\x00\xff\xff\x1a\x1e\x35\xbd\xc3\x0b\x00\x00")

func schema1_initSqlBytes() ([]byte, error) {
	return bindataRead(
		_schema1_initSql,
		"schema/1_init.sql",
	)
}

func schema1_initSql() (*asset, error) {
	bytes, err := schema1_initSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/1_init.sql", size: 3011, mode: os.FileMode(420), modTime: time.Unix(1485019493, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schema2_auto_update_columnSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\xcf\xc1\x6a\x83\x40\x10\xc6\xf1\xfb\x3e\xc5\x77\x10\x92\x5c\xfa\x02\xd2\x83\x59\x47\x2b\x94\x5d\x99\x28\xe9\x4d\x16\xdc\x4a\x41\xd4\x66\x56\xfa\xfa\xa5\xa6\xa1\x2d\xa5\x42\xc1\xb3\xff\xf1\xf7\xad\x66\x4a\x2a\x82\x65\x30\x95\x8f\x89\x26\x64\xb5\xd1\x55\x61\x0d\xc4\x87\x66\x9e\x5a\x17\x7c\xdb\xb8\xb0\x3f\x28\xa6\xaa\x66\x73\x42\xc5\x45\x9e\x13\x23\x39\x21\x8a\xd4\x91\xf2\xc2\x28\x00\x30\x74\xbe\xfb\x3a\xc0\x3d\x86\xf1\x6d\x7f\x88\x97\x6f\xd7\xdb\x8f\x24\x56\x64\xd2\x58\x45\x11\x7a\x37\x74\xb3\xeb\x3c\x76\x53\x3f\x75\xf2\xda\xef\x62\xa5\x3e\x07\xdd\x8c\xeb\xff\x9a\x59\xfc\x45\xbe\xad\xc1\x91\x32\xcb\x84\xba\x4c\x97\xf9\x06\x4b\x81\xcc\x32\x28\xd1\x0f\x60\x7b\x06\x3d\x91\xae\x2b\x42\xc9\x56\x53\x5a\x33\xfd\x7a\xd2\x9f\xde\xb3\xf7\xed\xba\xb7\x14\x9b\x79\x17\xef\xda\x46\x82\x0b\xb3\xf8\x75\xf7\x47\xb9\x99\x3f\x8d\x12\xd6\xdd\xa5\xd8\xcc\x13\x2f\xf2\x32\x0e\xeb\xe4\x2d\xfa\xbf\xfa\x1e\x00\x00\xff\xff\x0f\xb6\x02\xfa\xd6\x02\x00\x00")

func schema2_auto_update_columnSqlBytes() ([]byte, error) {
	return bindataRead(
		_schema2_auto_update_columnSql,
		"schema/2_auto_update_column.sql",
	)
}

func schema2_auto_update_columnSql() (*asset, error) {
	bytes, err := schema2_auto_update_columnSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/2_auto_update_column.sql", size: 726, mode: os.FileMode(420), modTime: time.Unix(1484486749, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schema3_favoritesSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x91\x4f\x4f\x83\x40\x14\xc4\xef\xfb\x29\xe6\x56\x48\x3c\x78\x6f\x3c\x6c\x97\x47\x4b\x0a\x2c\x79\xdd\x8d\x6d\x2f\x1b\x22\xab\xe9\x41\x68\xf8\xa3\x5f\xdf\x40\xa4\xa2\x31\x26\x5e\x5f\xe6\x37\x99\x99\xa7\x98\xa4\x21\x18\xb9\x49\x09\x43\xe7\x5b\xf7\x5c\xbe\x35\xed\xa5\xf7\x1d\x02\x01\x5c\x2a\x58\x9b\x44\x28\x38\xc9\x24\x9f\xb0\xa7\x13\x22\x8a\xa5\x4d\x0d\x5e\x7c\xed\xda\xb2\xae\x9a\x57\x37\x0c\x97\x2a\x08\xef\x84\x00\x9e\x5a\x5f\xf6\xbe\x72\x65\x0f\x93\x64\x74\x30\x32\x2b\xcc\x19\xb9\x36\xc8\x6d\x9a\xde\xe8\xba\x79\x1f\x09\x60\xb8\x56\xff\x01\x04\x70\x6d\xba\xde\xcd\xc9\x98\x62\x62\xca\x15\x1d\xa6\x7b\x77\x03\x27\xef\xb1\xd1\x2f\xca\xf1\xbe\x54\x0a\x40\xed\x48\xed\x03\x3a\x1a\x96\xca\x04\x63\x92\xb3\xce\x09\x31\xeb\x6c\xd1\x29\xc4\x03\x56\xf7\xab\x29\xf8\x5f\xc4\x57\xa9\x99\x10\xe1\x5a\x88\x79\x6e\x4e\xb6\x5b\xe2\x4f\x95\xfb\xbe\xbb\x5b\x0c\xb2\xa1\x58\x33\xc1\x16\xd1\x88\xe9\xfc\xe7\x8b\x62\xcd\x20\xa9\x76\x60\xfd\x08\x3a\x92\xb2\x86\x50\xb0\x56\x14\x59\x26\x74\xbe\x5f\xb8\x05\xe1\x5a\x7c\x04\x00\x00\xff\xff\xdf\x15\xbe\xe2\xf0\x01\x00\x00")

func schema3_favoritesSqlBytes() ([]byte, error) {
	return bindataRead(
		_schema3_favoritesSql,
		"schema/3_favorites.sql",
	)
}

func schema3_favoritesSql() (*asset, error) {
	bytes, err := schema3_favoritesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/3_favorites.sql", size: 496, mode: os.FileMode(420), modTime: time.Unix(1484486749, 0)}
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
	"schema/1_init.sql": schema1_initSql,
	"schema/2_auto_update_column.sql": schema2_auto_update_columnSql,
	"schema/3_favorites.sql": schema3_favoritesSql,
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
	"schema": &bintree{nil, map[string]*bintree{
		"1_init.sql": &bintree{schema1_initSql, map[string]*bintree{}},
		"2_auto_update_column.sql": &bintree{schema2_auto_update_columnSql, map[string]*bintree{}},
		"3_favorites.sql": &bintree{schema3_favoritesSql, map[string]*bintree{}},
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

