package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	id int
	ins []int
	outs []int
}

func addNode(G* []Node, newNode Node) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i, _ := range *G {
		p := 0.6
		rp := float64(r.Intn(100)) / 100.0
		if rp < p {
			inOrOut := float64(r.Intn(100)) / 100.0
			if inOrOut < 0.5 {
				(*G)[i].outs = append((*G)[i].outs, newNode.id)
				newNode.ins = append(newNode.ins, (*G)[i].id)
			} else {
				(*G)[i].ins = append((*G)[i].ins, newNode.id)
				newNode.outs = append(newNode.outs, (*G)[i].id)
			}
		}
	}
	*G = append(*G, newNode)
	
}

func printGraph(G []Node) {
	fmt.Println("Graph:")
	for _, n := range G {
		fmt.Println("id: ", n.id, ", ins: ", n.ins, ", outs: ", n.outs)
	}
}

func histogram(G []Node) ([]int, []int) {
	size := len(G)
	histIn := make([]int, size)
	histOut := make([]int, size)

	for _, n := range G {
		histIn[len(n.ins)]++
		histOut[len(n.outs)]++
	}

	return histIn, histOut
}

func FloydWarshall(G []Node) [][]int {
	inf := 1000000
	N := len(G)
	A := make([][]int, N)
	d := make([][]int, N)
	for i := range A {
		A[i] = make([]int, N)
		d[i] = make([]int, N)
		for j := range A[i] {
			for _, out := range G[i].outs {
				if out == j {
					A[i][j] = 1
				}
			}
		}
	}

	for i := range A {
		for j := range A[i] {
			if i != j {
				if A[i][j] == 1 {
					d[i][j] = A[i][j]
				} else {
					d[i][j] = inf
				}
			}
		}
	}

	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if (d[i][k] + d[k][j]) < d[i][j] {
					d[i][j] = d[i][k] + d[k][j]
				}
			}
		}
	}

	return d
}

func main() {
	Graph := make([]Node, 0)

	for i := 0; i < 4; i++ {
		n := Node{i, []int{}, []int{}}
		addNode(&Graph, n)
		printGraph(Graph)
	}
	fmt.Println()

	histIn, histOut := histogram(Graph)
	fmt.Println("histIn: ", histIn)
	for i, n := range histIn {
		fmt.Println(n, " nodes has ", i, " ins")
	}
	fmt.Println("histOut: ", histOut)
	for i, n := range histOut {
		fmt.Println(n, " nodes has ", i, " outs")
	}
	fmt.Println()

	d := FloydWarshall(Graph)
	fmt.Println(d)
}