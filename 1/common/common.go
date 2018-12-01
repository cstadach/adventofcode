package common

import (
	"bufio"
	"os"
	"strconv"
)

func LoadFrequencyFile() []int {
	file, err := os.Open("input")
	Check(err)
	defer file.Close()

	var frequencyDeltas []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		Check(err)
		frequencyDeltas = append(frequencyDeltas, number)
	}
	return frequencyDeltas
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
