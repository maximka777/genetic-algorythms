package main

import (
	"flag"
	"fmt"
	"image"
)

type Size image.Point

func (s Size) String() string {
	return fmt.Sprintf("%vx%v", s.X, s.Y)
}

type Config struct {
	popSize   int
	maxGen    int
	fieldSize Size
}

func (conf Config) String() string {
	return fmt.Sprintf("Config {popSize: %v, maxGen: %v, fieldSize: %v}", conf.popSize, conf.maxGen, conf.fieldSize)
}

func prepareConfig() Config {
	popSize := flag.Int("popSize", 10, "Population size")
	maxGen := flag.Int("maxGen", 100, "Maximal count of generations")
	// fieldSizeString := flag.String("fieldSize", "32*32", "Field size")

	flag.Parse()

	fieldSize := Size{32, 32}

	return Config{*popSize, *maxGen, fieldSize}
}

func main() {
	config := prepareConfig()
	fmt.Println(config)
}
