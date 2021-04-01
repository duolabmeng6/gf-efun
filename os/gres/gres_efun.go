// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gres provides resource management and packing/unpacking feature between files and bytes.
package gres



// Add解包并将<内容>添加到默认资源对象中。
//不必要的参数<前缀>表示前缀
//每个文件存储到当前资源对象。
func E添加(内容 string, 前缀 ...string) error {
	return defaultResource.Add(内容, 前缀...)
}

// Load加载、解包并将<路径>中的数据添加到默认资源对象中。
//不必要的参数<前缀>表示前缀
//每个文件存储到当前资源对象。
func E加载(路径 string, 前缀 ...string) error {
	return defaultResource.Load(路径, 前缀...)
}

// Get返回给定路径的文件。
func E获取(路径 string) *File {
	return defaultResource.Get(路径)
}

// GetWithIndex搜索带有<路径>的文件，如果文件是directory
//然后在这个目录下索引文件搜索。
//
// GetWithIndex通常用于http静态文件服务。
func E获取索引文件(路径 string, indexFiles []string) *File {
	return defaultResource.GetWithIndex(路径, indexFiles)
}

// GetContent直接返回默认资源对象中<路径>的内容。
func E获取内容(路径 string) []byte {
	return defaultResource.GetContent(路径)
}

//包含检查<路径>是否存在于默认资源对象中。
func E是否存在(路径 string) bool {
	return defaultResource.Contains(路径)
}

// IsEmpty检查并返回资源管理器是否为空。
func E是否为空() bool {
	return defaultResource.tree.IsEmpty()
}

// ScanDir返回在给定路径下的文件，参数<路径>应该是文件夹类型。
//
//模式参数<匹配规则>支持多个文件名模式，
//使用'，'符号来分隔多个模式。
//
//如果给定参数<递归>为true，则递归扫描目录。
func E搜索目录和文件(路径 string, 匹配规则 string, 递归 ...bool) []*File {
	return defaultResource.ScanDir(路径, 匹配规则, 递归...)
}

// ScanDirFile返回所有子文件，其绝对路径为给定的<路径>，
//如果给定参数<递归>为true，则递归扫描目录。
//
//注意它只返回文件，而不返回目录。
func E搜索文件(路径 string, 匹配规则 string, 递归 ...bool) []*File {
	return defaultResource.ScanDirFile(路径, 匹配规则, 递归...)
}

// Dump打印默认资源对象的文件。
func E调试() {
	defaultResource.Dump()
}
