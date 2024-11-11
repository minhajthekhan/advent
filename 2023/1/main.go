package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func buildInput() []string {
	b, err := os.ReadFile("2023/1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	values := []string{}
	v := ""
	for _, str := range string(b) {
		if str != '\n' {
			v += string(str)
		} else {
			values = append(values, v)
			v = ""
		}
	}

	return append(values, v)
}

func main() {

	input := buildInput()
	fmt.Println(getCalibration(input))
}

func getCalibration(input []string) int {
	sum := 0
	for _, str := range input {
		sum += getCalibrationValue(str)
	}
	return sum
}

func getCalibrationValue(str string) int {
	first := -1
	second := -1
	for _, c := range str {
		candidate := string(c)
		number, err := strconv.Atoi(candidate)
		if err != nil {
			continue
		}
		if first == -1 {
			first = number
			continue
		}

		second = number
	}

	if second == -1 {
		second = first
	}
	return (first * 10) + second
}
