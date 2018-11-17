package genetic

type Chromosome struct {
	Genes [][]byte
}

type Population struct {
	Individuals []Chromosome
}
