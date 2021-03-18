package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	c := make(chan string)
	go count("sheep", c)

	msg := <-c
	fmt.Println(msg)
	msg = <-c
	fmt.Println(msg)
}

func count(thing string, c chan string) {
	for i := 1; i <= 10; i++ {
		c <- strconv.FormatInt(int64(i), 16) + " " + thing
		time.Sleep(time.Millisecond * 500)
	}
}
