package ananb

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue().SetCap(100).Init()

	// err := q.Push([]byte{'1'})
	// if err != nil {
	// 	t.Error("push_err", err)
	// }

	rlt, err := q.Pop()
	if err != nil {
		t.Error("pop_err", err)
	}
	fmt.Println(rlt)
}
