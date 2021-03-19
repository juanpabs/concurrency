package channels

import (
	"time"
)

func ChanWait(duration string) chan string {
	c := make(chan string)

	go func() {
		for {
			c <- duration
			duration, _ := time.ParseDuration(duration)
			time.Sleep(duration)
		}
	}()

	return c
}
