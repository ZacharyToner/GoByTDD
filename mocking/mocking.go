package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

//Sleeper is a custom sleep interface
type Sleeper interface {
	Sleep()
}

//ConfigurableSleeper allows users to modify the durations
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

//Sleep calls the configured sleep function using the configured duration
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

//Countdown sends a 3 second countdown and Go! to output
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
