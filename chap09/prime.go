package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(x int) bool {
	if x <= 3 {
		return x > 1
	}

	if x%2 == 0 || x%3 == 0 {
		return false
	}

	for i := 5; i*i <= x; i += 6 {
		if x%i == 0 || x%(i+2) == 0 {
			return false
		}
	}

	return true
}

func primeTest(v chan int) {
	x := <-v

	if isPrime(x) {
		v <- 1
	} else {
		v <- 0
	}
}

func main() {
	goal := 100

	if len(os.Args) > 1 {
		g, err := strconv.Atoi(os.Args[1])
		if err != nil {
			goal = g
		}
	}

	chs := make([]chan int, goal)

	for _, ch := range chs {
		go prime(ch)
	}

	for i, ch := range chs {
		ch <- i
	}

	for i, ch := range chs {
		b := <-ch
		if b == 1 {
			fmt.Printf("%v ", i)
		}
	}

	fmt.Println()
}
