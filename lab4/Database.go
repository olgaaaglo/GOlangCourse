package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Database struct {
	elements []Elements
	m map[string]int
	files []string 
}

func (database* Database) initialize(objs ...Elements) {
	for _, obj := range objs {
		database.elements = append(database.elements, obj)
	}
	database.m = make(map[string]int, 0)
}

func (database* Database) fill(names ...string) {
	j := 0
	for _, name := range names {
		database.files = append(database.files, name)
		file, err := os.Open(name)
		if err != nil {
			fmt.Print(err)
		}
		data := make([]byte, 1000)
		count, err := file.Read(data)
		if err != nil {
			fmt.Println("error:", err)
		}

		elements := database.elements[j]
		err2 := elements.createElements(data, count)
		if err2 == nil {
			database.m[fmt.Sprintf("%T", elements)] = j
		} else {
			fmt.Println("error: ", err2)
		}
		j++
		defer file.Close()
	}
	fmt.Println("Fill Database: ")
	for _, e := range database.elements {
		fmt.Println(e.ToString())
	}
}

func (database Database) print() {
	fmt.Println("Database: ")
	for i, v := range database.m {
		fmt.Println("    Type: ", i)
		fmt.Println("    Index: ", v)
		fmt.Println(database.elements[v].ToString())
	}
}

func (database* Database) addToDatabase(elements Elements) {
	fmt.Println("Input:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := scanner.Bytes()
		count := len(data) + 1

		err := elements.createElements(data, count)
		if err == nil {
			index := database.m[fmt.Sprintf("%T", elements)]
			elementsInDatabase := database.elements[index]
			
			switch t := (elements).(type) { 
			case *PeopleStruct:
				for _, val := range elements.(*PeopleStruct).People {
					elementsInDatabase.(*PeopleStruct).People = append(elementsInDatabase.(*PeopleStruct).People, val)
				}
			case *OffersStruct:
				for _, val := range elements.(*OffersStruct).Offers {
					elementsInDatabase.(*OffersStruct).Offers = append(elementsInDatabase.(*OffersStruct).Offers, val)
				}
			default:
				fmt.Printf("unexpected type %T", t)
			} 
		}	
		break
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func (database* Database) saveToFiles() {
	for i, el := range database.elements {
		json, _ := json.MarshalIndent(el, "", " ")
	
		file, err := os.OpenFile(database.files[i], os.O_CREATE|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Print(err)
		}
		_, err2 := file.Write(json)
		if err2 != nil {
			fmt.Println("error:", err2)
		}
		defer file.Close()
	}
}