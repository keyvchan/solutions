package main

import "sort"

func main() {

}

type Interval struct {
	Start int
	End   int
}

type MyCalendar struct {
	// intervals is a list of intervals
	intervals []Interval
}

func Constructor() MyCalendar {
	return MyCalendar{}

}

func (this *MyCalendar) Book(start int, end int) bool {
	for _, interval := range this.intervals {
		// check every interval
		if (start >= interval.Start && start < interval.End) || (end > interval.Start && end <= interval.End) {
			// check start
			return false
		}
		if start < interval.Start && end > interval.End {
			// includes
			return false
		}
	}
	// merge the interval
	for i := 0; i < len(this.intervals); i++ {
		// check end < start
		if end == this.intervals[i].Start {
			this.intervals[i].Start = start
			return true
		} else if start == this.intervals[i].End {
			this.intervals[i].End = end
			return true
		}
	}
	this.intervals = append(this.intervals, Interval{start, end})
	// sort
	sort.Slice(this.intervals, func(i, j int) bool {
		return this.intervals[i].Start < this.intervals[j].Start
	})
	return true

}
