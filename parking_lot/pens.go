package main

import "sync"

type Pen[] struct {
	Item
	name     string
	price    float32
	quantity int
	lock     *sync.Mutex
}

func (p *Pen) hold() {
	p.lock.Lock()
}
func (p *Pen) release() {
	p.lock.Unlock()
}

func GetPens(name string, count int) float32 {
	ci := GetInstance[Pen]("pen")
	ci.GetValue(name)

}
