package config

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"../errors"
	"../utility"
)

type Config struct {
	PopSize   int
	MaxGen    int
	FieldSize utility.Size
}

func (conf Config) String() string {
	return fmt.Sprintf("Config {PopSize: %v, MaxGen: %v, FieldSize: %v}\n", conf.PopSize, conf.MaxGen, conf.FieldSize)
}

func isAsterisk(r rune) bool {
	return r == '*'
}

func PrepareConfig() (Config, error) {
	popSize := flag.Int("popSize", 100, "Population size")
	maxGen := flag.Int("maxGen", 10000, "Maximal count of generations")
	fieldSizeString := flag.String("fieldSize", "80*80", "Field size")

	flag.Parse()

	if *popSize < 2 {
		return Config{}, errors.GeneticError{"Population size must be higher than 1"}
	}

	if *maxGen < 1 {
		return Config{}, errors.GeneticError{"Maximal count of generations must be higher than 0"}
	}

	fieldSizeRegExp := regexp.MustCompile(`\d+\*\d+`)
	if !fieldSizeRegExp.MatchString(*fieldSizeString) {
		return Config{}, errors.GeneticError{"Incorrect format of field size"}
	}
	fieldSizeFields := strings.FieldsFunc(*fieldSizeString, isAsterisk)
	fieldSizeWidth, _ := strconv.Atoi(fieldSizeFields[0])
	fieldSizeHeight, _ := strconv.Atoi(fieldSizeFields[1])
	fieldSize := utility.Size{fieldSizeWidth, fieldSizeHeight}

	return Config{*popSize, *maxGen, fieldSize}, nil
}
