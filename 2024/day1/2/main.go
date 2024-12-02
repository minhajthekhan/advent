package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// input := "./sample.txt"
	input := "./input.txt"
	d1, l2 := parseInput(input)
	fmt.Println(d1.score(l2))

}

func parseInput(s string) (distinct, list) {
	b, err := os.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}

	d1 := make(distinct, 0)
	l2 := make(list, 0)
	for _, v := range strings.Split(string(b), "\n") {
		nums := strings.Split(v, "   ")
		first, second := nums[0], nums[1]
		d1 = d1.add(first)
		l2 = l2.add(second)
	}

	return d1, l2
}

type list []int

func (l list) add(s string) []int {
	num, _ := strconv.Atoi(s)
	l = append(l, num)
	return l
}

type distinct map[int]int

func (d distinct) add(v string) map[int]int {
	num, _ := strconv.Atoi(v)
	if _, ok := d[num]; !ok {
		d[num] = 0
	}
	x := d[num]
	x++
	d[num] = x
	return d
}

func (d distinct) score(l list) int {
	score := 0
	for key, occurrence := range d {
		count := 0
		for _, e := range l {
			if key == e {
				count++
			}
		}
		score += key * count * occurrence
	}
	return score
}
