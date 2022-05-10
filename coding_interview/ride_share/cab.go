package main

import "sync"

type cabi interface {
	isAvailable() bool
	startRide()
	completeRide()
	clearBooking()
	book(chan int, *sync.WaitGroup) bool
	getId() int
}

type cab struct {
	sync.Mutex
	capacity int
	cabId    int
	bookings []*booking
}

func (c *cab) book(ch <-chan int) {
	bm := getBookingManagerIns()
	rid := <- ch 
	booking := bm.getBookingById(rid)
	user := bm.getBookingById(booking.userId)
	// use  fan out such that each car will give it the  effinity score

}
func (c *cab) getId() int {
	return c.cabId
}

func (c *cab) clearBooking() {
	c.bookings = make([]*booking, 0, 4)
}

func (c *cab) completeRide() {
	if c.isAvailable() == false {
		c.Lock()
		defer c.Unlock()
		if c.isAvailable() == false {
			c.clearBooking()

		}
	}

}

func (c *cab) startRide() {
	if c.isAvailable() == false {
		c.Lock()
		defer c.Unlock()
		if c.isAvailable() == false {
			c.clearBooking()
		}
	}

}

func (c *cab) isAvailable() bool {
	return c.capacity > len(c.bookings)

}

type cabManager struct {
	cabs map[int]*cab
}

var cmlock = &sync.Mutex{}
var singletonCabManagerInst *cabManager

func (cm *cabManager) getCabById(id int) *cab {
	return cm.cabs[id]
}

func (cm *cabManager) registerCab(cab *cab) {
	cm.cabs[cab.getId()] = cab

}

func getCabMangerIns() *cabManager {
	if singletonCabManagerInst == nil {

		cmlock.Lock()
		defer cmlock.Unlock()
		if singletonCabManagerInst == nil {
			singletonCabManagerInst = &cabManager{cabs: make(map[int]*cab, 0)}
		}

	}
	return singletonCabManagerInst
}
