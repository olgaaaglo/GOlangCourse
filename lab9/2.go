package main 

import (
	"fmt"
)

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return fibonacci(n - 1) + fibonacci(n - 2)
	}
}

func getFibonacciChan(ch chan int) {
	defer close(ch)
	i := 0
	for {
		ch <- fibonacci(i)
		i++
	}
}

func getEvensSum(to int, chEven chan int) {
	fibCh := make(chan int)
	go getFibonacciChan(fibCh)

	sum := 0
	for {
		fibVal := <-fibCh
		if fibVal >= to {
			break
		}
		if fibVal % 2 == 0 {
			sum += fibVal
		}
	}
	chEven <- sum
}

func getOddSum(to int, chOdd chan int) {
	fibCh := make(chan int)
	go getFibonacciChan(fibCh)

	sum := 0
	for {
		fibVal := <-fibCh
		if fibVal >= to {
			break
		}
		if fibVal % 2 == 1 {
			sum += fibVal
		}
	}
	chOdd <- sum
}

func main() {
	chEven := make(chan int)
	chOdd := make(chan int)

	go getEvensSum(100, chEven)
	go getOddSum(100, chOdd)

	fmt.Println("Sum of even values: ", <-chEven)
	fmt.Println("Sum of odd values: ", <-chOdd)
}