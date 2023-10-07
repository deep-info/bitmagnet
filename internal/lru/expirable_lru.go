// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lru

import (
	"sync"
	"time"
)

// EvictCallback is used to get a callback when a cache entry is evicted
type EvictCallback[K comparable, V any] func(key K, value V)

// Expirable implements a thread-safe Expirable with expirable entries.
type Expirable[K comparable, V any] struct {
	size      int
	evictList *LruList[K, V]
	items     map[K]*Entry[K, V]
	onEvict   EvictCallback[K, V]

	// expirable options
	mu   sync.Mutex
	ttl  time.Duration
	done chan struct{}

	// buckets for expiration
	buckets []bucket[K, V]
	// uint8 because it's number between 0 and numBuckets
	nextCleanupBucket uint8
}

// bucket is a container for holding entries to be expired
type bucket[K comparable, V any] struct {
	entries     map[K]*Entry[K, V]
	newestEntry time.Time
}

// noEvictionTTL - very long ttl to prevent eviction
const noEvictionTTL = time.Hour * 24 * 365 * 10

// because of uint8 usage for nextCleanupBucket, should not exceed 256.
// casting it as uint8 explicitly requires type conversions in multiple places
const numBuckets = 100

// NewExpirable returns a new thread-safe cache with expirable entries.
//
// Size parameter set to 0 makes cache of unlimited size, e.g. turns Expirable mechanism off.
//
// Providing 0 TTL turns expiring off.
//
// Delete expired entries every 1/100th of ttl value. Goroutine which deletes expired entries runs indefinitely.
func NewExpirable[K comparable, V any](size int, onEvict EvictCallback[K, V], ttl time.Duration) *Expirable[K, V] {
	if size < 0 {
		size = 0
	}
	if ttl <= 0 {
		ttl = noEvictionTTL
	}

	res := Expirable[K, V]{
		ttl:       ttl,
		size:      size,
		evictList: NewList[K, V](),
		items:     make(map[K]*Entry[K, V]),
		onEvict:   onEvict,
		done:      make(chan struct{}),
	}

	// initialize the buckets
	res.buckets = make([]bucket[K, V], numBuckets)
	for i := 0; i < numBuckets; i++ {
		res.buckets[i] = bucket[K, V]{entries: make(map[K]*Entry[K, V])}
	}

	// enable deleteExpired() running in separate goroutine for cache with non-zero TTL
	//
	// Important: done channel is never closed, so deleteExpired() goroutine will never exit,
	// it's decided to add functionality to close it in the version later than v2.
	if res.ttl != noEvictionTTL {
		go func(done <-chan struct{}) {
			ticker := time.NewTicker(res.ttl / numBuckets)
			defer ticker.Stop()
			for {
				select {
				case <-done:
					return
				case <-ticker.C:
					res.deleteExpired()
				}
			}
		}(res.done)
	}
	return &res
}

// Purge clears the cache completely.
// onEvict is called for each evicted key.
func (c *Expirable[K, V]) Purge() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.items {
		if c.onEvict != nil {
			c.onEvict(k, v.Value)
		}
		delete(c.items, k)
	}
	for _, b := range c.buckets {
		for _, ent := range b.entries {
			delete(b.entries, ent.Key)
		}
	}
	c.evictList.Init()
}

// Add adds a value to the cache. Returns true if an eviction occurred.
// Returns false if there was no eviction: the item was already in the cache,
// or the size was not exceeded.
func (c *Expirable[K, V]) Add(key K, value V) (evicted bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()

	// Check for existing item
	if ent, ok := c.items[key]; ok {
		c.evictList.MoveToFront(ent)
		c.removeFromBucket(ent) // remove the entry from its current bucket as expiresAt is renewed
		ent.Value = value
		ent.ExpiresAt = now.Add(c.ttl)
		c.addToBucket(ent)
		return false
	}

	// Add new item
	ent := c.evictList.PushFrontExpirable(key, value, now.Add(c.ttl))
	c.items[key] = ent
	c.addToBucket(ent) // adds the entry to the appropriate bucket and sets entry.expireBucket

	evict := c.size > 0 && c.evictList.Length() > c.size
	// Verify size not exceeded
	if evict {
		c.removeOldest()
	}
	return evict
}

// Get looks up a key's value from the cache.
func (c *Expirable[K, V]) Get(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	var ent *Entry[K, V]
	if ent, ok = c.items[key]; ok {
		// Expired item check
		if time.Now().After(ent.ExpiresAt) {
			return value, false
		}
		c.evictList.MoveToFront(ent)
		return ent.Value, true
	}
	return
}

// Contains checks if a key is in the cache, without updating the recent-ness
// or deleting it for being stale.
func (c *Expirable[K, V]) Contains(key K) (ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok = c.items[key]
	return ok
}

// Peek returns the key value (or undefined if not found) without updating
// the "recently used"-ness of the key.
func (c *Expirable[K, V]) Peek(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	var ent *Entry[K, V]
	if ent, ok = c.items[key]; ok {
		// Expired item check
		if time.Now().After(ent.ExpiresAt) {
			return value, false
		}
		return ent.Value, true
	}
	return
}

