package genetic

import (
	"fmt"

	"../utility"
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
	Genes [][]Gene
}

func (chr *Chromosome) Random(width, height int) {
	chr.Genes = make([][]Gene, width)
	for i := 0; i < width; i++ {
		chr.Genes[i] = make([]Gene, height)
		for j := 0; j < height; j++ {
			chr.Genes[i][j].Random()
		}
	}
}

func (chr *Chromosome) FitnessFunction() int {
	return 0
}

type Population struct {
	Individuals []Chromosome
}

func (pop Population) String() string {
	s := "Population.Individuals {\n"
	for _, chr := range pop.Individuals {
		s += fmt.Sprintf("%v", chr) + "\n"
	}
	s += "}"
	return s
}

func (pop *Population) Initialize(size int, fieldWidth, fieldHeight int) {
	pop.Individuals = make([]Chromosome, size)
	for i := 0; i < size; i++ {
		pop.Individuals[i].Random(fieldWidth, fieldHeight)
	}
}
