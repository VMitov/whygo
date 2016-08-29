package main

import "fmt"

var msgs = make(chan int)
var done = make(chan bool)

func produce() {
	for i := 0; i < 10; i++ {
		msgs <- i
	}
	done <- true
}

func consume() {
	for {
		msg := <-msgs
		fmt.Println(msg)
	}
}

func main() {
	go produce()
	go consume()
	<-done
}
