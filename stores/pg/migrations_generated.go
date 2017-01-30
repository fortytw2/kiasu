// Code generated by go-bindata.
// sources:
// schema/1_init.sql
// schema/2_auto_update_column.sql
// schema/3_favorites.sql
// schema/4_indexes.sql
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

var _schema1_initSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x56\xd1\x6e\xe2\x3a\x10\x7d\xcf\x57\xcc\x1b\x54\xb7\x95\xae\xee\xd3\x95\x56\xfb\x90\x52\xa3\x8d\x0a\xa1\x1b\x82\x44\xbb\x5a\x45\x56\x3c\x14\xab\xc1\xce\xda\x13\x68\xff\x7e\x65\x97\xb2\xd0\x04\x36\xa0\xa5\xda\xd7\x78\xc6\x73\xe6\xcc\x99\x13\xf7\x12\x16\xa6\x0c\xd8\x34\x65\xf1\x38\x1a\xc5\x10\xf5\x21\x1e\xa5\xc0\xa6\xd1\x38\x1d\x43\xf9\x98\x9b\x97\x92\xf4\xa7\x20\x58\x47\xa6\xe1\xf5\x80\xc1\x0c\x51\x58\xe8\x06\x00\x52\xc0\x64\x12\xdd\xc0\x5d\x12\x0d\xc3\xe4\x1e\x6e\xd9\x3d\xdc\xb0\x7e\x38\x19\xa4\xf0\x88\x2a\x33\x5c\x09\xbd\xc8\xaa\x4a\x8a\xee\xc5\x65\x10\x00\x94\x45\xf5\x28\x15\xa4\x6c\x9a\xfa\x52\xf1\x64\x30\xb8\x74\x37\x29\x49\x92\x17\x59\x65\x8a\xf7\x87\x01\x40\x6e\x90\x13\x8a\x8c\x13\xa4\xd1\x90\x8d\xd3\x70\x78\x97\x3e\x6c\x62\x36\x35\x95\x5e\xb9\x3a\x00\x55\x29\x8e\x4b\x28\xb8\xa5\xcc\xe0\xcc\xa0\x9d\xd7\xf2\x36\x01\xa8\x7e\x54\x58\x35\x9e\x2b\xbe\xc0\x5d\xe4\x30\x89\xa3\xaf\x13\xe6\x0e\x05\xda\xdc\xc8\x92\xa4\xae\xb5\x1e\x00\xcc\xf1\x39\xcb\x75\xa1\x4d\x03\x2d\xb9\x56\x7b\x38\xe9\x7d\x61\xbd\xdb\x2e\x9b\xa6\x49\xd8\x4b\xbb\x0e\xcd\xc3\x28\x66\xd0\x4f\x46\xc3\x2d\xbe\x2e\xe0\x33\x74\xfe\xed\xf8\x1e\x0f\x65\xfc\x22\xac\x6d\x46\x8d\xb1\xa3\x12\xb7\x98\x7c\xcb\x0b\x2e\xde\x0b\xad\xd4\x96\x4e\x12\x1a\x78\x8d\x66\x6f\x59\x09\xeb\xb3\x84\xc5\x3d\x36\x5e\x6b\xf7\x83\xc5\x75\x48\x57\xae\xc7\x86\x83\x00\x80\x24\x15\x58\x97\x44\x4d\x0d\x5b\x3a\xcb\xb5\x22\x54\xf4\x17\xaa\x65\xd3\x66\xdb\x84\x26\x65\xd5\x15\x52\x59\x34\xa7\x5a\xd1\xb9\xe7\x1e\x00\x70\xc5\x8b\x17\x92\xb9\x85\xeb\xd1\x68\xc0\xc2\xb8\x1e\x3c\xe3\x85\x45\x77\x3b\x2e\xb8\xdc\x3f\x59\x54\xde\x8b\x51\x64\x25\xb7\x76\xa5\x8d\x68\x18\x32\xcf\x49\x2e\x71\x7f\xad\x8e\x2f\xd6\x59\x2b\x65\x26\xcd\x02\xc5\x31\xd1\xdc\x19\x58\x46\xfa\x09\x1b\x1c\xdc\x7f\xce\x8e\x24\x35\x00\xb0\x64\x64\x89\x59\x5e\x59\xd2\x0b\x34\x6e\x69\x77\x49\xd8\xe0\xf1\x50\xae\xae\xc0\x22\x01\xcd\x11\x04\xce\x78\x55\x10\x90\x86\x95\x36\x4f\xb0\x92\x34\x5f\x5f\x07\x64\x24\x2f\xdc\x7a\x71\x29\xb2\x4a\x91\xa3\xf6\xb7\x78\xe0\x1f\x90\x8a\xd0\x2c\x79\x01\x9d\xff\xfe\x07\xc1\x5f\x6c\xc7\x83\x9c\xe9\x42\x78\x6c\x16\x08\x9f\xe9\xdb\xf7\xf6\xdb\xf5\x9e\x97\xd6\x1b\xb3\x41\xde\x36\xe3\x4f\xec\x71\x7d\xc5\x5e\x5b\x3f\xd1\x86\x3f\x60\xc7\xfc\x8f\xd7\x0d\x65\x47\x8d\x6b\xff\x3f\x7e\x5c\xe7\x21\xd1\x20\x17\x99\x25\x4e\x95\xc5\x13\xa9\x74\x56\xd7\xf4\x47\x7b\xb5\xc0\xed\xde\x9d\xd7\xba\xc8\x5a\xe8\xeb\xff\x74\x87\x0a\x0f\xac\xf5\x74\x04\x2e\x65\x8e\xb5\x15\xf5\xef\x23\x9d\xf3\x3d\xef\x9b\xc3\x36\xcf\x0f\x33\x67\xd1\x5a\xa9\xd5\x99\x49\xfb\x90\x67\x80\x54\x4b\x5e\xc8\xa6\x24\x0f\x00\x9f\x4b\x69\xd0\xb6\xbb\x6f\x8f\x55\xad\x6d\xf8\xec\x0f\x80\xdd\x56\xb6\xa7\xf7\x33\x00\x00\xff\xff\x38\x4e\xc2\xe5\x54\x0c\x00\x00")

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

	info := bindataFileInfo{name: "schema/1_init.sql", size: 3156, mode: os.FileMode(420), modTime: time.Unix(1485725125, 0)}
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

