package main

import (
	"fmt"
	"sync"
)

type Test struct {
	sync.RWMutex
	val int
}

func (t *Test) increment(wg *sync.WaitGroup) {
	// fmt.Println("WAITING FOR LOCK")
	t.Lock()
	// fmt.Println("GOT THE LOCK")
	// time.Sleep(time.Millisecond*100)
	defer t.Unlock()
	defer wg.Done()
	t.val++

}

func (t *Test) decrement(wg *sync.WaitGroup) {
	// fmt.Println("WAITING FOR LOCK")
	t.Lock()
	// fmt.Println("GOT THE LOCK")
	// time.Sleep(time.Millisecond*110)

	defer t.Unlock()
	defer wg.Done()
	t.val--

}

func (t *Test) print(i int) {
	// fmt.Println("WAITING FOR LOCK")
	// t.Lock()
	// fmt.Println("GOT THE LOCK")

	// time.Sleep(time.Millisecond*180)

	fmt.Println("the value is ", t.val)
	fmt.Println("the diff is  ", t.val - i )

}
func main() {
	t := &Test{}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	val :=10000000
	go func(*sync.WaitGroup) {
		
		for i:= 0; i <val; i++{
			wg.Add(4)
			go t.increment(wg)
			go t.decrement(wg)
			go t.increment(wg)
			go t.decrement(wg)
		}
		wg.Done()
	}(wg)

	wg.Wait()
	t.print(val)

}
