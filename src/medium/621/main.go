package main

import "container/heap"

func main() {
}

type Task struct {
	value    byte
	priority *int
}
type PriorityQueue []Task

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return *pq[i].priority < *pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Task))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func leastInterval(tasks []byte, n int) int {
	priorities := map[byte]*int{}
	pq := PriorityQueue{}
	heap.Init(&pq)

	for _, task := range tasks {
		// check priority in map
		if _, ok := priorities[task]; !ok {
			// if not, create new priority
			priorities[task] = new(int)
		}
		task := Task{
			value:    task,
			priority: priorities[task],
		}
		heap.Push(&pq, task)
	}

	count := 0
	for pq.Len() > 0 {
		// check interval of most priority task
		if *pq[0].priority > 0 {
			// idle and all interval should minus 1
			for _, p := range priorities {
				*p -= 1
			}
			heap.Init(&pq)
		} else {
			// pop task and decrease priority
			task := heap.Pop(&pq).(Task)
			// idle and all interval should minus 1
			for t, p := range priorities {
				if t != task.value {
					*p -= 1
				}
			}
			*task.priority += n
			heap.Init(&pq)
		}
		count++
	}

	return count
}
