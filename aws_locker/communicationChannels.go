package main

import "sync"

type communicatioManager struct {
	addWorkerRequest  chan map[string]string
	addWorkerResponse chan map[string]string

	getStatusRequest  chan map[string]string
	getStatusResponse chan map[string]string
}

var comChannelsLock = &sync.Mutex{}

var communicatioManagerInstance *communicatioManager

func getComChannelsInstace() *communicatioManager {
	if communicatioManagerInstance == nil {
		comChannelsLock.Lock()
		defer comChannelsLock.Unlock()
		if communicatioManagerInstance == nil {
			communicatioManagerInstance = &communicatioManager{addWorkerRequest: make(chan map[string]string), addWorkerResponse: make(chan map[string]string), getStatusRequest: make(chan map[string]string), getStatusResponse: make(chan map[string]string)}
		}

	}
	return communicatioManagerInstance
}
