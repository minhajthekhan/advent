package main

import (
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// input := "./sample.txt"
	input := "./input.txt"

	sc := 0
	for _, list := range parseInput(input) {
		if list.isSafe() {
			sc++
			continue
		}

		if list.isSafeWithCombinations() {
			sc++
			continue
		}
	}

}

type report []int

func (r report) isSafe() bool {
	var inc *bool
	safe := true
	for i := 0; i < len(r)-1; i++ {
		ok, trend := isSafe(r[i], r[i+1], inc)
		inc = &trend
		if !ok {
			safe = false
			break
		}
	}

	return safe
}

func (r report) isSafeWithCombinations() bool {
	for _, c := range r.combinations() {
		if c.isSafe() {
			return true
		}
	}
	return false
}
func (r report) combinations() []report {
	var result []report
	for i := 0; i < len(r); i++ {
		tmp := make([]int, len(r))
		copy(tmp, r)
		result = append(result, slices.Delete(tmp, i, i+1))
	}
	return result
}

func isSafe(a, b int, inc *bool) (bool, bool) {
	if !isNear(a, b) {
		return false, false
	}
	if inc == nil {
		return true, a < b
	}
	if a < b && *inc {
		return true, *inc
	}

	if a > b && !*inc {
		return true, *inc
	}

	return false, *inc
}

func isNear(a, b int) bool {
	// Any two adjacent levels differ by at least one and at most three
	diff := int(math.Abs(float64(a - b)))
	if diff < 1 {
		return false
	}
	if diff > 3 {
		return false
	}

	return true
}

func parseInput(s string) []report {
	b, err := os.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}

	var reports []report

	for _, v := range strings.Split(string(b), "\n") {
		var levels []int
		for _, str := range strings.Split(v, " ") {
			lvl, _ := strconv.Atoi(str)
			levels = append(levels, lvl)
		}
		reports = append(reports, levels)
	}
	return reports
}
