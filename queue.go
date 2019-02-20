package queue

import (
	"errors"
)

// type queue interface {
// 	enqueue(int) int
// 	dequeue(int) int
// 	peek(int) int
// }

type intQueue struct {
	data []int
}

func (q *intQueue) enqueue(i int) (int, error) {
	(*q).data = append((*q).data, i)
	return len((*q).data), nil
}

func (q *intQueue) dequeue() (int, error) {
	if len((*q).data) == 0 {
		return 0, errors.New("Nothing in queue")
	}

	res := (*q).data[len((*q).data)-1]
	(*q).data = (*q).data[:len((*q).data)-1]
	return res, nil
}

func (q *intQueue) peek() (int, error) {
	if len((*q).data) == 0 {
		return 0, errors.New("Nothing in queue")
	}

	return (*q).data[len((*q).data)-1], nil
}
