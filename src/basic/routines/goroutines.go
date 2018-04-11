package routines

import (
	"fmt"
	"time"
)

func printNumber(from, to int, c chan int) {
	for x := from; x <= to; x++ {
		fmt.Printf("%d\n", x)
		time.Sleep(1 * time.Millisecond)
	}
	c <- 0
}
