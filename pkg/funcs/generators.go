package funcs

import (
	"crypto/rand"
	"math/big"
)

//RandomInt generates random int counts
func RandomInt(min, max int) int {
	v, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	v.Add(v, big.NewInt(int64(min)))
	return int(v.Int64())
}

//RandFloat generates random float64 counts
func RandFloat(min, max int) float64 {
	var1 := RandomInt(min, max)
	var2 := RandomInt(0, 100)
	return float64(var1) + float64(var2)/100
}
