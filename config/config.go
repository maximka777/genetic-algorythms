package config

import (
	"flag"
	"fmt"
	"image"
	"regexp"
	"strconv"
	"strings"

	"../errors"
)

type Size image.Point

func (s Size) String() string {
	return fmt.Sprintf("%v*%v", s.X, s.Y)
}

type Config struct {
	popSize   int
	maxGen    int
	fieldSize Size
}

func (conf Config) String() string {
	return fmt.Sprintf("Config {popSize: %v, maxGen: %v, fieldSize: %v}", conf.popSize, conf.maxGen, conf.fieldSize)
}

func isAsterisk(r rune) bool {
	return r == '*'
}

func PrepareConfig() (Config, error) {
	popSize := flag.Int("popSize", 10, "Population size")
	maxGen := flag.Int("maxGen", 100, "Maximal count of generations")
	fieldSizeString := flag.String("fieldSize", "32*32", "Field size")

	flag.Parse()

	fieldSizeRegExp := regexp.MustCompile(`\d+\*\d+`)
	if !fieldSizeRegExp.MatchString(*fieldSizeString) {
		return Config{}, errors.GeneticError{"Incorrect format of field size"}
	}

	fieldSizeFields := strings.FieldsFunc(*fieldSizeString, isAsterisk)

	fieldSizeWidth, _ := strconv.Atoi(fieldSizeFields[0])
	fieldSizeHeight, _ := strconv.Atoi(fieldSizeFields[1])

	fieldSize := Size{fieldSizeWidth, fieldSizeHeight}

	return Config{*popSize, *maxGen, fieldSize}, nil
}