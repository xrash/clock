package clock

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestClock(t *testing.T) {
	c := NewClock()
	assert.Equal(t, 0.0, c.Now().Seconds(), "New clock should have 0 seconds.")

	c.Stop()
	assert.Equal(t, 0.0, c.Now().Seconds(), "Stopped new clock should have 0 seconds.")

	c.Stop()
	c.Stop()
	c.Stop()
	assert.Equal(t, 0.0, c.Now().Seconds(), "Multiple times stopped new clock should have 0 seconds.")

	c.Start()
	time.Sleep(time.Second * 2)
	c.Stop()
	assert.Equal(t, 2.0, round(c.Now().Seconds()), "Clock should be around 2s.")

	time.Sleep(time.Second * 2)
	assert.Equal(t, 2.0, round(c.Now().Seconds()), "Stopped clock should keep stopped.")

	c.Start()
	time.Sleep(time.Second * 1)
	assert.Equal(t, 3.0, round(c.Now().Seconds()), "Nonstopped clock should run correctly.")

	time.Sleep(time.Second * 1)
	assert.Equal(t, 4.0, round(c.Now().Seconds()), "Nonstopped clock should keep running correctly.")

	c.Start()
	time.Sleep(time.Second * 1)
	assert.Equal(t, 5.0, round(c.Now().Seconds()), "Multiples times started clock should be OK.")

	c.Stop()
	time.Sleep(time.Second * 1)
	assert.Equal(t, 5.0, round(c.Now().Seconds()), "To stop again should be OK.")

	c.Stop()
	time.Sleep(time.Second * 1)
	assert.Equal(t, 5.0, round(c.Now().Seconds()), "To stop again multiple times should be OK.")
}
