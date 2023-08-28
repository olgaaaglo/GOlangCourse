package main

type Elements interface {
	createElements([]byte, int) error
	ToString() string
}