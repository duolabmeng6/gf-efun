// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package errors provides simple functions to manipulate errors.
//
// Very note that, this package is quite a base package, which should not import extra
// packages except standard packages, to avoid cycle imports.
package gerror

import (
	"fmt"
)



// New creates and returns an error which is formatted from given text.
func E新建(text string) error {
	return &Error{
		stack: callers(),
		text:  text,
		code:  -1,
	}
}

// Newf returns an error that formats as the given format and args.
func E新建格式化(format string, args ...interface{}) error {
	return &Error{
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  -1,
	}
}

// NewSkip创建并返回一个根据给定文本格式化的错误。
////参数<skip>指定了堆栈调用方的跳过数量。
func E新建跳过(skip int, text string) error {
	return &Error{
		stack: callers(skip),
		text:  text,
		code:  -1,
	}
}

// NewSkipf返回一个错误，格式化为给定的格式和参数。
//参数<skip>指定了堆栈调用方的跳过数量。
func E新建跳过格式化(skip int, format string, args ...interface{}) error {
	return &Error{
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  -1,
	}
}

// Wrap Wrap error with text。
//如果给定err为nil，则返回nil。
func E包(err error, text string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  text,
		code:  Code(err),
	}
}

// Wrapf returns an error annotating err with a stack trace
// at the point Wrapf is called, and the format specifier.
// It returns nil if given <err> is nil.
func E包格式化(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  Code(err),
	}
}

// WrapSkip wraps error with text.
// It returns nil if given err is nil.
// The parameter <skip> specifies the stack callers skipped amount.
func E包跳过(skip int, err error, text string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  text,
		code:  Code(err),
	}
}

// WrapSkipf wraps error with text that is formatted with given format and args.
// It returns nil if given err is nil.
// The parameter <skip> specifies the stack callers skipped amount.
func E包跳过格式化(skip int, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  Code(err),
	}
}

// NewCode creates and returns an error that has error code and given text.
func E新建代码(code int, text string) error {
	return &Error{
		stack: callers(),
		text:  text,
		code:  code,
	}
}

// NewCodef returns an error that has error code and formats as the given format and args.
func E新建代码格式化(code int, format string, args ...interface{}) error {
	return &Error{
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// NewCodeSkip creates and returns an error which has error code and is formatted from given text.
// The parameter <skip> specifies the stack callers skipped amount.
func E新建代码跳过(code, skip int, text string) error {
	return &Error{
		stack: callers(skip),
		text:  text,
		code:  code,
	}
}

// NewCodeSkipf returns an error that has error code and formats as the given format and args.
// The parameter <skip> specifies the stack callers skipped amount.
func E新建代码跳过格式化(code, skip int, format string, args ...interface{}) error {
	return &Error{
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// WrapCode wraps error with code and text.
// It returns nil if given err is nil.
func E包代码(code int, err error, text string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  text,
		code:  code,
	}
}

// WrapCodef wraps error with code and format specifier.
// It returns nil if given <err> is nil.
func E包代码格式化(code int, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// WrapCodeSkip wraps error with code and text.
// It returns nil if given err is nil.
// The parameter <skip> specifies the stack callers skipped amount.
func E包代码跳过(code, skip int, err error, text string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  text,
		code:  code,
	}
}

// WrapCodeSkipf用给定格式和参数格式化的代码和文本包装错误。
//如果给定err为nil，则返回nil。
//参数<skip>指定了堆栈调用方的跳过数量。
func E包代码跳过格式化(code, skip int, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

//如果没有错误代码或者没有实现接口代码，则返回-1。
func E获取代码(err error) int {
	if err != nil {
		if e, ok := err.(apiCode); ok {
			return e.Code()
		}
	}
	return -1
}

// Cause返回<err>的根本原因错误。
func E获取根错误(err error) error {
	if err != nil {
		if e, ok := err.(apiCause); ok {
			return e.Cause()
		}
	}
	return err
}

//堆栈以字符串形式返回堆栈调用者。
//如果<err>不支持栈，则直接返回错误字符串。
func E获取堆栈错误(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(apiStack); ok {
		return e.Stack()
	}
	return err.Error()
}

// Current创建并返回当前级别错误。
//如果当前级别错误为空，则返回空值。
func E获取当前错误(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(apiCurrent); ok {
		return e.Current()
	}
	return err
}

// Next返回下一个级别错误。
//如果当前级别错误或下一个级别错误为nil，则返回nil。
func E获取下一个错误(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(apiNext); ok {
		return e.Next()
	}
	return nil
}
