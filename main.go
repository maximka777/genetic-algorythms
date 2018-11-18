package main

import (
	"fmt"

	"./config"
	"./genetic"
)

func main() {
	config, err := config.PrepareConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Configuration:", config)

	var population genetic.Population

	population.Initialize(config.PopSize, config.FieldSize)

	population.CalculateFitness(config)

	for population.Fittest != 0 && population.Generation < config.MaxGen - 1 {
		population.NextGeneration()
		population.CalculateFitness(config)
	}

	fmt.Println(population)
	fittestChromosome := population.GetFittestIndividual()
	fmt.Println("Fittest chromosome:", fittestChromosome)
	fittestChromosome.Fitness = 1 << 32 - 1
	fittestChromosome.CalculateFitness(config, population.DoorPosition, true)
}
