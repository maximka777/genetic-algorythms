package utility

import (
	"fmt"
	"image"
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

func (r *Rand) Random(max int) int {
	if !r.seeded {
		r.initialize()
	}
	return r.seededRand.Int() % max
}

type Position image.Point

func (pos Position) Equals(pos1 Position) bool {
	return pos.X == pos1.X && pos.Y == pos1.Y
}

func RandomDoorPosition(size Size) Position {
	isWidth := Randomizer.Random(2) == 1
	isStart := Randomizer.Random(2) == 1

	if isWidth {
		if isStart {
			return Position{Randomizer.Random(size.X), 0}
		} else {
			return Position{Randomizer.Random(size.X), size.Y - 1}
		}
	} else {
		if isStart {
			return Position{0, Randomizer.Random(size.Y)}
		} else {
			return Position{size.X - 1, Randomizer.Random(size.Y)}
		}
	}
}

func RandomPosition(size Size) Position {
	return Position{Randomizer.Random(size.X), Randomizer.Random(size.Y)}
}

func CenterPosition(size Size) Position {
	return Position{size.X / 2, size.Y / 2}
}

type Size image.Point

func (s Size) String() string {
	return fmt.Sprintf("%v*%v", s.X, s.Y)
}
