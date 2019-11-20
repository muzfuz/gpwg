package trip

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// For each ticket select one of the three:
// -- Space Adventures
// -- SpaceX
// -- Virgin Galactic
//
// Mars is 62,100,000 KM from earth on the departure date
// The ship will travel between 16-30 km/s
// Journey will take at least 23.95 days (2,070,000 seconds)
// Journey will take at most 44.95 days (3,881,250) seconds
// Faster ships are more expensive
// Price ranges from 36 million to 50 million
// Double price for round trips
//

func Test_Spaceline(t *testing.T) {
	trip := New()
	assert.Contains(t, spacelines, trip.Spaceline(), "Trip spaceline should contain one of: %v", spacelines)
}

func Test_DurationSeconds(t *testing.T) {
	trip := Trip{travelSpeed: maxTravelSpeed}
	minDuration := distanceToMars / maxTravelSpeed
	assert.Equal(t, minDuration, trip.DurationSeconds(), "The fastest trip takes %d seconds", minDuration)

	trip = Trip{travelSpeed: minTravelSpeed}
	maxDuration := distanceToMars / minTravelSpeed
	assert.Equal(t, maxDuration, trip.DurationSeconds(), "The slowest trip takes %d seconds", maxDuration)
}

func Test_Price(t *testing.T) {
	trip := Trip{travelSpeed: minTravelSpeed}
	assert.Equal(t, minPrice, trip.Price(), "The slowest trip is cheapest")

	maxPrice := minPrice + ((maxTravelSpeed - minTravelSpeed) * costPerKMS)
	trip = Trip{travelSpeed: maxTravelSpeed}
	assert.Equal(t, maxPrice, trip.Price(), "The fastest trip is the most expensive")
}

func Test_Price_RoundTrip(t *testing.T) {
	trip := Trip{travelSpeed: minTravelSpeed, RoundTrip: true}
	assert.Equal(t, minPrice*2, trip.Price(), "A return trip costs twice as much")
}
