// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package glog

// Print prints <v> with newline using fmt.Sprintln.
// The parameter <v> can be multiple variables.
func E输出(v ...interface{}) {
	logger.Print(v...)
}

// Printf prints <v> with format <format> using fmt.Sprintf.
// The parameter <v> can be multiple variables.
func E输出自定义(format string, v ...interface{}) {
	logger.Printf(format, v...)
}

// See Print.
func E输出并换行(v ...interface{}) {
	logger.Println(v...)
}

// Fatal prints the logging content with [FATA] header and newline, then exit the current process.
func E致命(v ...interface{}) {
	logger.Fatal(v...)
}

// Fatalf prints the logging content with [FATA] header, custom format and newline, then exit the current process.
func E致命自定义(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}

// Panic prints the logging content with [PANI] header and newline, then panics.
func E恐慌(v ...interface{}) {
	logger.Panic(v...)
}

// Panicf prints the logging content with [PANI] header, custom format and newline, then panics.
func E恐慌自定义(format string, v ...interface{}) {
	logger.Panicf(format, v...)
}

// Info prints the logging content with [INFO] header and newline.
func E信息(v ...interface{}) {
	logger.Info(v...)
}

// Infof prints the logging content with [INFO] header, custom format and newline.
func E信息自定义(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

// Debug prints the logging content with [DEBU] header and newline.
func E调试(v ...interface{}) {
	logger.Debug(v...)
}

// Debugf prints the logging content with [DEBU] header, custom format and newline.
func E调试自定义(format string, v ...interface{}) {
	logger.Debugf(format, v...)
}

// Notice prints the logging content with [NOTI] header and newline.
// It also prints caller stack info if stack feature is enabled.
func E注意(v ...interface{}) {
	logger.Notice(v...)
}

// Noticef prints the logging content with [NOTI] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
func E注意自定义(format string, v ...interface{}) {
	logger.Noticef(format, v...)
}

// Warning prints the logging content with [WARN] header and newline.
// It also prints caller stack info if stack feature is enabled.
func E警告(v ...interface{}) {
	logger.Warning(v...)
}

// Warningf prints the logging content with [WARN] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
func E警告自定义(format string, v ...interface{}) {
	logger.Warningf(format, v...)
}

// Error prints the logging content with [ERRO] header and newline.
// It also prints caller stack info if stack feature is enabled.
func E错误(v ...interface{}) {
	logger.Error(v...)
}

// Errorf prints the logging content with [ERRO] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
func E错误自定义(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}

// Critical prints the logging content with [CRIT] header and newline.
// It also prints caller stack info if stack feature is enabled.
func E重要(v ...interface{}) {
	logger.Critical(v...)
}

// Criticalf prints the logging content with [CRIT] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
func E重要自定义(format string, v ...interface{}) {
	logger.Criticalf(format, v...)
}
