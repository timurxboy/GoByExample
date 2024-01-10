package main

import "time"

func main() {
	timeout := time.After(2 * time.Second)
	for i:= 0; i<5; i++ {
		select {
		case gopherId := <-c
		}
	}
}
