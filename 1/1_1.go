package main

import (
	"common/common"
	"fmt"
)

func First() {
	frequencyDeltas := common.LoadFrequencyFile()
	frequency := 0
	for _, frequencyDelta := range frequencyDeltas {
		frequency += frequencyDelta
	}
	fmt.Println(frequency)
}
