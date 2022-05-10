package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

var updateChannel = make(chan map[string]interface{})

var centralSysterm = make(map[string]interface{})

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s", name, elapsed)
}
func updateHandler(id string, val interface{}) {
	update := make(map[string]interface{})
	update[id] = val
	updateChannel <- update
}

func udpateWatcher() {
	i := 0
	for message := range updateChannel {
		for k, v := range message {
			centralSysterm[k] = v
			fmt.Printf("count %d  UPDATE for id : %s has now status %+v\n", i, k, v)
			i++
		}
	}
}

func main() {
	defer timeTrack(time.Now(), "factorial")
	wg := &sync.WaitGroup{}
	CashierObj := &Cashier{}
	DoctorObj := &Doctor{}
	MedicineObj := &Medicine{}
	ReceptionObj := &Reception{}
	ReceptionObj.setNext(DoctorObj).setNext(MedicineObj).setNext(CashierObj)
	patientList := make([]*Patient, 0)

	for i := 0; i < 5; i++ {
		patientList = append(patientList, &Patient{medicineDone: false, doctorCheckUpDone: false, paymentDone: false, id: uuid.New().String(), registrationDone: false})
	}

	for i := range patientList {
		wg.Add(1)
		go func(i *Patient) {
			ReceptionObj.execute(i, updateHandler)
			wg.Done()
		}(patientList[i])
	}

	go udpateWatcher()
	wg.Wait()
	close(updateChannel)

}
