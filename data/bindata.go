// Package main Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// .DS_Store
// api.txt
// api_test.txt
// configlocaldb.txt
// generate.txt
// main.txt
// sqlname.txt
// sqlname2.txt
package data

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _Ds_store = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\xd8\x31\x0a\x02\x31\x10\x85\xe1\x37\x31\x45\xc0\x26\xa5\x65\x1a\x0f\xe0\x0d\xc2\xb2\x9e\xc0\x0b\x58\x78\x05\xfb\x1c\x5d\x96\x79\x60\x60\xd5\x4e\x8c\xcb\xfb\x40\xfe\x05\x37\x2a\x16\x31\x23\x00\x9b\xee\xb7\x13\x90\x01\x24\x78\x71\xc4\x4b\x89\x8f\x95\xd0\x5d\x1b\x5f\x43\x44\x44\x44\xc6\x66\x9e\xb4\xff\xf5\x07\x11\x91\xe1\x2c\xfb\x43\x61\x2b\xdb\xbc\xc6\xe7\x03\x1b\xbb\x35\x99\x2d\x6c\x65\x9b\xd7\x78\x5f\x60\x23\x9b\xd8\xcc\x16\xb6\xb2\xcd\xcb\x4d\xcb\x38\x7c\x18\xdf\xd9\x38\xa1\x18\xa7\x10\x2b\x6c\xfd\xce\x77\x23\xf2\xef\x76\x9e\xbc\xfc\xfe\x9f\xdf\xcf\xff\x22\xb2\x61\x16\xe7\xcb\x3c\x3d\x07\x82\xf5\x0d\x00\xae\xdd\xf5\xa7\x43\x40\xf0\x3f\x0b\x0f\xdd\x5a\x1d\x04\x44\x06\xf3\x08\x00\x00\xff\xff\x6a\x00\x88\x6d\x04\x18\x00\x00")

func Ds_storeBytes() ([]byte, error) {
	return bindataRead(
		_Ds_store,
		".DS_Store",
	)
}

