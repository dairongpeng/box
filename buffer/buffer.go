package buffer

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Buffer struct {
	sync.Mutex

	data []byte

	fn CallBack

	maxByteSize int
	minSendTime time.Duration

	active         bool
	overflowStatus chan int
}

type CallBack func(b []byte) error

func NewBuffer(ctx context.Context, maxByteSize int,
	minSendTime time.Duration, callback CallBack) *Buffer {
	b := &Buffer{
		data: []byte{},

		fn: callback,

		maxByteSize: maxByteSize,
		minSendTime: minSendTime,

		active:         true,
		overflowStatus: make(chan int, 0),
	}

	go b.daemon(ctx)

	return b
}

func (buf *Buffer) Size() int {
	buf.Lock()
	defer buf.Unlock()

	return len(buf.data)
}

func (buf *Buffer) daemon(ctx context.Context) {
	ticker := time.NewTicker(buf.minSendTime)

	for {
		select {
		case <-buf.overflowStatus:
			buf.output()
		case <-ticker.C:
			buf.output()
		case <-ctx.Done():
			buf.output()
			return
		}
	}
}

func (buf *Buffer) Close() {
	buf.Lock()
	defer buf.Unlock()

	buf.active = false
}

// Input 输入数据到缓冲区中
func (buf *Buffer) Input(dt []byte) error {
	buf.Lock()
	defer buf.Unlock()

	if buf.active {
		buf.data = append(buf.data, dt...) // 数据添加到缓冲区
		if len(buf.data) >= buf.maxByteSize {
			buf.overflowStatus <- 1
		}
	} else {
		return fmt.Errorf("buffer closed")
	}

	return nil
}

// Output 输出数据到指定的目标
func (buf *Buffer) output() {
	buf.Lock()
	defer buf.Unlock()

	if len(buf.data) == 0 {
		return // 缓冲区为空，无需输出
	}

	err := buf.fn(buf.data)
	if err != nil {
		// 当回调失败时，数据无法选择清空
		return
	}

	buf.data = []byte{} // 清空缓冲区
}
