// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gcache

import (
	"context"
	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/os/gtimer"
	"github.com/gogf/gf/util/gconv"
	"time"
)

// New creates and returns a new cache object using default memory adapter.
// Note that the LRU feature is only available using memory adapter.
func E创建(lruCap ...int) *Cache {
	memAdapter := newAdapterMemory(lruCap...)
	c := &Cache{
		adapter: memAdapter,
	}
	// Here may be a "timer leak" if adapter is manually changed from memory adapter.
	// Do not worry about this, as adapter is less changed and it dose nothing if it's not used.
	gtimer.AddSingleton(time.Second, memAdapter.syncEventAndClearExpired)
	return c
}

// Clone returns a shallow copy of current object.
func (c *Cache) E克隆() *Cache {
	return &Cache{
		adapter: c.adapter,
		ctx:     c.ctx,
	}
}

// Ctx is a chaining function, which shallowly clones current object and sets the context
// for next operation.
func (c *Cache) E上下文(ctx context.Context) *Cache {
	newCache := c.Clone()
	newCache.ctx = ctx
	return newCache
}

// SetAdapter changes the adapter for this cache.
// Be very note that, this setting function is not concurrent-safe, which means you should not call
// this setting function concurrently in multiple goroutines.
func (c *Cache) E设置适配器(adapter Adapter) {
	c.adapter = adapter
}

// GetVar retrieves and returns the value of <键> as gvar.Var.
func (c *Cache) E获取变量(键 interface{}) (*gvar.Var, error) {
	v, err := c.Get(键)
	return gvar.New(v), err
}

// 在缓存中删除删除<键>。
// 已弃用，请使用Remove代替。
//func (c *Cache) E删除(键 []interface{}) error {
//	_, err := c.Remove(键...)
//	return err
//}

// KeyStrings returns all 键s in the cache as string slice.
func (c *Cache) E获取所有键文本() ([]string, error) {
	键, err := c.Keys()
	if err != nil {
		return nil, err
	}
	return gconv.Strings(键), nil
}

// KeyStrings returns all 键 in the cache as string slice.
func (c *Cache) E获取上下文() context.Context {
	if c.ctx == nil {
		return context.Background()
	}
	return c.ctx
}
