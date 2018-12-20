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

//DefaultSleeper is a wrapper around time.Sleep
type DefaultSleeper struct{}

//Sleep for DefaultSleeper that uses built-in time Sleep
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
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
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