var _schema4_indexesSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\xf5\xf3\x0c\x0c\x75\x55\xf0\xf4\x73\x71\x8d\x50\x28\xc8\x2f\x2e\x29\x8e\xcf\x4c\x89\xcf\x4c\xa9\x50\xf0\xf7\x83\xf0\x15\x34\x32\x53\x34\xad\xb9\xb8\xa0\x1a\x90\x55\xa6\xa5\xa6\x82\xd4\xc6\x83\x78\xa9\x29\xf1\x89\x25\x68\x1a\xa1\xf2\x3a\x0a\x70\x05\x18\x06\x15\xa7\x16\x17\x67\xe6\xe7\x15\xc7\x97\x16\xa7\x16\xc5\x67\xa6\x80\x74\xc3\xc4\x14\x34\xa0\x82\x9a\xd6\x5c\x80\x00\x00\x00\xff\xff\x15\x8e\xdc\xfa\xaf\x00\x00\x00")

func schema4_indexesSqlBytes() ([]byte, error) {
	return bindataRead(
		_schema4_indexesSql,
		"schema/4_indexes.sql",
	)
}

func schema4_indexesSql() (*asset, error) {
	bytes, err := schema4_indexesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/4_indexes.sql", size: 175, mode: os.FileMode(420), modTime: time.Unix(1485806571, 0)}
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
	"schema/4_indexes.sql": schema4_indexesSql,
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
		"4_indexes.sql": &bintree{schema4_indexesSql, map[string]*bintree{}},
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

