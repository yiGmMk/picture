package main

import (
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	go main()
	for {
		time.Sleep(time.Minute)
	}
}
