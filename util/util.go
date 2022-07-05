package util

import (
	"math/rand"
	"time"
)

func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
func RandomFloat64() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()
}
