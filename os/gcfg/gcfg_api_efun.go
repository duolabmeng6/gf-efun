// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gcfg

import (
	"errors"
	"time"

	"github.com/gogf/gf/encoding/gjson"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/os/gtime"
)

// Set sets value with specified `pattern`.
// It supports hierarchical data access by char separator, which is '.' in default.
// It is commonly used for updates certain configuration value in runtime.
func (c *Config) E设置(pattern string, value interface{}) error {
	if j := c.getJson(); j != nil {
		return j.Set(pattern, value)
	}
	return nil
}

// Get retrieves and returns value by specified `pattern`.
// It returns all values of current Json object if `pattern` is given empty or string ".".
// It returns nil if no value found by `pattern`.
//
// We can also access slice item by its index number in `pattern` like:
// "list.10", "array.0.name", "array.0.1.id".
//
// It returns a default value specified by `def` if value for `pattern` is not found.
func (c *Config) E获取(pattern string, def ...interface{}) interface{} {
	if j := c.getJson(); j != nil {
		return j.Get(pattern, def...)
	}
	return nil
}

// GetVar returns a gvar.Var with value by given `pattern`.
func (c *Config) E获取变量(pattern string, def ...interface{}) *gvar.Var {
	if j := c.getJson(); j != nil {
		return j.GetVar(pattern, def...)
	}
	return gvar.New(nil)
}

// Contains checks whether the value by specified `pattern` exist.
func (c *Config) E是否存在(pattern string) bool {
	if j := c.getJson(); j != nil {
		return j.Contains(pattern)
	}
	return false
}

// GetMap retrieves and returns the value by specified `pattern` as map[string]interface{}.
func (c *Config) E获取集合(pattern string, def ...interface{}) map[string]interface{} {
	if j := c.getJson(); j != nil {
		return j.GetMap(pattern, def...)
	}
	return nil
}

// GetMapStrStr retrieves and returns the value by specified `pattern` as map[string]string.
func (c *Config) E获取集合文本(pattern string, def ...interface{}) map[string]string {
	if j := c.getJson(); j != nil {
		return j.GetMapStrStr(pattern, def...)
	}
	return nil
}

// GetArray retrieves the value by specified `pattern`,
// and converts it to a slice of []interface{}.
func (c *Config) E获取数组(pattern string, def ...interface{}) []interface{} {
	if j := c.getJson(); j != nil {
		return j.GetArray(pattern, def...)
	}
	return nil
}

// GetBytes retrieves the value by specified `pattern` and converts it to []byte.
func (c *Config) E获取字节集(pattern string, def ...interface{}) []byte {
	if j := c.getJson(); j != nil {
		return j.GetBytes(pattern, def...)
	}
	return nil
}

// GetString retrieves the value by specified `pattern` and converts it to string.
func (c *Config) E获取文本(pattern string, def ...interface{}) string {
	if j := c.getJson(); j != nil {
		return j.GetString(pattern, def...)
	}
	return ""
}

// GetStrings retrieves the value by specified `pattern` and converts it to []string.
func (c *Config) E获取文本数组(pattern string, def ...interface{}) []string {
	if j := c.getJson(); j != nil {
		return j.GetStrings(pattern, def...)
	}
	return nil
}

// GetInterfaces is alias of GetArray.
// See GetArray.
//func (c *Config) EGetInterfaces(pattern string, def ...interface{}) []interface{} {
//	if j := c.getJson(); j != nil {
//		return j.GetInterfaces(pattern, def...)
//	}
//	return nil
//}

// GetBool retrieves the value by specified `pattern`,
// converts and returns it as bool.
// It returns false when value is: "", 0, false, off, nil;
// or returns true instead.
func (c *Config) E获取逻辑值(pattern string, def ...interface{}) bool {
	if j := c.getJson(); j != nil {
		return j.GetBool(pattern, def...)
	}
	return false
}

// GetFloat32 retrieves the value by specified `pattern` and converts it to float32.
func (c *Config) E获取浮点数32(pattern string, def ...interface{}) float32 {
	if j := c.getJson(); j != nil {
		return j.GetFloat32(pattern, def...)
	}
	return 0
}

// GetFloat64 retrieves the value by specified `pattern` and converts it to float64.
func (c *Config) E获取浮点数64(pattern string, def ...interface{}) float64 {
	if j := c.getJson(); j != nil {
		return j.GetFloat64(pattern, def...)
	}
	return 0
}

// GetFloats retrieves the value by specified `pattern` and converts it to []float64.
func (c *Config) E获取浮点数64数组(pattern string, def ...interface{}) []float64 {
	if j := c.getJson(); j != nil {
		return j.GetFloats(pattern, def...)
	}
	return nil
}

// GetInt retrieves the value by specified `pattern` and converts it to int.
func (c *Config) E获取整数型(pattern string, def ...interface{}) int {
	if j := c.getJson(); j != nil {
		return j.GetInt(pattern, def...)
	}
	return 0
}

// GetInt8 retrieves the value by specified `pattern` and converts it to int8.
func (c *Config) E获取整数型8(pattern string, def ...interface{}) int8 {
	if j := c.getJson(); j != nil {
		return j.GetInt8(pattern, def...)
	}
	return 0
}

