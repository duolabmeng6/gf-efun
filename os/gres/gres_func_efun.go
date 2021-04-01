// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gres

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/encoding/gcompress"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/os/gfile"
)


// Pack packs the path specified by <srcPaths> into bytes.
// The unnecessary parameter <keyPrefix> indicates the prefix for each file
// packed into the result bytes.
//
// Note that parameter <srcPaths> supports multiple paths join with ','.
func E打包(srcPaths string, keyPrefix ...string) ([]byte, error) {
	var (
		buffer       = bytes.NewBuffer(nil)
		headerPrefix = ""
	)
	if len(keyPrefix) > 0 && keyPrefix[0] != "" {
		headerPrefix = keyPrefix[0]
	}
	err := zipPathWriter(srcPaths, buffer, headerPrefix)
	if err != nil {
		return nil, err
	}
	// Gzip the data bytes to reduce the size.
	return gcompress.Gzip(buffer.Bytes(), 9)
}

// PackToFile packs the path specified by <srcPaths> to target file <dstPath>.
// The unnecessary parameter <keyPrefix> indicates the prefix for each file
// packed into the result bytes.
//
// Note that parameter <srcPaths> supports multiple paths join with ','.
func 打包到文件(srcPaths, dstPath string, keyPrefix ...string) error {
	data, err := Pack(srcPaths, keyPrefix...)
	if err != nil {
		return err
	}
	return gfile.PutBytes(dstPath, data)
}

// PackToGoFile packs the path specified by <srcPaths> to target go file <goFilePath>
// with given package name <pkgName>.
//
// The unnecessary parameter <keyPrefix> indicates the prefix for each file
// packed into the result bytes.
//
// Note that parameter <srcPaths> supports multiple paths join with ','.
func 打包到GO文件(srcPath, goFilePath, pkgName string, keyPrefix ...string) error {
	data, err := Pack(srcPath, keyPrefix...)
	if err != nil {
		return err
	}
	return gfile.PutContents(
		goFilePath,
		fmt.Sprintf(gstr.TrimLeft(packedGoSouceTemplate), pkgName, gbase64.EncodeToString(data)),
	)
}

// Unpack unpacks the content specified by <path> to []*File.
func 解包(path string) ([]*File, error) {
	realPath, err := gfile.Search(path)
	if err != nil {
		return nil, err
	}
	return UnpackContent(gfile.GetContents(realPath))
}

// UnpackContent unpacks the content to []*File.
func 解包内容(content string) ([]*File, error) {
	var data []byte
	var err error
	if isHexStr(content) {
		// It here keeps compatible with old version packing string using hex string.
		// TODO remove this support in the future.
		data, err = gcompress.UnGzip(hexStrToBytes(content))
		if err != nil {
			return nil, err
		}
	} else if isBase64(content) {
		// New version packing string using base64.
		b, err := gbase64.DecodeString(content)
		if err != nil {
			return nil, err
		}
		data, err = gcompress.UnGzip(b)
		if err != nil {
			return nil, err
		}
	} else {
		data, err = gcompress.UnGzip(gconv.UnsafeStrToBytes(content))
		if err != nil {
			return nil, err
		}
	}
	reader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, err
	}
	array := make([]*File, len(reader.File))
	for i, file := range reader.File {
		array[i] = &File{file: file}
	}
	return array, nil
}