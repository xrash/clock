package clock

import (
	"time"
)

type Clock struct {
	accumulator time.Duration
	lastPoint   time.Time
	stopped     bool
}

func NewClock() *Clock {
	return &Clock{
		stopped: true,
	}
}

func (c *Clock) Start() {
	if !c.stopped {
		return
	}

	c.lastPoint = time.Now()
	c.stopped = false
}

func (c *Clock) Stop() {
	if c.stopped {
		return
	}

	c.accumulator += time.Since(c.lastPoint)
	c.stopped = true
}

func (c *Clock) Reset() {
	c.accumulator = 0
	c.stopped = true
}

func (c *Clock) Now() time.Duration {
	if c.stopped {
		return c.accumulator
	}

	return c.accumulator + time.Since(c.lastPoint)
}
