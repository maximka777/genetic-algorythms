package experiment

import (
	"fmt"
	"image"
	"math"

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
	Step             int
	Size             config.Size
}

func (exp *Experiment) Initialize(size config.Size, chr genetic.Chromosome, doorPos Position) {
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
	exp.MaxSteps = CalculateMaxSteps(size)
	exp.Step = 0
	exp.Size = size
}

func (exp *Experiment) Draw() {
	experimentPicture := fmt.Sprintf("Experiment #%v\n", exp.Step)
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
	fmt.Println("Chromosome:", exp.Chromosome)
	fmt.Println("SmallestDistance:", exp.SmallestDistance)
}

func (exp *Experiment) MakeStep() {
	stepCommand := exp.Chromosome.Genes[exp.CurrentPosition.X][exp.CurrentPosition.Y]
	fmt.Println("Step command:", stepCommand)
	switch stepCommand {
	case 0:
		if exp.CurrentPosition.Y > 0 {
			exp.CurrentPosition.Y -= 1
		}
	case 1:
		if exp.CurrentPosition.X < exp.Size.X-1 {
			exp.CurrentPosition.X += 1
		}
	case 2:
		if exp.CurrentPosition.Y < exp.Size.Y-1 {
			exp.CurrentPosition.Y += 1
		}
	case 3:
		if exp.CurrentPosition.X > 0 {
			exp.CurrentPosition.X -= 1
		}
	}
	exp.Field[exp.CurrentPosition.X][exp.CurrentPosition.Y] = true
}

func (exp *Experiment) Evaluate() {
	for exp.Step < exp.MaxSteps {
		exp.Step++
		exp.MakeStep()
		exp.CalculateDistance()
		if exp.SmallestDistance == 0 {
			return
		}
		exp.Draw()
	}
}

func (exp *Experiment) CalculateDistance() {
	distance := int(math.Abs(float64(exp.DoorPosition.X-exp.CurrentPosition.X)) + math.Abs(float64(exp.DoorPosition.Y-exp.CurrentPosition.Y)))
	if distance < exp.SmallestDistance {
		exp.SmallestDistance = distance
	}
}

func CalculateMaxSteps(size config.Size) int {
	return (size.X/2 + 1) * (size.Y/2 + 1)
}
