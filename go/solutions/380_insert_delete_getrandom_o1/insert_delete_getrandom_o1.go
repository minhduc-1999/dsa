package insertdeletegetrandomo1

import "math/rand/v2"

// Title: Insert Delete GetRandom O(1)

// Problem link: https://leetcode.com/problems/insert-delete-getrandom-o1

// Testcases:
// case 0: ["RandomizedSet","insert","remove","insert","getRandom","remove","insert","getRandom"]	[[],[1],[2],[2],[],[1],[2],[]]

type RandomizedSet struct {
	h map[int]int
	v []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		h: make(map[int]int),
		v: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.h[val]
	if ok {
		return false
	}
	this.h[val] = len(this.v)
	this.v = append(this.v, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	i, ok := this.h[val]
	if ok {
		this.v[i] = this.v[len(this.v)-1]
		this.h[this.v[len(this.v)-1]] = i
		this.v = this.v[:len(this.v)-1]
		delete(this.h, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	r := rand.IntN(len(this.v))
	return this.v[r]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
