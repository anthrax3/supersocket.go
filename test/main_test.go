package test

import (
	"testing"
	"time"
)

func sum(n ...int) int {
	var c int
	for _, i := range n {
		c += i
	}
    return c
}

func TestSum(t *testing.T) {
	time.Sleep(time.Second * 2)

	if sum(1, 2, 3) != 6 {
		t.Fatal("sum error!")
	}
}