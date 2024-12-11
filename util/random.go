package util

import (
	"math/rand"
	"time"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWQYZabcdefghijklmnopqrstuvwxyz"

// Random Seed
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandomInt generates a random integer between min and max.
func RandomInt(min, max int32) int32 {
	return min + r.Int31n(max-min+1)
}

// RandomPrice generates random price for the rooms.
func RandomPrice() int32 {
	return RandomInt(5000, 10000)
}

// RandomBool randomly returns true or false.
func RandomBool() bool {
	return r.Intn(2) == 1
}

// RandomGuests generates random number of guests allowed for a room.
func RandomGuests() int32 {
	return RandomInt(2, 8)
}
