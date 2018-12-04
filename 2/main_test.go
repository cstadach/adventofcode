package main

import (
	"fmt"
	"testing"
)

func TestChecksum(t *testing.T) {
	words := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",}
	result := checksum(words)
	fmt.Println(result)
	if result != 12 {
		t.Error(
			"Expected", 12,
			"But received",  result,
		)
	}
}

func TestWord(t *testing.T) {
	words := []string{"abcde",
	"fghij",
	"klmno",
	"pqrst",
	"fguij",
	"axcye",
	"wvxyz",}
	result := findPairs(words)
	if result != "fgij" {
		t.Error(
			"Expected", "fgij",
			"But received",  result,
		)
	}
}
