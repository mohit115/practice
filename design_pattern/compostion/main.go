package main

import "fmt"

type container interface {
	getSize() int32
	iterate()
}

type file struct {
	name string
}

type folder struct {
	name    string
	objects []container
}

func (c *folder) add(o container) {
	c.objects = append(c.objects, o)
}

func (c *folder) iterate() {
	fmt.Println("----->", c.name)
	for _, f := range c.objects {
		f.iterate()
	}
}

func (c *folder) getSize() int32 {
	var size int32
	for _, f := range c.objects {
		size += f.getSize()

	}
	return size

}
func (c *file) getSize() int32 {
	return int32(len(c.name) * 1000)

}
func (c *file) iterate() {
	fmt.Println(c.name)

}

func main() {
	fo1 := &folder{name: "folder1", objects: make([]container, 0)}
	for i := 0; i < 10; i++ {
		v := &folder{name: fmt.Sprintf("%s_%d", "folder", i)}
		v.add(&file{name: fmt.Sprintf("%s_%d", "file", i)})
		fo1.add(v)
		for j := 0; j < 10; j++ {
			v := &folder{name: fmt.Sprintf("%s_%d_%d", "folder", i, j)}
			v.add(&file{name: fmt.Sprintf("%s_%d_%d", "file",i,j)})
			fo1.add(v)
		}

	}
	fmt.Println(fo1.getSize())
	fo1.iterate()
}
