package main

import (
	"container/list"
	"fmt"
	"io"
)

func main() {

	var n int

	for {
		_, err := fmt.Scan(&n)
		if err == io.EOF {
			break
		}
		A := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&A[i])
		}
		solution(A)

	}
}

func solution(A []int) {
	if len(A) <= 2 {
		// print it directly
		for i := 0; i < len(A); i++ {
			fmt.Print(A[i], " ")
		}
		fmt.Println()
	}
	fmt.Println(A)

	queue := list.New()
	queue.PushBack(A[len(A)-2])
	queue.PushBack(A[len(A)-1])

	for i := len(A) - 3; i >= 0; i-- {
		queue.PushFront(A[i])
		temp1 := queue.Back()
		queue.Remove(queue.Back())
		queue.PushFront(temp1.Value.(int))
		temp2 := queue.Back()
		queue.Remove(queue.Back())
		queue.PushFront(temp2.Value.(int))
	}

	for queue.Len() > 0 {
		element := queue.Remove(queue.Front())
		if element != 0 {
			fmt.Print(element.(int), " ")
		}
	}
	fmt.Println()

}
