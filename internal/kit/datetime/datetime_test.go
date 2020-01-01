package datetime

import (
	"testing"

	"github.com/jonboulle/clockwork"
)

func TestGetCurrentDateTime(t *testing.T) {
	c := clockwork.NewFakeClock()
	type args struct {
		clock clockwork.Clock
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testing get current date time package",
			args: args{
				clock: c,
			},
			want: c.Now().Format(DateAndTimeLayout),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCurrentDateTime(tt.args.clock); got != tt.want {
				t.Errorf("GetCurrentDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
