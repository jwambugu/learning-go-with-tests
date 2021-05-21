package main

import (
	"github.com/jwambugu/learning-go-with-tests/pkg/mocking"
	"os"
)

func main() {
	sleeper := &mocking.DefaultSleeper{}

	mocking.Countdown(os.Stdout, sleeper)
}
