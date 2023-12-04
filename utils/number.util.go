package utils

import (
	"math/rand"
	"time"
)

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func FPSToMilliseconds(fps int64) int64 {
	millis := time.Second.Milliseconds() / fps
	return millis
}
