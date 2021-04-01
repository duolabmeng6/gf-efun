// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gcache

import (
	"time"
)


// Set用<键>-<值> pair设置缓存，在<时间>后过期。
//
// 如果<时间> == 0，则它不会过期。
// 如果<时间> < 0，则删除<键>。
func (c *Cache) E设置(键 interface{}, 值 interface{}, 时间 time.Duration) error {
	return c.adapter.Set(c.getCtx(), 键, 值, 时间)
}

// 通过<data>设置批量设置缓存的键值对，该键值对在<时间>之后过期。
//
// It does not expire if <时间> == 0.
// It deletes the 键s of <data> if <时间> < 0 or given <值> is nil.
func (c *Cache) E设置集合(data map[interface{}]interface{}, 时间 time.Duration) error {
	return c.adapter.Sets(c.getCtx(), data, 时间)
}

// SetIfNotExist sets cache with <键>-<值> pair which is expired after <时间>
// if <键> does not exist in the cache. It returns true the <键> dose not exist in the
// cache and it sets <值> successfully to the cache, or else it returns false.
//
// The parameter <值> can be type of <func() interface{}>, but it dose nothing if its
// result is nil.
//
// It does not expire if <时间> == 0.
// It deletes the <键> if <时间> < 0 or given <值> is nil.
func (c *Cache) E设置不存在时写入(键 interface{}, 值 interface{}, 时间 time.Duration) (bool, error) {
	return c.adapter.SetIfNotExist(c.getCtx(), 键, 值, 时间)
}

// Get retrieves and returns the associated 值 of given <键>.
// It returns nil if it does not exist, its 值 is nil or it's expired.
func (c *Cache) E获取(键 interface{}) (interface{}, error) {
	return c.adapter.Get(c.getCtx(), 键)
}

// GetOrSet retrieves and returns the 值 of <键>, or sets <键>-<值> pair and
// returns <值> if <键> does not exist in the cache. The 键-值 pair expires
// after <时间>.
//
// It does not expire if <时间> == 0.
// It deletes the <键> if <时间> < 0 or given <值> is nil, but it does nothing
// if <值> is a function and the function result is nil.
func (c *Cache) E获取或设置(键 interface{}, 值 interface{}, 时间 time.Duration) (interface{}, error) {
	return c.adapter.GetOrSet(c.getCtx(), 键, 值, 时间)
}

// GetOrSetFunc retrieves and returns the 值 of <键>, or sets <键> with result of
// function <f> and returns its result if <键> does not exist in the cache. The 键-值
// pair expires after <时间>.
//
// It does not expire if <时间> == 0.
// It deletes the <键> if <时间> < 0 or given <值> is nil, but it does nothing
// if <值> is a function and the function result is nil.
func (c *Cache) E获取或设置函数(键 interface{}, f func() (interface{}, error), 时间 time.Duration) (interface{}, error) {
	return c.adapter.GetOrSetFunc(c.getCtx(), 键, f, 时间)
}

// GetOrSetFuncLock retrieves and returns the 值 of <键>, or sets <键> with result of
// function <f> and returns its result if <键> does not exist in the cache. The 键-值
// pair expires after <时间>.
//
// It does not expire if <时间> == 0.
// It does nothing if function <f> returns nil.
//
// Note that the function <f> should be executed within writing mutex lock for concurrent
// safety purpose.
func (c *Cache) E获取或设置函数锁(键 interface{}, f func() (interface{}, error), 时间 time.Duration) (interface{}, error) {
	return c.adapter.GetOrSetFuncLock(c.getCtx(), 键, f, 时间)
}

// Contains returns true if <键> exists in the cache, or else returns false.
func (c *Cache) E是否存在(键 interface{}) (bool, error) {
	return c.adapter.Contains(c.getCtx(), 键)
}

// GetExpire retrieves and returns the expiration of <键> in the cache.
//
// It returns 0 if the <键> does not expire.
// It returns -1 if the <键> does not exist in the cache.
func (c *Cache) E获取超时时间(键 interface{}) (time.Duration, error) {
	return c.adapter.GetExpire(c.getCtx(), 键)
}

// Remove从缓存中删除一个或多个键，并返回其值。
// 如果给出了多个键，则返回最后删除的项的值。
func (c *Cache) E删除(键 ...interface{}) (值 interface{}, err error) {
	return c.adapter.Remove(c.getCtx(), 键...)
}

// Update更新<键>的值而不改变它的过期时间，并返回旧的值。
// 如果<键>在缓存中不存在，则返回值<exist>为false。
//
// 如果给定<值>为nil，则删除<键>。
// 如果<键>不存在于缓存中，则不执行任何操作。
func (c *Cache) E更新(键 interface{}, 值 interface{}) (oldValue interface{}, exist bool, err error) {
	return c.adapter.Update(c.getCtx(), 键, 值)
}

// UpdateExpire更新<键>的过期时间，并返回旧的过期时间值。
//
// 如果<键>在缓存中不存在，则返回-1，不执行任何操作。
// 如果<时间> < 0，则删除<键>。
func (c *Cache) E更新过期时间(键 interface{}, 时间 time.Duration) (oldDuration time.Duration, err error) {
	return c.adapter.UpdateExpire(c.getCtx(), 键, 时间)
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