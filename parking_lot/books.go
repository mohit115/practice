package main

import "sync"

type Item struct {
	id       int32
	itemType string
}

type Book struct {
	Item
	name     string
	price    float32
	quantity int
	lock     *sync.Mutex
}

func Getbooks(name string, count int) float32 {
	// bs := GetInstance[Book]("books")
}
