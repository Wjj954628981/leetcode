import "container/list"

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 */
type LRUCache struct {
	HashMap    map[int]*list.Element
	LinkedList list.List
	Capacity   int
}

type Pair struct {
	k int
	v int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		HashMap:    map[int]*list.Element{},
		LinkedList: list.List{},
		Capacity:   capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	ele, ok := this.HashMap[key]
	if !ok {
		return -1
	}
	this.LinkedList.MoveToFront(ele)
	return ele.Value.(Pair).v
}

func (this *LRUCache) Put(key int, value int) {
	if ele, ok := this.HashMap[key]; ok {
		ele.Value = Pair{key, value}
		this.LinkedList.MoveToFront(ele)
		return
	}
	if this.LinkedList.Len() == this.Capacity {
		back := this.LinkedList.Back()
		this.LinkedList.Remove(back)
		delete(this.HashMap, back.Value.(Pair).k)
	}
	this.HashMap[key] = this.LinkedList.PushFront(Pair{key, value})
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

