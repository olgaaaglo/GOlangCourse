package main 

import (
	"fmt"
	"time"
)

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return fibonacci(n - 1) + fibonacci(n - 2)
	}
}

func printInfinity(t time.Duration) {
	str := `-\|/`
	i := 0
	for {
		time.Sleep(t)
		if i == 4 {
			i = 0
		}
		fmt.Print(string(str[i]))
		i++
	}
}

func main() {
	go printInfinity(250 * time.Millisecond)
	fmt.Println(fibonacci(45))
}