package syntax

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	arr := []int{1, 2, 3}
	array(arr)
	reverseArray(arr)
	for _, val := range arr {
		fmt.Printf("%v ", val)
	}
}
