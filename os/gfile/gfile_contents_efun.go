// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gfile

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"

	"github.com/gogf/gf/util/gconv"
)



// GetContents返回<path>的文件内容作为字符串。
//如果读取失败，返回一个空字符串。
func E读入文本(path string) string {
	return gconv.UnsafeBytesToStr(GetBytes(path))
}

// GetBytes返回<path>的文件内容作为[]字节。
//如果读取失败，返回nil。
func E读入文件(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	return data
}

// putContents将二进制内容放入<path>的文件中。
func E写到文件2(path string, data []byte, flag int, perm os.FileMode) error {
	// It supports creating file of <path> recursively.
	dir := Dir(path)
	if !Exists(dir) {
		if err := Mkdir(dir); err != nil {
			return err
		}
	}
	// Opening file with given <flag> and <perm>.
	f, err := OpenWithFlagPerm(path, flag, perm)
	if err != nil {
		return err
	}
	defer f.Close()
	if n, err := f.Write(data); err != nil {
		return err
	} else if n < len(data) {
		return io.ErrShortWrite
	}
	return nil
}

// Truncate通过<size>将<path>的文件截断到给定的大小。
func E截断文件(path string, size int) error {
	return os.Truncate(path, int64(size))
}

// PutContents将字符串<content>放到<path>的文件中。
//如果文件不存在，则递归创建<path>的文件。
//函数返回路径字符串，返回内容字符串
// return putContents(path， []byte(content)， os. o_wronly |os. o_create |os。DefaultPermOpen O_TRUNC)
// }

// PutContentsAppend将string <content>追加到<path>的文件。
//如果文件不存在，则递归创建<path>的文件。
func E写到文件追加文本(path string, content string) error {
	return putContents(path, []byte(content), os.O_WRONLY|os.O_CREATE|os.O_APPEND, DefaultPermOpen)
}

// PutBytes把<content>的二进制文件放到<path>的文件中。
//如果文件不存在，则递归创建<path>的文件。
func E写到文件(path string, content []byte) error {
	return putContents(path, content, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, DefaultPermOpen)
}

//将二进制<content>追加到<path>的文件。
//如果文件不存在，则递归创建<path>的文件。
func E写到文件追加字节集(path string, content []byte) error {
	return putContents(path, content, os.O_WRONLY|os.O_CREATE|os.O_APPEND, DefaultPermOpen)
}

// 返回给定<char>从<start>开始的文件偏移量。
func E寻找字节集(reader io.ReaderAt, char byte, start int64) int64 {
	buffer := make([]byte, DefaultReadBuffer)
	offset := start
	for {
		if n, err := reader.ReadAt(buffer, offset); n > 0 {
			for i := 0; i < n; i++ {
				if buffer[i] == char {
					return int64(i) + offset
				}
			}
			offset += int64(n)
		} else if err != nil {
			break
		}
	}
	return -1
}

// GetNextCharOffsetByPath返回给定的<char>从<start>开始的文件偏移量。
//打开一个<path>的文件供os读取。O_RDONLY标志和默认选项。
func E寻找字节集从路径(path string, char byte, start int64) int64 {
	if f, err := OpenWithFlagPerm(path, os.O_RDONLY, DefaultPermOpen); err == nil {
		defer f.Close()
		return GetNextCharOffset(f, char, start)
	}
	return -1
}

// getbytestlchar以[]字节的形式返回文件的内容
//直到下一个指定的字节<char>的位置。
//
//注意:返回值包含最后一个位置的字符。
func E取字节集从某字节开始(reader io.ReaderAt, char byte, start int64) ([]byte, int64) {
	if offset := GetNextCharOffset(reader, char, start); offset != -1 {
		return GetBytesByTwoOffsets(reader, start, offset+1), offset
	}
	return nil, -1
}

// GetBytesTilCharByPath返回<path>以[]字节的形式给出的文件内容
//直到下一个指定的字节<char>的位置。
//打开一个<path>的文件供os读取。O_RDONLY标志和默认烫发。
//
//注意:返回值包含最后一个位置的字符。
func E取字节集从某字节开始从文件(path string, char byte, start int64) ([]byte, int64) {
	if f, err := OpenWithFlagPerm(path, os.O_RDONLY, DefaultPermOpen); err == nil {
		defer f.Close()
		return GetBytesTilChar(f, char, start)
	}
	return nil, -1
}

//返回从<start>到<end>的[]字节的二进制内容。
//注意:返回值不包含最后一个位置的字符，这意味着
//它返回的内容范围为[start, end)。
func E取字节集中间(reader io.ReaderAt, start int64, end int64) []byte {
	buffer := make([]byte, end-start)
	if _, err := reader.ReadAt(buffer, start); err != nil {
		return nil
	}
	return buffer
}

// GetBytesByTwoOffsetsByPath返回二进制内容[]字节从<start>到<end>。
//注意:返回值不包含最后一个位置的字符，这意味着
//它返回的内容范围为[start, end)。
//打开一个<path>的文件供os读取。O_RDONLY标志和默认烫发。
func E取字节集中间从路径(path string, start int64, end int64) []byte {
	if f, err := OpenWithFlagPerm(path, os.O_RDONLY, DefaultPermOpen); err == nil {
		defer f.Close()
		return GetBytesByTwoOffsets(f, start, end)
	}
	return nil
}

// ReadLines逐行读取文件内容，它作为字符串传递给回调函数<callback>。
//它匹配以字符'\r'或'\n'分隔的每一行文本，并去除任何末尾的行结束标记。
//
//注意传递给callback函数的参数可能是一个空值，最后一个非空行
//将被传递给回调函数<callback>，即使它没有换行标记。
func E逐行读取(file string, callback func(text string) error) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if err = callback(scanner.Text()); err != nil {
			return err
		}
	}
	return nil
}


// ReadLinesBytes逐行读取文件内容，它作为[]字节传递给回调函数<callback>。
//它匹配以字符'\r'或'\n'分隔的每一行文本，并去除任何末尾的行结束标记。
//
//注意传递给callback函数的参数可能是一个空值，最后一个非空行
//将被传递给回调函数<callback>，即使它没有换行标记。
func E逐行读取字节集 (file string, callback func(bytes []byte) error) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if err = callback(scanner.Bytes()); err != nil {
			return err
		}
	}
	return nil
}
