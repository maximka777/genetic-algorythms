package experiment

import (
	"fmt"
	"image"

	"../config"
	"../genetic"
	"../utility"
)

type Position image.Point

func (pos Position) Equals(pos1 Position) bool {
	return pos.X == pos1.X && pos.Y == pos1.Y
}

func RandomDoorPosition(size config.Size) Position {
	isWidth := utility.Randomizer.Random(1) == 1
	isStart := utility.Randomizer.Random(1) == 1

	if isWidth {
		if isStart {
			return Position{utility.Randomizer.Random(size.X), 0}
		} else {
			return Position{utility.Randomizer.Random(size.X), size.Y - 1}
		}
	} else {
		if isStart {
			return Position{0, utility.Randomizer.Random(size.Y)}
		} else {
			return Position{size.X - 1, utility.Randomizer.Random(size.Y)}
		}
	}
}

func CenterPosition(size config.Size) Position {
	return Position{size.X / 2, size.Y / 2}
}

type Experiment struct {
	Field            [][]bool
	Chromosome       genetic.Chromosome
	CurrentPosition  Position
	DoorPosition     Position
	SmallestDistance int
	MaxSteps         int
	Size             config.Size
}

func (exp *Experiment) Initialize(size config.Size, chr genetic.Chromosome, doorPos Position, maxSteps int) {
	exp.Chromosome = chr
	exp.CurrentPosition = CenterPosition(size)
	exp.DoorPosition = doorPos
	exp.Field = make([][]bool, size.X)
	for i := 0; i < size.X; i++ {
		exp.Field[i] = make([]bool, size.Y)
		for j := 0; j < size.Y; j++ {
			exp.Field[i][j] = false
		}
	}
	exp.Field[exp.CurrentPosition.X][exp.CurrentPosition.Y] = true
	exp.MaxSteps = maxSteps
	exp.Size = size
}

func (exp *Experiment) Draw() {
	experimentPicture := ""
	for i := 0; i < exp.Size.X; i++ {
		for j := 0; j < exp.Size.Y; j++ {
			if exp.CurrentPosition.Equals(Position{i, j}) {
				experimentPicture += "C "
			} else if exp.DoorPosition.Equals(Position{i, j}) {
				experimentPicture += "D "
			} else if exp.Field[i][j] {
				experimentPicture += "X "
			} else {
				experimentPicture += "_ "
			}
		}
		experimentPicture += "\n"
	}
	fmt.Println(experimentPicture)
}

func (exp *Experiment) Evaluate() {

}
