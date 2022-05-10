package main

import (
	"fmt"
	"sync"
)

type booking struct {
	sync.Mutex
	bookingId string
	userId    string
	cabId     string
}

type bookingManager struct {
	bookings map[int]*booking
}

var bmlock = &sync.Mutex{}

var singletonBookingManagerInst *bookingManager

func (bm *bookingManager) book(uid string) bool {
	return false
}

func (bm *bookingManager) getStatus() {

}

func getBookingManagerIns() *bookingManager {

	if singletonBookingManagerInst == nil {
		bmlock.Lock()
		defer bmlock.Unlock()
		if singletonBookingManagerInst == nil {
			singletonBookingManagerInst = &bookingManager{bookings: make(map[int]*booking, 0)}
		}

	}
	return singletonBookingManagerInst
}

func (bm *bookingManager) getBookingById(id int) *booking {
	if b, ok := bm.bookings[id]; ok {
		return b
	} else {
		fmt.Printf("Booking %d not found\n", id)
	}
	return nil
}
