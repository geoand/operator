// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package template

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 10, 11, 9, 22, 1, 958566131, time.UTC),
		},
		"/innerloop": &vfsgen۰DirInfo{
			name:    "innerloop",
			modTime: time.Date(2018, 12, 14, 18, 28, 43, 861820694, time.UTC),
		},
		"/innerloop/deployment": &vfsgen۰CompressedFileInfo{
			name:             "deployment",
			modTime:          time.Date(2018, 12, 14, 17, 38, 14, 388654090, time.UTC),
			uncompressedSize: 1558,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x54\xc1\x52\xdc\x30\x0c\xbd\xef\x57\xf8\x07\xb2\x1e\x38\x9a\x53\x67\x69\x4f\x85\xd9\x19\x3a\xdc\x15\x47\x2c\x6a\x6d\xcb\xb5\x95\xd0\x9d\x0c\xff\xde\xf1\x06\x82\x53\x02\x74\x7a\xa8\x4e\x91\x94\xf7\xfc\xe4\xa7\x31\x44\xba\xc5\x94\x89\x83\x51\x10\x63\xd6\xc3\xd9\xe6\x07\x85\xce\xa8\x4b\x8c\x8e\x8f\x1e\x83\x6c\x3c\x0a\x74\x20\x60\x36\x4a\x39\x68\xd1\xe5\xf2\xa5\x0a\xc0\xa8\x71\xdc\x5e\x83\xc7\xc7\xc7\x8d\x52\x01\x3c\xd6\x95\x1c\xd1\x96\x5f\x13\x46\x47\x16\xb2\x51\x67\x1b\xa5\x32\x3a\xb4\xc2\x69\x22\xf1\x20\xf6\xfe\x6b\xc5\xba\xc2\x2b\xe8\xa3\x03\xc1\x27\x44\x25\xa7\x84\x5b\x80\x57\xe0\x25\xfe\x94\x56\x6a\xcf\xf2\x4a\x58\x0e\x02\x14\x30\xcd\x44\x8d\x82\x74\xa8\x68\x1b\xd5\xd8\x2a\xd1\x03\x24\xed\xa8\xd5\xb9\x8f\x98\x06\xca\x9c\x3a\x6d\x39\xdc\x55\x85\x6d\xc9\x67\x8c\x65\xef\x21\x74\xe6\x23\x92\x96\x42\x9d\xcf\xbf\x63\x18\x5e\xb0\xe3\xa8\x12\x84\x03\xaa\xed\x4d\x44\xbb\xfd\x1c\x86\xac\xe6\x61\x0b\xf5\xda\xc0\x53\x0c\xe0\xfa\xa9\x75\x5b\xbe\xaa\xde\x38\x62\xe8\xaa\x9c\x3c\x1c\xd0\xa8\x9f\x3d\x1c\xb7\xc4\x3a\x07\x7e\xe8\x12\x47\x9d\x63\xa2\x70\x68\x5a\x66\x69\xf2\x39\x99\xe2\x4d\x96\x25\x6c\xdf\x3b\xb7\x67\x47\xf6\x68\xd4\x27\xf7\x00\xc7\x3c\xf7\xdf\x92\x16\x39\xc9\xe2\xc2\x67\x53\xf6\x9c\xe4\x84\x38\x0d\x5b\xb2\xc5\x44\x31\xb1\xb0\x65\x67\xd4\xb7\xdd\x7e\xae\x0f\xec\x7a\x8f\x57\xdc\x87\x25\xab\x2f\x95\x3d\xc8\xbd\x59\xbd\xff\x8a\x77\x12\x9a\xef\x21\x61\xd7\x94\x95\x7b\x83\x45\x7c\xd4\x90\x84\xee\xc0\x4a\x7e\x85\x7f\x96\x7d\x23\x9c\xe0\x80\xcb\xa9\x29\x90\xec\x56\x76\x6f\x61\xf6\xb3\x9b\xbb\xab\xcb\x9b\xd7\x46\xa6\x3e\x34\xdf\x61\x00\xa3\xfb\x9c\xb4\x63\x0b\x4e\xe7\x73\xd2\xa9\x0f\x17\xa5\x17\xb8\xc3\xa7\x1e\xb5\xf8\x0b\xed\xdc\xb5\xec\x23\x39\x5c\x45\x43\xce\xe8\x5b\x87\x17\x6d\x4f\xae\x33\xba\x9b\xdf\x83\xac\x4f\x25\x88\xf1\xe3\x45\x79\xb9\xd6\x7f\x5b\x12\xcb\xf1\xd8\xac\x99\x23\x98\x3c\x05\x10\xe2\x70\x85\x39\x17\xae\xc9\x8b\x0e\x07\x5d\x35\x1b\xc7\x87\xf7\x40\x4f\x87\x7f\x21\x87\xff\x61\x71\x26\xe6\xda\x66\x1f\xe5\x78\x49\xc9\xa8\xf1\x71\xf3\x11\xba\xf9\x9b\x85\x52\x2a\x96\x07\x3d\x0b\x06\xb9\x3d\x1d\xb7\x73\x40\xde\x54\xea\x6c\x29\x5c\xbf\x43\xf5\x3b\x00\x00\xff\xff\x7e\xf0\xac\xb5\x16\x06\x00\x00"),
		},
		"/innerloop/deploymentconfig": &vfsgen۰CompressedFileInfo{
			name:             "deploymentconfig",
			modTime:          time.Date(2018, 12, 11, 11, 20, 19, 144666461, time.UTC),
			uncompressedSize: 2062,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x55\xcd\x6e\xe3\x38\x0c\xbe\xe7\x29\x84\xde\x1d\xa3\x7b\x54\x4f\x8b\x74\x17\xd8\x43\x8b\xa0\x29\x7a\x67\x64\xc6\xe1\x8e\xfe\x40\xd1\x9e\x09\x82\xbe\xfb\x40\x71\xea\xc8\x69\xd2\x4e\x0f\x33\x3a\x99\xa4\x48\x7e\xe4\x47\x53\x10\xe9\x05\x39\x51\xf0\x5a\x41\x8c\x69\x1e\x22\xfa\xb4\xa5\x8d\xcc\x29\xd4\xfd\xed\xec\x1b\xf9\x46\xab\x7b\x8c\x36\xec\x1c\x7a\x59\x04\xbf\xa1\x76\xe6\x50\xa0\x01\x01\x3d\x53\xca\xc2\x1a\x6d\xca\x5f\x2a\xc7\xd0\x6a\xbf\x9f\x3f\x82\xc3\xd7\xd7\x99\x52\x1e\x1c\x96\x9a\x14\xd1\xe4\xab\x8c\xd1\x92\x81\xa4\xd5\xed\x4c\xa9\x84\x16\x8d\x04\xbe\x12\x44\xa9\x66\x04\x60\x0e\x00\xa6\xf6\x24\x0c\x82\xed\x6e\x70\x97\x5d\x44\xad\x9e\x82\xb5\xe4\xdb\x99\x52\x82\x2e\x5a\x10\x1c\xac\x25\xf2\x7c\x4a\xf4\x57\x92\x7f\x0e\x20\x9f\xf3\x4a\xb3\xee\xad\xda\x7c\x4c\xf0\x02\xe4\x91\xc7\x64\x95\x02\x6e\x8b\xd4\x95\xaa\x4c\x21\xd4\x3d\x70\x6d\x69\x5d\xa7\x2e\x22\xf7\x94\x02\x37\x75\xce\x5e\x28\xe6\x59\x1e\x7d\x4c\x70\x0e\x7c\xa3\x3f\x0b\xb2\x26\x5f\xca\xe3\x75\xf4\xfd\xc9\x77\xbf\x57\x0c\xbe\x45\x35\x5f\x45\x34\xf3\x7f\x7c\x9f\x54\xd1\x90\xea\x62\xc1\xc3\xe9\xc1\x76\x83\xe9\x25\x7f\x15\xb6\xfd\x1e\x7d\x53\xc8\xe4\xa0\x1d\x6e\x1e\x92\x3c\x75\x5e\xc8\xe1\x10\x50\x67\xd2\x92\x4c\xef\x2e\x3b\x6b\x97\xc1\x92\xd9\x69\xf5\xb7\xfd\x0e\xbb\x34\xda\xaf\xe1\x89\x81\x65\xd2\xe5\x91\x89\x65\x60\x39\x25\xcf\xd2\xa4\x8c\xc8\x41\x82\x09\x56\xab\xe7\xc5\x72\xd4\xf7\xc1\x76\x0e\x1f\x42\xe7\xa7\x51\x5d\xd6\x2c\x41\xb6\xfa\x62\xd3\x8b\xb8\x03\xd0\xb4\x05\xc6\xa6\xca\xb3\x78\x25\x8a\xb8\x58\x03\x0b\x6d\xc0\x48\x7a\xe7\xff\x06\x7b\x25\x81\xa1\xc5\x69\xd5\xe4\x29\xff\xaa\xef\x06\x6e\xc2\xf0\x1b\x85\x8b\x87\xfb\xd5\x7b\xf6\xb8\xf3\xd5\xff\xd0\x83\xae\xbb\xc4\xb5\x0d\x06\x6c\x9d\xfe\xa2\x9a\x3b\x7f\x97\x6d\x3e\x34\x78\xb4\xd1\x1a\x7f\xa0\x19\xad\x26\xb8\x48\x16\x2f\x7a\x43\x4a\xe8\xd6\x16\xef\xd6\x1d\xd9\x46\xd7\xa7\xff\x2a\xd5\x07\x15\xc4\x78\x3e\x1d\x37\x63\xa9\xa7\x76\x4e\x46\xe4\xe6\x8b\x33\x62\x42\xdc\x55\x97\xb8\x61\x4c\xa1\x63\x83\x49\xab\xfd\x69\x12\x04\xd9\x91\x07\xa1\xe0\x1f\x30\xa5\x9c\x61\x20\xa8\xc1\xbe\x2e\x8c\x95\x0d\xed\x47\x4e\x47\x48\xff\x92\xc5\x3f\x30\x4d\x43\xe4\x92\x7b\x17\x65\x77\x4f\x3c\xa9\xed\x9a\x77\xf5\x2b\x53\xa6\x54\xcc\x8f\x47\x12\xf4\xf2\x72\x48\xb7\xb0\x40\x4e\x17\xe8\x4c\x56\x3c\x7e\x1c\x4a\x98\xda\xf6\x38\xa6\xd5\x71\x7d\xff\x97\x89\x5c\x6c\xf3\x02\x9a\x8d\xc4\x0e\xf2\x12\x18\xdc\x58\x17\x74\x12\x1c\x08\x19\xad\x84\x3b\x3c\x5f\xb6\x39\x49\xd1\x83\x2b\xcc\x6f\x38\x14\xa8\x87\x37\xef\x80\x60\x25\x8c\xe0\x9e\xa1\xfd\x64\x82\x4e\xbb\xea\x37\x17\x70\xbe\xde\xbe\x8c\xfc\xa3\x55\xfb\x33\x00\x00\xff\xff\x9b\x1f\x53\x16\x0e\x08\x00\x00"),
		},
		"/innerloop/imagestream": &vfsgen۰CompressedFileInfo{
			name:             "imagestream",
			modTime:          time.Date(2018, 11, 20, 19, 22, 4, 196654178, time.UTC),
			uncompressedSize: 646,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x41\x8b\xdb\x30\x10\x85\xef\xfe\x15\xc3\xd2\x6b\x6c\xb6\x47\x85\x1e\x96\xf6\x52\x28\xa5\x34\xd0\xfb\x44\x1e\xa7\x6a\x24\xcd\xa0\x91\x43\x83\xf0\x7f\x2f\x96\x15\xef\x16\xf6\x64\xfb\x7b\xf3\x66\x9e\x5f\x29\x1f\x22\x06\x7a\x11\x01\xf3\x09\xfa\xef\x18\x08\x96\x05\xc5\xfd\xa2\xa4\x8e\xa3\x81\xdb\x73\x77\x75\x71\x34\xf0\xcd\x69\xee\x5c\xa6\xa0\xa6\x2b\x05\x12\xc6\x0b\x41\x7f\x12\xb2\xfd\xd7\x80\x17\x52\x58\x96\xee\x00\x6f\xbd\x6e\xe5\x3d\x0b\x45\xfd\xed\xa6\xdc\x3b\x1e\x6e\xcf\x1d\xc0\xb6\xb0\xba\x4e\x39\x11\x86\x0e\x20\x50\xc6\x11\x33\x9a\x0e\x00\xc0\xe3\x99\xbc\x6e\xef\x00\x28\x62\xe0\x35\xea\xb2\x54\xbe\x7e\xad\xb8\xa6\xae\x4c\x85\x6c\xf3\x33\x5f\x67\xf9\xc1\xde\xd9\xfb\x63\x8b\x67\x8b\xde\xc0\x84\x5e\xa9\xa2\x52\xdc\x04\xfd\x17\xb6\x57\x4a\x35\x4c\x5b\x9c\xf1\xd2\x4e\x1f\xda\xcc\x4b\x8c\x9c\x31\x3b\x8e\x9f\xc3\xa8\x6d\x0c\x00\x77\xbc\x47\x05\xb0\x61\x54\x03\x4f\x69\x8e\x87\x3f\x78\x43\x33\xcc\x9a\x86\x7a\x7b\xd0\x8f\x6e\x48\x73\x3c\xae\x5a\xe4\x91\x9a\xe6\xce\xf4\x97\xec\xaa\x1e\x2d\x07\x71\x9e\xde\x75\xa2\x2a\x85\xb3\xa7\xe3\x79\x76\x7e\x34\xc3\x48\xe2\xf9\x1e\x28\x66\x1d\x2a\x42\x91\xa7\x16\xa3\x14\x8a\xe3\x9e\x73\x4a\x1c\x5e\x03\x6e\xed\xbf\xf9\xed\x5d\xd9\x1b\xfd\x49\xc2\xbb\xdb\x05\xe1\x94\x5b\x97\x50\x1e\x78\x1b\xf6\x98\x49\x73\x43\x89\x26\x4a\x14\x2d\xfd\x5f\x3c\x40\xbe\x0b\x19\x38\xf1\x9c\xec\xa3\xfa\x9a\xaf\x3d\xfe\x05\x00\x00\xff\xff\x5d\x16\x95\x07\x86\x02\x00\x00"),
		},
		"/innerloop/ingress": &vfsgen۰CompressedFileInfo{
			name:             "ingress",
			modTime:          time.Date(2018, 12, 14, 18, 28, 43, 860373667, time.UTC),
			uncompressedSize: 253,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\xc1\xaa\x83\x30\x10\x45\xf7\xf9\x8a\xfb\x03\x2a\x6e\xf3\x07\x6f\x23\x0f\x1e\xbc\xfd\x18\x87\x1a\xd4\x18\x32\x53\x29\x48\xfe\xbd\x44\xdb\x52\xda\xdd\x3d\x73\xce\x62\x28\xfa\x7f\x4e\xe2\xd7\x60\xc1\x37\xe5\x50\xa6\x34\x5b\xdb\xb3\x52\x6b\x26\x1f\x06\x8b\x9f\x70\x49\x2c\x62\x16\x56\x1a\x48\xc9\x1a\x20\xd0\xc2\x16\xfb\x5e\x77\xb4\x70\xce\x06\x98\xa9\xe7\x59\x8a\x03\x28\xc6\x77\x29\x91\x5d\x11\xe9\x3a\xf3\x51\x54\x18\x57\xd1\xb3\x1d\x55\xe3\xb9\x80\x48\x3a\xca\x13\xaa\x03\x2d\x9a\x07\x03\x3d\xb9\x89\xc3\x60\x5f\x07\x40\x38\x6d\xde\x71\xf7\xf5\xce\x47\xf0\xbb\x26\x3d\x82\xbf\xc8\xae\x2e\x94\xf3\x3d\x00\x00\xff\xff\x47\xfe\x76\xde\xfd\x00\x00\x00"),
		},
		"/innerloop/pvc": &vfsgen۰CompressedFileInfo{
			name:             "pvc",
			modTime:          time.Date(2018, 10, 4, 9, 16, 23, 850123135, time.UTC),
			uncompressedSize: 250,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\xbd\x8a\xc3\x30\x10\x84\x7b\x3d\xc5\xbe\xc0\x19\xae\x55\xeb\xfa\x8e\x03\x83\xfb\x3d\x79\x08\x22\xfa\x8b\x76\x1d\x08\x46\xef\x1e\xe4\x98\x10\x70\xb9\x33\xdf\xec\x0c\x17\x3f\xa3\x8a\xcf\xc9\xd2\xfd\xdb\x5c\x7d\x5a\x2c\xfd\x75\x45\x14\x49\xe7\x1c\xd6\x88\x31\xb0\x8f\x26\x42\x79\x61\x65\x6b\x88\x12\x47\x58\xda\xb6\x61\x2a\x70\xc3\xa4\xb9\xf2\x05\xc3\x2f\x47\xb4\x66\x88\x02\xff\x23\x48\x07\x89\xb8\x94\x9d\x7c\x9b\x1f\xe9\x43\x93\x02\xd7\x61\x76\x0e\x22\x3f\x79\xc1\x9e\xfd\x3a\x15\x74\x6b\xff\x51\x21\x79\xad\x0e\x47\x47\xc5\x6d\x85\xe8\x71\x11\xc9\x8b\x3f\x2f\x1c\xb9\xb0\xf3\xfa\x68\xcd\x3c\x03\x00\x00\xff\xff\x76\xbd\xed\x76\xfa\x00\x00\x00"),
		},
		"/innerloop/route": &vfsgen۰CompressedFileInfo{
			name:             "route",
			modTime:          time.Date(2018, 8, 14, 17, 14, 16, 0, time.UTC),
			uncompressedSize: 172,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xcd\x31\x0e\xc2\x30\x0c\x85\xe1\x3d\xa7\xf0\x09\x82\x58\x73\x08\x06\x90\xd8\x4d\xfb\x10\x16\x4d\x6c\x25\xa6\x4b\xd5\xbb\xa3\x28\x13\xa2\xf3\xff\x3e\x3d\x36\xb9\xa3\x36\xd1\x92\xa8\xea\xc7\x11\xd5\x50\xda\x4b\x9e\x1e\x45\x4f\xeb\x39\xbc\xa5\xcc\x89\xae\xbd\x85\x0c\xe7\x99\x9d\x53\x20\x2a\x9c\x91\x68\xdb\xe2\x85\x33\xf6\x3d\x10\x2d\xfc\xc0\xd2\x7a\x23\x62\xb3\xdf\xf8\x0f\x9a\x61\xea\x63\xd7\x41\xc6\xd1\x0d\x75\x95\x09\x47\xe2\x1b\x00\x00\xff\xff\xdd\xc7\x7e\xf9\xac\x00\x00\x00"),
		},
		"/innerloop/service": &vfsgen۰CompressedFileInfo{
			name:             "service",
			modTime:          time.Date(2018, 10, 4, 7, 33, 27, 665611911, time.UTC),
			uncompressedSize: 257,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\x31\x8a\xc3\x30\x10\x45\x7b\x9d\x62\x2e\xb0\x86\x6d\xd5\x6e\xbf\x18\xbc\x6c\x3f\x91\x7f\x8c\x88\xa4\x19\xa4\xc1\x10\x8c\xef\x1e\x2c\x52\x24\x24\xe9\x86\x79\xef\xc1\x67\x8d\xff\xa8\x2d\x4a\xf1\xb4\x7e\xbb\x4b\x2c\xb3\xa7\x09\x75\x8d\x01\x2e\xc3\x78\x66\x63\xef\x88\x0a\x67\x78\xda\xb6\xe1\x97\x33\xf6\xdd\x11\x25\x3e\x21\xb5\x83\x11\xb1\xea\x33\x7c\x0d\x9a\x22\x1c\xb2\x4a\xb5\x5e\x7d\xf5\xb3\x2b\x93\x22\x0c\xa3\x54\xbb\xb7\x5a\xc5\x24\x48\xf2\xf4\xf7\x33\xf6\x8f\x71\x5d\x60\xe3\xfb\xa0\x21\x21\x98\xd4\x8f\x5b\x66\x68\x92\x6b\x46\xb1\x20\xe5\x1c\x97\x07\x7e\x0b\x00\x00\xff\xff\x7c\x58\xa5\x67\x01\x01\x00\x00"),
		},
		"/servicecatalog": &vfsgen۰DirInfo{
			name:    "servicecatalog",
			modTime: time.Date(2018, 11, 19, 8, 35, 54, 754265971, time.UTC),
		},
		"/servicecatalog/servicebinding": &vfsgen۰CompressedFileInfo{
			name:             "servicebinding",
			modTime:          time.Date(2018, 10, 11, 9, 25, 12, 782687771, time.UTC),
			uncompressedSize: 273,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8f\x31\x6e\xc3\x30\x0c\x45\x77\x9d\x82\x43\x67\x15\xd9\x0a\x01\x1d\xda\x03\x64\xa8\x81\xee\x8c\xfc\x6b\x10\xb1\x69\xc1\x14\xb2\x08\xbc\x7b\x21\xb7\xa9\x9b\x51\xff\xe9\xf1\x93\xad\x3d\x29\x2f\x78\x2b\x85\xd2\x2b\xc5\x33\x2f\x20\xf7\xd0\x1a\x6d\xac\x13\x28\x0e\x05\x39\x0e\xd8\x6e\x92\x61\xe4\xce\x45\x3e\xb1\x99\xac\x9a\xc8\x7e\xe2\xcc\x95\xe7\x75\x8a\xd7\x17\x8b\xb2\x3e\xdf\x4e\x17\x54\x3e\x85\xab\xe8\x98\xe8\x57\x7d\x17\x1d\x45\xa7\xb0\xa0\xf2\xc8\x95\x53\x20\xea\xc5\x89\x5a\xdb\x5b\xdd\x03\xd1\xcc\x17\xcc\xd6\x19\x11\x97\xd2\xe1\x7d\xbd\x9d\x1f\xce\xbf\xd8\x0a\x72\x57\x44\xad\xb2\x66\x7c\xe0\x2b\x3d\xfc\x3d\xe6\x1b\xf2\x86\x7a\xbe\xe7\xc3\xdf\x73\x3f\x19\x3a\xba\x7f\x07\x00\x00\xff\xff\x85\x6f\x46\xd9\x11\x01\x00\x00"),
		},
		"/servicecatalog/serviceinstance": &vfsgen۰CompressedFileInfo{
			name:             "serviceinstance",
			modTime:          time.Date(2018, 11, 19, 8, 35, 54, 753585081, time.UTC),
			uncompressedSize: 333,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xbd\x6a\x04\x21\x14\x85\x7b\x9f\xe2\x16\xa9\x0d\xdb\x05\x21\x45\x08\x29\x92\x62\x09\x0c\xa4\x3f\xeb\x5c\x16\x59\xe7\x8e\x78\xcd\x12\x10\xdf\x3d\xe8\xee\xe4\x87\x2d\xfd\xce\x39\x7e\x62\xad\x77\x82\x85\x9f\x52\x22\xf7\x48\x76\x8f\x85\xa9\x35\x53\x2b\x65\xc8\x91\xc9\x4e\x89\xbd\x9d\x38\x9f\x83\x67\xa5\xd6\x90\xc2\x07\x67\x0d\xab\x38\xd2\x0b\xf6\x28\x88\xeb\xd1\x9e\x1e\xd4\x86\xf5\xfe\xbc\x3b\x70\xc1\xce\x9c\x82\xcc\x8e\xae\xd3\x57\xd1\x02\xf1\x6c\x16\x2e\x98\x51\xe0\x0c\x51\x37\x3b\xaa\x75\x68\x5b\x33\x44\x11\x07\x8e\xda\x33\x22\xa4\xd4\xc3\xed\x7d\x23\xff\xdd\xfc\xc1\x9a\xd8\xf7\x89\x8f\x9f\x5a\x38\x5f\x8d\xcf\x11\xaa\x2f\x5f\x85\xb3\x20\xee\x37\xd3\xa0\xe3\xaa\xff\xed\xf7\x08\xb9\x29\x77\x38\xba\x09\x19\x0b\x17\xce\x7a\xe1\x3f\xc7\xb7\x69\x95\xf1\x5d\x2c\x73\x6b\xdf\x01\x00\x00\xff\xff\xdd\xed\x80\x29\x4d\x01\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/innerloop"].(os.FileInfo),
		fs["/servicecatalog"].(os.FileInfo),
	}
	fs["/innerloop"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/innerloop/deployment"].(os.FileInfo),
		fs["/innerloop/deploymentconfig"].(os.FileInfo),
		fs["/innerloop/imagestream"].(os.FileInfo),
		fs["/innerloop/ingress"].(os.FileInfo),
		fs["/innerloop/pvc"].(os.FileInfo),
		fs["/innerloop/route"].(os.FileInfo),
		fs["/innerloop/service"].(os.FileInfo),
	}
	fs["/servicecatalog"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/servicecatalog/servicebinding"].(os.FileInfo),
		fs["/servicecatalog/serviceinstance"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
