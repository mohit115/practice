package main

import (
	"errors"
	"fmt"
	"math"
)

type catagory struct {
	id   int32
	name string
}

type post struct {
	id        int32
	name      string
	catagorys []catagory
	title     string
	body      string
}

type cacheable interface {
	post | catagory
}

type cache[t cacheable] struct {
	data map[string]t
}

func (c *cache[t]) set(key string, val t) {
	c.data[key] = val
}

func (c *cache[t]) get(key string) (t, error) {
	v, ok := c.data[key]
	if ok {
		return v, nil
	}
	return v, errors.New("empty name")
}

func main() {
	catagoryCache := cache[catagory]{data: make(map[string]catagory)}
	postCache := cache[post]{data: make(map[string]post)}
	catagoryObj := catagory{id: 1, name: "cat1"}
	postObj := post{id: 1, name: "post1", catagorys: []catagory{catagoryObj}}
	postCache.set("first", postObj)
	catagoryCache.set("first", catagoryObj)
	{
		val, err := catagoryCache.get("first1")
		if err == nil {
			fmt.Println(val)
		} else {
			fmt.Println("not found")
		}
	}

	{
		val, err := postCache.get("first")
		if err == nil {
			fmt.Println(val)
		} else {
			fmt.Println("not found")
		}
	}
	numbers := []float64{4, 9, 16, 25}
	fmt.Println(createMap(numbers, math.Sqrt))
	fmt.Println(createMap([]float64{4, 9, 16, 25}, func(k float64) float64 { return k * k * k }))

}
