package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// input := "./sample.txt"
	input := "./input.txt"
	l1, l2 := parseInput(input)
	l1.sort()
	l2.sort()
	distance := l1.distance(l2)
	fmt.Println(distance)

}

func parseInput(s string) (list, list) {
	b, err := os.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}

	l1 := make(list, 0)
	l2 := make(list, 0)
	for _, v := range strings.Split(string(b), "\n") {
		nums := strings.Split(v, "   ")
		first, second := nums[0], nums[1]
		l1 = l1.add(first)
		l2 = l2.add(second)
	}

	return l1, l2
}

type list []int

func (l list) add(s string) []int {
	num, _ := strconv.Atoi(s)
	l = append(l, num)
	return l
}

func (l list) sort() {
	sort.Ints(l)
}

func (l list) distance(against list) int {
	total := 0
	for i := range l {
		distance := l[i] - against[i]
		total += int(math.Abs(float64(distance)))
	}
	return total
}
