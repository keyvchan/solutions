package main

import (
	"container/list"
)

func main() {
	obj := Constructor("jrbilt.com")
	obj.Visit("uiza.com")

}

type BrowserHistory struct {
	// the browser history should have a queue, and we have current pointer

	current *list.Element

	history *list.List
}

func Constructor(homepage string) BrowserHistory {

	BrowserHistory := BrowserHistory{}
	BrowserHistory.history = list.New()
	BrowserHistory.history.PushBack(homepage)
	BrowserHistory.current = BrowserHistory.history.Back()
	return BrowserHistory

}

func (this *BrowserHistory) Visit(url string) {

	// how remove part of the list?

	// clear the history after current pointer
	temp := this.current.Next()

	for temp != nil {
		this.history.Remove(temp)
		temp = this.current.Next()
	}

	// put the url in the queue at current pointer
	this.history.PushBack(url)
	this.current = this.history.Back()

}

func (this *BrowserHistory) Back(steps int) string {
	// back from current pointer
	for i := 0; i < steps; i++ {
		if this.current.Prev() != nil {
			this.current = this.current.Prev()
		}

	}
	return this.current.Value.(string)
}

func (this *BrowserHistory) Forward(steps int) string {
	// forward from current pointer
	for i := 0; i < steps; i++ {
		if this.current.Next() != nil {
			this.current = this.current.Next()
		}
	}
	return this.current.Value.(string)
}
