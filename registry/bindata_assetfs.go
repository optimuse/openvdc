// Code generated by go-bindata.
// sources:
// ../schema/none.json
// ../schema/v1.json
// ../schema/vm/lxc.json
// ../schema/vm/null.json
// ../schema/vm/qemu.json
// DO NOT EDIT!

package registry

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

	"github.com/elazarl/go-bindata-assetfs"
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

var _schemaNoneJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8e\xb1\x4e\x03\x31\x0c\x86\xf7\x3c\x85\x15\x18\x40\x6a\x39\x40\x4c\x59\x41\xac\x48\x08\xb1\xa0\x0a\x85\x9c\xb9\xa6\x6a\xec\xe0\xf8\x86\xaa\xea\xbb\xa3\x5c\xe0\x54\x26\x3a\x64\xc8\xe7\xef\xf7\xef\xbd\x01\xb0\xe7\x25\xac\x31\x79\xeb\xc0\xae\x55\xb3\xeb\xba\x4d\x61\x5a\x36\x7a\xc5\x32\x74\xbd\xf8\x4f\x5d\x5e\xdf\x75\x8d\x9d\xd9\x45\xcd\xf5\x58\x82\xc4\xac\x91\xa9\x66\x9f\x32\xd2\xeb\xc3\x3d\x3c\x63\xe1\x51\x02\xc2\x0b\xa6\xbc\xf5\x8a\x0e\x88\x09\xe1\xe2\x91\x05\x14\x8b\x46\x1a\x80\x69\xbb\xbb\x6c\x6b\x74\x97\xb1\xe6\xf9\x63\x83\x41\x1b\x13\xfc\x1a\xa3\x60\x6f\x1d\xbc\x19\x80\x5f\xcb\x00\xac\xa6\x79\x16\xce\x28\x1a\xb1\x58\x07\xfb\x66\xbc\x07\x4e\x09\x49\x67\x72\xb4\xbb\xa8\x44\x1a\xec\x84\x0f\x0b\x73\x3c\x9b\x5d\xa4\x31\xcd\x7d\x13\xa9\x67\xdb\x9f\xef\xea\x4f\x36\x7b\xf1\xe9\xe6\xd4\xa6\xc9\xbe\xfd\xd7\x36\xf5\x1d\xcc\x77\x00\x00\x00\xff\xff\x57\x3a\x39\x38\x94\x01\x00\x00")

func schemaNoneJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaNoneJson,
		"schema/none.json",
	)
}

func schemaNoneJson() (*asset, error) {
	bytes, err := schemaNoneJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/none.json", size: 404, mode: os.FileMode(420), modTime: time.Unix(1499405161, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaV1Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x91\x41\x4b\xc3\x40\x10\x85\xef\xf9\x15\xcb\xea\xb1\xcd\x2a\x78\xca\x55\xef\x05\x15\x2f\x52\x24\xdd\x4c\x9b\x2d\xd9\x99\xed\xec\xa4\x56\x4a\xfe\xbb\x6c\x92\x86\x8a\x82\xa5\x39\xed\xbe\x79\xdf\xdb\x3c\xe6\x98\x29\xa5\x5d\xa5\x0b\xa5\x6b\x91\x10\x0b\x63\xb8\xfc\xcc\x37\x4e\xea\x76\xd5\x46\x60\x4b\x28\x80\x92\x5b\xf2\xa6\x3c\xc4\xda\x50\x00\xdc\x57\xd6\xf8\x32\x0a\xb0\x89\xb6\x06\x5f\x9a\xfd\x7d\xbe\x8d\x84\x37\x7a\x96\x02\x6f\x07\xf5\x94\x5a\x18\x93\x86\xf3\x41\xcd\x89\x37\xa6\xe2\x72\x2d\xf3\xbb\x87\x91\x1f\xb9\x0a\xa2\x65\x17\xc4\x11\x26\x76\x11\x00\xdf\x9e\x1e\xd5\x33\x44\x6a\xd9\x82\x7a\x05\x1f\x9a\x52\x40\xbd\x0c\xf9\x3d\x24\x5f\x01\x92\x9b\x56\x5b\xb0\x32\x68\x0c\xbb\xd6\x31\xa4\x5e\xef\x99\x52\xc9\xe5\xa4\x81\x7e\x98\x2e\x63\x8e\xce\x94\x5a\xf6\x40\x60\x0a\xc0\xe2\x20\xea\x42\x1d\x07\xd7\x87\x25\xef\x01\x65\x52\xce\x1e\x8b\xc2\x0e\x37\xba\x97\xbb\xd9\xf9\x13\x17\x9a\x7f\x56\xbd\x2c\xff\xf4\xd7\x7f\xf8\xcf\xca\xf7\x3a\x21\x2c\xd6\x53\xfb\xf4\x1d\xa7\x53\x5a\x10\x43\x9a\xea\xdc\x20\x21\x8c\xab\x9b\x0c\xdd\xec\x3f\x6a\xef\x4d\x73\xb0\x57\x71\xd8\x36\xcd\x55\xe0\x0e\x7c\xfb\x0b\x1c\x4f\xcb\xec\x74\xeb\xb2\x2e\xfb\x0e\x00\x00\xff\xff\x9d\x38\xf1\x55\xd5\x02\x00\x00")

