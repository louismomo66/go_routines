package main

import (
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	main()

	if len(orderOfFinish) != 5 {
		t.Error("wrong number of entries in the slice")
	}
}

// Inn0vexug@2024!
