package utility

import (
	"math/rand"
	"time"
)

var Randomizer Rand

type Rand struct {
	seededRand *rand.Rand
	seeded     bool
}

func (r *Rand) initialize() {
	r.seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	r.seeded = true
}

func (r Rand) Random(max int) int {
	if !r.seeded {
		r.initialize()
	}
	return r.seededRand.Int() % max
}
