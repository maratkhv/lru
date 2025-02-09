package lru

import (
	"github.com/maratkhv/lru/list"
)

type Cache[K comparable, V any] interface {
	Put(key K, v V)
	Get(key K) (v V, found bool)
}

type node[K comparable, V any] struct {
	key   K
	value V
}

type cache[K comparable, V any] struct {
	linkedList *list.List[node[K, V]]
	keyToValue map[K]*list.Element[node[K, V]]
	capacity   int64
}

func New[K comparable, V any](cap int64) Cache[K, V] {
	return &cache[K, V]{
		linkedList: list.New[node[K, V]](),
		keyToValue: make(map[K]*list.Element[node[K, V]]),
		capacity:   cap,
	}
}

func (c *cache[K, V]) Put(key K, v V) {
	n := node[K, V]{
		key:   key,
		value: v,
	}

	e := c.linkedList.PushFront(n)
	c.keyToValue[key] = e

	if c.linkedList.Len() > c.capacity {
		e = c.linkedList.Pop(c.linkedList.Back())
		delete(c.keyToValue, e.Value.key)
	}
}

func (c *cache[K, V]) Get(key K) (v V, found bool) {
	e, ok := c.keyToValue[key]
	if !ok {
		var empty V
		return empty, false
	}
	c.linkedList.Pop(e)
	e = c.linkedList.PushFront(e.Value)
	c.keyToValue[key] = e
	return e.Value.value, true
}

// NewAuto is the same as New but doesnt require to state K and V
// It takes key and value as a params but does nothing with them
func NewAuto[K comparable, V any](cap int64, _key K, _value V) Cache[K, V] {
	return New[K, V](cap)
}
