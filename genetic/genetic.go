package genetic

import (
	"../config"
	"../utility"
	"fmt"
	"math"
	"sort"
)

/**
 * 0 - Up
 * 1 - Right
 * 2 - Down
 * 3 - Left
 */

type Gene byte

func (g *Gene) Random() {
	*g = Gene(utility.Randomizer.Random(4))
}

type Chromosome struct {
	Genes   [][]Gene
	Fitness int
}

func (chr *Chromosome) Random(width, height int) {
	chr.Genes = make([][]Gene, width)
	for i := 0; i < width; i++ {
		chr.Genes[i] = make([]Gene, height)
		for j := 0; j < height; j++ {
			chr.Genes[i][j].Random()
		}
	}
	chr.Fitness = 1<<32 - 1
}

func (chr Chromosome) String() string {
	s := "Chromosome {\n"
	s += "\tGenes: {\n"
	for j := 0; j < cap(chr.Genes[0]); j++ {
		s += "\t\t[ "
		for i := 0; i < cap(chr.Genes); i++ {
			s += fmt.Sprintf("%v ", chr.Genes[i][j])
		}
		s += "]\n"
	}
	s += "\t}\n"
	s += fmt.Sprintf("\tFitness: %v\n", chr.Fitness)
	s += "}\n"
	return s
}

func (chr *Chromosome) CalculateFitness(config config.Config, doorPos utility.Position, draw bool) {
	var experiment Experiment
	experiment.Initialize(config.FieldSize, *chr, doorPos)
	experiment.Evaluate(draw)
	chr.Fitness = experiment.SmallestDistance
}

type Population struct {
	Individuals  Chromosomes
	Fittest      int
	Generation   int
	DoorPosition utility.Position
}

func (pop Population) String() string {
	s := "Population {\n"
	s += fmt.Sprintf("\tGeneration: %v\n", pop.Generation)
	s += fmt.Sprintf("\tFittest: %v\n", pop.Fittest)
	s += "}\n"
	return s
}

func (pop *Population) CalculateFitness(config config.Config) {
	for idx := range pop.Individuals {
		pop.Individuals[idx].CalculateFitness(config, pop.DoorPosition, false)
	}
	Fittest := 1<<32 - 1
	for _, chr := range pop.Individuals {
		if chr.Fitness < Fittest {
			Fittest = chr.Fitness
		}
	}
	pop.Fittest = Fittest
}

func (pop *Population) Initialize(individualsCount int, size utility.Size) {
	pop.Individuals = make([]Chromosome, individualsCount)
	for i := 0; i < individualsCount; i++ {
		pop.Individuals[i].Random(size.X, size.Y)
	}
	pop.Generation = 0
	pop.Fittest = 1<<32 - 1
	pop.DoorPosition = utility.RandomDoorPosition(size)
}

type Chromosomes []Chromosome

func (c Chromosomes) Len() int {
	return len(c)
}
func (c Chromosomes) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c Chromosomes) Less(i, j int) bool {
	return c[i].Fitness < c[j].Fitness
}

func (a Chromosome) Crossover(b Chromosome) Chromosome {
	result := Chromosome{}
	crossOverPosition := utility.RandomPosition(utility.Size{cap(a.Genes), cap(a.Genes[0])})
	result.Genes = make([][]Gene, cap(a.Genes))

	for i := 0; i < cap(a.Genes); i++ {
		result.Genes[i] = make([]Gene, cap(a.Genes[0]))
		for j := 0; j < cap(a.Genes[0]); j++ {
			if j < crossOverPosition.Y {
				result.Genes[i][j] = a.Genes[i][j]
			} else if j == crossOverPosition.Y {
				if i < crossOverPosition.X {
					result.Genes[i][j] = a.Genes[i][j]
				} else {
					result.Genes[i][j] = b.Genes[i][j]
				}
			} else {
				result.Genes[i][j] = b.Genes[i][j]
			}
		}
	}

	needMutation := utility.Randomizer.Random(2) == 1

	if needMutation {
		result.Mutate()
	}

	return result
}

func (chr *Chromosome) Mutate() {
	mutationCount := utility.Randomizer.Random(cap(chr.Genes) * cap(chr.Genes[0]) / 2)

	for i := 0; i < mutationCount; i++ {
		mutationPosition := utility.RandomPosition(utility.Size{cap(chr.Genes), cap(chr.Genes[0])})
		mutationResult := Gene(utility.Randomizer.Random(4))
		chr.Genes[mutationPosition.X][mutationPosition.Y] = mutationResult
	}
}

func (pop *Population) NextGeneration() {
	pop.Generation += 1
	sort.Sort(pop.Individuals)
	bestChromosome := pop.Individuals[0]
	bestSecondChromosome := pop.Individuals[1]
	for i := 0; i < cap(pop.Individuals); i++ {
		pop.Individuals[i] = bestChromosome.Crossover(bestSecondChromosome)
	}
}

func (pop Population) GetFittestIndividual() Chromosome {
	sort.Sort(pop.Individuals)
	return pop.Individuals[0]
}

type Experiment struct {
	Field            [][]bool
	Chromosome       Chromosome
	CurrentPosition  utility.Position
	DoorPosition     utility.Position
	SmallestDistance int
	MaxSteps         int
	Step             int
	Size             utility.Size
	CycleStep        bool
}

func (exp *Experiment) Initialize(size utility.Size, chr Chromosome, doorPos utility.Position) {
	exp.Chromosome = chr
	exp.CurrentPosition = utility.CenterPosition(size)
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
	exp.SmallestDistance = 1<<62 - 1
}

func (exp *Experiment) Draw() {
	experimentPicture := fmt.Sprintf("Experiment\nStep#%v\n", exp.Step)
	for j := 0; j < exp.Size.Y; j++ {
		for i := 0; i < exp.Size.X; i++ {
			if exp.CurrentPosition.Equals(utility.Position{i, j}) {
				if exp.CurrentPosition.Equals(exp.DoorPosition) {
					experimentPicture += "W "
				} else {
					experimentPicture += "C "
				}
			} else if exp.DoorPosition.Equals(utility.Position{i, j}) {
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

func (exp *Experiment) MakeStep() {
	stepCommand := exp.Chromosome.Genes[exp.CurrentPosition.X][exp.CurrentPosition.Y]
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
	if exp.Field[exp.CurrentPosition.X][exp.CurrentPosition.Y] {
		exp.CycleStep = true
	} else {
		exp.Field[exp.CurrentPosition.X][exp.CurrentPosition.Y] = true
	}
}

func (exp *Experiment) Evaluate(draw bool) {
	for exp.Step < exp.MaxSteps {
		exp.Step++
		exp.MakeStep()
		exp.CalculateDistance()
		// NOTE: Uncomment me if you want to watch the whole experiment
		//if draw {
		//	exp.Draw()
		//}
		if exp.SmallestDistance == 0 {
			if draw {
				exp.Draw()
			}
			return
		}
	}
}

func (exp *Experiment) CalculateDistance() {
	distance := int(math.Abs(float64(exp.DoorPosition.X-exp.CurrentPosition.X)) + math.Abs(float64(exp.DoorPosition.Y-exp.CurrentPosition.Y)))
	if distance < exp.SmallestDistance {
		exp.SmallestDistance = distance
	}
}

func CalculateMaxSteps(size utility.Size) int {
	return (size.X/2 + 1) * (size.Y/2 + 1)
}
