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

func primeTest(v <-chan int, o chan<- int) {
	x := <-v

	if isPrime(x) {
		fmt.Printf("%v ", x)
	}

	o <- 0
}

func main() {
	goal := 100

	if len(os.Args) > 1 {
		g, err := strconv.Atoi(os.Args[1])
		if err == nil {
			goal = g
		}
	}

	ich := make(chan int, goal+1)
	och := make(chan int, goal+1)

	for i := 1; i <= goal; i++ {
		go primeTest(ich, och)
	}

	for i := 1; i <= goal; i++ {
		ich <- i
	}

	for i := 1; i <= goal; i++ {
		<-och
	}

	fmt.Println()
}
