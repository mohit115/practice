package main

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
)

type locker struct {
	loakcerId  int
	allocated  bool
	orderId    int
	lockerType string
}

type lockerManager struct {
	smallLoackers      []*locker
	mediumLockers      []*locker
	largeLockers       []*locker
	xlargeLockers      []*locker
	allocatedLockerMap map[int]*locker
	lockerCount        int
}

var lockerManagerInstance *lockerManager

var atomicLock = &sync.Mutex{}

func getLockerManger() *lockerManager {
	if lockerManagerInstance == nil {
		atomicLock.Lock()
		defer atomicLock.Unlock()
		if lockerManagerInstance == nil {
			lockerManagerInstance = &lockerManager{smallLoackers: make([]*locker, 0), mediumLockers: make([]*locker, 0), largeLockers: make([]*locker, 0), xlargeLockers: make([]*locker, 0), lockerCount: 0, allocatedLockerMap: make(map[int]*locker)}
		}
	}
	return lockerManagerInstance
}

func (lm *lockerManager) addLocker(req <-chan map[string]string, resp chan<- map[string]string) {
	var tempLockerList *[]*locker
	for r := range req {
		m := make(map[string]string)
		m["id"] = r["id"]
		fmt.Println("adding lock ", r["type"])
		switch r["type"] {
		case "S":
			tempLockerList = &lm.smallLoackers
		case "M":
			tempLockerList = &lm.mediumLockers
		case "L":
			tempLockerList = &lm.largeLockers
		case "X":
			tempLockerList = &lm.xlargeLockers
		default:
			m["status"] = "failed"
			m["message"] = "invalid locker type"
			resp <- m
			continue
		}
		val, err := strconv.Atoi(r["count"])
		if err != nil {
			m["status"] = "failed"
			m["message"] = err.Error()
			resp <- m
			continue
		}
		for i := 0; i < val; i++ {
			lm.lockerCount++
			*tempLockerList = append(*tempLockerList, &locker{loakcerId: lm.lockerCount, allocated: false, orderId: 0, lockerType: r["type"]})

		}
		m["status"] = "ok"
		resp <- m
	}

}

func (lm *lockerManager) getStatus(req <-chan map[string]string, res chan<- map[string]string) {
	for r := range req {
		if r["request"] == "status" {
			payload := make(map[string]string)
			payload["id"] = r["id"]
			payload["SMALL"] = strconv.Itoa(len(lm.smallLoackers))
			payload["MEDIUM"] = strconv.Itoa(len(lm.mediumLockers))
			payload["LARGE"] = strconv.Itoa(len(lm.largeLockers))
			payload["XLARGE"] = strconv.Itoa(len(lm.xlargeLockers))
			payload["ALLOCATED"] = strconv.Itoa(len(lm.allocatedLockerMap))

			res <- payload
		}

	}

}

func (lm *lockerManager) getLocker(lockerType string) (*locker, error) {
	var lockerInstance *locker
	switch lockerType {
	case "S":
		if len(lm.smallLoackers) != 0 {
			lockerInstance = lm.smallLoackers[0]
			lm.smallLoackers = lm.smallLoackers[1:]
			break
		}
	case "M":
		if len(lm.mediumLockers) != 0 {
			lockerInstance = lm.mediumLockers[0]
			lm.mediumLockers = lm.mediumLockers[1:]
			break
		}
	case "L":
		if len(lm.largeLockers) != 0 {
			lockerInstance = lm.largeLockers[0]
			lm.largeLockers = lm.largeLockers[1:]
		}
	}
	if lockerInstance == nil {
		return nil, errors.New("locker not available ")
	}
	return lockerInstance, nil
}

var lockerAtomicLock = &sync.Mutex{}

func (lm *lockerManager) consumeLocker(lockerInstance *locker, oderId int) bool {
	if lockerInstance.allocated == false {
		lockerAtomicLock.Lock()
		defer lockerAtomicLock.Unlock()
		if lockerInstance.allocated == false {
			lockerInstance.orderId = oderId
			lockerInstance.allocated = true
			lm.allocatedLockerMap[lockerInstance.orderId] = lockerInstance
			return true
		}

	}
	return false
}
