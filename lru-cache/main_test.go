package lru_cache

import (
	"container/list"
	"testing"
)

func TestLRUCache(t *testing.T) {

	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	val := cache.Get(1) // returns 1
	if val != 1 {
		t.Errorf("expect %d, got %d", 1, val)
	}
	cache.Put(3, 3)    // evicts key 2
	val = cache.Get(2) // returns -1 (not found)
	if val != -1 {
		t.Errorf("expect %d, got %d", -1, val)
	}
	cache.Put(4, 4)    // evicts key 1
	val = cache.Get(1) // returns -1 (not found)
	if val != -1 {
		t.Errorf("expect %d, got %d", -1, val)
	}
	val = cache.Get(3) // returns 3
	if val != 3 {
		t.Errorf("expect %d, got %d", 3, val)
	}
	val = cache.Get(4) // returns 4
	if val != 4 {
		t.Errorf("expect %d, got %d", 4, val)
	}

}

func TestLRUCache_Get(t *testing.T) {
	type fields struct {
		list     *list.List
		vmap     map[int]*list.Element
		capacity int
	}
	type args struct {
		key int
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   int
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &LRUCache{
				list:     tt.fields.list,
				vmap:     tt.fields.vmap,
				capacity: tt.fields.capacity,
			}
			if got := this.Get(tt.args.key); got != tt.want {
				t.Errorf("LRUCache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
