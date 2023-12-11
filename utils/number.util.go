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

func RandomColor0n255() uint8 {
	r := rand.Intn(2)

	if r == 0 {
		return 0
	}

	return 255
}

func RandomColor255() uint8 {
	return uint8(rand.Intn(255))
}
