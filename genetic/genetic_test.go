package genetic

import (
	"testing"
)

func TestChromosomeCrossover(t *testing.T) {
	var chrA Chromosome
	chrA.Random(3, 3)
	var chrB Chromosome
	chrB.Random(3, 3)
	chrA.Crossover(chrB)
}
