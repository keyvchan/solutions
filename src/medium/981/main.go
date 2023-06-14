package main

import "fmt"

func main() {
}

type TimeStamp struct {
	timestamp int
	value     [2]string
}

type TimeMap struct {
	time []TimeStamp
}

func Constructor() TimeMap {
	return TimeMap{}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	// set key to value at timestamp
	this.time = append(this.time, TimeStamp{timestamp: timestamp, value: [2]string{key, value}})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	// find the largest timestamp smaller than or equal to timestamp

	left, right := 0, len(this.time)-1
	for {
		fmt.Println(left, right)
		if left > right {
			// we could not find the key, return largest value
			// find back
			for i := left; i >= 0; i-- {
				if this.time[i].value[0] == key {
					return this.time[i].value[1]
				}
			}
			return ""
		}
		// check left and right
		mid := left + (right-left)/2
		if this.time[mid].timestamp > timestamp {
			right = mid - 1
		} else if this.time[mid].timestamp < timestamp {
			left = mid + 1
		} else {
			return this.time[mid].value[1]
		}
	}
}
