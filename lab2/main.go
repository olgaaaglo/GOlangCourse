package main

import (
	"fmt"
	"math/rand"
	"time"
	"errors"
	"strings"
)

func addToArr(tab [4]int, x int) ([]int, error) {
	sli := make([]int, len(tab))
	if tab[len(tab) - 1] != 0 {
		return tab[:], errors.New("Out of range!")
	}
	j := 0
	for j < len(tab) {
		if tab[j] == 0 {
			sli[j] = x
			return sli, nil
		} else if x < tab[j] {
			sli[j] = x
			for k := j; k < len(sli) - 1; k++ {
				sli[k + 1] = tab[k]
			}
			return sli, nil
		} else {
			sli[j] = tab[j]
		}
		j++
	}

	return sli, nil
}

func main() {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// fmt.Printf("%T\n", letters)
	// fmt.Printf("%T\n\n", letters[0])

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
		switch {
		case len(s) >= 3 && len(s) < 7:
			sliShort = append(sliShort, s)
		case len(s) >= 7 && len(s) < 11:
			sliMedium = append(sliMedium, s)
			if cap(sliMedium) > sliMediumCap {
				// fmt.Println(sliMedium, len(sliMedium))
			}
			sliMediumCap = cap(sliMedium)
		case len(s) >= 11 && len(s) <= 13:
			sliLong = append(sliLong, s)
		}

		i++
	}

	// fmt.Printf("%T\n ", sliShort)
	// fmt.Println("Short strings: ", sliShort, len(sliShort))
	// fmt.Println("Medium strings: ", sliMedium, len(sliMedium))
	// fmt.Println("Long strings: ", sliLong, len(sliLong))

	fmt.Println("1.")
	m := make(map[int] []string)
	for j := 0; j < len(sliShort); j++ {
		m[len(sliShort[j])] = append(m[len(sliShort[j])], sliShort[j])
	}
	for j := 0; j < len(sliMedium); j++ {
		m[len(sliMedium[j])] = append(m[len(sliMedium[j])], sliMedium[j])
	}
	for j := 0; j < len(sliLong); j++ {
		m[len(sliLong[j])] = append(m[len(sliLong[j])], sliLong[j])
	}
	// fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println()

	for k, v := range m {
		fmt.Println(k, v)
	}

//len string w unicode utf8 

	fmt.Println("\n2.")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	sli := make([][]int, 10)
	set := map[int]bool{}
	for j := 0; j < 5; j++ {
		size := 1 + r.Intn(10)
		sli[j] = make([]int, size)
		for i := 0; i < size; i++ {
			sli[j][i] = r.Intn(6)
		}
		fmt.Println(sli[j])
	}

	i = 0
next:
    for i < 6 {
		for j := 0; j < 5; j++ {
			iExists := false
			for k := 0; k < len(sli[j]); k++ {
				if sli[j][k] == i {
					iExists = true
				}
			}
			if !iExists {
				i++
				goto next
			}
		}
		set[i] = true
        i++
    }
		
	fmt.Println("Set:")
	for j := 0; j <= 5; j++ {
		_, ok := set[j]
		if ok == true {
			fmt.Println(j)
		}
	}

	fmt.Println("\n3.")
	tab3 := [4]int{0, 0, 0, 0}
	fmt.Println(tab3)

	values := [5]int{1, 3, 2, 1, 4}

	for i := range values {
		sli3, err := addToArr(tab3, values[i])
		copy(tab3[:], sli3[:])
		fmt.Println(tab3)
		if err != nil {
			fmt.Println(err)
		}
	}
	
	fmt.Println("\n4.")
	fmt.Println(toLower(addDotToStr)("UpperOrLower"))
	fmt.Println(removeWhiteSpace(addDotToStr)("No white space at the end  "))
	fmt.Println(addPrefix(addDotToStr)(" before string"))
	decors := []funcfuncString{removeWhiteSpace, toLower, addPrefix}
	fmt.Println(aggregateDecorators(addDotToStr, decors...)("All Decorators!  "))
	
}

type funcString func(string) string
type funcfuncString func(funcString) funcString

func toLower(f funcString) funcString {
	return func(s string) string { return strings.ToLower(f(s)) }
}

func removeWhiteSpace(f funcString) funcString {
	return func(s string) string { 
		for s[len(s) - 1] == ' ' {
			s = s[:len(s) - 1]
		}
		return f(s)
	}
}

func addPrefix(f funcString) funcString {
	return func(s string) string { return "prefix" + f(s) }
}

func aggregateDecorators(f funcString, decors ...funcfuncString) funcString {
	return func(s string) string {
		var fs funcString = f
		for _, decor := range decors {
			fs = decor(fs)
		}
		return fs(s)
	}
}

func addDotToStr(s string) string {
	return s + "."
}