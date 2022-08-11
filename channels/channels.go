package channels

import (
	"fmt"
	"time"
)

func CountFunc() {
	ch := make(chan string)
	hi := "hi"

	go count(ch, hi)
	for msg := range ch {
		fmt.Println(msg)
	}
}

func count(ch chan string, thing string) {
	for i := 0; i < 5; i++ {
		ch <- thing
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}
