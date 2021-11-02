package main

import "log"

func main() {
	ch := make(chan int, 5)
	done := make(chan struct{})

	go func() {
		//values := []int{1, 2, 3, 4, 5}
		//for _, val := range values {
		//	ch <- val
		//}

		ch <- 1
		ch <- 2
		//close(ch)
	}()

	go func() {
		for v := range ch {
			log.Println(v)
		}
		done <- struct{}{}
	}()

	<-done
}
