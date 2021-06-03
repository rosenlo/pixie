// Code generated for package schema by go-bindata DO NOT EDIT. (@generated)
// sources:
// schema.graphql
package schema

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

// Mode return file modify time
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

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x58\x5f\x6f\x23\xb7\x11\x7f\xd7\xa7\x18\xc7\x0f\x67\x03\xaa\x1a\x14\x4d\x50\xe8\x29\x1b\x49\x97\xdb\xda\x96\x55\x4b\xbe\x6b\x70\x30\x2e\xd4\x72\xa4\x25\xb4\x4b\x6e\x48\xae\x6c\x35\xb8\xef\x5e\x0c\xc9\xfd\x43\x49\xbe\xeb\x05\x7d\x92\x96\x1c\xce\x3f\xce\xfc\x66\x86\x97\xb0\xca\x85\x81\x8d\x28\x10\x38\x9a\x4c\x8b\x35\x1a\xb0\x39\x82\xc9\x72\x2c\x19\x6c\xb4\x2a\xdd\x77\xb2\x48\xc1\xa0\xde\x8b\x0c\x47\x83\xcb\xc1\x25\xa4\xf6\x8d\x01\xa9\x2c\x08\x8e\xac\x18\xc2\xba\xb6\xf0\x8c\x20\x11\x39\x58\x05\x25\x93\x35\x2b\x8a\x03\x6c\x51\xa2\x66\x16\xc1\x1e\x2a\x34\xb0\x51\xda\xf1\x5b\x1d\x2a\x5c\x66\x5a\x54\x16\x1e\xd3\xc1\x25\x3c\xe7\x28\xc1\xb6\xca\x08\x03\x75\xc5\x99\x45\x3e\xf2\x2a\x66\x4c\xc2\x1a\x81\x2b\x89\xb0\x3e\x80\xae\xa5\x14\x72\x0b\xbf\x79\xaa\x91\xc9\x7f\x03\x11\x18\x70\xa1\x31\xb3\x4a\x1f\x46\x83\x60\x5f\xa3\x83\x09\x66\x8d\xf8\xc8\x9a\x53\xfa\x21\x30\xc9\x21\x53\x95\x40\x03\xc2\x92\x19\xa4\xea\x4f\x95\x78\x11\xf8\x97\x82\xad\xcd\x5f\x59\x25\xa0\x62\xd9\x8e\x6d\xd1\x99\x52\x1b\xa7\x4e\x67\xcd\x68\x30\x08\xae\xfb\x63\x00\xf0\x7b\x8d\xfa\x30\x86\x7f\xd1\xcf\x00\xa0\xac\x2d\xb3\x42\xc9\x31\xdc\x85\x7f\x83\xcf\x83\x01\x39\x06\x1e\x0d\xea\x54\x6e\x94\x3b\x26\xf8\x18\xd2\xe9\xc5\x00\x40\xb2\x12\xc7\xb0\xb4\x5a\xc8\x2d\x7d\x63\xc9\x44\xd1\x5f\xa8\x44\x66\x6b\x1d\xd1\x28\xbd\x9d\x1f\x1d\x53\x7a\x9b\x4e\xfb\x0b\xc2\x24\x55\xa5\xd5\x1e\xf9\x18\x7e\x56\xaa\x40\x26\x2f\x5a\x5d\xee\xf5\xf6\x7f\x51\x45\xb2\x75\x81\x9e\x0d\x2b\xcc\x19\x3e\x64\xd3\x12\xad\xa5\x8b\x22\x5e\x3b\x3c\xf4\x19\xec\x59\x51\xf7\x38\x7e\x1e\x0c\x50\xd6\x25\x24\xda\x8a\x0d\xcb\x2c\xf9\xd4\x1d\x03\x48\x56\x9f\x1e\xe7\x37\xf3\xfb\x0f\xf3\xe6\xf3\x36\x9d\x3f\xfe\xfb\x53\x72\x37\xfd\xf1\xef\xcd\xd2\x34\x79\xf8\x90\xce\xe3\xb5\xc9\xfd\x7c\x95\xa4\xf3\xd9\xc3\xa7\xe5\x6c\xf5\xe9\xd7\xe4\xee\x76\x79\x7e\xeb\x0c\xbf\x98\x60\x35\xbb\x5b\xdc\x26\xab\x59\x60\xd2\x2a\x5b\x5b\x95\xa9\xb2\x2a\xd0\xe2\x4c\x5a\x61\x0f\x4b\x4b\xc1\x4e\x7a\x27\xb3\x65\x4f\x6d\xfa\x5a\xcc\xe6\xd3\x74\xfe\x4b\xf8\x7a\x78\x9c\xcf\xbb\xaf\xb7\x49\x7a\x3b\x9b\x86\x8f\xd5\xec\xe1\x2e\x9d\x27\xab\xd9\xf4\xac\xa4\x24\xa3\xd8\x69\x1d\x94\x44\xfe\xb9\x84\x44\x02\x72\x61\x81\x39\x32\x50\x59\x56\x6b\x03\x62\x03\x8c\xe2\x55\x43\xce\x0c\x94\x8a\x8b\x8d\xa0\x5c\xcd\x11\x84\xac\x6a\x0b\x16\x5f\x2c\x05\xb3\x90\x06\xb5\xbb\x34\xa5\x81\x63\x81\xee\x7f\x96\x33\xcd\x32\x8b\xda\x8c\x9c\x10\x97\x58\x42\x66\x45\xcd\x09\x32\x0e\x95\x3b\x20\x1d\xbf\x1d\x1e\xd6\x8a\x69\xee\xf2\xa9\x62\xc6\x33\x50\x65\xc9\x24\x77\xc7\x49\xe3\xd9\x34\x5d\x79\x75\xc1\x60\x81\x59\xa7\xaf\x2c\x0e\xe7\x95\xce\x72\x65\x50\x02\x93\xc0\x7a\xde\x00\x53\x6f\xb7\x68\xe8\xec\xa8\x51\x8b\x8b\xcc\xe5\xbb\x55\x4e\x04\x29\x15\x1d\xb1\x39\xb3\x94\xe2\x26\x57\x75\xc1\xa1\x54\x7b\x74\x44\x24\xea\x8d\x01\x92\x4d\x40\xe5\x01\x40\x92\x63\x18\x05\x7a\xa5\x85\x83\x32\xb6\x6e\xac\x58\xce\x6e\x67\x93\xd5\x17\xe2\xe1\x46\x48\x1e\xc2\xe1\x26\x0a\x87\x9b\x4f\x8b\xfb\x69\xf8\xb7\x7c\x3f\x69\xfe\x4d\x1e\xd2\xc5\x2a\x7c\xcc\x93\xbb\xd9\x72\x91\x4c\x66\x6d\x4a\x4d\xb1\x2a\xd4\xa1\x44\x69\x6f\xf0\x70\x94\xa0\x47\xe9\x95\x69\x24\xfc\x4c\xec\x9d\x19\xc3\xdb\x42\x31\x4b\xab\x84\xf2\x51\xce\x39\xb6\xc9\x22\xfd\xbf\xf2\x73\x80\xe7\xd8\x91\x43\xc7\x2d\xba\x05\x2c\x1a\x37\x10\x73\x11\x28\x02\x4a\x98\xab\x1d\x1e\xcc\x18\x3e\x06\x7e\x4f\x17\xd7\x63\xf8\xd8\x83\x91\x8b\xa7\xc0\x81\xd6\x4c\xd8\x23\x46\x6e\x3d\x2b\x6a\x63\x51\x5f\x05\x13\xae\xc7\x30\xf1\x2b\x8d\xa8\x40\xf0\xf3\x81\xf0\xf1\x2a\x02\xb4\x57\x89\x49\x4a\x7f\xa7\x2f\x69\xa2\xa4\x44\x17\xb4\xa7\x32\xbb\xbd\x8e\xa1\x68\xb0\xed\x8a\xf5\x40\x6e\x1c\x41\x1e\x71\xb8\x4d\x9b\x15\x3a\xd7\xd0\x9a\xf6\x54\x1f\xde\xaf\xbb\xe3\xa6\x91\xd4\x0f\xf5\x2b\x97\xdc\x0d\xf5\x30\x84\xf6\x42\x99\x31\xa4\xd2\x0e\x43\xd2\x8d\x5f\xc1\x97\x61\x63\xea\x63\x57\x3f\xae\x63\xe2\x07\x34\x75\x61\x8f\xc5\xbe\x15\x58\xf0\x63\xd9\x1b\x5a\x0c\x26\x9f\x4d\x94\xa1\x03\x5f\x8d\xbf\xd7\x42\x23\x4f\xf4\x96\x88\xe9\x06\xce\x93\x3f\xbd\xa2\x5e\x44\xbe\x6c\xd1\xe1\x69\xe0\xc0\xc0\x17\xea\x72\xab\x01\x25\xaf\x94\x90\xd6\x0c\x41\xe3\x06\x5d\xc2\x73\x95\xb9\xd6\x20\x2b\x54\xcd\x59\x25\x46\x95\x56\x0e\x45\x0a\xb1\xc7\xf7\x02\x9f\x49\x9d\xdb\xf0\xff\x0e\x2d\xe3\xcc\x32\x1f\x15\x0d\xc5\x44\x49\x8b\xd2\x9a\x5e\x50\xdc\x1e\x6d\x11\xb9\x71\x7a\xb8\x70\x77\xff\x62\x66\x7e\xf7\x0c\xab\x65\xb4\x71\xe1\x6d\xf2\xb8\x40\x29\x6b\x5c\x52\xf6\x50\x82\x04\x44\xb0\xe1\xf9\x47\x34\x3d\xf6\x31\xa9\xe7\x4e\xcd\x5f\x60\xcd\x2a\x11\x78\x7a\xcc\xf0\xcc\xfc\x6a\x8f\x4b\xd8\xec\x00\xe6\x24\x60\x1c\x3a\x50\x19\x71\x8d\x54\xc9\xac\x45\x1e\x0a\x91\x30\xbd\xaa\x64\x42\xec\x3c\xe7\x22\xcb\x5d\x15\x58\x23\x4a\xa8\x98\x36\xc8\xa9\x5c\x9d\x62\xbb\x6a\x0b\x80\x07\x7f\xb6\x5e\x5a\x55\x41\xa5\x8c\xa0\x30\x70\x15\xa8\x95\x99\xf6\x43\xd4\xd1\x7f\xc8\xd1\xe6\xa8\x4f\x74\x20\xbd\x18\x75\x2d\x82\x0f\x01\x5f\x30\xab\x2d\x75\x40\x4d\x61\x1b\xb9\xb6\x6a\xd6\xae\xb7\x0d\x91\x77\x61\x51\xf4\xea\x94\xef\x83\x91\x65\x39\xa8\x8d\x13\x14\x94\x74\xba\xd1\xff\x8e\x74\x0c\x1f\x57\xfd\x85\xa7\xd6\xa9\xd1\x72\xcf\x9f\x42\x72\x7c\xe9\x31\xf6\xd5\xce\xe6\x68\x30\xd2\x81\x69\xe7\xfb\x20\x32\xa5\x53\x0e\x14\x22\x2f\xf8\xda\x4c\xe6\xb3\xde\xe1\xd0\xc7\xd3\x4d\xb1\x75\x10\xf8\x2c\x8a\x02\x4a\xb6\xf3\x95\x34\x78\xa5\xe7\x28\x92\xd3\x7d\x25\x1b\x4b\xb0\x4e\xcc\xfb\x9e\x32\x91\xe1\xaf\xe5\xf1\xb9\xb0\x3a\x72\xc5\x4e\x48\xfe\x1a\xcc\x1c\x75\xb3\xa1\x8c\x51\x5a\x39\x28\x6c\x57\x4b\x66\xb3\x9c\x42\x84\xe3\x8b\x83\xa1\x54\xda\x27\x52\x92\x7a\xbc\x73\xcc\x5d\xf3\xd7\x36\x03\xa1\x12\xd0\x62\x6d\x7a\xf7\xc3\x71\xc3\x28\x03\x1c\x1b\x6a\x71\xa4\xb2\x79\x08\xb0\x9d\x54\xcf\x92\x3c\x35\x59\x46\x3d\x1d\x9d\x0b\xf4\x06\x72\x64\x85\xcd\x0f\x74\x34\x47\xa6\xed\x1a\x99\xf5\xd7\xa9\x31\x43\xb1\x47\x4e\x9d\x98\xc6\x6d\x5d\x30\x0d\x42\x5a\xd4\xd4\xa2\xbb\x76\xcc\xe6\x3e\x2b\x02\x70\x12\x3b\x8d\xa6\x52\x92\x93\x06\x56\x39\xf4\x45\x63\x4d\x50\xe2\xdd\x2c\xb9\x5d\xbd\xfb\xf5\x54\x89\x5a\xf6\xd4\x70\x37\xde\x71\xcc\x7c\xf1\xf3\xa3\xe0\x82\xc6\x27\x98\x10\x9e\x3a\x0d\x84\x01\x9a\x19\x04\x6f\x12\xae\xb3\xc1\x8f\x91\xc2\xc8\x37\xd6\xcf\x4e\x2e\xc1\x28\xd6\x8c\x2a\xc9\x3a\x66\xda\x26\x4f\xa3\xc1\x72\x5d\xa0\x81\x77\xab\xd5\xe2\x8d\x81\x1f\xbe\xff\xde\x69\x57\x9b\xd6\x7f\xe7\x95\x77\x81\xba\x55\x0e\xed\x85\xe9\x74\x0d\x76\xfc\xf2\xb0\x98\x34\x16\x50\x34\xad\x35\xb2\x9d\x19\x39\x06\xb9\xaa\xd0\xe3\x13\xb3\x6d\x67\xd9\x18\xee\xf8\x66\xa4\xe8\x9a\x65\x3b\xea\x63\x85\x44\x67\xb2\x46\x53\x97\x94\x98\x10\x34\xf2\x9a\x04\x3d\xa7\xe9\x72\x72\x3f\x9f\xcf\x26\x2b\x37\x00\x1c\xfb\x99\x26\x5c\xba\x9b\x30\x22\x47\x8e\x16\x7e\xa5\xd2\x2a\x43\x63\x28\xe7\x1b\xf2\xc6\x07\x8b\x69\xb2\xf2\x53\x86\xe7\xbb\x17\xff\x11\x4d\x3b\xdd\x58\xee\xdd\x4e\x4b\x34\xd0\x1b\x94\x16\x98\x3c\x80\x72\x00\xb0\xa9\xb5\x07\x02\x1f\xc6\x7e\xe0\x36\xc0\xd6\xaa\xf6\x2e\x78\x0e\x48\x21\x6c\x3f\x36\x69\x38\x8e\x55\x39\xb5\x31\xe8\xf2\xcc\x0c\x58\x7d\x08\xf1\xe7\x05\x78\x95\x36\x4c\x14\xd8\x46\x8d\x54\xcf\x64\x30\x83\x35\xe3\x91\x03\x9d\x91\xb3\x66\x84\x6a\xc0\xe1\xbd\xe3\x3e\x51\x72\x23\xfc\xfc\x59\x31\x63\x6c\xae\x55\xbd\xcd\x67\x6e\x78\xed\x86\xdf\xf6\x10\xd5\x55\x26\x64\x94\xb5\xc7\x93\xef\xf9\x4e\x38\x80\x42\x47\x56\xa2\x31\x6c\xdb\x47\x19\x1f\xbf\xed\x42\x23\xf3\xe6\x1f\x66\xb6\x27\xa7\xff\x71\xf6\xd4\x46\x68\x63\x57\xa2\xc4\x48\x5c\xc1\x4e\x16\x1b\x7e\x0b\xc5\xff\x94\xf6\xb5\xf9\x26\xf5\x81\x02\xc8\x3b\xcb\x35\xc9\xb1\xe7\x7c\x63\x80\x64\x17\xed\x36\x36\xd2\x72\xeb\xeb\xae\xad\x3e\x1a\x3d\x1a\x6d\x22\x04\x6d\xac\x7e\xd7\xc0\x45\x64\xc0\xbe\x77\xdb\xe3\xe8\xee\xbb\xdd\xf7\xa8\x4d\x8c\xf0\x21\x95\x5e\xdd\x98\xc7\x65\xa2\xd2\x68\xed\x61\x72\x76\xef\xb4\x13\x0d\x1e\xd2\xaa\x58\x14\x4c\x62\x7b\x2d\xae\x94\xb4\x5f\xde\x51\xb2\x2e\xe7\x8a\xa3\xef\xca\xc3\x42\x2a\x8d\xd5\x35\xb5\x62\xc8\xfb\x9b\x47\xfe\x8b\x27\x0d\xef\xc9\x2a\xe1\x5c\xa3\x89\x2e\xd4\xaa\x1d\xca\xd3\x59\xcd\x8f\x50\x7b\x11\x5e\x2d\x4e\x5e\x98\x84\xdb\xbb\x15\x72\x77\x7a\xb6\x79\xc6\x72\x27\x27\x2e\xb0\x82\x52\xd1\x38\x05\x3f\x71\xac\x34\x52\x47\xc6\xaf\x9a\x28\xfa\x2e\x10\xf8\xc2\x45\xb9\x1d\x22\x13\xf6\x82\x41\xf5\x12\x3a\xd4\xef\xae\x07\x00\x8f\x0e\x13\xfa\x97\x7a\x15\xdc\x4d\xde\x4e\xa7\x17\xc3\x2f\x25\xf7\x75\xef\x6d\xaa\x51\x33\x6a\x73\x4f\xba\x5e\x80\x29\x52\x55\x9f\xbe\xd2\x24\x9f\xb0\xf3\xfd\x6e\xd7\xf7\x36\x0c\xfc\xf7\xf9\x93\xde\xa8\xc7\x2f\x8d\xbf\x43\xff\x48\xf6\xf5\x91\xd8\x5f\x20\x2d\x5f\xc5\x17\x38\xf4\xf0\x11\x3d\x08\x0e\x5d\x1a\xcd\x8f\xc6\xdf\x2e\x0e\x62\xe5\x16\xa8\x4b\x61\x28\x3d\xcc\x15\xcd\xea\xad\xc7\xeb\x78\x73\x0c\x33\x2e\x5c\x6b\x77\x74\xaa\x63\xee\xa7\x53\xcf\xfa\x5e\x6f\x5b\xb3\xc3\xf3\xa4\xe3\xaa\xba\xf5\x8e\x63\x8f\x98\xb8\xb5\xaf\x07\x9f\x07\x03\xdf\xa1\xbf\x22\xda\x27\xc3\xe9\x5b\xe7\xe9\xc1\x9e\x04\x9f\x06\xaf\xbc\x6e\x76\xc9\xd7\x0d\xe9\xfe\xa1\x43\x47\x59\x63\x72\xf6\xb7\x1f\x7e\x3c\xf3\xd2\xd2\x9f\xd5\xbd\x7a\x16\x4b\xd7\xeb\x86\x9d\xa7\x13\x5a\x47\xb6\x8f\x21\xca\x21\x79\xce\xe4\x16\x0b\xb5\x8d\xb2\x5c\x94\x68\x2c\x2b\xab\xb8\x30\x5c\xc2\xc3\x57\x26\x5c\x27\xf2\x78\xb0\xfd\xca\x1b\xf0\xc9\xeb\xcf\x37\x8a\x69\xa6\xd8\x50\xf9\xbc\xcc\xf1\x89\x16\xee\xa1\xfb\xa5\x68\xa8\xa3\x47\x64\x61\xfe\xb9\xbc\x9f\xff\x19\x25\xe2\xa9\xfb\x9b\x2c\x05\x6a\x96\x1a\x2d\xe3\x97\xef\x6f\x12\xfe\x8a\xfd\x47\xef\x01\xa1\x8e\xc4\xa6\x7f\x1e\xfc\x37\x00\x00\xff\xff\x2d\x24\x10\x1e\xbd\x19\x00\x00")

func schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaGraphql,
		"schema.graphql",
	)
}

func schemaGraphql() (*asset, error) {
	bytes, err := schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.graphql", size: 6589, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
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
	"schema.graphql": schemaGraphql,
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
	"schema.graphql": &bintree{schemaGraphql, map[string]*bintree{}},
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
