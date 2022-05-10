package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type singleton struct {
	obj int
}

var lock = &sync.Mutex{}

var singletonInst *singleton

func getInstance() *singleton {

	if singletonInst == nil {
		lock.Lock()
		defer lock.Unlock()
		if singletonInst == nil {
			singletonInst = &singleton{obj: rand.Intn(100)}
		}

	}
	fmt.Println("singleton is already creted ")
	return singletonInst
}

func test(c chan int) {
	for i := 0; i < 100; i++ {
		c <- getInstance().obj
	}
	close(c)

}

func main() {
	c := make(chan int)
	go test(c)
	go test(c)
	go test(c)
	go test(c)

	k := 1
	for i := range c {
		k++
		fmt.Println(i, k)
	}

}
