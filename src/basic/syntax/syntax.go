package syntax

import (
	"fmt"
)

func array(arr []int) {
	if arr == nil {
		return
	}
	fmt.Printf("%v\n", len(arr))
	for index, val := range arr {
		fmt.Printf("index=%v,val=%v\n", index, val)
	}
}

func reverseArray(arr []int) {
	if arr == nil {
		return
	}
	var i, j int
	for i, j = 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
