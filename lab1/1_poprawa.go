package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Printf("%T\n", letters)
	fmt.Printf("%T\n\n", letters[0])

	i := 0
	sliShort := make([]string, 0)
	sliMedium := make([]string, 0)
	sliLong := make([]string, 0)
	for i < 100 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		size := 3 + r.Intn(11)
		
		genString:
		var s string = ""
		prev := ""
		for j := 0; j < size; j++ {
			idx := r.Intn(len(letters))
			char := string(letters[idx])
			if char == prev {
				goto genString
			}
			prev = char
			s += char
		}
		sliMediumCap := cap(sliMedium)
		// switch len(s) {
		// case 3, 4, 5, 6:
		// 	sliShort = append(sliShort, s)
		// case 7, 8, 9, 10:
		// 	sliMedium = append(sliMedium, s)
		// 	if cap(sliMedium) > sliMediumCap {
		// 		fmt.Println(sliMedium, len(sliMedium))
		// 	}
		// 	sliMediumCap = cap(sliMedium)
		// case 11, 12, 13:
		// 	sliLong = append(sliLong, s)
		// }
		switch {
		case len(s) >= 3 && len(s) < 7:
			sliShort = append(sliShort, s)
		case len(s) >= 7 && len(s) < 11:
			sliMedium = append(sliMedium, s)
			if cap(sliMedium) > sliMediumCap {
				fmt.Println(sliMedium, len(sliMedium))
			}
			sliMediumCap = cap(sliMedium)
		case len(s) >= 11 && len(s) <= 13:
			sliLong = append(sliLong, s)
		}

		i++
	}

	fmt.Println("\nResult: ")
	fmt.Println("Short strings: ", sliShort, len(sliShort))
	fmt.Println("Medium strings: ", sliMedium, len(sliMedium))
	fmt.Println("Long strings: ", sliLong, len(sliLong))
}