package main 

import (
	"fmt"
	"strconv"
)


func f1[T any](sli []T, fun func(t T) bool) []T {
	newSli := make([]T, 0)
	for _, val := range sli {
		if fun(val) {
			newSli = append(newSli, val)
		}
	}
	return newSli
}

func f12[T ~int | ~float64 | string](t T) bool {
	var p interface{}
	p = *new(T)

	switch p.(type) {
	case int:
		return t > T(0)
	case float64:
		return t < T(0)
	default:
		return true
	}
}



func f2[T any](sli []T, fun func(t T) T) []T {
	newSli := make([]T, 0)
	for _, val := range sli {
		newSli = append(newSli, fun(val))
	}
	return newSli
}

func f22[T ~int | ~float64](t T) T {
	var p interface{}
	p = *new(T)

	switch p.(type) {
	case int:
		return t * t
	case float64:
		return t * t * t
	default:
		return t
	}
}



func f3[T any](sli []T, fun func(t1 T, t2 T) T) T {
	if len(sli) > 0 {
		result := sli[0]
		for i := 1; i < len(sli); i++ {
			result = fun(result, sli[i])
		}
		return result
	}
	return *new(T)
}

func f32[T int | float64 | string](t1 T, t2 T) T {
	if t1 < t2 {
		return t1
	} else {
		return t2
	}
}



func f4[T comparable](m map[T]T) ([]T, []T) {
	keys := make([]T, 0)
	values := make([]T, 0)

	for key, val := range m {
		keys = append(keys, key)
		values = append(values, val)
	}

	return keys, values
}



func f5[T comparable](keys []T, values []T) map[T]T {
	m := make(map[T] T)
	if len(keys) < len(values) {
		for i, key := range keys {
			m[key] = values[i]
		}
	} else {
		for i, val := range values {
			m[keys[i]] = val
		}
	}
	return m
}



func f6[T int | float64](s []string) []T {
	newSli := make([]T, 0)
	var p interface{}
	p = *new(T)

	switch p.(type) {
	case int:
		for _, val := range s {
			intVal, err := strconv.Atoi(val)
			if err == nil {
				newSli = append(newSli, T(intVal))
			}
		}
	case float64:
		for _, val := range s {
			floatVal, err := strconv.ParseFloat(val, 64)
			if err == nil {
				newSli = append(newSli, T(floatVal))
			}
		}
	default:
	}
	return newSli
}

func main() {
	fmt.Println("1.")
	sli := []int{-1, 2, 3}
	fmt.Println(f1(sli, f12[int]))

	sli2 := []float64{1.0, -2.0, 3.0}
	fmt.Println(f1(sli2, f12[float64]))

	fmt.Println("\n2.")
	sli3 := []int{1, 2, 3}
	fmt.Println(f2(sli3, f22[int]))
	fmt.Println(f2([]float64{1.1, 2.1, 3.1}, f22[float64]))

	fmt.Println("\n3.")
	fmt.Println(f3([]int{10, 2, 31}, f32[int]))

	fmt.Println("\n4.")
	m := make(map[int]int)
	m[1] = 2
	m[3] = 4
	fmt.Println(f4(m))

	fmt.Println("\n5.")
	fmt.Println(f5([]int{1, 2, 3}, []int{4, 5, 6}))

	fmt.Println("\n6.")
	fmt.Println(f6[int]([]string{"1", "2", "3"}))
	fmt.Println(f6[float64]([]string{"1.1", "2.2", "3.3"}))
}