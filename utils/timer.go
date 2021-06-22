package utils

import "time"

type Timer struct {
	tms time.Time
	tmd time.Duration
}

func NewTimer(tmd time.Duration, now ...bool) *Timer {
	c := &Timer{
		tmd: tmd,
	}
	if len(now) > 0 && now[0] {
		c.tms = time.Now()
	}
	return c
}
func (c *Timer) Tick() bool {
	if time.Since(c.tms) < c.tmd {
		return false
	} else {
		c.tms = time.Now()
		return true
	}
}
func (c *Timer) Reset(now bool, tmd ...time.Duration) {
	if len(tmd) > 0 {
		c.tmd = tmd[0]
	}
	if now {
		c.tms = time.Now()
	} else {
		c.tms = time.Time{}
	}
}
