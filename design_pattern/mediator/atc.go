package main

import "sync"

type atc interface {
	canLand(flyingObj) bool
	canTakeoff(flyingObj) bool
	clearRunway() bool
}

type variable int

const (
	propalerPlane variable = iota
	helicoptor
	jetplane
	cargoplane
)

type flyingObj interface {
	id() string
	vehicleType() variable
	needLandingSpace() float32
	needTakeOffSpace() float32
	hasAltitude() float32
	hasEmergency() bool
	source() string
	destination() string
}

type flight struct {
	id string

	vehicleType  variable
	landingSpace float32
	takeoffSpace float32
	altitude     float32
	emergency    bool
	source       string
	destination  string
}

type runway struct {
	sync.Mutex
	id            int
	runway_length float32
	available     bool
	flight        flyingObj
}

func (r *runway) setAvailability(flag bool) {
	r.available = flag
}

func (r *runway) getAvailability() bool {
	return r.available
}

func (r *runway) clearRunway() {
	r.flight = nil
}

func (r *runway) hasFlight() flyingObj {
	return r.flight
}

type blrAtc struct {
	name         string
	runways      []*runway
	landingQueue []flyingObj
	takeoffQueue []flyingObj
}

func (a *blrAtc) canLand(f flyingObj) bool {

	for i, _ := range a.runways {
		if  a.runways[i].available {
			a.runways[i].Lock()
			defer  a.runways[i].Unlock()
			if a.runways[i].available {
				a.runways[i].setAvailability(false)
			}
			return true
		}
	}
	return false
}

func (a *blrAtc) canTakeoff(f flyingObj) bool {

	for i, _ := range a.runways {
		if a.runways[i].available {
			a.runways[i].Lock()
			defer a.runways[i].Unlock()
			if a.runways[i].available {
				a.runways[i].setAvailability(false)
			}
			return true
		}
	}
	return false
}

func (a *blrAtc) clearRunway(f flyingObj) {
	for i, _ := range a.runways {
		if f.id() == a.runways[i].flight.id() {
			a.runways[i].flight = nil
			a.runways[i].available = true
		}
	}

}
