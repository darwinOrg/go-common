package utils

import "math/rand"

func RandomIntInRange(min, max int) int {
	return min + rand.Intn(max-min)
}
