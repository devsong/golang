package routines

import (
	"testing"
)

func TestRoutines(t *testing.T) {
	c := make(chan int, 3)
	go printNumber(1, 10, c)
	go printNumber(11, 20, c)
	_, _ = <-c, <-c
}