func schemaV1JsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaV1Json,
		"schema/v1.json",
	)
}

func schemaV1Json() (*asset, error) {
	bytes, err := schemaV1JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/v1.json", size: 725, mode: os.FileMode(420), modTime: time.Unix(1500457433, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaVmLxcJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x4b\x6f\x9b\x40\x10\xbe\xf3\x2b\x10\xc9\xad\x76\x49\xaa\x5c\xea\x5b\xd5\xaa\x52\x4f\xad\xaa\xaa\x97\x88\xa2\x65\x19\x60\x13\xf6\x91\x7d\xb8\x46\x16\xff\xbd\x5a\x8c\xcd\x6b\x49\x9c\x3a\x8e\x72\xb0\x64\xbe\x99\xf9\x66\x76\xe7\x63\x76\xd9\x7a\xbe\x1f\x5c\x2a\x5c\x00\x45\xc1\xca\x0f\x0a\xad\xc5\x2a\x0c\xef\x14\x67\xcb\x1d\xfa\x9e\xcb\x3c\x4c\x25\xca\xf4\xf2\xea\x26\xdc\x61\x17\xc1\xc2\xc6\xa5\xa0\xb0\x24\x42\x13\xce\x6c\xec\x77\x01\xec\xf7\x97\xcf\xfe\x4f\x50\xdc\x48\x0c\xfe\x2f\xa0\xa2\x44\x1a\x56\xfe\x9a\x86\xe5\x06\xef\xa2\x74\x25\xc0\xba\xf3\xe4\x0e\xb0\xde\x61\x12\x1e\x0c\x91\x90\x06\x2b\xff\xd6\xf3\xfd\xbd\x97\xe7\xfb\x51\x63\x17\x92\x0b\x90\x9a\x80\x0a\x56\xfe\x76\xe7\x11\x63\x4e\x29\x30\x7d\x40\x7a\xdc\x4a\x4b\xc2\xf2\xa0\x81\xeb\x85\xd7\xb7\x1d\x7c\x81\x19\x7a\xc8\xd7\x20\x6d\x95\x2d\x10\x0d\xa2\x29\x61\xf1\x1a\x0b\xe3\xca\x46\x98\x86\x1c\x64\xb0\xd8\x1b\x52\xc8\x90\x29\x6d\x65\xd7\x13\x12\x0a\x94\xcb\x2a\xce\x93\x93\x98\x4e\x2f\xe5\x45\xca\x60\x3c\x85\x38\x97\xdc\x08\xe5\xe2\x41\x52\xa2\xaa\x63\x31\x8c\x3c\x18\xf8\xa6\x81\x5a\x6f\x2d\x0d\x1c\x4c\xa4\x05\xbb\x6e\x6c\xc7\xcd\xac\x9d\x7d\xb1\xa5\xca\x0c\x61\x38\xa6\x80\x7d\x96\x6d\xd7\x73\x87\x1a\x5b\xcb\x44\x93\xfd\x88\x03\x10\xf5\x22\x1c\x2a\x1d\x66\xe9\x63\x53\xb5\x2e\x86\xd6\x89\x3e\xf7\x2a\x05\x5d\x8c\x7c\x9b\x4e\x30\x98\xa2\x40\x85\xae\xa6\xf0\xba\x44\x6c\x8a\x52\x84\xdd\x06\x51\x54\x2a\x18\x80\x51\xef\xa9\xee\xfb\x5b\x12\x94\xa6\xf2\xb9\x8b\x15\x48\x6b\x90\xcd\x18\xf9\x73\x7b\xb5\xfc\x88\x96\xd9\xa7\xe5\xd7\x68\xfb\xa1\xee\x9e\x56\xd1\xbb\xcb\x60\x36\x31\x11\xeb\x9b\xff\xc9\x9c\x71\x49\x91\x6e\x54\x2f\xd6\x37\x03\x7e\x6f\xfc\xaf\x1e\x48\xaf\xdc\xe0\x58\xb7\x23\xce\x25\xbe\x91\xa4\xec\x4b\x44\x18\xb1\xd3\x72\xa4\xc1\x04\x29\xe2\xa4\x7a\x8c\xce\x7f\x54\x71\xee\xf9\x38\xb7\x25\x03\x73\x3d\xda\x20\x09\x25\x20\x35\x95\xef\x33\x69\x90\xc4\xc5\xe9\x1c\xf9\x74\xa1\x4f\x71\x3c\xde\xd1\x43\x77\x28\xda\xfc\xe8\xef\xe6\x75\x67\x21\x6c\xc6\x32\xb3\xfd\x41\xca\xff\xb2\x92\xa3\xf4\x39\xad\x9c\x19\x37\x96\x8d\x28\x2d\xf9\x58\xb7\xfb\xb6\xf4\xd0\xe8\xd5\xb5\xd1\x96\xf6\x36\x14\xb6\x46\x92\xa0\xd3\xd7\xf4\xc6\x84\xda\x3b\x5f\x6c\x65\x25\x61\x66\x33\xd6\xd5\xa5\x84\xcc\xd2\x5e\x84\x5d\xd3\xc3\xfe\x78\x0a\x7b\xb3\x27\x1c\x8d\x1b\x67\xaa\xc4\xa8\x2a\xe1\xaf\x90\x08\x03\xd3\x7c\x72\x5e\xbe\x7c\x9e\x14\x12\x82\xd8\xf9\xf3\x64\x90\x72\x89\xce\x9f\xc7\x24\x86\x69\x73\xa6\x3c\xb3\xc7\x1d\xa1\x28\x3f\xea\xac\x9b\x9b\x8c\x33\x83\x67\xf6\xc5\x18\xb4\xb0\x9d\xaa\xb1\x91\xe5\x13\xf1\x83\x41\xd8\x9d\xf0\x46\x92\x19\x15\x16\x80\xef\x95\xa1\xb1\xe3\xa2\x76\x54\x6d\x7b\x82\xa3\x63\x27\xe7\x8f\x73\xfe\x0f\xd7\xec\xbc\x00\x23\xa3\x0b\x60\x9a\x60\x64\x9b\x3a\x5e\xc0\xdc\xc6\xf4\x6e\xf4\xc3\xab\xa3\xe3\xb3\x68\x74\xb5\x0c\x84\x49\xe2\x7b\xa8\xdc\xe5\x28\x55\xc4\xc2\x24\x25\xc1\x8d\xcf\x53\xdf\x66\x9e\xfd\xd5\xde\xbf\x00\x00\x00\xff\xff\x45\x38\x91\x7a\x87\x0e\x00\x00")

func schemaVmLxcJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaVmLxcJson,
		"schema/vm/lxc.json",
	)
}

func schemaVmLxcJson() (*asset, error) {
	bytes, err := schemaVmLxcJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/vm/lxc.json", size: 3719, mode: os.FileMode(420), modTime: time.Unix(1499405184, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaVmNullJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x41\x8b\xdb\x40\x0c\x85\xef\xfe\x15\x62\xda\x43\x0b\x4e\xdd\x42\x4f\xbe\xb6\x14\x7a\x2a\x94\x65\x2f\x4b\x30\x93\xb1\xe2\x4c\xf0\x8c\x26\x1a\x4d\xc0\x84\xfc\xf7\xc5\x71\xe2\x38\xbb\xde\xbd\x64\x0f\x3e\xf8\x7b\x7a\x4f\xb2\xa5\x43\x06\xa0\x3e\x47\xb3\x41\xa7\x55\x09\x6a\x23\x12\xca\xa2\xd8\x46\xf2\x8b\x81\x7e\x23\x6e\x8a\x9a\xf5\x5a\x16\xdf\x7f\x16\x03\xfb\xa4\xf2\xde\x57\x63\x34\x6c\x83\x58\xf2\xbd\xf7\x5f\x40\xff\xf8\xfb\x17\xfc\xc7\x48\x89\x0d\xc2\x03\xba\xd0\x6a\xc1\x12\xf6\xae\xf0\xa9\x6d\xe1\xcb\x1f\x62\x10\x8c\x62\x7d\x03\xe4\xdb\xee\xeb\x90\x24\x5d\xc0\x3e\x82\x56\x5b\x34\x32\x30\xc6\x5d\xb2\x8c\xb5\x2a\xe1\x29\x03\xb8\x54\x65\x00\xcb\x93\x1e\x98\x02\xb2\x58\x8c\xaa\x84\xc3\x50\x51\x19\x72\x0e\xbd\x8c\x64\x92\x1d\x85\xad\x6f\xd4\x09\x1f\xf3\x6c\xaa\x8d\xb5\xe8\x93\x1b\xfb\x9d\xc8\x79\x72\x75\x26\xcb\x1b\xbb\xb3\xbe\xda\x9b\x90\xe6\xda\x59\x2f\xd8\x20\xab\xfc\x22\xd4\xb8\xd6\xa9\xed\x47\xfb\xf1\x2a\xc4\xa1\x23\xee\xaa\x66\x75\x57\xd2\xfd\xa3\x7c\xc8\x18\x9e\x6a\xac\x1a\xa6\x14\xe2\x5c\x8e\x66\xd6\xdd\x35\x25\x79\xbb\x4b\xf8\x57\xd0\xf5\xd5\xc2\x09\x47\xc9\x9e\xe1\x75\x1d\x87\x97\xdb\x3c\xce\xee\xc5\xb0\x8e\x9b\x2a\x8a\x6e\xf0\x9d\x4b\xc8\xdf\xde\xba\x27\x8f\xa3\x0e\xbd\x43\xb3\xdc\x02\x0a\xd3\x77\xc3\xa8\xe5\xc6\x52\x63\x14\xa6\x6e\x8a\x18\x57\x44\x32\x9e\xd2\xcc\x8f\x1c\x1a\x0f\xdf\x92\xf5\xcf\x31\x7b\x0e\x00\x00\xff\xff\xaa\xf3\x11\xf0\xa2\x03\x00\x00")

func schemaVmNullJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaVmNullJson,
		"schema/vm/null.json",
	)
}

func schemaVmNullJson() (*asset, error) {
	bytes, err := schemaVmNullJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/vm/null.json", size: 930, mode: os.FileMode(420), modTime: time.Unix(1499405161, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaVmQemuJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\xcd\x8e\x9b\x3c\x14\xdd\xf3\x14\x96\xbf\xd9\x7d\x49\x33\x1d\x65\x53\x76\x55\xab\x4a\x5d\x55\xaa\xaa\x6e\x46\x14\x39\xe6\x42\x3c\x83\x7f\x72\x6d\x13\x45\x11\xef\x5e\x41\x02\x31\xc1\xa3\xa6\x33\xdd\xc1\xb9\x7f\xe7\x72\x0e\xf7\x98\x10\x42\xef\x2c\xdf\x82\x64\x34\x25\x74\xeb\x9c\x49\x57\xab\x27\xab\xd5\xf2\x84\xbe\xd3\x58\xad\x0a\x64\xa5\x5b\xde\xaf\x57\x27\xec\x3f\xba\xe8\xea\x0a\xb0\x1c\x85\x71\x42\xab\xae\xf6\x9b\x01\xf5\xf3\xf3\x27\xf2\x1d\xac\xf6\xc8\x81\xfc\x00\x69\x6a\xe6\x20\x25\x8d\x5c\xed\x40\xfa\x53\x99\x3b\x18\xe8\xf2\xf5\xe6\x09\xb8\x3b\x61\x08\x3b\x2f\x10\x0a\x9a\x92\xc7\x84\x90\x21\x2b\x21\x24\xeb\xe3\x06\xb5\x01\x74\x02\x2c\x4d\xc9\xf1\x94\x91\x73\x2d\x25\x28\x37\x22\x41\x6f\xeb\x50\xa8\x8a\xf6\x70\xbb\x48\xc2\xd8\x98\x0b\xca\xcb\x71\x5e\x8f\x0c\x34\xcf\x48\x36\x29\x97\x42\xe5\x0d\x37\x3e\x36\x4e\x28\x07\x15\x20\x5d\x0c\x81\x02\x4a\xe6\xeb\x8e\xda\xfb\x59\x13\x09\x52\xe3\x21\xaf\x36\x6f\xea\xf4\x76\x2a\xff\x84\x86\xd2\x05\xe4\x15\x6a\x6f\x6c\xac\x0f\x43\x64\x87\x4b\x17\xaf\xc4\xce\xc3\x57\x07\xb2\xcb\x76\xe8\x61\x0c\x89\x33\x78\x91\xe3\x78\xad\x66\x1b\xd5\xa5\xa3\x8a\x25\xe3\x70\x0b\x81\x61\xca\xf1\x22\x7a\xc4\x8e\xe7\xc8\xcc\x94\x61\xc5\x08\x64\x41\x45\xc4\xa6\xd3\x29\x21\x36\xb7\xeb\x62\x1a\x9d\x19\x74\xb0\x29\xb8\xed\x55\x6e\xaf\x84\x82\x39\x0a\xd2\xb8\xc3\x1c\x6e\x6a\xa6\xe6\xa8\x64\x3c\x1e\x30\xdb\x83\xa5\x13\x30\x0b\xde\xda\x30\xbf\x6b\xc2\x8a\x02\xff\x76\x59\xc3\x9c\x03\xec\x0f\xc9\xaf\xc7\xfb\xe5\x07\xb6\x2c\x3f\x2e\xbf\x64\xc7\x87\xf6\xf2\x96\x66\xff\xdf\xd1\x17\x07\x0b\xd3\xac\x5f\x33\xb9\xd4\x28\x99\xeb\x5d\x6f\x9a\xf5\xa4\x7f\x72\xfd\xd4\x4e\xac\xd7\x1d\x8b\x5c\x48\x56\x41\xcc\x7a\x57\x86\x7a\xc9\x1c\xb1\x3b\x16\x63\x9d\x44\xb6\xa6\x85\xde\xab\x5a\xb3\x22\xf7\x58\xff\xa1\x7e\xf2\xad\x2e\x3b\x7b\x14\xf1\xd6\x7c\x0b\xfc\xd9\x7a\x99\x47\xac\x7b\x13\xb7\xa1\xc1\x6b\x6a\x47\x7e\x37\x6f\x14\xfd\x59\x28\xb2\xfd\xb5\xdc\x3b\xae\xf7\x0f\xa1\xca\xd9\x5c\xe5\x51\xb3\xe8\x09\x98\x7e\xf5\x39\xeb\xe8\x99\xf2\x16\xf2\xe7\x46\xc6\x8c\xb2\xd1\xba\x86\xe0\xb7\x0b\x8e\x6d\xc9\x6a\x0b\xc9\xc0\xad\x4d\xda\xe4\x77\x00\x00\x00\xff\xff\x0e\xae\x27\x8a\xb2\x07\x00\x00")

func schemaVmQemuJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaVmQemuJson,
		"schema/vm/qemu.json",
	)
}

func schemaVmQemuJson() (*asset, error) {
	bytes, err := schemaVmQemuJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/vm/qemu.json", size: 1970, mode: os.FileMode(420), modTime: time.Unix(1500457433, 0)}
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
	"schema/none.json":    schemaNoneJson,
	"schema/v1.json":      schemaV1Json,
	"schema/vm/lxc.json":  schemaVmLxcJson,
	"schema/vm/null.json": schemaVmNullJson,
	"schema/vm/qemu.json": schemaVmQemuJson,
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
		"none.json": &bintree{schemaNoneJson, map[string]*bintree{}},
		"v1.json":   &bintree{schemaV1Json, map[string]*bintree{}},
		"vm": &bintree{nil, map[string]*bintree{
			"lxc.json":  &bintree{schemaVmLxcJson, map[string]*bintree{}},
			"null.json": &bintree{schemaVmNullJson, map[string]*bintree{}},
			"qemu.json": &bintree{schemaVmQemuJson, map[string]*bintree{}},
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

func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
