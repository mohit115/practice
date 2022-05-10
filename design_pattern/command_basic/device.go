package main

type device interface {
	on()
	off()
	tune(int)
	scan()
	volumeUp()
	volumeDown()
	channelNext()
	channelPrev()
}
type tunedChannel struct {
	channelNumber    int
	channelFrequency float32
	channelName      string
}
