package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputs := LoadInputFile()
	fmt.Println("checksum:", checksum(inputs))
	fmt.Println(findPairs(inputs))
}

func findPairs(words[]string) string {
    for index, word := range words {
        for checkIndex := index +1; checkIndex < len(words); checkIndex++ {
        	var inequalCount, inequalPos int
			for charPos := 0; charPos < len(word); charPos++ {
				if words[checkIndex][charPos] != word[charPos] {
					inequalPos = charPos
					inequalCount++
				}
			}
			if inequalCount == 1 {
				beginning := strings.Split(word,"")[:inequalPos]
				end := strings.Split(word,"")[inequalPos+1:]
				result := strings.Join(beginning,"") + strings.Join(end, "")
				return result
			}
		}
	}
	return ""
}

func checksum (words []string) int {
	var twos, threes int
	for _, word := range words {
		letterCount := countLetters(word)
        containsTwo := false
		containsThree := false
		for _, count := range letterCount {
			containsTwo = containsTwo || count == 2
			containsThree = containsThree || count == 3
		}
		if containsTwo {twos++}
		if containsThree {threes++}
	}
	return twos * threes
}

func countLetters(word string) map[string]int {
	charMap := map[string]int{}
	wordChars := strings.Split(word, "")
	for i, char := range wordChars {
		if charMap[char] == 0 { charMap[char]++ }

		if strings.Contains(strings.Join(wordChars[i+1:],""), char) {
			charMap[char] += 1
		}
	}
	return charMap
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