// GetInt16 retrieves the value by specified `pattern` and converts it to int16.
func (c *Config) E获取整数型16(pattern string, def ...interface{}) int16 {
	if j := c.getJson(); j != nil {
		return j.GetInt16(pattern, def...)
	}
	return 0
}

// GetInt32 retrieves the value by specified `pattern` and converts it to int32.
func (c *Config) E获取整数型32(pattern string, def ...interface{}) int32 {
	if j := c.getJson(); j != nil {
		return j.GetInt32(pattern, def...)
	}
	return 0
}

// GetInt64 retrieves the value by specified `pattern` and converts it to int64.
func (c *Config) E获取整数型64(pattern string, def ...interface{}) int64 {
	if j := c.getJson(); j != nil {
		return j.GetInt64(pattern, def...)
	}
	return 0
}

// GetInts retrieves the value by specified `pattern` and converts it to []int.
func (c *Config) E获取整数型数组(pattern string, def ...interface{}) []int {
	if j := c.getJson(); j != nil {
		return j.GetInts(pattern, def...)
	}
	return nil
}

// GetUint retrieves the value by specified `pattern` and converts it to uint.
func (c *Config) E获取正整数(pattern string, def ...interface{}) uint {
	if j := c.getJson(); j != nil {
		return j.GetUint(pattern, def...)
	}
	return 0
}

// GetUint8 retrieves the value by specified `pattern` and converts it to uint8.
func (c *Config) E获取正整数8(pattern string, def ...interface{}) uint8 {
	if j := c.getJson(); j != nil {
		return j.GetUint8(pattern, def...)
	}
	return 0
}

// GetUint16 retrieves the value by specified `pattern` and converts it to uint16.
func (c *Config) E获取正整数16(pattern string, def ...interface{}) uint16 {
	if j := c.getJson(); j != nil {
		return j.GetUint16(pattern, def...)
	}
	return 0
}

// GetUint32 retrieves the value by specified `pattern` and converts it to uint32.
func (c *Config) E获取正整数32(pattern string, def ...interface{}) uint32 {
	if j := c.getJson(); j != nil {
		return j.GetUint32(pattern, def...)
	}
	return 0
}

// GetUint64 retrieves the value by specified `pattern` and converts it to uint64.
func (c *Config) E获取正整数64(pattern string, def ...interface{}) uint64 {
	if j := c.getJson(); j != nil {
		return j.GetUint64(pattern, def...)
	}
	return 0
}

// GetTime retrieves the value by specified `pattern` and converts it to time.Time.
func (c *Config) E获取时间(pattern string, format ...string) time.Time {
	if j := c.getJson(); j != nil {
		return j.GetTime(pattern, format...)
	}
	return time.Time{}
}

// GetDuration retrieves the value by specified `pattern` and converts it to time.Duration.
func (c *Config) E获取秒数(pattern string, def ...interface{}) time.Duration {
	if j := c.getJson(); j != nil {
		return j.GetDuration(pattern, def...)
	}
	return 0
}

// GetGTime retrieves the value by specified `pattern` and converts it to *gtime.Time.
func (c *Config) E获取GTime(pattern string, format ...string) *gtime.Time {
	if j := c.getJson(); j != nil {
		return j.GetGTime(pattern, format...)
	}
	return nil
}

// GetJson gets the value by specified `pattern`,
// and converts it to a un-concurrent-safe Json object.
func (c *Config) E获取Json(pattern string, def ...interface{}) *gjson.Json {
	if j := c.getJson(); j != nil {
		return j.GetJson(pattern, def...)
	}
	return nil
}

// GetJsons gets the value by specified `pattern`,
// and converts it to a slice of un-concurrent-safe Json object.
func (c *Config) E获取Jsons(pattern string, def ...interface{}) []*gjson.Json {
	if j := c.getJson(); j != nil {
		return j.GetJsons(pattern, def...)
	}
	return nil
}

// GetJsonMap gets the value by specified `pattern`,
// and converts it to a map of un-concurrent-safe Json object.
func (c *Config) E获取JsonMap(pattern string, def ...interface{}) map[string]*gjson.Json {
	if j := c.getJson(); j != nil {
		return j.GetJsonMap(pattern, def...)
	}
	return nil
}

// GetStruct retrieves the value by specified `pattern` and converts it to specified object
// `pointer`. The `pointer` should be the pointer to an object.
func (c *Config) E获取Struct(pattern string, pointer interface{}, mapping ...map[string]string) error {
	if j := c.getJson(); j != nil {
		return j.GetStruct(pattern, pointer, mapping...)
	}
	return errors.New("configuration not found")
}

// GetStructDeep does GetStruct recursively.
// Deprecated, use GetStruct instead.
func (c *Config) E获取StructDeep(pattern string, pointer interface{}, mapping ...map[string]string) error {
	if j := c.getJson(); j != nil {
		return j.GetStructDeep(pattern, pointer, mapping...)
	}
	return errors.New("configuration not found")
}

