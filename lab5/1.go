package main 

import (
	"bufio"
	"io"
	"fmt"
	"flag"
	"log"
	"os"
	"strings"
)

var readFile = flag.Bool("rf", false, "readFile")
var removeDuplicate = flag.Bool("rd", false, "removeDuplicate") //words
var filterBy = flag.String("fb", "", "filterBy")

func readAndTransform(reader io.Reader)(string) {

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		log.Println("Inputed text = ", text)
		sli := strings.Split(text, " ")

		log.Println("removeDuplicate = ", *removeDuplicate)
		if *removeDuplicate {
			begin:
			for i, s1 := range sli {
				for j, s2 := range sli {
					if i != j {
						if strings.Compare(s1, s2) == 0 {
							sli = append(sli[:j], sli[j+1:]...)
							goto begin
						}
					}
				}
			}
			log.Println("After removeDuplicate: ", sli)
		}

		log.Println("filterBy = ", *filterBy)
		if *filterBy != "" {
			for i, s1 := range sli {
				if !strings.HasPrefix(s1, *filterBy) {
					sli = append(sli[:i], sli[i+1:]...)
				}
			}
			log.Println("After filterBy ", *filterBy, ": ", sli)
		}

		textNew := strings.Join(sli, " ")
		return textNew
	}
	return ""
}

func init() {
	log.SetPrefix("LOG: ")
}

func main() {
	flag.Parse()
	if len(os.Args) == 1 {
		log.Println("Flags set to default values")
	}
	log.Println("readFile = ", *readFile)

	if *readFile {
		file, _ := os.Open("file.txt")
		log.Println("After readAndTransform: ", readAndTransform(file))
	} else {
		log.Println("After readAndTransform: ", readAndTransform(os.Stdin))
	}
	
}