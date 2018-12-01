package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main () {
	file, err := os.Open("./input")
	check(err)
	defer file.Close()

	frequency := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		number, err := strconv.Atoi(scanner.Text())
		check(err)
		frequency += number
	}
	fmt.Print(frequency)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