// GetStructs converts any slice to given struct slice.
func (c *Config) E获取Structs(pattern string, pointer interface{}, mapping ...map[string]string) error {
	if j := c.getJson(); j != nil {
		return j.GetStructs(pattern, pointer, mapping...)
	}
	return errors.New("configuration not found")
}

// GetStructsDeep converts any slice to given struct slice recursively.
// Deprecated, use GetStructs instead.
func (c *Config) E获取StructsDeep(pattern string, pointer interface{}, mapping ...map[string]string) error {
	if j := c.getJson(); j != nil {
		return j.GetStructsDeep(pattern, pointer, mapping...)
	}
	return errors.New("configuration not found")
}

// GetMapToMap retrieves the value by specified `pattern` and converts it to specified map variable.
// See gconv.MapToMap.
func (c *Config) E获取MapToMap(pattern string, pointer interface{}, mapping ...map[string]string) error {
	if j := c.getJson(); j != nil {
		return j.GetMapToMap(pattern, pointer, mapping...)
	}
	return errors.New("configuration not found")
}

// GetMapToMaps retrieves the value by specified `pattern` and converts it to specified map slice
// variable.
// See gconv.MapToMaps.
func (c *Config) E获取MapToMaps(pattern string, pointer interface{}, mapping ...map[string]string) error {
	if j := c.getJson(); j != nil {
		return j.GetMapToMaps(pattern, pointer, mapping...)
	}
	return errors.New("configuration not found")
}

// GetMapToMapsDeep retrieves the value by specified `pattern` and converts it to specified map slice
// variable recursively.
// See gconv.MapToMapsDeep.
func (c *Config) E获取MapToMapsDeep(pattern string, pointer interface{}, mapping ...map[string]string) error {
	if j := c.getJson(); j != nil {
		return j.GetMapToMapsDeep(pattern, pointer, mapping...)
	}
	return errors.New("configuration not found")
}

// ToMap converts current Json object to map[string]interface{}.
// It returns nil if fails.
func (c *Config) E到Map() map[string]interface{} {
	if j := c.getJson(); j != nil {
		return j.ToMap()
	}
	return nil
}

// ToArray converts current Json object to []interface{}.
// It returns nil if fails.
func (c *Config) E到Array() []interface{} {
	if j := c.getJson(); j != nil {
		return j.ToArray()
	}
	return nil
}

// ToStruct converts current Json object to specified object.
// The `pointer` should be a pointer type of *struct.
func (c *Config) E到Struct(pointer interface{}, mapping ...map[string]string) error {
	if j := c.getJson(); j != nil {
		return j.ToStruct(pointer, mapping...)
	}
	return errors.New("configuration not found")
}

// ToStructDeep converts current Json object to specified object recursively.
// The `pointer` should be a pointer type of *struct.
func (c *Config) E到StructDeep(pointer interface{}, mapping ...map[string]string) error {
	if j := c.getJson(); j != nil {
		return j.ToStructDeep(pointer, mapping...)
	}
	return errors.New("configuration not found")
}

// ToStructs converts current Json object to specified object slice.
// The `pointer` should be a pointer type of []struct/*struct.
func (c *Config) E到Structs(pointer interface{}, mapping ...map[string]string) error {
	if j := c.getJson(); j != nil {
		return j.ToStructs(pointer, mapping...)
	}
	return errors.New("configuration not found")
}

// ToStructsDeep converts current Json object to specified object slice recursively.
// The `pointer` should be a pointer type of []struct/*struct.
func (c *Config) E到StructsDeep(pointer interface{}, mapping ...map[string]string) error {
	if j := c.getJson(); j != nil {
		return j.ToStructsDeep(pointer, mapping...)
	}
	return errors.New("configuration not found")
}

// ToMapToMap converts current Json object to specified map variable.
// The parameter of `pointer` should be type of *map.
func (c *Config) E到MapToMap(pointer interface{}, mapping ...map[string]string) error {
	if j := c.getJson(); j != nil {
		return j.ToMapToMap(pointer, mapping...)
	}
	return errors.New("configuration not found")
}

// ToMapToMaps converts current Json object to specified map variable slice.
// The parameter of `pointer` should be type of []map/*map.
func (c *Config) E到MapToMaps(pointer interface{}, mapping ...map[string]string) error {
	if j := c.getJson(); j != nil {
		return j.ToMapToMaps(pointer, mapping...)
	}
	return errors.New("configuration not found")
}

// ToMapToMapsDeep converts current Json object to specified map variable slice recursively.
// The parameter of `pointer` should be type of []map/*map.
func (c *Config) E到MapToMapsDeep(pointer interface{}, mapping ...map[string]string) error {
	if j := c.getJson(); j != nil {
		return j.ToMapToMapsDeep(pointer, mapping...)
	}
	return errors.New("configuration not found")
}

// Clear removes all parsed configuration files content cache,
// which will force reload configuration content from file.
func (c *Config) E清除() {
	c.jsonMap.Clear()
}

// Dump prints current Json object with more manually readable.
func (c *Config) E调试() {
	if j := c.getJson(); j != nil {
		j.Dump()
	}
}
