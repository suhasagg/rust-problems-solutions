import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	ll       *list.List
	cache    map[int]*list.Element
}

type entry struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		ll:       list.New(),
		cache:    make(map[int]*list.Element),
	}
}

func (c *LRUCache) Get(key int) int {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		return ele.Value.(*entry).value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		ele.Value.(*entry).value = value
		return
	}
	ele := c.ll.PushFront(&entry{key, value})
	c.cache[key] = ele
	if c.ll.Len() > c.capacity {
		c.removeOldest()
	}
}

func (c *LRUCache) removeOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
	}
}
