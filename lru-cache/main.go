package lru_cache

import "container/list"

type LRUCache struct {
	list     *list.List
	vmap     map[int]*list.Element
	capacity int
}

type Item struct {
	key int
	val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		list:     list.New(),
		vmap:     make(map[int]*list.Element, capacity*2),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {

	if val, ok := this.vmap[key]; ok {
		this.list.MoveToFront(val)
		return val.Value.(*Item).val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {

	if val, ok := this.vmap[key]; ok {
		val.Value.(*Item).val = value
		this.list.MoveToFront(val)
	} else {
		val = this.list.PushFront(&Item{
			key, value,
		})
		this.vmap[key] = val
		if this.list.Len() > this.capacity {
			ele := this.list.Back()
			delete(this.vmap, ele.Value.(*Item).key)
			this.list.Remove(ele)
		}
	}

}
