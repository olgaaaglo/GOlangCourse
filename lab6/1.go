package main 

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	//"crypto/sha256"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []*User  `xml:"user"`
}

type User struct {
	XMLName  xml.Name `xml:"user"`
	Login    string   `xml:"login"`
	Password string   `xml:"paswword"`
	Role     int      `xml:"role"`
}

type Person struct {
	XMLName    xml.Name  `xml:"person"`
	Id         int       `xml:"id"`
	FirstName  string    `xml:"firstName"`
	LastName   string    `xml:"lastName"`
	Age        int       `xml:"age"`
	Birth      time.Time `xml:"birth"`
	Death      time.Time `xml:"death"`
	Pesel      int       `xml:"pesel"`
	CreditCard int       `xml:"creditcard"`
	Gender     rune      `xml:"gender"`
}

type People struct {
	XMLName xml.Name  `xml:"persons"`
	People  []*Person `xml:"person"`
}

func decryptFile() {
	encrypted, err := ioutil.ReadFile("encrypted.xml")
	if err != nil {
		fmt.Println(err)
	}

	key, err := ioutil.ReadFile("key.txt")
	if err != nil {
		fmt.Println("read file err: %v", err.Error())
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("cipher err: %v", err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("cipher GCM err: %v", err.Error())
	}
	
	nonce := encrypted[:gcm.NonceSize()]
	encrypted = encrypted[gcm.NonceSize():]
	plainText, err := gcm.Open(nil, nonce, encrypted, nil)
	if err != nil {
		fmt.Println("decrypt file err: %v", err.Error())
	}
	
	err = ioutil.WriteFile("people.xml", plainText, 0777)
	if err != nil {
		fmt.Println("write file err: %v", err.Error())
	}
}

func addId(scanner *bufio.Scanner, person* Person) {
	fmt.Println("Enter id:")
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		id, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Wrong id value: ", err)
			return
		} else {
			person.Id = id
		}
		break
	}
}

func addFirstName(scanner *bufio.Scanner, person* Person) {
	fmt.Println("Enter first name:")
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		person.FirstName = text
		break
	}
}

func addBirth(scanner *bufio.Scanner, person* Person) {
	fmt.Println("Enter birth date:")
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		birth, err := time.Parse("2006-01-02", text)
		if err != nil {
			fmt.Println("Wrong birth value: ", err)
			return
		} else {
			person.Birth = birth
		}
		break
	}
}

func addGender(scanner *bufio.Scanner, person* Person) {
	fmt.Println("Enter gender:")
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		genderArr := []rune(text)
		gender := genderArr[0]
		
		if gender != 102 && gender != 109 {
			fmt.Println("Wrong gender value %s ", gender)
			return
		} else {
			person.Gender = gender
		}
		break
	}
}

func addPesel(scanner *bufio.Scanner, person* Person) {
	fmt.Println("Enter PESEL:")
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		pesel := string(text[:])
		year, month, day := person.Birth.Date()
		
		if string(pesel[0]) != strconv.Itoa((year % 100 - year % 10) / 10) ||
			string(pesel[1]) != strconv.Itoa(year % 10) {
			fmt.Println("Year doesn't match")
			// return
		}
		if string(pesel[2]) != strconv.Itoa(int(month) / 10) ||
			string(pesel[3]) != strconv.Itoa(int(month) % 10) {
			fmt.Println("Month doesn't match")
			// return
		}
		if string(pesel[4]) != strconv.Itoa(day / 10) ||
			string(pesel[5]) != strconv.Itoa(day % 10) {
			fmt.Println("Day doesn't match")
			// return
		}

		gender := pesel[6:10]
		// fmt.Println(gender)
		gen, err := strconv.Atoi(gender)
		if err != nil {
			fmt.Println("Wrong gender value")
		}
		if (gen % 2 == 0 && person.Gender != 102) ||
			(gen % 2 == 1 && person.Gender != 109) {
			fmt.Println("Gender doesn't match")
		}

		
		// peselInt, err := strconv.Atoi(pesel)
		// if err != nil {
		// 	fmt.Println("Wrong pesel value")
		// }

		// weights := []int{1, 3, 7, 9, 1, 3, 7, 9, 1, 3}
		// controlSum := 0
		// for i, w := range weights {
		// 	controlSum += w * (peselInt % 
		// }

		// if peselInt[10] != (10 - controlSum % 10) % 10 {
		// 	fmt.Println("Wrong control sum")
		// }
		// person.Pesel = pesel
		
		break
	}
}

func addToDatabase() {
	var person Person
	
	scanner := bufio.NewScanner(os.Stdin)
	
	// addId(scanner, &person)
	// addFirstName(scanner, &person)
	addBirth(scanner, &person)
	addGender(scanner, &person)
	addPesel(scanner, &person)
	fmt.Println("Person: ", person)
	
}

func main() {
	if _, err := os.Stat("users.xml"); errors.Is(err, os.ErrNotExist) {
		file, err2 := os.Create("users.xml")
		if err2 != nil {
			fmt.Println(err2)
		} else {
			
		}
		defer file.Close()
	} else {
		file, err2 := os.Open("users.xml")
		if err2 != nil {
			fmt.Println(err2)
		} else {

		}
		defer file.Close()
	}

	decryptFile()
	addToDatabase()
}