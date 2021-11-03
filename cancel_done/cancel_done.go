package main

import (
	"log"
	"time"
)

func main() {
	done := make(chan struct{})
	tick := time.Tick(500 * time.Millisecond)

	go func() {
		time.Sleep(time.Second * 10)
		close(done)
	}()

exit:
	for {
		select {
		case <-done:
			log.Println("done")
			break exit
		case <-tick:
			log.Println("tick")
		}
	}
}
