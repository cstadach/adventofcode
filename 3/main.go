package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs := LoadInputFile()
	fmt.Println(countOverlapping(buildMatrix(inputs)))
	fmt.Println(findLonelyField(inputs))
}

type field struct {
	id     string
	xCoord int
	yCoord int
	width  int
	height int
}

func countOverlapping(matrix map[string]int) int {
	result := 0
	for _, value := range matrix {
		if value > 1 {
			result++
		}
	}
	return result
}

func findLonelyField(inputs []string) string {
	matrix := buildMatrix(inputs)
	for _, input := range inputs {
		field := newField(input)
		result := true
		for x := field.xCoord; x < field.xCoord+field.width; x++ {
			for y := field.yCoord; y < field.yCoord+field.height; y++ {
				result = result && matrix[fmt.Sprintf("%dx%d", x, y)] == 1
			}
		}
		if result {
			return field.id
		}
	}
	return ""
}

func buildMatrix(inputs []string) map[string]int {
	matrix := map[string]int{}
	for _, input := range inputs {
		field := newField(input)
		for x := field.xCoord; x < field.xCoord+field.width; x++ {
			for y := field.yCoord; y < field.yCoord+field.height; y++ {
				matrix[fmt.Sprintf("%dx%d", x, y)] += 1
			}
		}
	}
	return matrix
}

//#123 @ 3,2: 5x4
func newField(input string) field {
	inputFields := strings.Split(input, " ")
	coordinates := strings.Split(inputFields[2][:len(inputFields[2])-1], ",")
	xCoord, _ := strconv.Atoi(coordinates[0])
	yCoord, _ := strconv.Atoi(coordinates[1])
	dimensions := strings.Split(inputFields[3], "x")
	width, _ := strconv.Atoi(dimensions[0])
	height, _ := strconv.Atoi(dimensions[1])

	return field{
		id:     inputFields[0],
		xCoord: xCoord,
		yCoord: yCoord,
		width:  width,
		height: height,
	}
}

func LoadInputFile() []string {
	file, err := os.Open("input")
	Check(err)
	defer file.Close()

	var inputs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		inputs = append(inputs, input)
	}
	return inputs
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
