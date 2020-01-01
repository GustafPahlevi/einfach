package datetime

import (
	"github.com/jonboulle/clockwork"
)

// DateAndTimeLayout constant date and time layout for YYYY-DD-MM hh:mm:ss
const DateAndTimeLayout = "2006-01-02 15:04:05"

// GetCurrentDateTime is a kit for gen current date time
func GetCurrentDateTime(clock clockwork.Clock) string {
	return clock.Now().Format(DateAndTimeLayout)
}
