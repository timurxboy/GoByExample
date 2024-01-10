package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptz(iptz *int) {
	*iptz = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptz(&i)
	fmt.Println("zeroptl:", i)

	fmt.Println("pointer:", &i)
}
