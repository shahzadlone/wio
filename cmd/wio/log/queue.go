package log

import (
    "fmt"
    "github.com/fatih/color"
)

type node struct {
    Value queueBuffer
}

// NewQueue returns a new queue with the given initial size.
func NewQueue(size int) *Queue {
    return &Queue{
        nodes: make([]*node, size),
        size:  size,
    }
}

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type Queue struct {
    nodes []*node
    size  int
    head  int
    tail  int
    count int
}

// Push adds a node to the queue.
func (q *Queue) push(n *node) {
    if q.head == q.tail && q.count > 0 {
        nodes := make([]*node, len(q.nodes)+q.size)
        copy(nodes, q.nodes[q.head:])
        copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
        q.head = 0
        q.tail = len(q.nodes)
        q.nodes = nodes
    }
    q.nodes[q.tail] = n
    q.tail = (q.tail + 1) % len(q.nodes)
    q.count++
}

// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) pop() *node {
    if q.count == 0 {
        return nil
    }
    node := q.nodes[q.head]
    q.head = (q.head + 1) % len(q.nodes)
    q.count--
    return node
}

type queueBuffer struct {
    text          string
    logType       string
    providedColor *color.Color
}

// creates a queue buffer from the logging information
func pushLog(queue *Queue, logType string, providedColor *color.Color, message string, a ...interface{}) {
    text := message

    if a != nil {
        text = fmt.Sprintf(message, a...)
    }

    buff := queueBuffer{logType: logType, providedColor: providedColor, text: text}
    queue.push(&node{buff})
}

// removes the queue buffer and returns the value
func popLog(queue *Queue) queueBuffer {
    return queue.pop().Value
}
