package mars

import (
	"fmt"

	"github.com/muzfuz/gpwg/mars/trip"
)

// GenerateTickets will generate a list of n length random tickets
func GenerateTickets(n int) string {
	ticketsStr := ""
	ticketsStr += "Spaceline        Days Trip type  Price\n"
	ticketsStr += "======================================\n"
	for i := 0; i < n; i++ {
		ticketsStr += ticket()
	}
	return ticketsStr
}

// ticket formats each ticket and presents
// the raw numbers in a human readable format
func ticket() string {
	t := trip.New()
	return fmt.Sprintf(
		"%-18s %d %-10s $%4d\n",
		t.Spaceline(),
		t.DurationSeconds()/86400, // number of seconds in a day
		t.Type(),
		t.Price()/1_000_000, // price is in millions
	)
}
