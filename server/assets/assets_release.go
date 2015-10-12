// Code generated by go-bindata.
// sources:
// assets/server/tls/snakeoil.crt
// assets/server/tls/snakeoil.key
// DO NOT EDIT!

// +build release

package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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
	name string
	size int64
	mode os.FileMode
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

var _assetsServerTlsSnakeoilCrt = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x55\xc7\x0e\xab\xda\x96\x9c\xf3\x15\x3d\x47\x2d\x32\xb6\x87\xc0\x26\xb3\xc9\x06\xc3\x8c\x9c\xb3\x31\x98\xaf\x6f\x9f\x73\xa5\xd6\xd3\x7d\xcc\x56\x2d\xa9\xa8\x5d\xaa\xd2\xfa\xdf\x3f\x1f\x2f\xca\xaa\xf9\x3f\x82\xe8\xfa\xaa\xa4\x0a\x9c\x2f\xfe\x45\x11\xa8\xaa\x12\x00\x82\xc0\x7d\x1c\x81\x73\xc4\x03\x38\x91\xa6\x4f\xb1\x5a\x7f\x32\xf3\x37\x4b\xbc\xc3\x1d\x3e\x10\x0d\xc8\x75\x32\x47\x3c\x45\xbe\x86\x42\x10\xc0\x53\xbc\x38\x97\xaf\xcc\x00\xe1\xb9\xca\x17\x3a\xb3\x4e\xe5\x7e\x48\x29\x6d\x4f\x42\xf1\x14\x5b\xce\xf9\xb3\xe4\xb9\xc9\x17\x42\x66\xcc\x86\xc7\x6a\x0c\xe6\x27\xf5\xc5\x27\xe4\xd5\xbf\x44\xdc\xe9\x18\x7a\xc3\xec\x48\x4c\x69\x9f\xe4\xcb\xb4\x29\x89\x1f\x4a\x9d\x99\xd0\x87\x87\xd9\x72\x17\x04\xce\x17\x5e\xdc\x11\xfe\xc1\xda\xbf\xd8\xf9\xff\x58\xcb\x43\xe8\x6c\x87\xe0\x44\x00\x09\x1c\x47\x16\x0f\x2d\x78\x5e\xa2\x0f\x7f\xcf\xfc\xc3\x2e\xf0\x50\x77\x48\x69\x4b\xc2\xf8\x93\x0d\xcc\x1c\xf9\xa2\x07\x79\xee\x9f\x5d\x0d\xb5\x74\xc8\xbf\x29\xb9\xed\x11\xf9\x78\x23\xd0\x75\x0e\xb1\x8a\xc0\x8f\x08\x48\xdc\xb6\x18\xc3\x7f\x4a\xfe\x79\xd3\xa8\xff\xf6\x85\xff\xf9\x02\xaa\x4a\xb5\xb9\x9f\x77\x08\x57\x4d\xc2\x6f\xe0\x39\x73\xda\x5f\x02\x7b\x68\x49\x66\xab\x52\x24\x36\xba\x2a\xcb\x75\xdf\x78\xf5\xb9\x7d\x89\x3b\x61\x68\x38\x4b\xc7\x3e\xa0\x38\x17\xfa\xcd\xc3\x32\xad\xcd\xed\x94\x97\x8f\x4c\x10\x17\x02\x22\x36\xc1\xd8\xd8\x33\xa5\xa9\x46\x23\x1b\x0c\x29\xd9\x33\x5e\xe7\x5a\x51\x39\xe8\x7c\x5d\x58\x05\xda\x17\x5e\x3f\xa0\x10\x16\x27\xa4\x7f\x32\xe7\x08\xf7\x4b\x71\xe9\x1e\x3d\xa2\x84\xba\xad\x55\xa8\x05\xf4\xfb\x91\x10\xd8\xb1\x3b\x96\x8e\xa2\x03\xf6\x8a\xa8\x23\x4b\xe3\xca\x68\x49\x4a\x54\x32\xf9\x9b\xf6\xd8\x29\x6b\x5d\x52\x24\x8b\x6c\x7e\x56\x6d\xc2\x68\xb7\xf2\x0c\x04\x6b\xfd\x79\xc6\x73\x10\xea\xcf\x41\xe7\x67\x5f\xa1\xed\x1d\xb0\x69\xd7\xdf\x17\xdd\xe1\x6c\xa2\xd8\x7b\x36\xf6\x50\x79\xc5\xfa\x5e\x8a\xb4\x70\x39\xaa\x9c\xf2\x00\x15\xd6\x18\x4b\x44\xf4\x88\xb0\x1b\xd8\xb7\x62\x20\x4d\xa7\x7f\x8f\xbd\x74\x14\x4c\xe0\x11\x46\x39\xcd\xc1\xdb\xa2\x0a\xd2\x64\xca\x4f\x21\x3e\xf3\xb4\x22\x23\xb0\xdd\x8a\x8f\x19\x05\x3a\x4a\x34\x78\xd7\x4e\xbf\xdf\x52\x00\x79\xa7\x43\x5d\x10\x94\xe5\x60\x69\xb2\x97\x67\xa6\x64\xc4\xd9\xbe\xe5\x4d\x63\x6e\xc7\x76\x1f\x0b\xdb\xa0\x8b\x09\xfd\x48\x3b\x60\x9c\x36\x7b\x88\x5e\x79\xed\xa3\x57\x0c\x78\x2d\x30\x9c\x56\x49\xc8\xba\x85\x97\x1a\x6d\xeb\xed\xb9\x35\xbb\x2e\xd3\x95\x93\x03\xe2\x66\xa0\x60\xaa\x02\x61\x71\xb3\x3c\x7a\xed\xd5\x6a\x8b\x9a\x65\x36\xf7\x48\xa6\x0a\xd7\xb5\x31\xe7\xd1\xf5\x2f\xda\x0e\xbf\x37\x0b\x89\xdb\x9b\xfe\xc2\x4e\x7e\xfd\x78\x8c\x5c\x00\x73\x59\x41\xb0\xd9\x06\x1e\x92\xb4\xa4\x61\x98\xb7\x6f\x37\xde\x93\xf0\x32\x63\x80\x17\x2d\x6f\x40\xc7\x55\x49\xc8\x6e\x83\xa3\xa7\x97\x50\xc3\xe3\x97\x03\xed\x08\x2d\x65\x60\x1b\x01\x1f\x94\x23\xcf\xa7\xcb\x36\x72\x54\xb0\x31\x78\xb3\xa3\xcb\x76\x41\x13\x46\x58\x51\x4f\x82\xd4\x2c\x2e\x57\xf5\x91\x85\x0f\xb5\xfe\xd5\x3e\x93\x6f\xb5\x0b\x8b\xf8\x7c\x14\xcd\xd8\xdd\xc8\xd2\xa5\x5c\x9e\x9d\xca\x02\x6f\x8c\xc9\xd2\xcc\x5f\x3a\xf8\x36\x2b\xf4\x2c\x9f\xd6\x89\xbd\x22\xa0\x92\x2c\xf7\x29\xdd\x33\x63\x39\x51\x46\x69\x0a\x26\x67\x05\xdd\x45\xaa\xd1\x24\x98\x4d\x3b\x64\x5d\xfd\x92\xb7\xb2\x62\x0d\xee\x6a\x08\x39\xdc\x85\x2b\x26\x3a\x65\x69\x77\xa7\xa6\xa6\xa0\xdc\x79\x7d\x94\xca\x2d\x37\x9e\xa3\xc4\x55\xbf\x66\x70\xff\x54\x1f\xf9\x57\xf7\xc1\xdf\x58\x83\xa4\xe8\x49\xd4\x73\xef\x86\xc0\x9c\xa6\xc0\x67\x3b\x5d\x2c\x43\xd3\x53\x78\x58\xa7\xb7\x65\x59\x09\x56\x54\x9e\x18\xf4\x9c\x14\x19\x07\x09\x9e\xfc\xe2\x27\x24\x7f\x77\x41\xac\xaa\x5d\x7a\x33\xf6\x4e\xb9\xaf\x1b\x97\x3f\xef\xfc\xd5\x65\xe4\xd7\xcd\xdb\x89\x2b\x45\x08\x33\x8d\xc3\x4a\x38\x1f\x7b\xf2\x52\x99\x6c\x4f\xf4\x00\x69\xf2\x58\x9a\xcd\xe7\xd7\x45\x75\xe6\x29\x63\xc2\xf8\xce\x84\xcf\x71\xc5\x33\x8d\x61\x15\x6a\x04\xfa\x63\xb1\x41\x4c\xf2\xda\xc0\x61\x36\x74\x59\x6b\xae\xdd\xc3\x3d\x64\x74\xdd\x3d\xe3\x29\x34\xc8\xf7\xa8\x24\xa9\x36\x7b\x1b\xce\x76\x4c\x3f\x3e\x9b\x34\x60\x0e\xce\xd1\xda\x60\xcc\xf1\x2b\xe1\xdc\x37\x43\x5d\x66\x6a\x36\x8a\x4f\x48\x95\x6f\x63\x0f\x83\x50\xd4\xb9\x4b\xcb\x49\x72\x7c\xd6\x45\x66\x1e\x0b\xa2\xe2\x6e\xa1\xf2\xb0\x1e\x89\x41\xf5\xa8\x71\x4f\x2c\xfc\xa6\xb9\x10\xdd\x71\x4b\x37\xed\xd3\xaa\xce\xb0\x12\x6f\xcf\xb1\x6a\xc2\xbd\x7f\xcd\xa3\xaa\x70\xaf\xe8\xf4\x03\xe8\xdc\x03\x84\x50\xa3\xfd\xfd\x79\xa4\x91\x86\xa1\x3e\xf7\x7d\xe8\xe7\x95\xd1\x0e\x45\xe9\x28\xb3\xfc\xb2\x75\x97\x05\xae\xb5\x24\x02\x37\x47\xa5\x6b\x7c\xc7\x4e\xeb\x67\xa8\x7a\x53\x4b\x12\x45\xac\x11\x75\x8a\xc4\x34\x9f\x0a\xcc\x85\x45\xf7\x9f\x50\x98\x4c\x38\x27\x48\xb8\x33\xf4\xe3\xab\x7f\xe3\xd1\xda\xb5\x5d\x16\x3b\xfb\x52\xac\xd7\x2c\xb0\x11\x3c\x28\xf2\xf3\xeb\x9e\x2b\x66\x0c\xf3\x7a\x3c\x1d\x88\x0c\x41\x3a\x82\xa4\xa6\xcf\xaf\x0d\xd1\xaf\xa9\x56\xd8\xae\x2b\x81\x1f\x69\x2c\x4a\xfb\xef\x77\x28\xce\x84\x26\xbb\x89\x4d\x8a\x56\xb4\xe2\x5f\xb9\xa0\x5f\x92\xb4\xea\xaa\xad\x90\xa0\xe7\x66\x88\x8c\x66\xb9\xb8\x28\x2c\xcf\x47\x18\x7b\xd4\xc8\x4a\x7b\x68\xcb\xe1\x13\xf0\xdf\x82\x2c\xf3\xd3\x6a\xc3\x87\x76\x64\x18\xaa\x01\x97\xe2\x37\x8f\x65\x8c\xfe\xb2\xd7\x8d\x15\xee\x7e\x66\x8e\x23\x40\xc4\x3c\xd4\xe8\xae\x8f\xae\xb0\xdb\xbd\xfc\xd9\x04\x84\xf8\x4a\x37\x9e\xd8\x3c\x55\x1f\x9e\x7b\x4e\x96\xf4\x47\xdb\xad\xa5\xe7\xb7\xca\xbb\x63\x83\x7f\xb6\x1d\x41\x52\xfb\x2b\x31\x1f\x57\x62\xb2\x08\xf7\xc8\x60\xe8\xa8\x24\x4c\xa7\x54\xc7\x15\x71\xed\x3a\xdc\xf9\x9a\xfe\x3b\xf0\xf5\x7d\x10\xc5\x6e\xc2\x32\xb2\x7b\x6f\xe3\xbd\x5f\x73\x2d\xd3\x17\xa1\x70\x1e\xa2\x7e\xcc\xec\x8a\x1b\xc3\x7b\x46\xfe\x1e\x38\xd1\x04\xff\x7d\xf4\xfe\x2f\x00\x00\xff\xff\x82\xb7\xc1\xb2\x11\x07\x00\x00"

