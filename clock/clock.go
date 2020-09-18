package clock

import "strconv"

// Clock struct definition for clock
type Clock struct {
	hour   int
	minute int
}

// New returns a new Clock
func New(hour int, minute int) Clock {
	clock := Clock{
		hour:   0,
		minute: 0,
	}
	if hour < 0 {
		clock = clock.subHour(-1 * hour)
	} else {
		clock = clock.addHour(hour)
	}
	if minute < 0 {
		clock = clock.Subtract(-1 * minute)
	} else {
		clock = clock.Add(minute)
	}
	return clock
}

// String string representation of clock
func (c Clock) String() string {
	return getPaddedString(c.hour) + ":" + getPaddedString(c.minute)
}

// getPaddedString returns string value of n with 2 digit formatting
func getPaddedString(n int) string {
	if n > 9 {
		return strconv.Itoa(n)
	}
	return "0" + strconv.Itoa(n)
}

// Add adds minutes in clock
func (c Clock) Add(minutes int) Clock {
	c.minute = c.minute + minutes
	hourAddition := 0
	for c.minute > 59 {
		c.minute -= 60
		hourAddition++
	}
	c = c.addHour(hourAddition)
	return c
}

// Subtract subtracts minutes in clock
func (c Clock) Subtract(minutes int) Clock {
	c.minute = c.minute - minutes
	hourSub := 0
	for c.minute < 0 {
		c.minute += 60
		hourSub++
	}
	c = c.subHour(hourSub)
	return c
}

// addHour adds hour in clock
func (c Clock) addHour(hour int) Clock {
	c.hour = (c.hour + hour) % 24
	return c
}

// subHour substracts hour in clock
func (c Clock) subHour(hour int) Clock {
	c.hour = (c.hour - hour)
	for c.hour < 0 {
		c.hour += 24
	}
	return c
}
