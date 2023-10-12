package test

import (
	"fmt"
	"testing"
)

func TestAppend(t *testing.T) {

	arr := make([]int, 0)
	last := 0
	// go1.13=> 1 2 4 8 16 32 64 128 256 512 1024 1280 1696 2304
	// go1.15=> 1 2 4 8 16 32 64 128 256 512 1024 1280 1696 2304
	// go1.19=> 1 2 4 8 16 32 64 128 256 512 848 1280 1792 2560
	// go1.21=> 1 2 4 8 16 32 64 128 256 512 848 1280 1792 2560
	for i := 0; i <= 2000; i++ {
		if cap(arr) != last {
			fmt.Print(cap(arr), " ")
			last = cap(arr)
		}
		arr = append(arr, i)
	}
}
