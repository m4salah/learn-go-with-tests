package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

var (
	finalWord  = "Go!"
	countStart = 3
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	timeToSleep := func() {
		time.Sleep(1 * time.Second)
	}
	for i := countStart; i > 0; i-- {
		timeToSleep()
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	timeToSleep()
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
func main() {
	defaultSleeper := &DefaultSleeper{}
	Countdown(os.Stdout, defaultSleeper)
}
