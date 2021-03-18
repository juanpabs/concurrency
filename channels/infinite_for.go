package channels

import "fmt"

//when reading a channel will couse a fatal error when the channer get out of info.
func Infinite_for(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
	}
}

func InfiniteForCheckChanStatus(c chan string) {
	for {
		msg, open := <-c
		if !open {
			break
		}
		fmt.Println(msg)
	}
}
