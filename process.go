package main

import (
	"github.com/sarabala1979/numa-sdk-code/test"
	"time"
)

func process(eventTime, waterMark time.Time, msg []byte) []byte {

	return test.Test(msg)
}