// Remove removes the provided key from the cache, returning if the
// key was contained.
func (c *Expirable[K, V]) Remove(key K) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if ent, ok := c.items[key]; ok {
		c.removeElement(ent)
		return true
	}
	return false
}

// RemoveOldest removes the oldest item from the cache.
func (c *Expirable[K, V]) RemoveOldest() (key K, value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if ent := c.evictList.Back(); ent != nil {
		c.removeElement(ent)
		return ent.Key, ent.Value, true
	}
	return
}

// GetOldest returns the oldest entry
func (c *Expirable[K, V]) GetOldest() (key K, value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if ent := c.evictList.Back(); ent != nil {
		return ent.Key, ent.Value, true
	}
	return
}

// Keys returns a slice of the keys in the cache, from oldest to newest.
func (c *Expirable[K, V]) Keys() []K {
	c.mu.Lock()
	defer c.mu.Unlock()
	keys := make([]K, 0, len(c.items))
	for ent := c.evictList.Back(); ent != nil; ent = ent.PrevEntry() {
		keys = append(keys, ent.Key)
	}
	return keys
}

// Values returns a slice of the values in the cache, from oldest to newest.
// Expired entries are filtered out.
func (c *Expirable[K, V]) Values() []V {
	c.mu.Lock()
	defer c.mu.Unlock()
	values := make([]V, len(c.items))
	i := 0
	now := time.Now()
	for ent := c.evictList.Back(); ent != nil; ent = ent.PrevEntry() {
		if now.After(ent.ExpiresAt) {
			continue
		}
		values[i] = ent.Value
		i++
	}
	return values
}

// Len returns the number of items in the cache.
func (c *Expirable[K, V]) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.evictList.Length()
}

// Resize changes the cache size. Size of 0 means unlimited.
func (c *Expirable[K, V]) Resize(size int) (evicted int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if size <= 0 {
		c.size = 0
		return 0
	}
	diff := c.evictList.Length() - size
	if diff < 0 {
		diff = 0
	}
	for i := 0; i < diff; i++ {
		c.removeOldest()
	}
	c.size = size
	return diff
}

// Close destroys cleanup goroutine. To clean up the cache, run Purge() before Close().
// func (c *Expirable[K, V]) Close() {
//	c.mu.Lock()
//	defer c.mu.Unlock()
//	select {
//	case <-c.done:
//		return
//	default:
//	}
//	close(c.done)
// }

// removeOldest removes the oldest item from the cache. Has to be called with lock!
func (c *Expirable[K, V]) removeOldest() {
	if ent := c.evictList.Back(); ent != nil {
		c.removeElement(ent)
	}
}

// removeElement is used to remove a given list element from the cache. Has to be called with lock!
func (c *Expirable[K, V]) removeElement(e *Entry[K, V]) {
	c.evictList.Remove(e)
	delete(c.items, e.Key)
	c.removeFromBucket(e)
	if c.onEvict != nil {
		c.onEvict(e.Key, e.Value)
	}
}

// deleteExpired deletes expired records from the oldest bucket, waiting for the newest entry
// in it to expire first.
func (c *Expirable[K, V]) deleteExpired() {
	c.mu.Lock()
	bucketIdx := c.nextCleanupBucket
	timeToExpire := time.Until(c.buckets[bucketIdx].newestEntry)
	// wait for newest entry to expire before cleanup without holding lock
	if timeToExpire > 0 {
		c.mu.Unlock()
		time.Sleep(timeToExpire)
		c.mu.Lock()
	}
	for _, ent := range c.buckets[bucketIdx].entries {
		c.removeElement(ent)
	}
	c.nextCleanupBucket = (c.nextCleanupBucket + 1) % numBuckets
	c.mu.Unlock()
}

// addToBucket adds entry to expire bucket so that it will be cleaned up when the time comes. Has to be called with lock!
func (c *Expirable[K, V]) addToBucket(e *Entry[K, V]) {
	bucketID := (numBuckets + c.nextCleanupBucket - 1) % numBuckets
	e.ExpireBucket = bucketID
	c.buckets[bucketID].entries[e.Key] = e
	if c.buckets[bucketID].newestEntry.Before(e.ExpiresAt) {
		c.buckets[bucketID].newestEntry = e.ExpiresAt
	}
}

// removeFromBucket removes the entry from its corresponding bucket. Has to be called with lock!
func (c *Expirable[K, V]) removeFromBucket(e *Entry[K, V]) {
	delete(c.buckets[e.ExpireBucket].entries, e.Key)
}

// Sample returns a random sample of the cache entries not exceeding size n.
// A check can be supplied to exclude any given item from the sample.
func (c *Expirable[K, V]) Sample(n int, check func(K, V) bool) []Entry[K, V] {
	c.mu.Lock()
	defer c.mu.Unlock()
	if n > len(c.items) {
		n = len(c.items)
	}
	sample := make([]Entry[K, V], 0, n)
	now := time.Now()
	for _, ent := range c.items {
		if now.After(ent.ExpiresAt) {
			continue
		}
		if check != nil && !check(ent.Key, ent.Value) {
			continue
		}
		sample = append(sample, *ent)
		if len(sample) == n {
			break
		}
	}
	return sample
}