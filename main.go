package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/concurrency/channels"
)

func main() {
	c1 := channels.ChanWait("500ms")
	c2 := channels.ChanWait("1s")

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

func count(thing string, c chan string) {
	for i := 10; i <= 20; i++ {
		c <- strconv.FormatInt(int64(i), 10) + " " + thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
