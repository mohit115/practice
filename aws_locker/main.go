package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

)

func customAddHandler(req chan<- map[string]string, resp <-chan map[string]string) func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hii there")
	return func(w http.ResponseWriter, r *http.Request) {
		var i = 0
		payload := make(map[string]string)
		payload["type"] = r.URL.Query().Get("type")
		payload["count"] = r.URL.Query().Get("count")
		payload["id"] = uuid.New().String()
		fmt.Println("<----- pushing messages for id ", i, "   ", payload["id"])
		i++
		req <- payload

		for message := range resp {
			result := make(map[string]string)
			result["requested_id"] = payload["id"]
			result["response_id"] = message["id"]
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(result)
		}
	}

}

func customStatusHandler(req chan<- map[string]string, resp <-chan map[string]string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqPayload := make(map[string]string)
		reqPayload["id"] = uuid.New().String()
		reqPayload["request"] = "status"
		req <- reqPayload

		for respayload := range resp {
			if respayload["id"] == reqPayload["id"] {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(respayload)
				return
			}

		}

	}
}

func main() {
	lm := getLockerManger()
	go lm.addLocker(getComChannelsInstace().addWorkerRequest, getComChannelsInstace().addWorkerResponse)
	go lm.addLocker(getComChannelsInstace().addWorkerRequest, getComChannelsInstace().addWorkerResponse)
	go lm.addLocker(getComChannelsInstace().addWorkerRequest, getComChannelsInstace().addWorkerResponse)
	go lm.addLocker(getComChannelsInstace().addWorkerRequest, getComChannelsInstace().addWorkerResponse)
	go lm.addLocker(getComChannelsInstace().addWorkerRequest, getComChannelsInstace().addWorkerResponse)
	go lm.addLocker(getComChannelsInstace().addWorkerRequest, getComChannelsInstace().addWorkerResponse)
	go lm.addLocker(getComChannelsInstace().addWorkerRequest, getComChannelsInstace().addWorkerResponse)

	go lm.getStatus(getComChannelsInstace().getStatusRequest, getComChannelsInstace().getStatusResponse)

	customAddLoackerHandlerFunction := customAddHandler(getComChannelsInstace().addWorkerRequest, getComChannelsInstace().addWorkerResponse)
	customGetStatusHandlerFunction := customStatusHandler(getComChannelsInstace().getStatusRequest, getComChannelsInstace().getStatusResponse)

	http.HandleFunc("/add", customAddLoackerHandlerFunction)
	http.HandleFunc("/status", customGetStatusHandlerFunction)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
