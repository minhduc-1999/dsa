package lrucache

// Title: LRU Cache

// Problem link: https://leetcode.com/problems/lru-cache

// Testcases:
// case 0: ["LRUCache","put","put","get","put","get","put","get","get","get"]	[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]

type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

type LRUCache struct {
	size int
	cap  int
	head *Node
	tail *Node
	hmap map[int]*Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:  capacity,
		hmap: map[int]*Node{},
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.hmap[key]
	if !ok {
		return -1
	}
	this.removeNode(node)
	this.add(node)
	return node.Val
}

func (this *LRUCache) add(node *Node) {
	if this.size == this.cap {
		delete(this.hmap, this.tail.Key)
		if this.tail.Prev != nil {
			this.tail = this.tail.Prev
			this.tail.Next = nil
		}
	} else {
		this.size++
	}
	if this.head != nil {
		node.Next = this.head
		this.head.Prev = node
	} else {
		this.tail = node
	}
	this.head = node
}

func (this *LRUCache) removeNode(node *Node) {
	this.size--
	if node == this.head {
		this.head = node.Next
	}
	if node == this.tail {
		this.tail = node.Prev
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	node.Prev = nil
	node.Next = nil
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.hmap[key]
	if !ok {
		node = &Node{
			Val: value,
			Key: key,
		}
		this.add(node)
		this.hmap[key] = node
	} else {
		node.Val = value
		this.removeNode(node)
		this.add(node)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
