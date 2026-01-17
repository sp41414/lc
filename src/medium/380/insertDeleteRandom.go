package insertDeleteRandom

import "math/rand/v2"

type RandomizedSet struct {
	valIndexes map[int]int
	vals       []int
}

func Constructor() RandomizedSet {
	vals := []int{}
	valIndexes := map[int]int{}

	return RandomizedSet{
		valIndexes: valIndexes,
		vals:       vals,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.valIndexes[val]; ok {
		return false
	}
	this.vals = append(this.vals, val)
	this.valIndexes[val] = len(this.vals) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.valIndexes[val]
	if !ok || len(this.vals) == 0 {
		return false
	}

	lastIndex := len(this.vals) - 1
	if index != lastIndex {
		this.vals[index], this.vals[lastIndex] = this.vals[lastIndex], this.vals[index]
		this.valIndexes[this.vals[index]] = index
	}

	delete(this.valIndexes, val)
	this.vals = this.vals[:len(this.vals)-1]

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.vals[rand.IntN(len(this.vals))]
}
