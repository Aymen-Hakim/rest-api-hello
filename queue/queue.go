package queue

import "errors"

type Queue struct {
	slice []string
}

func New() *Queue {
	return &Queue{[]string{}}
}

func (q *Queue) Enqueue(cmd string) {
	q.slice = append(q.slice, cmd)
}

func (q *Queue) Dequeue() (string, error) {
	if len(q.slice) == 0 {
		return "", errors.New("the queue is empty")
	}
	item := q.slice[0]
	q.slice = q.slice[1:]
	return item, nil
}
