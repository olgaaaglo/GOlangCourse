package main 

import (
	"fmt"
)


func main() {
	sli, err := split("aaa/bbb/ccc", "/")
	fmt.Println(sli, err)

	fmt.Println(split("aaa bbb ccc"))
	fmt.Println(split("aaa bbb ccc", "", " "))

	fmt.Println(splitAndSort("bbb ccc aaa"))

	fmt.Println(getNumberOfNonEmptyLines("aaa\n\nbbb\nccc\n\nd"))

	fmt.Println(getFrequencyForWords("aaa bbb aaa ccc bbb aaa"))

	fmt.Println(getSortedPalindroms("bccb cda aaa aba dcdc"))
}