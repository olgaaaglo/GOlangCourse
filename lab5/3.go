package main 

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"flag"
	"math/rand"
	"time"
	"sort"
)

var seedMath = flag.Bool("sm", true, "seedMath")
var min = flag.Int("min", 0, "min")
var max = flag.Int("max", 10, "max")

type Range struct {
	min int 
	max int 
}

func main() {
	flag.Parse()
	rangeForRand := Range{*min, *max}

	histogram := make(map[int] int, 0)
	i := 0
	for i < 100 {
		if *seedMath {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			randVal := rangeForRand.min + r.Intn(rangeForRand.max - rangeForRand.min)
			histogram[randVal]++
		} else {
			var b [8]byte
			var randVal int
			if _, err := crand.Read(b[:]); err != nil {
				randVal = 0
			}
			randVal = int(binary.LittleEndian.Uint32(b[:]))
			fmt.Println(randVal)
			histogram[randVal]++
		}
		i++
	}

	fmt.Println(histogram)

	keys := make([]int, 0)
 
    for k := range histogram{
        keys = append(keys, k)
    }
    sort.Ints(keys)
 
    for _, k := range keys {
        fmt.Println(k, histogram[k])
    }
}