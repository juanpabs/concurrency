package main

import (
	"strconv"
	"time"

	"github.com/concurrency/channels"
)

func main() {
	c := make(chan string)
	go count("sheep", c)

	channels.InfiniteForRangeChan(c)
}

func count(thing string, c chan string) {
	for i := 10; i <= 20; i++ {
		c <- strconv.FormatInt(int64(i), 10) + " " + thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