func assetsServerTlsSnakeoilCrtBytes() ([]byte, error) {
	return bindataRead(
		_assetsServerTlsSnakeoilCrt,
		"assets/server/tls/snakeoil.crt",
	)
}

func assetsServerTlsSnakeoilCrt() (*asset, error) {
	bytes, err := assetsServerTlsSnakeoilCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/server/tls/snakeoil.crt", size: 1809, mode: os.FileMode(420), modTime: time.Unix(1444379500, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _assetsServerTlsSnakeoilKey = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x97\xb7\xb2\x84\x48\x94\x6d\x7d\xbe\xa2\x7d\xe2\x05\x5a\x19\xcf\x40\x6b\xad\xf1\x0a\x5d\x68\x4d\xc1\xd7\xcf\xed\x8e\x31\x27\xdd\xb4\x72\xe5\x8e\x73\xd6\xfe\x7f\xff\x1e\x4e\x94\x55\xeb\x1f\xcf\x67\xff\x71\x3c\x35\x62\x03\xf1\x1f\x5d\x4c\xff\xbb\x01\x4c\x55\xd5\x74\x56\xe5\x58\x56\xe7\xd9\x46\x64\xd1\x2f\x51\x18\x1b\x3b\x2c\x37\xdd\x44\x8d\xaf\xce\x4d\xfa\x91\x63\x4d\x57\xe4\xd7\x48\xde\x70\x9f\x82\xad\x1d\x4c\xa7\xe0\xc4\x9f\x6d\xc0\x38\x0c\x90\x3f\xd9\x2d\x4b\xbb\x79\x5d\x2d\x89\x7a\xd8\xae\x54\x70\x2a\xfa\xe6\x54\xd3\x6b\x4a\x43\x70\xf2\x25\x78\x20\x39\x81\x9c\x1c\x07\x64\x6d\xab\xf6\x15\xf8\xd6\x32\x4a\x7c\xbf\xcf\x7c\x0e\x9d\x07\x58\xb4\xbf\x06\x68\x58\xa6\x33\x3d\x0a\x38\x6e\x6e\x2f\xb7\x25\x10\x4f\xb0\xf8\x46\x91\x39\x53\x76\x35\xf7\x3b\x46\xf6\x74\xf2\xc2\x2d\xef\x5c\x3b\x13\x48\x4c\x47\x6f\x21\xd6\x39\x45\xc1\x7d\x04\x3a\xa8\x95\x39\xf5\x02\x2d\x7b\x9c\x3c\xd8\x4a\x17\x3f\x9d\xe5\xc1\xac\x1b\x10\x77\x96\xd3\x4f\x9e\x79\x61\x59\x28\x22\xc9\x64\x19\x0c\xfc\x73\x81\x71\x18\x35\xc3\x67\xe3\x39\xac\x54\x9d\xf2\x03\x28\xd0\x16\x75\xdf\x7a\xbd\x6d\xf2\x47\xe6\x29\xc2\xc7\x68\x15\x26\x1c\x35\x44\x46\x78\x80\xdf\x21\x86\xa9\x92\x4a\xb1\x01\x64\x70\xcf\x43\x4e\x21\x6b\xec\xf7\x20\x69\xa4\x0d\x37\x2a\x36\x7c\x1f\xb0\x3b\x0f\x76\x0a\x1b\x39\x3f\x12\x95\x14\x84\xc0\x1c\x2b\x01\x49\x77\xf9\x26\xb2\x0d\x7f\xee\x69\xba\xf8\xf7\x29\x71\xfa\x6a\xc9\x8e\xa4\x63\xdc\x19\x78\xeb\x85\x3d\x0d\xb2\xc9\x52\x23\x73\x1f\x10\x8d\x9e\x1d\xd9\xf8\xfc\xe5\x66\xfb\x7d\xae\xc3\x7f\x74\x7c\xce\xbf\x1c\x07\x3b\xc9\x7e\xe1\xf6\x97\x0b\xf5\x45\xfa\x21\x6d\x45\xf2\x3b\xed\xf6\x38\x6a\xfc\x9a\xbc\x24\x24\x11\x62\x04\xd4\x8f\x80\xba\x61\x72\x63\x27\x46\x67\x5f\x6a\x48\x94\x4f\xc6\xe8\x33\xdc\x44\xd7\xdd\x8a\x6f\xfa\xf2\xf2\xbc\x09\x27\x88\xd1\xc9\x77\x97\xd4\xc4\x63\xde\xde\xd2\xbe\x2b\xec\xb4\x23\x07\x85\x99\x6c\x04\x00\x25\x69\x5b\x95\xc3\x3c\x3b\x71\x29\x5e\xe5\xab\x6a\xf8\x69\xcd\x21\x68\x67\xd2\x07\x86\xab\x10\xfd\xee\x4c\x67\x32\xa2\xad\x65\x1d\x43\xca\x8d\x1a\xeb\x8b\xc8\xf3\x51\x43\xf9\x1f\x71\x53\x47\x00\x6c\x4d\xd2\xde\x16\x53\x6a\xbf\x13\x74\x1f\xbf\x73\x01\x17\xdf\xed\x66\x4b\x2f\xa7\x8c\x46\x90\x62\x76\x74\x7d\xff\xa9\x84\xd3\x6c\x31\x29\x84\x77\xb9\x6e\x4f\x7b\x12\xf7\xb9\x29\x54\x5f\x69\x00\x35\x60\xb3\xba\xc8\x7c\x81\x25\xe3\x3d\xb1\xe9\xdc\x19\xf9\x01\x0a\xf3\x2e\xd8\x46\x98\x7d\x70\xc3\x9a\x62\xb7\x5a\x6d\x82\x57\xaa\x36\x02\x71\x71\x2d\x22\xf0\x07\x1f\x7e\x11\xf2\xec\x2d\xb2\x00\xeb\xfe\x1b\x61\x2e\x27\xbd\x93\x36\xe0\xe6\x08\xb9\x09\x53\x67\xc5\xac\xb1\xd8\xd1\x21\x9e\xae\x8c\x36\x58\xdf\x2d\x55\x63\x2c\x96\xd2\xbb\xc5\x51\x73\x8f\x47\x8c\xad\x6c\x94\xb4\x7c\xdf\x05\x6c\xd9\xe3\x93\xdd\xfe\xc3\xac\xe5\x60\x77\xc3\x47\x66\x1a\x21\xb2\x59\x7c\xc0\x8d\xcf\xcc\xa9\xdd\x97\x23\x3b\x18\x2f\x52\x94\x2b\x40\x03\x86\xea\x98\x47\x31\xd2\x16\xcc\x0d\xdb\x35\xc9\x83\x81\x55\xed\x55\x49\x29\x82\xb2\x5e\xa4\x53\x5c\x71\x1f\xf9\x09\x38\x24\x1f\x61\x16\x46\xc3\x45\x75\xb0\xa1\xcb\x74\x0e\xa7\x10\xa3\x74\xac\x37\x51\x6b\x7e\x43\xba\xd6\x56\x4a\xc5\xb7\x72\xde\x8a\x00\xde\xf4\x07\xc1\x75\xd4\x19\xd2\x54\x4b\x75\x0a\x7b\x2f\x3d\x09\x95\xc3\x0d\x57\x7d\x0a\xc3\x39\xd4\x9c\xbe\x59\x1a\x10\xa1\x23\x5f\x7a\x07\x17\x9c\x5c\xa9\x5d\xa9\x4e\xda\x77\x90\x72\x3d\x62\x80\x7f\x2b\x6c\xd0\xb5\x40\x98\xf4\xdc\xe0\x98\xae\x17\x7a\x5a\xe5\x11\xb5\xc9\x3b\x03\xae\x4d\x35\x81\x6f\x5c\x7e\x61\xc5\xd3\x6a\x1c\x0e\x88\xef\x2c\x9f\xff\xfe\x14\x24\x85\x13\x3f\x3e\x25\x09\xdc\xa4\x6b\xae\x22\xa4\x59\xd5\x1e\xd4\xd2\xfa\x5b\xbd\x37\xe3\xee\x40\x43\xc0\x5f\x9f\xf4\xbe\x34\x4a\x6d\xc6\x98\x45\xfd\x1b\xc5\x63\x7e\x3f\x27\x6c\x28\x64\x70\xf4\xf2\xa8\x5c\xf5\x38\x91\x10\xa0\xd2\xea\x38\xa8\xde\xa3\xc8\x76\x18\x29\xbb\xc8\x7a\x74\x41\x7d\x04\xa6\xfc\x34\xb0\x91\x27\x63\xf7\xcc\x43\xfd\x9d\x68\x61\xbe\x11\x0c\x86\x64\xac\x0d\x92\x74\x12\xc4\xe0\x47\x92\x3d\xe2\x05\x00\xd7\x92\xbd\x20\xa6\xd7\xf9\xbb\xa4\xb5\x4f\x52\xdd\x63\x4f\xe2\x44\x90\xcb\x65\x99\xb3\x4c\x57\x0b\xff\x82\xed\xe1\xa2\xb6\x9d\xda\xa2\x6e\x25\xb4\xfe\x7b\x08\x12\xdd\x92\xa6\xa4\x09\xf3\x09\x80\x2a\x30\x70\x61\x6c\x21\x52\xa3\xa2\x23\x96\xf1\xc9\x42\x65\x7a\x99\x3a\xc2\x68\x18\xc6\xa5\xeb\x57\x94\xdc\x51\x32\xa4\x42\x87\x7b\xed\x56\x4e\x18\x5f\xb4\xcf\xba\x29\x06\x78\x2a\x4a\x49\x04\x80\x6e\xd7\x0e\xbb\xc1\x4b\x98\x57\x27\xba\x47\x25\x53\xb4\x8d\xd0\xb2\x42\x28\x22\x10\x9d\x3a\xad\xcf\xef\xca\x59\x9a\xb8\x35\xe3\x87\xc0\x67\x91\xb8\xea\x1a\xb2\xc8\x06\xbf\x84\x37\xa6\x20\x17\xd0\xc2\x1f\x16\xe2\x08\x1a\x5d\x50\x5e\x96\x16\x79\x56\x27\xbd\x96\xf4\x2b\xcc\xd3\x03\x9d\x2a\x21\xcb\xdc\xa7\x19\x53\x3e\x54\x9e\x7b\xb5\x7e\xee\x7f\xc9\x75\x45\x16\x52\x60\x93\xd7\xd2\x08\x88\x3c\xcf\x38\xeb\xac\x7f\xb1\x7a\xe1\xc1\x3b\x00\x8b\x79\xcd\x9b\xf3\xde\x02\x34\x5f\x22\xee\xb4\xca\xeb\x11\xf4\x54\xed\xf3\xc2\x17\x0e\x35\xb2\x76\x1b\x42\x16\x04\xe2\xe5\x79\x1b\x85\x61\x00\x7a\xa4\xb3\xa4\xd6\x03\x87\x95\xe3\x68\xd7\xad\x73\x24\xab\x0e\x67\x6e\x7f\xaf\x63\x4d\x85\xf3\xd5\x9e\x25\x67\xc2\xef\x75\x52\xb2\x8c\x98\x56\xa8\x68\x61\x55\x48\x07\x7d\xaa\x41\x91\x32\xd8\x0e\x84\x8d\xc7\x45\x3b\xbf\x18\x4d\xe8\x50\xbe\x44\x8e\x88\xb9\x9c\x28\xf7\x49\x5c\x68\x17\x4d\xba\x4c\x79\x7f\x3c\x70\xe9\xf3\xa3\x8c\x18\x86\x0b\x75\x31\x86\x2a\x1b\xf9\x64\x2e\xa0\x6e\x0f\x42\x13\xf8\xfe\xd2\xc0\x24\x36\x22\x8c\x13\xae\x1d\x10\x86\x34\xea\x13\x8c\xbb\x60\x9c\x7f\x72\x23\x55\xc1\x2d\xd1\x3f\xe4\x3e\x67\x6a\x74\x2f\xa7\xde\xbe\xcb\x15\x26\x50\xde\xc9\x1e\x3e\x13\x9a\x6b\xdf\x00\xa2\xc2\x89\x8e\x60\xcf\xd6\x4a\xbd\xc8\xaa\x11\x89\x4c\x0a\x27\x7f\x9f\x76\xd4\xfc\xd6\xf9\xf8\x2a\x7a\x0a\x67\xf9\x83\xb3\xbe\x3e\x43\xd5\x44\x42\xe9\x5a\xab\x35\xde\x3b\x36\xf8\x18\xa3\x1d\x03\x09\x35\x88\xf0\x6a\x35\x14\x31\xec\xde\xff\x12\xc6\x82\x62\x6d\xae\xef\xbe\x33\x1f\xdc\xbe\x37\x98\xfb\x34\x43\x65\xe7\xe1\x14\x98\xd2\xdb\xd3\x79\xb3\x54\xa3\x55\xbe\x5f\x9c\xaa\x64\x80\xe2\x09\x0a\x8d\x7f\xb4\xbc\x93\x49\x63\xe5\x36\x4a\x0f\x8b\x08\x3b\x02\xa5\x87\x90\x77\x1d\x18\x79\x3b\xcc\xf8\xca\x5f\x07\x36\xb1\x6c\xc3\x5b\x1a\x61\x6b\xe3\xaa\x8f\xd9\x34\x77\xf2\x39\x7a\xe0\x56\x34\xbb\x3c\x9c\x90\xcf\xd6\x6b\xca\x87\x4b\xb7\x55\x1c\x66\x26\x79\x9a\x43\x01\xc4\xa6\x5a\x81\xd6\x55\xe9\xf1\xcc\xc0\xea\x09\xbf\x94\xfc\x77\x9b\x73\x99\xd5\x76\xe2\x6d\x0a\xda\x6a\x1b\x20\x46\x2b\x3b\x4d\xd5\xef\xcb\xf8\xb0\x96\x6b\xda\xaf\x2c\xfc\x81\x33\x55\xe2\x72\x4d\xfd\x2f\x70\x1d\x81\xf2\xe8\x58\x17\x12\x98\x78\xe6\x29\xd1\x74\x7e\xd7\xda\xdf\x0a\xca\xb1\xa7\x7d\x2e\xfa\x0b\x2c\x89\xa2\x64\x19\xe3\xea\xe2\xdf\x83\x28\x32\xb4\x41\x51\xd6\x5f\xba\x99\xb0\x3b\xd6\xcf\x96\x50\x92\x05\x53\xe6\x47\xff\x96\xcd\x9b\x1c\x8e\x9a\x9f\x74\x5d\xe2\xe1\x9b\x0c\x10\xf7\x87\xa4\x00\xba\x98\x95\xe1\xb4\xd8\x47\xa2\x0c\xe6\xdc\x85\xd3\x19\x4e\xaa\xee\x24\x94\x9d\x70\xf9\x43\x3d\x52\xca\x66\xa5\x5d\x81\x9a\xff\x00\x73\x1c\x1f\xf0\x4c\x98\x58\x9d\xcd\xdc\xd9\x92\xe6\x33\x20\x94\x33\xdb\xfb\x13\x2b\x1a\x7f\x34\xe6\x11\xfd\x35\x3e\x58\xc1\x38\xae\x20\x5e\x7f\x91\x27\x95\xd9\xa8\x8a\x14\x5a\xa0\x54\x54\xad\xf3\x6d\xd9\xa7\xd6\x6b\xc4\xe8\x30\x12\xb0\x2b\x80\x48\x0d\xd4\x1e\x5e\x6b\x8c\x8c\xf0\xf7\xe0\x8a\x8b\x42\xbe\x9b\xbb\x52\x5c\x3e\xdd\x79\x8b\xab\x76\xb2\x99\xe5\xbc\xca\xf9\xdc\x1d\xab\x34\x74\xd9\xce\xac\x82\xd5\xcc\xf5\x4f\x0f\x59\x4e\xfd\x86\x98\x0e\x2c\xa5\xb8\xa4\x8b\x25\xf8\x9e\x9d\x3d\xa1\x51\xe6\xf4\xc0\x1c\xe7\x49\x68\x5e\x75\x17\x67\xdb\xb8\x4d\x35\xa0\x7d\x8f\x6e\xf6\x6b\xd2\xdc\x02\x9e\x0b\xa5\x9f\x29\x37\x4e\x8f\xab\xf1\x55\x38\x03\x15\x4c\x3c\xd0\x5e\xd3\xc8\x7e\x81\x8d\xbc\x2d\xdc\x6b\x1c\xf7\x57\x7c\xdf\x9f\x84\x16\x90\xb4\x4e\x93\x7c\xff\xa4\x0b\xcb\x30\x6e\xeb\xa3\x91\xa8\xf2\x66\xae\x32\x68\x80\xd9\xa4\x7f\xcd\x81\x07\xb0\x3b\x99\x9d\x5c\xd2\xf6\x9f\x94\x7d\x64\x8d\x7a\x1d\x14\x2d\x39\x41\x6e\x2c\x7c\x8e\xcc\x89\xbf\xf7\x05\x63\xf5\xd0\xa2\x4f\xba\xa4\xc7\x2e\x1c\x84\xf2\xcb\x64\x0a\x11\xf0\x64\x22\xbd\x5c\x00\x10\x2f\x2b\xb1\xf3\x9f\xa5\x49\x7c\x37\x24\x63\x01\x9b\x75\x04\x33\xe4\x97\x4b\xd3\x27\x81\x3f\x56\xb0\x40\x93\x8b\x3f\x86\x57\x4c\x94\x5a\x8f\x7b\x9b\x0a\xad\x0c\x5e\xf3\x54\xf7\xba\x64\x1e\x33\x80\xd0\x08\xd3\x28\x9f\x8f\xfc\x8b\xcd\x43\x0a\xa5\x1a\x8c\x6d\x33\x25\x53\x9d\xe1\x6e\x2a\x4e\x65\x5f\x57\xb4\x38\x31\xd6\x91\xb3\x10\x3e\x6c\x29\x25\x5a\xe1\xcb\x3c\x9e\x84\xbc\x8e\x08\x72\x5d\x20\x0c\x1b\x6a\x9c\xaf\xbf\x81\xcf\x9d\xa7\xb2\xd0\xd3\x22\x0a\x6e\x1b\x84\x76\x68\xc9\xb0\x6d\x81\x38\x9f\xa7\x19\x56\xea\x79\x7c\xb0\x59\x54\x29\xd6\xa7\x90\x71\xda\x47\xd6\x5b\xcc\x1c\x03\x39\x81\xf0\x1b\x06\xfb\x11\x58\x2b\x93\xfa\x4d\xaf\x82\xa9\xd4\xdc\x15\xcf\xca\xb2\x53\xcc\xa7\xc7\x21\x12\x5f\xae\xc2\xeb\x28\x7d\x71\x45\x50\x42\x57\x75\xe6\x86\x6a\xb0\xfb\xb2\x09\xde\x53\x12\xc3\x80\xdc\xd1\x15\x46\x67\x85\x76\xc5\x2a\x8c\x8f\xce\x2c\xc0\x0e\x13\x7f\x8c\x16\xf2\x71\x7c\x01\x3d\x71\xe8\x42\xe6\x90\xb5\x21\x79\x57\x03\x1d\xad\x71\x28\x9e\x3a\x7a\x84\x77\x6b\xc1\x86\xd5\x1c\xe0\x4d\x57\xf2\x64\x93\xe2\xb5\xee\x1a\x0a\x25\xa8\x23\x67\x5e\xa3\xd1\x4f\x84\xbe\x70\x74\x8b\x7c\xd3\x88\x1c\xab\x35\x1f\x54\xcf\x91\xf9\xb2\x0c\xd3\xa5\xa3\x73\x44\x0e\x91\xe6\xec\xfe\x64\x04\xa0\x0d\xa2\xe5\x05\x77\xe9\x6c\xb2\x33\x25\x84\x77\x60\x76\x47\x0e\xb0\x30\x91\xbb\x2c\xa3\x17\x59\x94\x17\xc3\x71\x75\x34\x75\x88\xa7\x00\xab\x96\x9f\x1a\xac\x72\x8e\x24\x82\x77\x27\xf4\x88\x18\x40\xa8\xee\x9b\x9b\xa8\x96\x4a\xbd\xc7\x78\x1a\xcf\xb3\x74\x01\x82\xa8\x7d\xff\xb2\x20\xd2\xfe\x46\x79\xd8\x2b\x32\x55\x2c\x03\xe1\x1a\xf9\x6b\x6b\x6c\x2c\x53\xf4\x9f\x29\x7f\xad\x35\x13\x1f\x1e\x06\xbc\x3d\xd5\x9d\x4f\x49\x56\x70\x35\xa5\x87\x73\xa5\xc3\xc4\xa7\x1b\x45\xa5\xb8\x6e\x91\x63\x48\xd8\x07\x66\xbf\xab\x5c\x39\x38\x08\x96\xa1\xf9\x4e\x06\x8a\x5c\x39\x7a\x8f\x7f\x66\xab\x18\x7f\x86\x62\xde\x1f\xf2\xa6\x6a\x35\x49\x2d\xd6\x3e\xf9\x1e\x96\xcd\x2a\x7d\x7d\x25\x0b\xa8\x2c\x7d\x8e\x30\x6c\xbf\xdf\x20\xca\x95\xec\xfc\xbc\x9a\x47\x6a\x8f\x47\x3d\x35\x56\x4e\x5d\xef\xd4\x5e\xa4\x54\xc0\x6a\xec\xfa\xdf\xbe\xc9\x9f\xf8\xf3\xe7\xd5\x24\x52\x5b\x59\x5c\x5b\xb8\xd4\x18\xe6\x18\x0a\x2f\x6f\x59\xc2\xcb\xec\x52\x02\x33\xda\x88\xb8\x1f\x71\x06\x85\x73\xa2\x84\x9b\xf6\xf7\xff\x0f\xfc\x57\x39\x44\x4b\xf8\xbf\xab\xc8\xff\x04\x00\x00\xff\xff\x14\x22\x96\x1d\xab\x0c\x00\x00"

func assetsServerTlsSnakeoilKeyBytes() ([]byte, error) {
	return bindataRead(
		_assetsServerTlsSnakeoilKey,
		"assets/server/tls/snakeoil.key",
	)
}

func assetsServerTlsSnakeoilKey() (*asset, error) {
	bytes, err := assetsServerTlsSnakeoilKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/server/tls/snakeoil.key", size: 3243, mode: os.FileMode(420), modTime: time.Unix(1444379500, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"assets/server/tls/snakeoil.crt": assetsServerTlsSnakeoilCrt,
	"assets/server/tls/snakeoil.key": assetsServerTlsSnakeoilKey,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"assets": &bintree{nil, map[string]*bintree{
		"server": &bintree{nil, map[string]*bintree{
			"tls": &bintree{nil, map[string]*bintree{
				"snakeoil.crt": &bintree{assetsServerTlsSnakeoilCrt, map[string]*bintree{
				}},
				"snakeoil.key": &bintree{assetsServerTlsSnakeoilKey, map[string]*bintree{
				}},
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

