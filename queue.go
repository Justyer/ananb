package ananb

import (
	"errors"
	"fmt"
	"sync"
)

var (
	//ErrQueueFull is null
	ErrQueueFull = errors.New("queue has been full")
	//ErrQueueEmpty is null
	ErrQueueEmpty = errors.New("queue has no data")
)

// qnode : 传输字节内容的载体
type qnode struct {
	Ctx []byte
}

// Queue : 任务队列
type Queue struct {
	Chl  chan *qnode
	cap  int
	lock *sync.Mutex
}

// NewQueue : 实例化一个任务队列
func NewQueue() *Queue {
	return &Queue{}
}

// Init : 初始化队列参数
func (q *Queue) Init() *Queue {
	q.Chl = make(chan *qnode, q.cap)
	q.lock = &sync.Mutex{}
	return q
}

// SetCap : 设置队列最大长度
func (q *Queue) SetCap(cap int) *Queue {
	q.cap = cap
	return q
}

// Push : 往队列中打入数据
func (q *Queue) Push(c []byte) error {
	if q.Len() < q.Cap() {
		q.Chl <- &qnode{Ctx: c}
		return nil
	}
	return ErrQueueFull
}

// PushMany : 往队列中打入一批数据
func (q *Queue) PushMany(cs [][]byte) error {
	var errs []error
	for i := 0; i < len(cs); i++ {
		err := q.Push(cs[i])
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("%s", errs)
	}
	return nil
}

// MustPush : 往队列中打入数据，如果队列满了则阻塞，直到队列中有空位
func (q *Queue) MustPush(c []byte) {
	q.Chl <- &qnode{Ctx: c}
}

// MustPushMany : 往队列中打入一批数据，如果队列满了则阻塞，直到队列中有空位
func (q *Queue) MustPushMany(cs [][]byte) {
	for i := 0; i < len(cs); i++ {
		q.MustPush(cs[i])
	}
}

// Pop : 从队列中取出数据
func (q *Queue) Pop() ([]byte, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.pop()
}

func (q *Queue) pop() ([]byte, error) {
	if q.Len() > 0 {
		return (<-q.Chl).Ctx, nil
	}
	return nil, ErrQueueEmpty
}

// PopMany : 从队列中取出多个数据
func (q *Queue) PopMany(count int, mustGet bool) ([][]byte, error) {
	if !mustGet && count > q.Len() {
		return nil, errors.New("task not enough")
	}

	q.lock.Lock()
	defer q.lock.Unlock()

	var rlts [][]byte
	for i := 0; i < count; i++ {
		rlt, err := q.pop()
		if err != nil {
			return rlts, nil
		}
		rlts = append(rlts, rlt)
	}
	return rlts, nil
}

// Cap : 队列容量
func (q *Queue) Cap() int {
	return cap(q.Chl)
}

// Len : 队列容量
func (q *Queue) Len() int {
	return len(q.Chl)
}
