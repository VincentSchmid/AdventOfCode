package day03

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strconv"
)

const (
	SEPARATOR    = '.'
	GEAR         = '*'
	SYMBOL_REGEX = `[^A-Za-z0-9.]`
	NUMBER_REGEX = `[0-9]+`
)

type Vector struct {
	x int
	y int
}

func getValue(line string, index int) (int, int) {
	digits := string(line[index])
	numDigits := 1
	i := index - 1

	for i >= 0 && IsDigit(line[i]) {
		digits = string(line[i]) + digits
		i--
		numDigits++
	}

	i = index + 1
	for i < len(line) && IsDigit(line[i]) {
		digits += string(line[i])
		i++
		numDigits++
	}

	number, _ := strconv.Atoi(digits)

	return number, numDigits
}

func ValueTouchesSymbol(valueArea []string) bool {
	re := regexp.MustCompile(SYMBOL_REGEX)

	for _, s := range valueArea {
		if re.MatchString(s) {
			return true
		}
	}

	return false
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func ExpandArea(topLeft Vector, bottomRight Vector, maxTopLeft Vector, maxBottomRight Vector, extend int) (Vector, Vector) {
	newTopLeft := Vector{x: utils.Max(maxTopLeft.x, topLeft.x-extend), y: utils.Max(maxTopLeft.y, topLeft.y-extend)}
	newBottomRight := Vector{x: utils.Min(maxBottomRight.x, bottomRight.x+extend+1), y: utils.Min(maxBottomRight.y, bottomRight.y+extend+1)}

	return newTopLeft, newBottomRight
}

func getValueArea(arr *[]string, x int, y int, width int, maxTopLeft Vector, maxBottomRight Vector, extend int) []string {
	newTopLeft, newBottomRight := ExpandArea(Vector{x: x, y: y}, Vector{x: x + width, y: y}, maxTopLeft, maxBottomRight, extend)

	valueArea := make([]string, newBottomRight.y-newTopLeft.y)
	copy(valueArea, (*arr)[newTopLeft.y:newBottomRight.y])

	for i, row := range valueArea {
		valueArea[i] = row[newTopLeft.x:newBottomRight.x]
	}

	return valueArea
}

func Problem01(lines []string) int {
	sum := 0
	lineLength := len(lines[0])
	maxTopLeft := Vector{x: 0, y: 0}
	maxBottomRight := Vector{x: len(lines[0]), y: len(lines)}

	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			if IsDigit(line[x]) {
				val, numDigits := getValue(line, x)
				valueArea := getValueArea(&lines, x, y, numDigits-1, maxTopLeft, maxBottomRight, 1)

				if ValueTouchesSymbol(valueArea) {
					sum += val
				}

				x += utils.Min(numDigits, lineLength)
			}
		}
	}

	return sum
}

func Problem02(lines []string) int {
	sum := 0
	maxTopLeft := Vector{x: 0, y: 0}
	maxBottomRight := Vector{x: len(lines[0]), y: len(lines)}
	re := regexp.MustCompile(NUMBER_REGEX)

	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			if line[x] == GEAR {
				newTopLeft, newBottomRight := ExpandArea(Vector{x: x, y: y}, Vector{x: x, y: y}, maxTopLeft, maxBottomRight, 1)

				valueArea := make([]string, newBottomRight.y-newTopLeft.y)
				copy(valueArea, lines[newTopLeft.y:newBottomRight.y])

				vals := make([]int, 0)
				for yn, row := range valueArea {
					results := re.FindAllStringIndex(row[newTopLeft.x:newBottomRight.x], -1)

					for _, match := range results {
						val, _ := getValue(valueArea[yn], x+match[0]-1)
						vals = append(vals, val)
					}
				}

				if len(vals) > 1 {
					sum += vals[0] * vals[1]
				}
			}
		}
	}

	return sum
}

func Run() {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(3))
	fmt.Println(Problem01(lines))
	fmt.Println(Problem02(lines))
}
