package clock

import "fmt"

const testVersion = 4
const day = 24 * 60

type Clock int

func (clock Clock) adjust() Clock {
	if clock < 0 {
		return clock + day
	}
	return clock
}

func New(hour, minute int) Clock {
	clock := Clock((hour*60 + minute) % day)
	return clock.adjust()
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock/60, clock%60)
}

func (clock Clock) Add(minutes int) Clock {
	return ((clock + Clock(minutes)) % day).adjust()
}
