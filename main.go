package main

import (
	"fmt"

	"./config"
	"./experiment"
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

	population.Initialize(config.PopSize, config.FieldSize.X, config.FieldSize.Y)

	fmt.Println("Initial population:", population)

	doorPosition := experiment.RandomDoorPosition(config.FieldSize)

	var experiment experiment.Experiment

	experiment.Initialize(config.FieldSize, population.Individuals[0], doorPosition)

	experiment.Draw()
	experiment.Evaluate()
}
