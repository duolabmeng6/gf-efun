// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gcache

import (
	"time"
)


// Set用<key>-<value> pair设置缓存，在<duration>后过期。
//
// 如果<duration> == 0，则它不会过期。
// 如果<duration> < 0，则删除<键>。
func (c *Cache) E设置(key interface{}, value interface{}, duration time.Duration) error {
	return c.adapter.Set(c.getCtx(), key, value, duration)
}

// 通过<data>设置批量设置缓存的键值对，该键值对在<duration>之后过期。
//
// It does not expire if <duration> == 0.
// It deletes the keys of <data> if <duration> < 0 or given <value> is nil.
func (c *Cache) E设置集合(data map[interface{}]interface{}, duration time.Duration) error {
	return c.adapter.Sets(c.getCtx(), data, duration)
}

// SetIfNotExist sets cache with <key>-<value> pair which is expired after <duration>
// if <key> does not exist in the cache. It returns true the <key> dose not exist in the
// cache and it sets <value> successfully to the cache, or else it returns false.
//
// The parameter <value> can be type of <func() interface{}>, but it dose nothing if its
// result is nil.
//
// It does not expire if <duration> == 0.
// It deletes the <key> if <duration> < 0 or given <value> is nil.
func (c *Cache) E设置不存在时写入(key interface{}, value interface{}, duration time.Duration) (bool, error) {
	return c.adapter.SetIfNotExist(c.getCtx(), key, value, duration)
}

// Get retrieves and returns the associated value of given <key>.
// It returns nil if it does not exist, its value is nil or it's expired.
func (c *Cache) E获取(key interface{}) (interface{}, error) {
	return c.adapter.Get(c.getCtx(), key)
}

// GetOrSet retrieves and returns the value of <key>, or sets <key>-<value> pair and
// returns <value> if <key> does not exist in the cache. The key-value pair expires
// after <duration>.
//
// It does not expire if <duration> == 0.
// It deletes the <key> if <duration> < 0 or given <value> is nil, but it does nothing
// if <value> is a function and the function result is nil.
func (c *Cache) E获取或设置(key interface{}, value interface{}, duration time.Duration) (interface{}, error) {
	return c.adapter.GetOrSet(c.getCtx(), key, value, duration)
}

// GetOrSetFunc retrieves and returns the value of <key>, or sets <key> with result of
// function <f> and returns its result if <key> does not exist in the cache. The key-value
// pair expires after <duration>.
//
// It does not expire if <duration> == 0.
// It deletes the <key> if <duration> < 0 or given <value> is nil, but it does nothing
// if <value> is a function and the function result is nil.
func (c *Cache) E获取或设置函数(key interface{}, f func() (interface{}, error), duration time.Duration) (interface{}, error) {
	return c.adapter.GetOrSetFunc(c.getCtx(), key, f, duration)
}

// GetOrSetFuncLock retrieves and returns the value of <key>, or sets <key> with result of
// function <f> and returns its result if <key> does not exist in the cache. The key-value
// pair expires after <duration>.
//
// It does not expire if <duration> == 0.
// It does nothing if function <f> returns nil.
//
// Note that the function <f> should be executed within writing mutex lock for concurrent
// safety purpose.
func (c *Cache) E获取或设置函数锁(key interface{}, f func() (interface{}, error), duration time.Duration) (interface{}, error) {
	return c.adapter.GetOrSetFuncLock(c.getCtx(), key, f, duration)
}

// Contains returns true if <key> exists in the cache, or else returns false.
func (c *Cache) E是否存在(key interface{}) (bool, error) {
	return c.adapter.Contains(c.getCtx(), key)
}

// GetExpire retrieves and returns the expiration of <key> in the cache.
//
// It returns 0 if the <key> does not expire.
// It returns -1 if the <key> does not exist in the cache.
func (c *Cache) E获取超时时间(key interface{}) (time.Duration, error) {
	return c.adapter.GetExpire(c.getCtx(), key)
}

// Remove从缓存中删除一个或多个键，并返回其值。
// 如果给出了多个键，则返回最后删除的项的值。
func (c *Cache) E移除(keys ...interface{}) (value interface{}, err error) {
	return c.adapter.Remove(c.getCtx(), keys...)
}

// Update更新<key>的值而不改变它的过期时间，并返回旧的值。
// 如果<key>在缓存中不存在，则返回值<exist>为false。
//
// 如果给定<value>为nil，则删除<key>。
// 如果<key>不存在于缓存中，则不执行任何操作。
func (c *Cache) E更新(key interface{}, value interface{}) (oldValue interface{}, exist bool, err error) {
	return c.adapter.Update(c.getCtx(), key, value)
}

// UpdateExpire更新<key>的过期时间，并返回旧的过期时间值。
//
// 如果<key>在缓存中不存在，则返回-1，不执行任何操作。
// 如果<duration> < 0，则删除<键>。
func (c *Cache) E更新过期时间(key interface{}, duration time.Duration) (oldDuration time.Duration, err error) {
	return c.adapter.UpdateExpire(c.getCtx(), key, duration)
}

// Size返回缓存中的项数。
func (c *Cache) E大小() (size int, err error) {
	return c.adapter.Size(c.getCtx())
}

// Data返回缓存中所有键值对的副本作为map类型。
// 注意这个函数可能会占用大量的内存，你可以实现这个函数
// 如果必要的话。
func (c *Cache) E数据() (map[interface{}]interface{}, error) {
	return c.adapter.Data(c.getCtx())
}

// Keys以slice的形式返回缓存中的所有键。
func (c *Cache) E获取所有键() ([]interface{}, error) {
	return c.adapter.Keys(c.getCtx())
}

// Values以slice的形式返回缓存中的所有值。
func (c *Cache) E获取所有值() ([]interface{}, error) {
	return c.adapter.Values(c.getCtx())
}

// Clear 清除缓存中的所有数据。
// 注意，这个函数是敏感的，应该谨慎使用。
func (c *Cache) E清除() error {
	return c.adapter.Clear(c.getCtx())
}

// Close 在必要时关闭缓存。
func (c *Cache) E关闭() error {
	return c.adapter.Close(c.getCtx())
}