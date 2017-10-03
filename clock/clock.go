package clock

import (
	"fmt"
)

const testVersion = 4

type Clock struct {
	hour, minute int
}

func New(hour, minute int) Clock {
	var totalMinutes int = (hour*60 + minute) % (24 * 60)
	if totalMinutes < 0 {
		totalMinutes += 24 * 60
	}
	return Clock{hour: totalMinutes / 60, minute: totalMinutes % 60}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}
