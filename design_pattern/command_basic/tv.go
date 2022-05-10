package main

import (
	"fmt"
	"math/rand"
)

type tv struct {
	power         bool
	channels      map[int]interface{}
	tunedCh       tunedChannel
	minVolume     int
	maxVolume     int
	currentVolume int
}

func (r *tv) on() {
	if r.power == false {
		r.power = true
		fmt.Println("power on the tv")

	}

}

func (r *tv) off() {
	if r.power == true {
		r.power = false
		fmt.Println("power on the tv")

	}

}

func (r *tv) tune(channelId int) {
	val, exist := r.channels[channelId]

	if exist == false {
		fmt.Println("channel does not exist")
	} else {
		if tc, ok := val.(tunedChannel); ok == true {
			r.tunedCh = tc
		} else {
			fmt.Println("channel does not exist")
		}

	}
}

func (r *tv) volumeUp() {
	if r.currentVolume < r.maxVolume {
		r.currentVolume++
		fmt.Println("volume is at", r.currentVolume)
	} else {
		fmt.Println("volume is maxed out at", r.currentVolume)
	}

}
func (r *tv) volumeDown() {
	if r.currentVolume > r.minVolume {
		r.currentVolume--
		fmt.Println("volume is at", r.currentVolume)
	} else {
		fmt.Println("minimum volume cannot be lower than ", r.currentVolume)
	}

}

func (r *tv) channelNext() {
	r.tune(r.tunedCh.channelNumber + 1%len(r.channels))

}

func (r *tv) channelPrev() {
	if r.tunedCh.channelNumber == 0 {
		r.tune(len(r.channels) - 1)
	} else {
		r.tune(r.tunedCh.channelNumber - 1)
	}

}
func (r *tv) scan() {
	fmt.Println("scanning the channels in radio")

	for i := 0; i < 100; i++ {
		r.channels[i] = &tunedChannel{channelNumber: i, channelName: fmt.Sprintf("%s_%d", "Channel", i), channelFrequency: rand.Float32()}
		fmt.Printf("%+v\n", r.channels[i])

	}

}
