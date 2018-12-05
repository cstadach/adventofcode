package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOverlappingFields(t *testing.T) {
	coordinates := []string{
		"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2",}
	result := buildMatrix(coordinates)
	if !(assert.Condition(t, func() bool {return result["3x3"] > 1}) &&
		assert.Condition(t, func() bool {return result["4x3"] > 1}) &&
		assert.Condition(t, func() bool {return result["3x4"] > 1}) &&
		assert.Condition(t, func() bool {return result["4x4"] > 1})) {
		t.Error( "Expected overlapping fields did not match" )
	}
}

func TestOverlappingFieldsCount(t *testing.T) {
    coordinates := []string{
		"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2",}
    result := countOverlapping(buildMatrix(coordinates))
    if !(assert.Equal(t, 4, result)) {
		t.Error( "Expected overlapping fields did not match" )
	}
}
