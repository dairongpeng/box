package buffer

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"testing"
	"time"
)

//=== RUN   TestBuffer
//buffer_test.go:50: sending ....
//receive msg because data len = 1058
//receive msg because time len = 173
//receive msg because data len = 1097
//receive msg because data len = 1193
//receive msg because time len = 189
//receive msg because data len = 1150
//receive msg because data len = 1226
//receive msg because time len = 23
//receive msg because time len = 849
//buffer closed | size=0
//input buffer err because buffer closed | size=0
//input buffer err because buffer closed | size=0
//input buffer err because buffer closed | size=0
//input buffer err because buffer closed | size=0
//input buffer err because buffer closed | size=0
//buffer_test.go:57: exited!
//--- PASS: TestBuffer (46.06s)
//PASS
//
//Process finished with the exit code 0

func TestBuffer(t *testing.T) {
	ctx := context.Background()
	// 1k && 5s
	buffer := NewBuffer(ctx, 1024, 5*time.Second, receiveMsg)

	go func() {
		for i := 0; i < 65; i++ {
			// 验证buffer关闭的逻辑
			if i == 60 {
				buffer.Close()
				fmt.Printf("buffer closed | size=%d \n", buffer.Size())
			}
			go func() {
				bts := generateData()
				err := buffer.Input(bts)
				if err != nil {
					fmt.Printf("input buffer err because buffer closed | size=%d \n", buffer.Size())
				}
			}()
			time.Sleep(500 * time.Millisecond)
		}
	}()

	t.Log("sending ....")

	exit := make(chan os.Signal, 0)
	signal.Notify(exit, os.Interrupt)

	<-exit

	t.Log("exited!")
}

func receiveMsg(bytes []byte) error {
	if len(bytes) < 1024 {
		fmt.Printf("receive msg because time len = %d \n", len(bytes))
	} else {
		fmt.Printf("receive msg because data len = %d \n", len(bytes))
	}
	return nil
}

func generateData() []byte {
	rand.Seed(time.Now().UnixNano())
	dataSize := rand.Intn(256)

	data := make([]byte, dataSize)
	rand.Read(data)

	return data
}
