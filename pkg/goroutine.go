package pkg

import "log"

func Go(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("goroutine panic: %v\n", err)
			}
		}()
		f()
	}()
}
