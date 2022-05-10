package main

func main() {
	blr := &blrAtc{name: "blr", runways: make([]*runway, 0), landingQueue: make([]flyingObj,0), takeoffQueue: make([]flyingObj, 0)}
	blr.runways = append(blr.runways, &runway{id: 1, runway_length: 5000, available: true,flight: nil })
	blr.runways = append(blr.runways,&runway{id: 2, runway_length: 4500, available: true,flight: nil } )

	flight = &flight{vehicleType: jetplane,landingSpace: 3000, takeoffSpace: 3500, altitude: 0, emergency: false,source: }

}
