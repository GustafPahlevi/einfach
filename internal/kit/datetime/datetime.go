package datetime

import (
	"github.com/jonboulle/clockwork"
)

const DateAndTimeLayout = "2006-01-02 15:04:05"

func GetCurrentDateTime(clock clockwork.Clock) string {
	return clock.Now().Format(DateAndTimeLayout)
}
