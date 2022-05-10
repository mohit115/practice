package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func test(c chan string, resp chan string) {
	for r := range c {
		res, err := http.Get(r)
		if err != nil {
			fmt.Println("No response from request")
			return
		}

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body) // response body is []byte
		if err != nil {
			fmt.Println("No able to parse the response")
		} else {
			resp <- string(body)
		}

	}

}

func main() {

	regularMap := make(map[int]interface{}, 0)
	syncMap := sync.Map{}
	wg := &sync.WaitGroup{}
	for i :=0 ; i<100;i++{
		wg.Add(1)
		go func (i int)  {
			regularMap[0]=i
			defer wg.Done()	
		}(i)
	}
	wg.Wait()
}
