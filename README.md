# LRU Cache Implementation

This is a simple implementation of an LRU (Least Recently Used) cache built on top of a doubly linked list in Go. It was created as an educational exercise and is not intended for production use.

## Features
- Stores a limited number of items, automatically removing the least recently used item when capacity is exceeded.
- Designed to demonstrate basic cache eviction logic using a linked list and a map for O(1) access.

## Usage
```go
cache := lru.New[keyType, valueType](10)  // Initialize cache with capacity of 10
cache.Put("key1", "value1")
value := cache.Get("key1")
