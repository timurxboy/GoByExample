package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	// for i := 0; i < 5; i++ {
	// 	go sleepyGopher(i, c)
	// }

	// for i := 0; i < 5; i++ {
	// 	gopherId := <-c
	// 	fmt.Println("gopher", gopherId, "has finished sleeping")
	// }

	timeout := time.After(3 * time.Second)

	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)

		select {
		case gopherId := <-c:
			fmt.Println("gopher ", gopherId, "has finished sleeping")
		case <-timeout:
			fmt.Println("my patience ran out")
			return
		}
		go sleepyGopher()
	}
}

// func sleepyGopher(id int, c chan int) {
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("... ", id, "snore ...")
// 	c <- id
// }

func sleepyGopher(id int, c chan int) {
	// time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	time.Sleep(time.Second)
	c <- id
}
