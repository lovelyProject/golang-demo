package main

import (
	"fmt"
	"math/rand"
)

func main() {
	chanOne := make(chan int, 10)
	chanTwo := make(chan int, 10)
	go createAndFill(chanOne)
	go mathPow(chanOne, chanTwo)

	for num := range chanTwo {
		fmt.Println(num)
	}
}

func createAndFill(chanOne chan int) {
	randSlice := make([]int, 10, 10)
	for i := range randSlice {
		randSlice[i] = rand.Intn(100)
	}

	fmt.Println(randSlice)

	for _, v := range randSlice {
		chanOne <- v
	}

	close(chanOne)
}

func mathPow(chanOne chan int, chanTwo chan int) {
	for num := range chanOne {
		num = num * num
		chanTwo <- num
	}

	close(chanTwo)
}
