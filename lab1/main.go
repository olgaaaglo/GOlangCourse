// /usr/lib/go-1.18/bin/go run main.go
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello!")

	x := 7
	var pi *int
	pi = &x
	fmt.Println(*pi, pi)
	fmt.Printf("%T %T\n", *pi, pi)

	var c float64 = 3
	var d float64 = c * 4
	fmt.Println(d)

	var e int = 4
	var f float64 = c * float64(e)
	fmt.Println(f)

	fmt.Println(math.Sqrt(4))

	// // sli := []int{}
	// sli := make([]int, 0)
	// fmt.Println(sli)
	// fmt.Printf("%T\n", sli)

	// // sli[0] = 1

	// sli2 := []int{9: 0}
	// fmt.Println(sli2)
	// fmt.Println(len(sli2), cap(sli2));

	// sli = append(sli, 1)
	// fmt.Println(len(sli), cap(sli));

	// sli = append(sli, 1)
	// fmt.Println(len(sli), cap(sli));

	// sli = append(sli, 1)
	// fmt.Println(len(sli), cap(sli));


	// sli3 := make([]int, 5, 10)
	// fmt.Println(sli3, len(sli3), cap(sli3));

	// sli3[4] = 1
	// // sli3[5] = 1



	tab := [5]int{1, 2, 3}
	sli := tab[:]
	sli[1] = 7
	sli = append(sli, 9)
	fmt.Println(sli, tab)

	sla := make([]int, len(sli))
	copy(sla, sli)

	sla[1] = 8
	sla = append(sla, 10)
	fmt.Println(sli, sla)

	// sla2 := make([]int, len(sli))
	// copy(sla, tab)


	fmt.Printf("%v %T\n", os.Args, os.Args)
	fmt.Println(len(os.Args))

	fp, err := strconv.Atoi(os.Args[1])
	fmt.Printf("%T\n", fp)
	fmt.Println(fp, err)


	if p := 2; p < 5 {
		fmt.Println("if")
	} else {
		fmt.Println("else")
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	a := r.Intn(10)
	// a := 1
	b := 2

	// var w int

	fmt.Println(a)
	if w := a; a < b {
		w = b
		fmt.Println(w)
	}
	// fmt.Println(w)
}

func init() {
	if len(os.Args) != 4 {
		fmt.Println("USAGE: ...")
		os.Exit(-1)
	}
}