func Ds_store() (*asset, error) {
	bytes, err := Ds_storeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".DS_Store", size: 6148, mode: os.FileMode(420), modTime: time.Unix(1582475885, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _apiTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x55\x5f\x6f\xda\x30\x10\x7f\x4e\x3e\xc5\x4d\xea\x03\xa9\xaa\xe4\x9d\x25\x48\x0c\x36\x69\x0f\xab\xba\xb2\x3d\x4d\xd3\x64\x92\x03\xa2\x39\x8e\xe7\x5c\x06\x14\xf1\xdd\x27\xdb\x49\x08\xd0\xd0\x52\x5a\xed\x2d\xbe\xdc\xdd\xef\xcf\xf9\xe4\x20\x18\x96\x94\xc3\x1c\x05\x2a\x46\x98\xc0\x32\xa5\x05\x7c\x41\x62\x43\x99\xc2\x82\x48\x16\xfd\x20\x98\xa7\xb4\x28\xa7\x7e\x9c\x67\x01\xae\xd6\x0f\x0f\xeb\x20\x43\x62\x4c\xa6\xae\x64\xf1\x6f\x36\x47\x08\x43\xf0\xef\xaa\xef\xc1\xc0\x75\xf5\xf9\x73\x26\x73\x45\xf5\x51\x31\x31\x47\xb8\x4a\x45\x82\xab\x1b\xb8\x22\x36\xe5\x08\xfd\x08\xfc\x6f\xfa\xab\xd0\x69\x41\x30\x52\xc8\x08\xc1\x84\xdc\x59\x29\x62\xb0\x11\x13\x08\xc3\xaa\xcc\x1f\x31\x79\xcb\x32\x8d\xd4\x4b\xa6\x70\x5d\xfc\xe1\xfe\xf8\x83\x07\x3d\x54\x0a\x50\xa9\x5c\x79\xb0\x71\x9d\x5f\x37\xfa\x00\x11\x24\x53\xff\xe3\x0a\xe3\x5d\xfd\x58\xe5\xd2\xb4\x9c\x10\x23\xcc\x50\x18\x96\x4e\x3a\x33\x05\xef\x22\x10\x29\xd7\x1d\x1c\x85\x54\x2a\xe1\x3a\xdb\x53\xdd\x5a\x14\xf7\xfb\x55\xd5\x5b\xd7\x0d\x02\x0d\xd9\xd6\xd5\x50\x78\x73\x55\x6d\x16\x13\x52\x65\x4c\x2e\xad\xa5\x19\xd9\x0e\x78\x92\x8a\x79\x05\x0e\x85\x49\x82\x8d\xbb\xcb\xb0\x75\x9f\x52\xe4\x89\x99\x94\x69\x66\x75\x5b\x39\xbd\x56\x6e\xab\xd5\x75\x17\x88\x57\x0d\xb6\xeb\xff\xbe\x03\x0a\x8b\x92\x53\x27\x63\xeb\xc8\xce\xa2\x82\x32\xb2\xb1\xbe\xb1\xe9\x4e\xa1\x64\x0a\x0f\x27\x76\xc6\xf0\x13\x9c\xa1\x02\xdd\xd7\x1f\xf1\xbc\xc0\x9e\xe7\x3a\x76\x0a\x26\xf6\xb5\x44\xb5\xbe\xcf\x97\x87\x08\x26\x7e\x34\x84\x7b\x24\x95\xe2\xdf\x17\x3b\x57\xd7\xbf\x89\x77\x55\x7a\xd4\x59\xb0\xd9\xd6\xd2\x93\xe9\x23\xc2\x6b\x72\xdd\x77\xb0\xce\x18\x72\x6e\x1d\x68\x05\x9e\x5e\x86\x5d\x46\xed\xd4\x8f\x9f\xcf\xd6\x96\x2f\x8b\xf6\xbd\x30\xec\x8f\xa9\x0f\x39\x3f\xe3\x6a\xcc\x72\x05\xba\xb1\x7f\x8b\x2b\xea\x79\xd5\x6f\x63\x62\xff\xb4\x8b\x75\xdf\xc8\xd4\xb7\xae\x41\xcc\xc4\x90\x73\x18\x0c\xde\x1f\xe2\x36\xc0\x1a\xd9\x39\xf6\x22\x02\x26\x25\x8a\xe4\xd8\xa6\x1b\xb0\xa4\x3c\x43\xda\x10\x6e\x6e\x72\x7b\x3a\xdf\x65\x72\xc1\x52\xdb\xea\xff\xb3\xd4\x16\xfb\xb2\xa5\x7e\x62\xab\x2d\xc4\xe3\x5b\x3d\x46\x8e\x2f\x37\xce\x56\x3f\xcf\xb8\x33\x7d\xb1\xad\x2f\xf5\xa5\x79\x76\x4c\xfc\xe0\xe1\x31\x08\xa7\x6c\x69\x76\xbd\x39\x76\x49\x2d\x5e\x41\x6b\x7b\x81\x5f\x4b\xed\xde\x9e\x84\x21\xa0\x48\xb4\xd6\x7f\x01\x00\x00\xff\xff\xac\x1f\x93\x53\x40\x09\x00\x00")

func apiTxtBytes() ([]byte, error) {
	return bindataRead(
		_apiTxt,
		"api.txt",
	)
}

func apiTxt() (*asset, error) {
	bytes, err := apiTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api.txt", size: 2368, mode: os.FileMode(420), modTime: time.Unix(1582830093, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _api_testTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5a\xdb\x72\xe3\x36\xd2\xbe\x16\x9f\xa2\xc3\x8a\xa7\xc8\x8c\x86\x9c\x99\xff\xdf\x1b\xc5\x52\x95\x63\x6b\x2a\xb3\x1b\x1f\x56\xd2\x24\x5b\xeb\x75\x4d\x41\x64\x4b\xc6\x9a\x04\x68\x00\xb4\xec\x71\xf9\xdd\xb7\x00\xf0\x24\x8b\x3a\x58\x9e\xad\xad\xe4\x22\x16\x71\xe8\xfe\xba\xf1\x75\xa3\x1b\x49\x18\x1e\xe5\x8a\xc3\x1c\x19\x0a\xa2\x30\x86\x05\x55\xd7\x70\x8a\x8a\x1c\x65\x14\xae\x95\xca\x64\x2f\x0c\xe7\x54\x5d\xe7\xd3\x20\xe2\x69\x88\xf7\x0f\xdf\xbe\x3d\x84\x29\x2a\x42\x32\xea\x64\x24\xba\x21\x73\x84\xc3\x43\x08\x2e\x8a\xdf\x83\x81\xe3\xd0\x34\xe3\x42\x81\xe7\x74\xdc\x98\x28\x32\x25\x12\x43\x79\x9b\xb8\x4e\xc7\x45\x16\xf1\x98\xb2\x79\xf8\x6f\xc9\x99\x1e\x98\xa5\x4a\xff\x49\xf8\x5c\xff\xe1\x52\xff\x5b\xe0\x2c\xc1\xc8\x8c\x4b\x25\x28\x9b\x9b\x51\x85\x52\x51\x66\x96\x29\x9a\xa2\xeb\xf8\x8e\x73\x47\x04\xe8\xf1\x93\x29\xfc\x24\x6f\x93\xe0\xe4\x17\x33\x14\x71\x36\xa3\xf3\x78\x0a\x29\xc9\x2e\xad\x88\x2b\xca\x14\x8a\x19\x89\xf0\xf1\xc9\x89\x38\x93\xaa\xd8\x78\x46\x52\x84\x3e\x18\xf1\xda\x90\x4f\x34\xc1\x0b\x81\x33\x7a\x0f\x83\x81\xeb\x38\x61\x08\x7d\xfb\x0f\x5c\x63\x92\xa1\x90\x7a\x8c\x48\x99\xa7\x28\x81\x14\xba\x12\x1e\x91\x24\x9e\x06\xda\x2c\x98\xd1\x04\x81\xc8\x9e\x13\x86\x8f\x7a\x3f\x00\xb8\xbf\x72\xa9\xdc\x1e\xb8\x66\xe1\xb5\xfe\xe8\x96\x53\x17\x5c\x98\xa9\xbf\xfc\xff\xff\x7d\xac\x47\xbf\x48\x14\x7a\x34\x9e\x32\x92\x62\x63\x35\x91\xb2\x6d\x5c\xdb\xd1\x36\x3e\x1e\xff\x76\xca\x63\x3b\x45\x25\x99\x26\xe8\x3a\x61\xf8\xe4\xcc\x72\x16\x41\xc2\x49\x7c\x6c\x0c\xf0\x7c\x78\x74\x3a\xb3\x54\x05\x17\x82\x32\x95\x30\xcf\x85\xc6\xb4\xeb\x3b\x1d\x6d\x56\x17\x50\x08\xe8\xf5\x81\xcb\xe0\x3c\x43\xe6\xb9\xab\xf6\xeb\xb5\x74\x66\x16\xfe\xd0\x07\x46\x13\x2d\xb9\x93\xf0\x79\x70\x41\x18\x8d\xb4\xe8\x63\xc2\x18\x57\xc0\x33\x64\xcb\x0e\x34\xbe\x73\x8d\x96\x60\x28\x04\x17\x9e\xef\x3b\x9d\x27\xa7\x13\x63\xc4\x63\x34\xaa\xb5\x8e\xe0\x0c\x17\x27\x76\xc8\xd3\x5b\x7c\xa7\xa3\xf5\xf5\xa1\x58\x17\xd8\x49\xef\x4d\x49\x85\x9d\x31\xcd\x51\x81\x01\x53\x00\xcb\x05\x51\x94\x33\x0d\x4d\xf0\x74\x1d\xbe\x27\xc7\x3a\x34\x12\x48\x14\x9e\x4c\xbd\xb8\x62\x64\x17\x62\x4b\x32\x4b\xc4\x2e\xf0\x05\x43\x51\x7c\xf9\xe0\x69\x4c\xa8\x45\x99\x13\x90\x52\x9b\xa8\xcf\x61\x9c\xe9\x83\x98\x79\xee\xf1\x68\x78\x34\x19\xc2\xc9\xd1\xe4\xe8\x97\xa3\xf1\x10\x0e\x24\x9c\xff\x71\x36\x1c\xc1\x81\x74\x4b\xe1\x85\x54\x7f\xe5\x08\x5d\x78\x0b\x52\xfa\x4e\xe7\xab\x3d\xba\x3e\xc4\xd3\x60\x78\x8f\x91\x67\x46\x05\xaa\x5c\xb0\x0a\xbe\x44\x35\xf9\xb6\x84\x7e\x17\x80\xe3\xe1\x04\x26\x9f\x4f\x87\xf0\xcf\xf3\xb3\x21\x7c\x99\x1c\xbb\xaf\xc5\x11\x0b\x9e\x6d\x72\xe2\x2e\xa8\x4e\x46\xe7\x17\x4d\xa7\x55\xbe\x7a\x2d\x38\xc1\x17\xc3\x7b\x2a\x95\x5c\xc2\x77\x9b\xa3\x78\xa8\xce\x98\x88\xb9\x84\x20\x08\x1a\x59\x47\x63\x36\xdb\x60\xca\x79\x62\x15\xd5\x06\xd8\xed\x2b\x8e\xfd\x6d\x78\x3c\x81\xe1\x3f\x3e\x8f\x27\x63\xf0\x0e\xa4\xef\x16\x8a\xda\x6d\x28\xa7\x2a\x1b\xfe\xae\x07\x46\x7c\xe1\x99\x19\x0b\x2b\x08\x02\x3f\x18\x47\x84\x79\x6f\x2c\x9e\x55\x0b\x95\x4e\x13\x2d\x36\x9a\xf1\xfa\x08\xee\x48\x42\xe3\x36\x6b\x9c\x8e\x99\x2a\x9d\xd9\x74\x58\x17\x4a\xab\x3e\xc0\xa7\xd1\xf9\x29\x64\xf3\xaf\x46\xac\x84\x3f\x7e\x1d\x8e\x86\x56\x07\xb3\x49\xf9\xc7\x0f\x6e\xa1\x74\x15\x23\x65\x54\x4d\x4c\x06\xf7\x56\xd8\xd0\x4c\x6b\x4e\x27\x93\xb7\xc9\x67\x36\xe3\x2b\x0c\xd1\x69\xb8\x7f\x20\x41\xdf\x56\xfa\x6f\x2e\x51\x98\x6f\x22\xe5\x82\x8b\x58\xff\x76\xdf\x3a\x9d\x8e\x2b\x65\x92\xf2\x18\xfb\x86\x45\x65\x4a\xb9\xb4\x49\xfd\xaa\x39\x62\x72\xf9\xd2\x88\xc9\xe3\xcb\x6b\x74\x06\x5f\x1a\x29\x73\xf4\x95\xef\x74\xec\xad\x54\xba\x4e\x7b\xde\xe6\xd9\x8c\x4b\x35\x17\xa8\x01\x94\x06\xad\x3a\x45\x3b\xe4\x94\x50\xe6\xa5\xf0\x53\x71\x5f\x06\xa7\xc6\x25\x61\xa8\xbf\x75\x80\xe7\x99\x25\x48\xaf\xbf\xe4\xc3\xad\xf9\x31\xb2\xf9\xb1\xde\x03\x2d\x79\xb0\xe4\x5e\x95\x07\x4b\x73\xea\xcb\x76\xd5\x39\x81\x57\x50\x6a\x67\x10\xc7\x85\xf8\x4d\x10\xca\x5c\x66\x35\xef\x2c\xba\xd8\xd6\x2e\x39\x0c\x45\xce\x8c\x2d\xd2\xe9\xe0\x3d\x55\xbf\x93\x44\x3b\x32\x0d\x46\x39\xf3\x7c\xa7\xf2\xb3\x42\x22\x62\xbe\x60\x55\x2c\xda\x84\xb6\xea\x8d\x9d\x71\x9d\x18\x09\x6d\xb0\x3a\x5c\x06\xc3\x7b\xaa\xbc\x02\x90\xaf\xe9\xa0\x1e\x32\x84\x88\xa7\x19\x11\x38\xd1\xbf\x35\x3f\xbc\x46\x3a\xea\xc2\x52\x6e\xd2\x51\x5c\x90\x88\xf1\x63\xbb\xcf\x13\x28\xf3\x44\x75\x01\xef\x33\x8c\xd4\xea\x86\xa5\x92\x61\xa6\xb3\x50\xb5\xb7\x07\x07\x77\x5d\x38\xb8\x83\x77\x00\x07\x93\x2e\x1c\x4c\xe0\x5f\xcc\xed\xc2\xb2\xc8\xe7\xdf\x15\xa3\xc1\x53\x22\x47\xbf\xbe\x10\x70\x46\xf2\x44\xed\x0f\x6c\x59\xc0\xab\xd1\xd9\x09\xe8\xf7\xab\xa9\x12\xa9\xae\x4e\xf6\x87\xd9\xd8\xbd\x2f\x46\xcd\xc1\xa2\x0c\xc5\x24\x96\x10\x11\x06\x53\x04\xc2\x1e\x80\x0b\x5d\x42\x91\x99\x42\x01\xf1\x14\xac\x31\x5d\x90\x1c\x04\x92\x58\xe3\xe3\xad\x25\x33\x10\x16\x43\xc2\xf9\x0d\xe8\xe4\x61\x15\x9e\x92\xcc\x30\x9f\xdc\xa0\xd7\x5e\x66\xfb\x55\x9e\x31\x15\xdb\x17\x96\x12\x21\xaf\x49\xe2\x5d\x5e\x4d\x1f\x54\xe9\x9d\x3a\xf8\xbb\xf0\xa6\x12\xbd\x39\x2c\x74\xba\xb7\xd4\xb7\x46\xef\x8c\x65\x1d\x14\x2b\x66\x09\x4a\x25\x79\x47\x28\x4e\x67\xc6\x05\xdc\x74\xe1\x4e\x43\x11\x84\xcd\x11\x6a\x74\x7a\x0f\x9d\xc1\x9d\x96\x51\x19\x79\x79\x73\x65\x26\x96\x38\xf0\x37\x7c\xb0\x27\x3f\x32\xcb\xec\xef\xa1\x11\xa4\x7f\xbb\x5d\xad\xa3\x29\xa2\x0b\x77\xbe\x16\x52\x50\x73\x46\x12\x89\x4e\x47\x3b\xe7\xa9\xe2\xab\x0e\xa6\xba\xc8\x33\x46\xbe\x80\xa3\xe5\x91\xff\x55\x72\x56\x35\x00\xc6\x91\xa7\x85\x1b\xed\xfc\xcb\x0e\x6d\xbd\xb4\x2a\xda\x76\x92\xb6\x1c\x40\x4b\xc6\xad\x0b\x21\xbb\xc8\xab\xad\xf2\xab\xb1\x1a\x9b\xbf\x21\xfc\x8b\xb6\x34\x98\x08\x9a\x8e\x33\x12\xa1\xb7\x2a\xd2\xd7\xd9\x61\xed\xc2\x86\x1e\xdf\x24\x8f\x30\xa4\xca\xb4\xe1\xa0\xae\x05\xcf\xe7\xd7\x80\x24\xba\xb6\x41\x0c\x7c\xa6\x25\xe5\x91\x32\x91\x48\xb2\x2c\x79\x00\x75\x5d\x25\x78\x93\xdc\x75\x97\x02\x8a\x37\xb7\xe9\xce\x3b\x06\xce\x96\x2e\x82\x94\x64\x96\x07\x78\x9b\x93\xe4\x93\x5e\xb8\x81\x04\x5d\xb3\x57\x53\xb8\x11\x58\x0d\x71\xbe\x2d\xbb\x0c\x4b\x72\xc3\x7c\xdb\xbc\x07\xbf\x93\x24\xc7\xf3\x59\x7d\x98\x77\x6d\xb3\x15\x71\xf4\x7d\x75\x3e\x1b\xeb\x35\x79\xa0\x05\x9b\xdb\x54\x87\x14\xd5\x63\xef\x7f\x06\x0a\x87\x90\x07\x67\x79\x6a\x21\xfb\x3f\x03\x7d\xfb\xd6\xe8\xd5\x3c\xf9\xc1\x2b\x70\x5e\x16\x92\x02\xbb\x8c\xfa\x81\xbe\x66\xaf\x7c\xef\xae\x1e\xf9\x5c\xda\xe7\xf9\x5d\xc8\x5b\xc7\x7d\x1b\x99\x65\x50\xa5\xca\x5e\xb9\x33\xcf\x35\xab\x7b\x70\x20\x37\x84\x68\x2b\x86\x2e\xbc\x10\xc2\xf3\x28\x66\x34\xd1\x4c\x71\xc2\xb0\x28\xc4\x33\x8c\xe8\x8c\x46\xe0\x38\x87\x87\x45\xca\xf9\xd1\x4e\xf5\xfa\x10\x4c\x6c\x59\x3d\x18\x14\xef\x1d\x87\x87\xc5\xac\x41\x03\x83\x81\xf9\x28\xdf\x3f\x56\x66\xcd\xdb\x87\x29\x66\x62\xa2\x08\xbc\x03\xc6\x15\xf6\x20\xe6\xb9\x56\x30\x15\x24\xba\x41\x25\x81\xda\x7a\xc8\x2e\x62\x88\x31\x48\xcd\x73\x98\xa2\x5a\x20\x32\xe0\xea\x1a\xc5\x82\x4a\x04\x4d\x54\xc3\xac\x4c\xa0\xc2\x18\x88\x04\x85\x69\x96\x68\xce\x13\x43\xdf\xea\x3d\xa7\x06\x73\x4c\xb2\x31\x65\xf3\x02\x13\xf4\xe1\xf2\xe3\xd5\xba\xd9\xc7\x86\x89\xa6\x44\xd5\x90\x3e\xc0\x60\xd0\x6d\x9b\xf8\xa8\xe5\x3d\xd9\x37\xa4\x3c\x8b\x89\xc2\x0d\x5a\xd7\x4d\xad\x51\xa8\x5d\xf7\x3c\x36\x65\xf1\x34\x65\x46\x97\xe4\xd5\x6a\xda\x63\xec\xd1\x69\x2c\xb7\xc3\xa7\x24\xfb\x64\xaf\xf7\xc1\xa0\x6e\x4c\xf1\x0e\x85\x6c\x95\xed\xad\x1c\x2f\x5c\xae\xf5\xa3\x5f\xd5\x37\x9b\xd6\x3c\x36\xe3\x33\x41\xb6\xaa\xc2\x87\x77\xf0\x41\x07\xee\xc0\x06\xf0\xbb\x77\x26\xa8\xca\xda\x49\xe7\x31\x64\x75\xf2\x59\xd9\x7f\x49\xaf\xfc\x06\xff\x6d\x96\xac\x5e\xe3\x4c\x15\xde\x5b\x7b\x32\xe5\xba\x46\x67\x64\xbb\x06\x13\x16\xad\x3e\x52\x75\xd3\x34\x59\x7d\x0e\xeb\xf7\xb7\xec\x77\x7d\xa7\xaa\x78\xb6\xaa\x5a\xdf\x94\xd4\x99\x66\xa9\xd9\x59\x2b\xcb\x74\xfe\xcf\x7a\x02\xc0\x44\xa2\x91\xf6\xec\x95\xe0\x84\x33\xec\x6d\x93\xe8\x96\xd7\xb4\xee\xd9\xab\x2b\xba\xf9\x28\x50\xf6\x31\xeb\x53\xca\x6e\xa6\x35\x64\xb6\x99\x61\x64\xfc\x50\xbc\x9a\x2c\x0b\x68\xa2\x69\x21\x9e\x65\x0c\xc6\xb0\x10\x9c\xcd\x41\x2a\xa2\x72\x09\x11\x8f\xb1\x07\x73\xae\x74\x4d\xb0\x20\x4c\xd9\x6c\x5d\x1a\x6a\x5b\x8e\xad\xde\x6b\xa8\x76\x97\x1e\xff\x6a\x8e\xad\x23\xe5\x8e\x14\x5b\xb7\xdd\x30\xac\xba\xab\x8b\x63\xd9\x90\x2e\x2f\xdf\x5f\x05\xdb\x00\xbd\x98\x88\x6b\x03\x6e\x5f\x1e\x6e\xb0\xb6\xd1\xce\xb7\x54\x2b\x5b\x4c\xef\x6e\x4a\xb6\x5b\x2c\x26\x19\xad\x39\x94\x33\x5b\xc4\x60\x5c\x54\x84\x41\x2b\x59\x9b\x34\x18\xa1\x12\x14\xef\xf6\x27\xc2\x36\x01\xfb\x50\x61\x3b\xa8\x17\x90\x61\x9b\xb0\x3d\xe8\xb0\x83\xcd\x26\x2f\xfd\x89\xf9\x70\x94\x24\x7b\x5d\x3d\x5b\xf6\x1b\x36\x7c\xdd\x8d\x08\x1f\xfe\xfc\x39\xe1\x39\xf1\xb7\x7a\x77\x0f\x5e\xaf\x91\xf5\x0a\x5a\xaf\x3f\xbb\xf2\x6d\xd1\xd6\x6f\x30\xc5\x88\xe4\xba\x60\xce\x28\x50\x09\x27\xc3\xf1\x71\x17\x2e\x7b\x57\xfa\x43\x26\x34\x42\xdd\x0f\x92\x24\x01\x22\x04\x79\x00\x4c\x30\x45\x66\x5f\x24\x4d\xfb\x66\xfa\xac\xf5\xa5\xe0\x46\x76\xf4\x74\xc9\x65\xca\xba\x2e\x7c\x7d\xfe\x92\x61\x6c\x5c\x13\x7e\x97\xf4\xaa\xec\x1f\xcd\xcf\xcd\xb1\xd6\x72\x12\xaf\x8a\x36\xe3\xc1\x27\xc7\xb4\x47\x55\xd0\x7d\xd9\x58\xd5\x6f\x8d\xb9\xcd\xdb\xdb\x12\xf0\xe6\x36\x22\xd8\x86\xe7\x05\x2c\xdd\x2c\x6a\x0f\x92\x6e\x35\x76\x63\xe6\xdd\x6c\xf8\xff\x20\xf1\x1e\x1e\x02\xb2\x58\x37\x47\x61\x18\x63\x82\xba\xc5\x4c\x12\xdb\xa4\x52\x56\x86\x47\xf1\x20\xaa\x38\x90\x28\xe2\x69\xca\xb5\x15\x30\xe3\x02\xe9\x9c\xc1\x0d\x3e\xc8\xd6\xc6\x7a\x64\x77\xd7\xfd\x75\x4d\xb9\x13\xa3\x6b\x6f\xca\x6d\xde\xde\x6c\x30\xb6\xdd\xf5\xdb\x80\xbc\x80\x6b\x9b\x45\xed\xc1\xb5\xad\x56\x6a\xae\x55\xff\x29\xf6\xbb\x57\x35\xfd\x97\xf3\xcc\x36\x0d\x23\xbe\xb0\x5d\xc3\x19\x1f\xf1\x85\xbb\x6c\x69\x2d\x5c\xde\x26\x5a\xac\x59\x24\x9f\x3d\xf0\xbe\xa4\xdc\xb1\xff\x3b\x0e\xe3\x05\x06\xd7\xa4\xb9\x5a\xe1\x77\x29\xc9\xec\x03\xd3\xd3\x0a\x83\xf7\xad\x53\x36\xee\x6e\xf2\x77\x8b\x9a\x17\xd3\xf3\xfb\x5d\xd7\xdb\x4c\x78\x45\xfd\xf1\xdf\x2a\x40\x8c\xdc\x04\x59\xf9\x9c\x0a\x03\x78\xff\xbd\xf9\xbd\x4f\x69\xd3\xc2\x60\x4d\xb6\x2a\x39\xff\x27\x00\x00\xff\xff\xba\x7f\xcd\x53\x81\x26\x00\x00")

func api_testTxtBytes() ([]byte, error) {
	return bindataRead(
		_api_testTxt,
		"api_test.txt",
	)
}

func api_testTxt() (*asset, error) {
	bytes, err := api_testTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api_test.txt", size: 9857, mode: os.FileMode(420), modTime: time.Unix(1583197411, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configlocaldbTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\x52\x50\x50\x50\x50\xf2\xc8\x2f\x2e\x51\xb2\x52\x50\xca\xc9\x4f\x4e\xcc\xc9\x00\x71\x74\x20\x12\x01\xf9\x45\x60\x09\x53\x13\x63\x23\x98\x58\x68\x71\x6a\x11\x48\xac\xba\x5a\x41\x2f\xa0\x28\x3f\xcb\x2f\x31\x37\x55\xa1\xb6\x16\xae\x25\xb1\xb8\x18\x8f\x34\x88\x8b\x47\x3a\x38\xd8\xc7\x37\x3f\x05\xac\x22\x25\xb3\x38\x31\x29\x27\x55\x89\xab\x16\x10\x00\x00\xff\xff\xa7\xbe\x44\xac\xa6\x00\x00\x00")

func configlocaldbTxtBytes() ([]byte, error) {
	return bindataRead(
		_configlocaldbTxt,
		"configlocaldb.txt",
	)
}

func configlocaldbTxt() (*asset, error) {
	bytes, err := configlocaldbTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "configlocaldb.txt", size: 166, mode: os.FileMode(420), modTime: time.Unix(1582911461, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _generateTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x3b\x6b\x53\xe3\x46\xb6\x9f\xad\x5f\x71\xd2\xb5\x37\x2b\x2d\x1e\xd9\x30\x2c\x99\xeb\x1d\xa7\xca\x63\x04\xe3\x2c\x18\xd6\x36\x49\x76\x31\x35\x69\x4b\x6d\xd3\x19\x3d\x9c\xee\xd6\x0c\xc4\xe1\xbf\xdf\xea\x87\xe4\x96\x30\x60\x42\xa6\x6e\xee\x9d\x4a\x61\xa9\x4f\x9f\x67\x9f\x97\x8e\x94\x25\x0e\x3f\xe2\x05\x81\xd5\x0a\xfc\x73\x73\x7d\x77\xe7\x38\x34\x59\x66\x4c\x80\xeb\x34\x10\x49\xc3\x2c\xa2\xe9\xa2\xf5\x33\xcf\x52\x24\x17\x18\xcb\x18\x97\x57\x71\xb6\x90\x3f\x09\x16\xd7\x2d\x86\xd3\x48\xde\x30\x32\x8f\x49\x28\xe4\xa5\xa0\x09\x41\x8e\xd3\x40\xf3\x44\xdd\xd3\xac\x45\xb3\x5c\xd0\x58\xde\x64\x8a\x02\x17\x2c\xcc\xd2\x4f\xe6\x92\xa6\x0b\xb5\x2a\xc8\x8d\x68\x09\x92\x2c\x63\x2c\x24\x05\xcf\x71\xc4\xed\x92\x40\x3f\x8b\xf3\x24\x05\x2e\x58\x1e\x0a\x58\x39\x8d\x21\x4e\x08\x68\x44\xa7\x31\x91\x5b\x8a\x9b\x11\x99\x03\xcc\xb2\x2c\x86\x56\x8b\xce\x41\xb0\x9c\x00\xe5\x30\xcf\x18\xa1\x8b\x14\x3e\x92\x5b\xe7\xce\x50\x9d\xe0\x59\x4c\xea\x44\x01\x4a\x52\xff\xca\x09\xbb\xb5\xee\xb5\x14\x1c\x2e\xaf\xf4\x55\x49\x68\x2c\xb0\x20\xa7\x38\xbc\xa6\xa9\x4d\xef\xa8\x4e\xb0\x9f\x33\xb5\x15\x68\x2a\x9c\x86\x62\xcf\x01\xe0\xf2\x4a\x5d\x4a\x72\xf3\x3c\x0d\x61\x44\x70\x34\x3e\x75\x3d\x70\x79\x02\x7f\xb3\x89\x37\x81\x30\x06\xea\x1c\x3c\xc9\x61\x76\x2b\x08\xd7\x8b\x9d\x2e\x68\x1b\xfb\x12\xbd\x17\xc7\x6e\xc6\xfd\xb1\x88\x68\xea\x39\x0d\x3a\x57\x7b\xbe\xea\x42\x4a\x63\x89\xd8\x60\x44\xe4\x2c\x75\x1a\x77\x4e\x43\x42\xba\x20\x0f\xd9\xbf\x48\x13\xcc\xf8\x35\x8e\x5d\x43\xf9\x6b\x9e\x78\x4e\xb1\xf9\xce\x71\x5a\xad\x63\x92\x12\x26\x75\xc0\x9c\xe7\x09\xe1\x20\xae\xb1\x00\x71\x4d\x60\xc9\x68\x82\xd9\x2d\x0c\x0e\xa5\xc1\x69\xaa\x16\xe7\x94\x71\x01\xa1\x3e\x3f\x97\xa6\x11\xb9\x81\xb6\xa7\xf5\x2c\x48\xb9\x91\xc0\xd2\x24\x84\xcd\x71\x48\x56\x77\x4d\x10\x37\xe2\x88\xc6\xc5\xa1\x7a\x5a\x63\x58\x39\x4e\xe3\x13\x66\xc0\x93\x8a\xc5\xf5\xe2\x92\x91\x39\xbd\x29\x4d\xad\xf6\xe5\x73\x7b\xa9\xd5\x82\x79\x22\xfc\x73\x46\x53\x11\xa7\x2e\x3a\x0e\x86\xc1\xa8\x37\x09\x3a\x80\x9a\x60\xbc\xd7\x97\xbe\x74\x36\x97\x12\x79\xbe\x3c\x3e\xd7\xd3\xe6\x7b\x18\x0e\xdd\x2e\x20\x5b\x1e\xa4\x0c\xcc\x13\xe8\x42\x24\xb0\xef\xda\x30\xcf\x71\x1a\x92\x1c\x4f\x7c\xed\x1d\x12\x59\x23\x18\x2b\x83\xab\xc3\xcc\x1f\x92\xcf\x2e\x1a\x66\x30\x97\x86\x48\x71\x42\x90\x14\x45\x1e\x58\xc3\xe8\xda\x55\x64\x68\x4c\xce\xd5\xbd\xeb\xa9\xe3\x94\x4e\x6f\xcc\xf7\x55\x49\x3c\xca\x84\xf4\x10\x13\x6a\xfe\x40\x9e\x83\x6b\xb6\x35\x01\xf9\xc8\xd3\x72\xc9\x7d\xdf\x42\x5b\x0b\x64\xec\xd7\x2d\xe8\x5d\x76\xa2\x4c\x5c\x49\x19\x80\xc4\x9c\x6c\xde\xa4\x44\x94\x1c\x09\x57\x2c\x8d\xac\x3b\x80\x3e\x2c\xcc\x79\x47\x1f\x10\xec\x14\xa7\xb3\x03\xc8\x5f\x64\x32\x5b\x34\x5a\xad\x08\x8b\xd2\x9b\x23\x2c\xb0\xdf\xe3\x9c\x88\x42\x50\x29\xa3\xbd\xc3\xf2\x77\x09\x76\x91\xdf\x92\x84\xad\xdd\xf7\xfd\xbe\xb0\x32\x61\xcc\x18\xd3\x2c\x14\xc2\x29\x4a\x8a\xcb\xd7\x3c\x69\x82\x54\x43\xdb\xd5\xec\x4b\x69\x5c\x06\x6a\x05\xa7\x48\x5b\x9c\x85\x70\x79\x25\xe3\xa7\xa9\x74\xa8\x7a\xb6\x32\x4b\xdd\xad\x1b\x42\x99\xaa\xa0\xe0\x9f\xe6\x5c\x94\xf4\xb4\x23\x48\x2f\x40\x9e\x7f\x48\x62\x9a\x70\x17\xbd\x7d\x8b\x9a\x80\xbe\xfd\x16\x79\xfe\x39\x66\x9c\xb8\x9a\xa4\x2d\x84\x27\xdd\x65\xae\xce\xd7\xd8\x2b\xe3\x7e\x9f\x11\x15\x70\x5a\xab\x87\xd2\x82\xb6\x4e\x99\x1a\x84\xf0\x83\x1b\x12\xe6\x82\xb8\x9a\x9e\x54\xcb\x10\xf7\xfb\x71\xc6\x89\x5b\x26\x09\x85\xaa\x12\x45\xd7\xfc\x33\xda\xc2\x35\x89\x97\x84\x71\x09\xe2\xd7\x59\x1e\x47\x90\x73\xd2\x81\x6b\x21\x96\xbc\xd3\x6a\x2d\xa8\xb8\xce\x67\x7e\x98\x25\xad\x59\x8c\x3f\x12\x92\xcc\x18\xb9\x6d\x2d\xe3\x9c\xe1\x98\xfe\x4a\xb4\xc5\x39\x4d\x17\x79\x8c\x19\xfd\x95\xb8\xbc\xb4\xa3\xe1\xb0\x52\x1a\x15\x3e\xfe\x1e\xf3\xb1\x72\x31\xb7\x58\x99\x64\x27\xd9\x67\xc2\x5c\xee\x35\x01\x71\xe4\xd9\x1a\x97\x7b\x18\x4d\x9e\x40\x73\x6c\xff\xaf\xa3\x97\x7b\xa5\xfd\x0a\x3f\x09\xf1\x92\x0a\xa5\xc5\x46\xa1\xeb\x34\xa8\x88\x89\xa4\x50\xa0\xc7\x92\xe4\xb6\xc8\x6b\x01\xd4\x29\x48\x20\x07\x9e\xcf\x12\x2a\x64\xe4\x61\x21\xfd\x31\x83\x71\xb1\xd2\x13\x4d\xc0\x69\x04\x99\xb8\x26\xec\x33\xe5\xc4\x92\x96\x17\xe2\x27\x64\x5b\xe1\x47\x64\x19\xe3\x90\xc8\xfa\x53\xd3\xe7\x91\x0d\xb6\x81\x3f\x48\xc7\x06\xe4\x79\xea\xa7\x09\x08\xad\x0d\x11\x66\x49\x82\x5d\x2a\x35\x68\x42\x4c\xd2\x85\xb8\x96\xd7\x35\x07\xa0\xf0\x16\x5c\x03\x7d\x05\xbb\x95\x63\x46\x4d\xb4\xe9\xf8\x10\x32\xc7\x65\xb9\x6d\x11\x4a\x90\x10\x71\x9d\x45\xdc\xc8\x70\xaf\x26\x7b\x30\x22\x9f\x08\xe3\x44\xd7\x73\x59\xb6\x19\xe1\x79\x2c\x8a\xba\xae\x04\x98\x67\x0c\xa8\x0c\xc2\x98\xa4\x2e\x4f\x7c\xbd\xd9\x93\x02\xfe\x03\x28\x7c\xdb\x85\xf6\x3f\x80\xbe\x7a\x65\xa4\x52\xf8\x5d\xc0\xcb\x25\x49\x23\x43\xaf\x09\x25\xde\x25\xbd\xb2\x93\x52\x69\xa0\x0d\xc2\xd9\xf5\xc1\xb2\xd3\xc6\x82\x50\x54\xa5\xa2\x22\x6c\xa8\xab\xb5\x1a\x51\xa9\x44\x12\xb5\xa8\x11\x96\x8d\xef\xed\xb1\xd3\xa9\x06\x3e\xa6\x80\xe9\x4e\xdd\x0d\x4e\x97\x71\xff\x98\x08\x92\x7e\x72\xd1\xf1\xd9\x79\xaf\xff\xcf\xde\x71\x80\x8c\xe3\xc3\x0f\x8c\x0a\xb9\x9d\x0a\x10\x19\xcc\x08\x90\x1b\x41\xd2\x88\x44\x0f\x72\x1a\xa8\xde\xd7\x66\x64\xfa\x88\x4a\x57\x41\xd3\x30\xce\x23\x32\xa1\x09\x51\x7d\xa6\xe3\x34\xec\xa5\x2e\xcc\x71\xcc\x89\x3e\xf1\x0f\x4d\x10\xaa\xc9\xec\x74\x81\xe1\x74\x41\xd6\x47\xa8\x4c\x63\xf6\x98\xf6\xa8\xdc\xa4\x70\xfc\xa2\xdd\xd4\x65\xf6\x33\x15\xe1\xb5\xd9\xa9\xba\x10\xbd\x1e\x62\x4e\x00\x1d\xf6\x26\x81\x8c\x95\xc9\xe0\xb4\xfc\x1d\x4f\x7a\xa7\xe7\x95\x9b\xc9\x7f\xe4\xed\x60\x38\x09\x46\xdf\xf7\x4e\x50\x47\xe2\xd7\x64\x97\xed\xb2\x5c\x8e\xc8\x1c\xe7\xb1\x50\x5b\xee\x74\xa1\xbc\x73\x1a\x1c\x76\xba\x80\x8a\x47\x84\x69\x3a\x15\x53\x24\x6b\xc1\x0c\x73\xd2\xe2\xbf\xc4\x53\x34\x95\x0f\x0a\x7a\xdb\x54\x7c\x80\x29\xb2\xd2\x7a\x4c\x67\xad\xe5\x2f\x7a\x8f\x0e\x54\x8b\xb5\x6a\x99\x0c\xde\x54\x3d\x40\x18\x62\x25\x5b\x0f\xad\x93\xcd\xda\x5f\xb4\x79\x4d\x9c\x8d\x69\xba\x30\x4d\xd9\xfd\x04\x65\x55\x0e\x6d\x5f\xb9\xd3\x7b\x80\x52\x1f\x2f\x1f\x22\x64\x25\xf3\x32\x31\x5b\x04\x1f\xa1\xf8\x98\x78\x76\x89\xd8\x2c\xe9\x43\x84\x0f\x59\xb6\x54\x97\xca\x99\x13\x92\x56\x3c\xb8\xe6\xc0\xda\x94\xee\x14\x1d\x8e\xce\xce\x61\xd2\x7b\x77\x12\xc0\xe0\x08\x82\x1f\x07\xe3\xc9\x18\x54\xf7\x54\x72\x94\xad\x19\xf4\x7b\xe3\x7e\xef\x30\x98\x6e\x63\x7e\xdd\x5c\x3c\x57\x96\x9f\xd6\x5c\xf5\x03\xd7\x0e\xa0\x9f\xb6\x39\x6c\xf5\x9c\x75\x44\x49\x1c\xf1\xfb\x31\x2b\x9f\xcb\x4e\xf1\x12\xba\x90\xe0\xe5\xa5\x06\x5e\xe9\x1f\xe9\x6a\xe8\xdd\xd9\xd9\x49\xd0\x1b\xa2\x8e\x7c\x3e\x03\x24\x03\x19\x35\x0b\x80\x59\xad\x02\xfa\xef\x7b\xa3\xc1\x61\x01\x32\x4f\xad\x1a\xf4\x7d\x6f\x64\x41\x2b\xa0\x49\xf0\xe3\xc4\x22\x67\x83\xc6\xa7\xbd\x93\x93\xc1\xd0\x80\x11\x4d\xc5\xee\x81\x86\x94\x8b\x50\x40\x5e\xef\x95\x90\xe0\x38\x18\x15\x62\x5b\x90\x77\x83\x63\x0b\x4d\x42\x0e\xf6\x2d\x3e\xe3\x60\x34\x90\x31\x5f\xe1\x53\x2e\x6e\xa2\x66\x01\x6d\x6a\x47\x27\x67\xbd\x49\x69\x07\x34\x8f\x33\x5c\xc2\x46\x41\xcf\xb6\x9d\x82\x15\x14\x15\xde\x9b\x92\x57\x05\x76\x18\xf4\x07\xa7\x25\x6a\x95\xe6\xf0\xe2\x34\x18\x0d\xfa\x8f\xc1\x0a\xbb\x57\x60\xe7\xa3\xa0\x3f\x18\x0f\xce\x86\x35\x18\xb4\x5a\x87\x67\x17\xd2\xef\xcb\x1d\x4a\x04\x99\x41\xd7\xa2\xcb\x04\xe4\xcb\xc4\x64\x4e\x51\xe6\xd5\x47\xa1\x45\x82\xed\x3c\x0c\xd5\x82\xd4\xa0\xeb\x6c\x7c\xcf\x3f\xbe\x1b\x9f\x0d\x1f\x70\x1d\x09\x7a\x57\xc2\x2a\xa0\x8b\x8b\xb5\x93\x56\x40\x77\xb5\xf8\x73\xb6\xaf\x3e\x45\x62\x96\x91\x5a\x36\x83\xa6\x14\xe9\x2c\x5a\xec\xd1\x29\x44\xc7\xde\xa5\x55\xac\xae\xca\x1d\x3f\xdd\x24\x71\x67\x8a\x1e\x24\x25\xc3\x7f\x8a\xd4\x38\xc2\x6c\x2b\xb3\xec\xc6\xfd\x06\xe1\x27\xb4\x96\xd3\x94\x8d\x27\xf3\x07\x66\x8f\xd5\x7a\xd5\xb1\x6d\x6f\x1d\xa5\xb9\x25\x58\x01\x30\x3d\xab\xea\x57\xdd\x0a\xb2\xe7\x6d\x25\x66\x88\x53\xd9\x2c\x3f\x22\xa9\x16\xc0\x97\x3b\x5d\xf4\x7c\xc1\xbf\xd6\xed\xa5\xff\xe4\xe9\x3e\xa9\xca\xd6\x85\x5a\x57\x8a\x67\x15\xac\xc1\x70\x1c\x8c\x26\x30\x18\x4e\xce\x36\xd5\x29\x17\x39\xcf\xd0\x5c\x3d\x25\x74\xbb\xc5\xa0\x23\xcc\x52\x41\xd3\x9c\x98\x89\xc0\x0b\x0f\xd4\x58\x01\xbe\xef\x9d\x5c\x04\x63\x29\x59\x43\xcf\xbd\x3a\x5d\xd8\x2d\x85\xfc\xf0\x52\xf9\xfe\x52\x7a\xbc\x99\xa5\xfa\x03\x91\x61\x3d\x62\xdb\xf2\xc8\xb4\x5c\x3b\x3b\x15\xb1\x47\xc1\xe4\x62\x34\x1c\x0c\x8f\x7f\x87\x2b\xbd\xcc\x64\xcf\x68\x32\x54\xa7\xf0\x74\x4c\x3c\x2b\x1c\xb6\x76\x0a\x8d\x6a\x75\x73\x3b\xf2\x79\xe9\x0f\x8c\x1e\x75\x65\xb8\x14\xe1\xff\xa4\x61\x46\x44\x30\x4a\x3e\x3d\x2f\xaa\xc6\xc1\x49\xd0\x9f\x58\x4a\xe9\xa4\x28\x23\xea\x68\x74\x76\xba\x29\xd2\x7e\x78\x1f\x8c\x02\x1d\x6f\x1b\xdc\xfa\xd9\x66\x36\xb6\x29\xed\xa5\x78\x74\xe1\x2f\x6a\x44\xb8\xd1\xb1\x0b\x43\x4d\xd5\xa4\xe0\xb9\xc7\x21\xa1\xd2\xc6\xea\x50\x67\x8c\xe0\x8f\x6b\xe3\xff\x5e\x93\xf7\xe2\xf8\x4b\x5b\xfd\x6c\x74\x18\x8c\xe0\xdd\xbf\xe1\x25\x79\xae\x6e\xeb\xcd\x56\x40\x70\x18\x8c\xfb\x5b\x05\xe3\xc5\x32\x7a\x6e\x1e\xbf\x38\x97\xad\xd6\x26\x15\xc7\xc1\xa4\xe2\x54\x7b\x5f\x38\xa1\x3f\xed\x69\x65\x7a\x7c\x56\x04\xaf\x23\xc4\xe8\xf2\xbf\x14\x1f\x95\x4c\x5e\x3b\xe8\x3f\x61\x62\xd7\xbe\xf4\x25\x12\xfb\x9f\x3e\x75\x1f\x92\x98\x3c\x33\x8c\x0e\x83\x93\x60\x12\xfc\x79\x93\xf4\xa6\xbc\xe2\x6d\xe5\x07\xda\x18\x4f\xfa\xc1\x8b\x54\x41\xee\x1f\x56\x38\xb6\xd2\xe6\xb9\x05\xe2\xf1\xd3\xbd\x6f\xc6\x56\x6b\x12\x8c\x27\x30\x3e\x0f\xfa\x83\xa3\x41\xdf\xbc\x96\x2e\xde\xaf\x1e\x49\xb1\xa4\x6c\xae\x1a\x9c\xab\x3f\xda\x42\x5e\xf9\x30\xa8\x10\x04\xe1\x42\x6e\xe6\xd6\x4b\xec\x82\xc8\x21\x16\xb8\x42\xd1\x69\xf4\xb3\x64\x89\x99\x86\x14\x84\xee\x1c\x47\xea\x14\x61\x81\xef\x8d\x5f\x4a\xf2\x2b\xa7\x3e\x80\x59\xcd\xb2\x2c\x9e\x10\x2e\x24\xb1\x26\x20\x33\x7f\x34\x1c\xd0\x5d\xd3\xa9\x4f\x66\xb6\xc1\xa8\x8e\x6c\x56\x5a\x8c\xc7\x70\xa0\xd5\x1a\x12\x12\x71\x08\x73\x2e\xb2\x44\x06\xbe\x84\x80\xc8\xe0\x73\xc6\x3e\x82\x74\x3a\x1c\xc7\x10\x62\x4e\xb8\x53\x1f\xfc\x6c\xc1\xc0\xa9\x4f\x84\xb6\xc3\xa9\x8e\x8a\x56\x6a\x84\xf3\x04\x4a\x75\x86\xb4\x52\xb3\x9d\xa7\x51\xac\xe1\xd2\x56\x28\xd5\xa9\xd3\x56\x28\xd5\x71\xd4\x8a\x13\x46\xf1\x53\x07\x59\x1d\x54\x6d\x87\x53\x1d\x61\x6d\x87\x53\x1d\x6e\xad\xcc\xd0\xe8\x8f\xf4\x98\xea\x88\x6c\x65\xc6\x60\xdb\x88\x55\xce\xce\xb6\x44\xaa\x0e\xd5\xb6\xd1\xc5\xa9\x4f\xdb\x9e\x87\x54\x44\xc1\x17\xb0\x5a\x75\x98\xb7\x25\x87\xfb\x63\xbe\xda\x94\x6f\x25\x1b\x0e\x8b\x88\x0e\x45\x8b\x86\xfc\xd7\x6a\xe5\xbc\xf8\x10\xc5\x96\x73\x41\x04\x60\x96\xe5\x69\x04\xcb\x5f\x64\xbe\x23\x40\x39\x97\x2d\x67\x6d\x5a\xb8\x12\x34\x79\x8c\x8b\x53\x1f\x20\x2a\x04\x2e\x70\xb2\xdc\x12\x4b\x1b\x65\x4b\xac\xea\xc0\x71\x15\xe5\x0c\x0b\x9a\xa5\x4f\x1c\x71\x75\x16\xb9\xfa\x99\x57\x30\xe4\x6d\xc5\xf4\xd2\x64\xd6\x22\x70\x9a\x86\x04\x3e\x92\x5b\x0e\x21\x4e\x61\x46\x00\xa7\xb7\x90\xb1\x88\x30\xa7\x3e\xcd\x54\xc4\x67\x0f\x52\x77\xea\x23\xce\x55\x9e\xd3\xe8\x51\xf1\x1f\x18\x24\xe8\x0d\xa7\x78\x79\xff\x1d\xc2\x8b\x27\xa5\x4f\x0e\x38\x3b\x50\x8e\x6f\x4c\xad\xac\x8c\x4b\x7d\xbb\xb6\xee\x00\x6a\x6e\x39\xd4\x2c\xec\xe0\x4a\xa2\x34\xaa\xbf\x1d\x67\x38\x8d\xfc\x31\x21\x91\xab\xc6\xd0\xc3\xec\xb3\xeb\xf9\x17\x29\xbd\x19\xe2\x34\x73\x3d\xaf\xde\x65\xc9\x96\xa4\x0b\x68\x65\xba\x6d\x2d\x21\x8d\x9e\xd9\x74\x6f\x54\xd0\xee\x2b\x8c\xb4\xcd\x7b\x0c\x6a\x1d\xf8\x1a\xfc\x78\x23\x7e\xb7\xa9\xcb\xb4\xfb\x05\xcb\x3c\x6b\xa6\xf6\x5d\xad\x3f\xb2\x5e\xd4\xb9\x45\xbf\x7b\x94\xb1\x04\x8b\x77\x59\x16\xbb\xca\xac\x83\x54\xa4\xee\x9e\x07\x5f\x75\xa1\xad\x5e\xd6\xe9\x0f\x56\x2a\xe5\xfd\x25\x7c\x91\xf6\x29\xc9\x6b\xac\x3f\xf6\xd9\x3d\x30\xce\x84\x4a\x76\x95\xce\xe0\xf7\x70\xa3\xf3\xf5\x4e\xd5\x34\xff\xf6\x5b\xd1\xf8\x8f\xc8\x1c\x56\xd0\x6a\xe9\x0f\xff\x40\x17\xd4\xf5\x37\x14\x35\xc3\x0c\x52\xe1\xaa\xd7\x46\x46\x0a\xaf\x09\xbb\x6d\x6f\xe3\xc7\x33\x0f\x62\xae\xed\xfa\x7a\xef\x9b\x83\x6f\xbc\x35\x0d\x4b\xdf\x75\x21\xfc\x7f\xa3\xef\xeb\x5d\x77\xb3\xae\xeb\x9a\xf7\x7f\x58\xd7\x42\xcb\x83\xd7\xee\x7d\x25\xab\x6d\xda\x0b\xe2\xa5\xf2\x58\x6a\xc4\x2c\x98\xd4\xfa\x87\x3f\x2e\x1b\x1c\x49\xc2\x5a\xc1\x61\xc6\x92\x23\xcd\x47\xaa\xf9\xd7\xf9\x5f\x9b\xf0\x6a\xb7\x09\x07\xfb\x5e\x55\x8e\x97\xf9\xef\x63\x72\x18\x35\xb5\x3c\x47\x9a\x97\x72\xac\x52\x98\xd7\x7b\x6b\x61\xec\x36\xe5\x05\x92\xe8\x17\x9c\x87\x58\x10\xb7\xdd\x6e\xb7\x9b\x8a\xae\xff\x1d\x4e\x73\xcc\x6e\x9b\xb0\x6b\x16\x4c\xe9\x99\xf4\x5d\xcf\x7f\x9f\xe5\x4c\xda\xe8\x1e\xe0\x94\xa6\xb9\x20\x1b\x41\x63\x12\x66\x69\xb4\x11\x24\x6b\x19\xaf\x82\x2f\x26\x7d\xcf\x9f\xb0\x3c\x0d\xa5\x5c\x6a\xe9\x94\x86\xcc\x6c\xf3\x90\x6d\x83\x4a\x0f\xf5\x62\x43\xd8\x82\x6d\x21\x40\xbd\x1d\x7b\x09\xff\x29\xda\xdd\xeb\xbc\xde\xef\xec\xff\x7d\x8a\xd6\x0c\xac\x96\xf7\x0b\x2a\x27\x8f\x14\xfe\x06\x7b\xfb\x6b\xd5\xec\x9e\xf1\x05\x9c\xa5\x33\x7f\xc7\xb3\xd4\xf5\x6c\xc2\xb3\x2f\x42\xd9\xee\x2c\x5f\x76\x12\x45\xe5\x96\xad\xab\x5b\x54\xed\xb2\x35\xb1\x6a\xfa\xe6\xaf\x1a\xc3\x2c\xe5\x02\xc2\x6b\xcc\x38\x11\xb2\x23\xc3\xb3\x30\x22\xf3\xc5\x35\xfd\xf9\x63\x9c\xa4\xd9\xf2\x17\xc6\x45\xfe\xe9\xf3\xcd\xed\xaf\x08\x76\x9c\x46\x03\xf5\xde\xf5\x0f\x83\xa3\xe3\xf7\x83\xef\xfe\x79\x72\x3a\x3c\x3b\xff\xd7\x68\x3c\xb9\xf8\xfe\x87\x1f\xff\xfd\x9f\xf6\xee\xde\xeb\xfd\xbf\x1f\x7c\xf3\xe6\xbf\x91\xe3\x34\x66\xb2\x87\x4b\xf0\x47\xe2\x16\xdf\x2c\x6b\x09\x3c\xeb\x93\x45\xdd\xe2\xcd\x54\x5a\x9f\x5d\xd2\x2b\xe8\x16\xa2\x5c\xae\x2b\xb4\xec\xc8\xcc\xaa\xe7\x5d\x55\x7a\x55\xad\xd9\xcc\xab\xe8\xab\x2d\xbd\xd1\x58\xab\xe9\x74\x8a\x52\x9c\x10\xf9\xdb\x01\xf9\xf7\x81\xce\x67\x3a\x45\x4d\x05\xc7\x0b\xb3\xd9\x9e\x48\x3e\x5d\x5e\x55\x6b\xad\x08\x84\x54\xdc\x3e\xc8\x6e\xaf\x5d\xb2\xbb\xab\x9f\x9b\x39\x51\x57\xfa\x4a\xf9\x79\xec\xca\x69\xe4\xd2\x72\x29\xf9\xec\x5e\xee\x1e\x28\xcb\x7a\x4e\xe3\x43\xf9\x29\xb6\x12\x65\x44\x70\xe4\xe6\x97\x9d\xab\x8d\x9f\x62\xc7\xd9\xc2\x3f\xc7\x29\x0d\xe3\xd4\x45\x7d\x9c\xa6\x99\x28\x3f\x38\x07\xf5\x08\xa4\xa8\xf9\x01\x63\x19\x73\x4d\xff\x9b\x5f\xbe\x91\xc7\xe3\xaa\xdf\xdf\xa0\x7d\xb3\xdf\xf6\xe0\x6b\x68\xdf\x7c\x73\x24\x81\x07\x06\x78\x70\xa5\x16\x8f\x3c\xf8\x0d\xdc\xf6\xcd\x3e\xbc\x7d\x0b\xfb\x9e\xd3\x50\x4a\x74\xd5\xff\x26\x31\x5e\x32\x9a\x8a\xb9\x8b\xfe\xeb\xe6\x95\xf5\x1f\x6a\x42\x7e\xd9\xee\xec\x5f\xc9\xdf\xfd\xce\x81\xfa\x3d\xe8\xbc\x51\xbf\x6f\x3a\xbb\x6d\x75\xb1\xdb\x56\x4a\x95\x5f\xa7\xfe\x4f\x00\x00\x00\xff\xff\xa5\x35\x9d\xa9\x6e\x34\x00\x00")

func generateTxtBytes() ([]byte, error) {
	return bindataRead(
		_generateTxt,
		"generate.txt",
	)
}

func generateTxt() (*asset, error) {
	bytes, err := generateTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "generate.txt", size: 13422, mode: os.FileMode(420), modTime: time.Unix(1583264104, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mainTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x90\xc1\x8a\xdc\x30\x10\x44\xcf\xea\xaf\xa8\xd5\x49\x1a\x8c\x9d\x73\xc0\x97\x40\x36\x39\x24\x61\xc9\xe4\x07\x84\xa7\x65\x44\x64\xcb\x48\x32\x31\xcc\xcc\xbf\x07\xc9\x36\x84\xdc\xf6\x62\x8b\xa2\xaa\xba\x5f\x2f\x66\xf8\x6d\x46\xc6\x64\xdc\x4c\xe4\xa6\x25\xc4\x0c\x45\x42\x5a\x6f\x46\x49\x42\xfa\x50\x7f\x21\x95\x6f\xca\xd1\xcd\x63\x92\xa4\x89\xec\x3a\x0f\x35\xa6\x34\xee\x24\xf2\x96\xdf\x72\xc4\xc7\x1e\x25\xd9\x5e\xab\x53\xc9\xbc\x65\xd9\x40\x9a\xc5\xb5\xc7\x73\x0c\xc8\x3c\x2d\xde\x64\x86\x49\x28\x32\xac\xf3\x2c\x35\x89\xc5\x2d\xfc\x6f\xcb\xa7\x10\xbc\x92\x45\x95\x0d\xac\xf1\x89\x1b\xc8\x35\x31\x8a\x74\xc3\xc8\x33\x47\x93\x5d\x98\x4b\xb8\x26\xde\x4c\x4c\xac\x34\x09\x67\x71\x39\xeb\xee\xe8\xba\xc4\x19\xeb\x02\x1b\x22\xa6\x10\x19\x43\xb8\x31\x89\xb2\xf6\xab\xf3\x5c\x26\x1e\x70\xed\xaf\xf0\x2d\xfc\xe1\xa8\x2e\x3b\x92\x26\x51\xca\xd4\xe9\xec\x7b\x48\xa9\xf1\x78\x40\xbd\x9c\x91\xaf\x26\x5d\x57\x6b\xdd\x76\xba\x1a\xc8\xca\xab\xeb\x6d\x84\xf0\xa1\xac\x36\xbb\x41\xc9\x1f\x61\x67\x7e\x3d\x98\xc5\x93\x48\x88\x34\x35\xe0\x58\xc9\x7f\xb2\xb9\x5d\xbf\xab\x63\x6e\x11\x5f\x7a\xcc\xce\xff\xdf\xc4\x31\xee\x71\x21\x8a\xa9\xc7\x97\xfd\x1c\xac\x2e\xa5\xed\xd8\xe4\x3d\x35\x21\xb5\x9f\x37\x97\xd5\x07\x4d\xe2\x49\x00\xd0\x75\xf5\x58\x29\xaf\xd6\xb6\x6d\x4b\x4f\xfa\x1b\x00\x00\xff\xff\x5d\xc5\xaf\x2c\x31\x02\x00\x00")

func mainTxtBytes() ([]byte, error) {
	return bindataRead(
		_mainTxt,
		"main.txt",
	)
}

func mainTxt() (*asset, error) {
	bytes, err := mainTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "main.txt", size: 561, mode: os.FileMode(420), modTime: time.Unix(1581363047, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sqlnameTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x8e\x41\x6e\x83\x40\x0c\x45\xf7\x73\x0a\x5f\x00\xbc\x47\x4d\x97\x5d\x46\x91\x7a\x80\xca\x90\x9f\xc9\xb4\xd0\x19\x6c\x23\x91\x44\xdc\xbd\x22\x45\x5d\xb0\xac\x57\x96\xad\xf7\xdf\x67\x8e\xb9\x89\xf8\x86\x8a\x83\x68\x80\x8b\x94\x44\x95\x8d\xfd\xe1\xf1\xa0\xfa\x7d\xec\xdf\x52\x0f\x5a\x16\xa2\xca\x67\x3f\x48\x49\xb5\xcf\x1e\xfe\x03\x7e\x38\xcc\x37\xfa\x98\x1d\xa4\x18\xa7\xa4\xb0\x26\x30\xd3\x73\xae\xee\xc5\x1a\xe6\x98\xfc\x3a\xb5\x75\x97\x07\xc6\x7c\xbb\xdf\x6f\xbc\x19\x42\x91\xee\x4b\x22\x68\x95\x9c\xb6\x7d\x59\x42\x60\x6e\x71\xc9\x0a\xba\x24\x35\xa7\x98\x69\xb5\xad\xc9\x9d\x42\x1c\x93\x41\xa9\x3a\x51\x75\xfe\x45\x35\x7f\x1e\x65\x78\x16\x7c\x29\x62\xd6\xec\xcf\xaf\x7f\xe8\xb9\xdd\xff\xc2\x4f\x00\x00\x00\xff\xff\xe5\x96\x7b\x3d\x38\x01\x00\x00")

func sqlnameTxtBytes() ([]byte, error) {
	return bindataRead(
		_sqlnameTxt,
		"sqlname.txt",
	)
}

func sqlnameTxt() (*asset, error) {
	bytes, err := sqlnameTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sqlname.txt", size: 312, mode: os.FileMode(420), modTime: time.Unix(1582911461, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sqlname2Txt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x8f\xb1\x6e\xc3\x30\x0c\x44\x77\x7d\x05\x7f\xc0\xe6\x6e\xd4\x1d\x3b\x06\x01\xfa\x01\x05\xed\x5c\x14\xb5\x76\x25\x53\x34\xe0\x24\xf0\xbf\x17\x4a\x8d\xa2\x68\x96\x2e\xd1\x24\x1e\x79\xb8\x77\xcc\x3e\x36\x1e\x9f\x50\x31\x10\xa5\x90\x40\x23\x4c\x24\x05\xaa\xca\xd4\x9a\xce\xa0\x2a\x4f\x43\x7b\xbd\x52\xfd\x3a\x0d\x2f\x61\x00\xad\x2b\x51\xd3\x50\x91\xf6\x1a\xdf\x77\x32\xde\xb4\xdf\x16\x5b\xac\x95\x14\x6a\x5b\xcc\x3d\x3e\xe6\xcd\x90\x6d\xcb\xda\x45\x03\x29\xa6\x39\x28\x72\xe3\x98\xe9\xf6\x4e\x66\x29\x37\xcc\x3e\xd8\x69\xee\xea\x3e\x8e\x8c\xe5\x7c\xb9\x9c\x79\x43\xf9\xc7\x65\x09\x76\x2e\x49\xff\x21\x1e\xdf\x5c\xdb\x7f\x5d\x9d\x63\xee\x70\x8c\x0a\x3a\x06\xcd\x46\x3e\x52\xa1\x2a\x04\xbd\x42\x0c\x73\x86\x52\xb5\xa7\xea\x70\x57\xe9\x29\x49\xce\x77\x4d\x9f\x7f\xac\x87\xee\xef\xee\x2b\x00\x00\xff\xff\x1a\x72\xca\x44\xbb\x01\x00\x00")

func sqlname2TxtBytes() ([]byte, error) {
	return bindataRead(
		_sqlname2Txt,
		"sqlname2.txt",
	)
}

func sqlname2Txt() (*asset, error) {
	bytes, err := sqlname2TxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sqlname2.txt", size: 443, mode: os.FileMode(420), modTime: time.Unix(1582911461, 0)}
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
	".DS_Store":         Ds_store,
	"api.txt":           apiTxt,
	"api_test.txt":      api_testTxt,
	"configlocaldb.txt": configlocaldbTxt,
	"generate.txt":      generateTxt,
	"main.txt":          mainTxt,
	"sqlname.txt":       sqlnameTxt,
	"sqlname2.txt":      sqlname2Txt,
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
	".DS_Store":         &bintree{Ds_store, map[string]*bintree{}},
	"api.txt":           &bintree{apiTxt, map[string]*bintree{}},
	"api_test.txt":      &bintree{api_testTxt, map[string]*bintree{}},
	"configlocaldb.txt": &bintree{configlocaldbTxt, map[string]*bintree{}},
	"generate.txt":      &bintree{generateTxt, map[string]*bintree{}},
	"main.txt":          &bintree{mainTxt, map[string]*bintree{}},
	"sqlname.txt":       &bintree{sqlnameTxt, map[string]*bintree{}},
	"sqlname2.txt":      &bintree{sqlname2Txt, map[string]*bintree{}},
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
