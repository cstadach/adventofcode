package main

import (
	"common/common"
	"fmt"
)

func Second () {
	result := loopDeltas(common.LoadFrequencyFile())
    fmt.Println(result)
}

func loopDeltas(frequencyDeltas []int) int {
	var frequency int
	var frequencyList = []int{0}
	var repetition = false

	for repetition == false {
		for _, frequencyDelta := range frequencyDeltas {
			frequency += frequencyDelta
			repetition = frequencyRepeated(frequency, frequencyList)
			if repetition {break}

			frequencyList = append(frequencyList,frequency)
		}
	}
	return frequency
}

func frequencyRepeated(currentFrequency int, frequencyList []int) bool {
	for _, freq := range frequencyList {
		if freq == currentFrequency {
			return true
		}
	}
	return false
}

