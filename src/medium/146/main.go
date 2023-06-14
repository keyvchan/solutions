package main

import "container/list"

type LRUCache struct {
	least       *list.List
	least_index map[int]*list.Element
	capacity    int
	store       map[int]int
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		least:       list.New(),
		least_index: map[int]*list.Element{},
		capacity:    capacity,
		store:       map[int]int{},
	}
	return cache
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.store[key]; ok {
		this.least.MoveToFront(this.least_index[key])
		return this.store[key]
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// if key is in the cache, update the value and move the key to the front
	if _, ok := this.store[key]; ok {
		this.store[key] = value
		this.least.MoveToFront(this.least_index[key])
		return
	}
	if len(this.store) == this.capacity {
		// replace the least with the new key
		delete(this.store, this.least.Remove(this.least.Back()).(int))
		this.least.PushFront(key)
		this.least_index[key] = this.least.Front()
		this.store[key] = value
	} else {
		this.least.PushFront(key)
		this.least_index[key] = this.least.Front()
		this.store[key] = value
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {

}
