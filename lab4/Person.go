package main

import (
	"fmt"
	"encoding/json"
	"errors"
)

type Person struct {
	Id int
	Name string
	Age int
	Education string
}

func (person Person) ToString() string {
	return fmt.Sprintf("Person: id=%d, name=%s, age=%d, ed=%s", person.Id, person.Name, person.Age, person.Education);
}

type PeopleStruct struct {
	People []*Person
}

func (people* PeopleStruct) createElements(data []byte, count int) error {
	err := json.Unmarshal(data[:count], people)
    if err != nil {
		return err
    }
	if len(people.People) == 0 {
		return errors.New("Error creating people")
	}
	fmt.Println("Create People: ")
	for _, p := range people.People {
		fmt.Println(*p)
	}
	return nil
}

func (people PeopleStruct) ToString() string {
	ret := "People: "
	for _, p := range people.People {
		ret += p.ToString() + "\n        "
	}
	return ret
}
