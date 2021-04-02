// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package glog

import (
	"io"
)

// SetConfig set configurations for the logger.
func E设置配置(config Config) error {
	return logger.SetConfig(config)
}

// SetConfigWithMap set configurations with map for the logger.
func E设置配置Map(m map[string]interface{}) error {
	return logger.SetConfigWithMap(m)
}

// SetPath sets the directory path for file logging.
func E设置输出目录(path string) error {
	return logger.SetPath(path)
}

// GetPath returns the logging directory path for file logging.
// It returns empty string if no directory path set.
func E获取路径() string {
	return logger.GetPath()
}

// SetFile sets the file name <pattern> for file logging.
// Datetime pattern can be used in <pattern>, eg: access-{Ymd}.log.
// The default file name pattern is: Y-m-d.log, eg: 2018-01-01.log
func E设置文件(pattern string) {
	logger.SetFile(pattern)
}

// SetLevel sets the default logging level.
func E设置级别(level int) {
	logger.SetLevel(level)
}

// GetLevel returns the default logging level value.
func E获取级别() int {
	return logger.GetLevel()
}

// SetWriter sets the customized logging <writer> for logging.
// The <writer> object should implements the io.Writer interface.
// Developer can use customized logging <writer> to redirect logging output to another service,
// eg: kafka, mysql, mongodb, etc.
func E设置io流(writer io.Writer) {
	logger.SetWriter(writer)
}

// GetWriter returns the customized writer object, which implements the io.Writer interface.
// It returns nil if no customized writer set.
func E获取io流() io.Writer {
	return logger.GetWriter()
}

// SetDebug enables/disables the debug level for default logger.
// The debug level is enabled in default.
func E设置调试级别(debug bool) {
	logger.SetDebug(debug)
}

// SetAsync enables/disables async logging output feature for default logger.
func E设置异步(enabled bool) {
	logger.SetAsync(enabled)
}

// SetStdoutPrint sets whether ouptput the logging contents to stdout, which is true in default.
func E设置堆栈输出(enabled bool) {
	logger.SetStdoutPrint(enabled)
}

// SetHeaderPrint sets whether output header of the logging contents, which is true in default.
func E设置头输出(enabled bool) {
	logger.SetHeaderPrint(enabled)
}

// SetPrefix sets prefix string for every logging content.
// Prefix is part of header, which means if header output is shut, no prefix will be output.
func E设置前缀(prefix string) {
	logger.SetPrefix(prefix)
}

// SetFlags sets extra flags for logging output features.
func E设置标志(flags int) {
	logger.SetFlags(flags)
}

// GetFlags returns the flags of logger.
func E获取标志() int {
	return logger.GetFlags()
}

// SetCtxKeys设置logger的上下文键。键用于检索值
//从上下文和打印到日志内容。
//
//注意，多次调用这个函数将覆盖之前的set上下文键。
func E设置上下文键(keys ...interface{}) {
	logger.SetCtxKeys(keys...)
}

// GetCtxKeys retrieves and returns the context keys for logging.
func E获取上下文键() []interface{} {
	return logger.GetCtxKeys()
}

// PrintStack prints the caller stack,
// the optional parameter <skip> specify the skipped stack offset from the end point.
func E输出Stack(skip ...int) {
	logger.PrintStack(skip...)
}

// GetStack returns the caller stack content,
// the optional parameter <skip> specify the skipped stack offset from the end point.
func E获取Stack(skip ...int) string {
	return logger.GetStack(skip...)
}

// SetStack enables/disables the stack feature in failure logging outputs.
func E设置堆栈(enabled bool) {
	logger.SetStack(enabled)
}

// SetLevelStr sets the logging level by level string.
func E设置级别文本(levelStr string) error {
	return logger.SetLevelStr(levelStr)
}

// SetLevelPrefix sets the prefix string for specified level.
func E设置级别前缀(level int, prefix string) {
	logger.SetLevelPrefix(level, prefix)
}

// SetLevelPrefixes sets the level to prefix string mapping for the logger.
func E设置级别前缀map(prefixes map[int]string) {
	logger.SetLevelPrefixes(prefixes)
}

// GetLevelPrefix returns the prefix string for specified level.
func E获取级别前缀(level int) string {
	return logger.GetLevelPrefix(level)
}
