package main

import (
	"sort"
)

func main() {
}

type TimeMap struct {
	time_index    []int
	time_to_index map[int]int
	// this should be a ordered map
	value_store []map[string]string
}

func Constructor() TimeMap {
	return TimeMap{
		value_store:   []map[string]string{},
		time_to_index: map[int]int{},
		time_index:    []int{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	// time_index store all index of timestamp
	this.time_index = append(this.time_index, timestamp)
	this.value_store = append(this.value_store, map[string]string{key: value})

	// time_to_index store the index of corresponding timestamp in value_store
	this.time_to_index[len(this.time_index)-1] = len(this.value_store) - 1
}

func (this *TimeMap) Get(key string, timestamp int) string {

	// find the index of timestamp in time_index
	index := sort.SearchInts(this.time_index, timestamp)
	if index == len(this.time_index) {
		index = len(this.time_index) - 1
	} else {
		if this.time_index[index] != timestamp {
			if index > 0 {
				index -= 1
			}
		}

	}

	return this.value_store[this.time_to_index[index]][key]
}
