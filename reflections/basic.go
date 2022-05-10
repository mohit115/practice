package main

import (
	"fmt"
	"reflect"
)

var x interface{}

func main() {
	x = &struct{ name string }{}
	t0 := reflect.TypeOf(x)
	v0 := reflect.ValueOf(x)
	fmt.Println(t0, v0, t0.Kind())
	// y := [5]int{}
	// t0 = reflect.TypeOf(y)
	// v0 = reflect.ValueOf(y)
	// fmt.Println(t0, v0, t0.Kind())
	// y := make([]int, 3)
	var z [10]int
	t0 = reflect.TypeOf(z)
	v0 = reflect.ValueOf(z)
	fmt.Println(t0, v0, t0.Kind())
}
