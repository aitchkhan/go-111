package channels

import (
	"fmt"
	"time"
)

func SelectStatement() {
	ch1 := make(chan string, 2)
	ch2 := make(chan string, 2)
	go func() {
		for {
			ch2 <- "msg1"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			ch2 <- "msg2"
			time.Sleep(time.Second * 2)
		}
	}()

	// for msg3:= range <-ch1{
	// 	fmt.Println(msg3)
	// 	// msg4 := <-ch2
	// 	// fmt.Println(msg4)
	// }
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}

	}
}
