package lfu_cache

import (
	"container/list"
)

type Item struct {
	key    int
	val    int
	parent *ParentItem
}

type ParentItem struct {
	ele       *list.Element
	sublist   *list.List //sublist 指针
	visit_num int        //访问次数
}

type LFUCache struct {
	list     *list.List //最外层的list
	vmap     map[int]*list.Element
	capacity int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		list:     list.New(),
		vmap:     make(map[int]*list.Element, capacity*2),
		capacity: capacity,
	}
}

func (this *LFUCache) Get(key int) int {

	if v, ok := this.vmap[key]; ok {
		item := v.Value.(*Item)
		this.updateVisitNum(v)
		return item.val
	}

	return -1

}

func (this *LFUCache) updateVisitNum(v *list.Element) {
	item := v.Value.(*Item)

	//找到parentItem的下一个元素A，如果不存在就创建一个
	//把当前元素移动到A的sublist的头部
	item.parent.sublist.Remove(v)

	next := item.parent.ele.Next()
	var nextp *ParentItem
	if next == nil {
		nextp = &ParentItem{
			sublist:   list.New(),
			visit_num: item.parent.visit_num + 1,
		}
		nextp.ele = this.list.InsertAfter(nextp, item.parent.ele)

	} else {
		nextp = next.Value.(*ParentItem)
		if nextp.visit_num != item.parent.visit_num+1 {
			nextp = &ParentItem{
				sublist:   list.New(),
				visit_num: item.parent.visit_num + 1,
			}
			nextp.ele = this.list.InsertAfter(nextp, item.parent.ele)
		}
	}
	this.vmap[item.key] = nextp.sublist.PushFront(item)

	if item.parent.sublist.Len() == 0 {
		this.list.Remove(item.parent.ele)
	}
	item.parent = nextp
}

func (this *LFUCache) Put(key int, value int) {

	if v, ok := this.vmap[key]; ok {
		item := v.Value.(*Item)
		item.val = value

		this.updateVisitNum(v)

		return
	}

	if this.capacity == 0 {
		return
	}

	if len(this.vmap) >= this.capacity {
		//需要移除this.list头节点的头
		back := this.list.Front()
		pi := back.Value.(*ParentItem)
		item := pi.sublist.Back()
		key2 := item.Value.(*Item).key
		//log.Println(key,value, len(this.vmap),this.capacity,this.list.Len(),"delete key:",key2)
		delete(this.vmap, key2)
		pi.sublist.Remove(item)
		if pi.sublist.Len() == 0 {
			//如果sublist为空，就把当前parentItem从this.list中移除
			this.list.Remove(back)
		}
	}

	back := this.list.Front()
	var pi *ParentItem
	if back != nil {
		pi = back.Value.(*ParentItem)
		if pi.visit_num != 0 {
			pi = &ParentItem{
				sublist:   list.New(),
				visit_num: 0,
			}
			back = this.list.PushFront(pi)
			pi.ele = back
		}
	} else {
		pi = &ParentItem{
			sublist:   list.New(),
			visit_num: 0,
		}
		back = this.list.PushFront(pi)
		pi.ele = back
	}

	item := &Item{
		key:    key,
		val:    value,
		parent: pi,
	}
	val := pi.sublist.PushFront(item)
	this.vmap[key] = val
}
