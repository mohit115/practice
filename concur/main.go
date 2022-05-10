package main

import (
	"fmt"
	"sync"
	"time"
)

func f1(x int, res chan int) {
	time.Sleep(time.Second * 2)
	res <- x
}

func main() {
	wg := &sync.WaitGroup{}
	res := make(chan int)
	l := make([]int, 0)
	go func() {
		//con
		//lis
		for i := range res {
			l = append(l, i)
		}
	}()
	for i := 0; i < 100; i++ {
		go func(i int) {
			//prod
			wg.Add(1)
			f1(i, res)
			wg.Done()
		}(i)

	}
	wg.Wait()
	close(res)
	fmt.Println(l)
	fmt.Print(len(l))
}
