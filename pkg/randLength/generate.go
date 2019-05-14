package randLength

import (
	"math/rand"
	"time"
)

func Generate(min, max int) int {
	rand.Seed(time.Now().UTC().Unix())
	return rand.Intn(max-min) + min
}
