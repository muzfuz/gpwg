package trip

import (
	"math/rand"
	"time"
)

const (
	distanceToMars = 62_100_000 // km
	minTravelSpeed = 16         // km/s
	maxTravelSpeed = 30         // km/s
	minPrice       = 36_000_000 // Currency
	costPerKMS     = 1_000_000  // Currency - cost of each km/s
)

var (
	r          = rand.New(rand.NewSource(time.Now().Unix()))
	spacelines = [3]string{"Space Adventures", "SpaceX", "Virgin Galactic"}
)

// Trip is a single instance of a journey to mars
type Trip struct {
	spaceline   string
	travelSpeed int
	roundTrip   bool
}

// New returns an instance of a single Trip,
// with all random values pre-populated.
func New() Trip {
	return Trip{
		spaceline:   randomSpaceline(),
		travelSpeed: randomTravelSpeed(),
		roundTrip:   randomBool(),
	}
}

// Spaceline returns which line the trip is with
func (t Trip) Spaceline() string {
	return t.spaceline
}

// DurationSeconds returns the length of time the trip will take
func (t Trip) DurationSeconds() int {
	return distanceToMars / t.travelSpeed
}

// Price defines how much the trip will cost.
// It is calculated based on travel speed, at a cost
// of costPerKMS for each additional km/s from baseline (minTravelSpeed)
func (t Trip) Price() int {
	additionalKMS := t.travelSpeed - minTravelSpeed
	price := minPrice + (additionalKMS * costPerKMS)
	if t.roundTrip {
		price = price * 2
	}
	return price
}

// Type returns a One-way or Round-trip string
func (t Trip) Type() string {
	if t.roundTrip {
		return "Round-trip"
	}
	return "One-way"
}

// randomSpaceline chooses one of the configured spacelines at random
func randomSpaceline() string {
	i := r.Intn(len(spacelines))
	return spacelines[i]
}

// randomTravelSpeed generates random int between minTravelSpeed and maxTravelSpeed
func randomTravelSpeed() int {
	return r.Intn(maxTravelSpeed-minTravelSpeed+1) + minTravelSpeed
}

// randomBool returns either true or false
func randomBool() bool {
	bools := []bool{true, false}
	i := r.Intn(len(bools))
	return bools[i]
}
