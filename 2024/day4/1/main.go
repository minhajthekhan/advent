package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	b, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	x := (strings.Split(string(b), "\n"))

	// x := []string{
	// 	"MMMSXXMASM",
	// 	"MSAMXMSMSA",
	// 	"AMXSXMAAMM",
	// 	"MSAMASMSMX",
	// 	"XMASAMXAMM",
	// 	"XXAMMXXAMA",
	// 	"SMSMSASXSS",
	// 	"SAXAMASAAA",
	// 	"MAMMMXMMMM",
	// 	"MXMXAXMASX",
	// }

	c := 0
	for i := range x {
		c1, s := leftToRight(x[i])
		c += c1
		// fmt.Println(s)
		c2, s := topBottom(x[i], x[i+1:], s)
		c += c2
		c3, s := rightToLeft(x[i], s)
		c += c3
		c4, s := bottomUp(x[len(x)-i-1], x[0:len(x)-i-1], s)
		c += c4
	}
	fmt.Println(c)
}

func leftToRight(r string) (int, string) {
	c := 0
	completeStr := ""
	// xStartsFrom := -1
	for i := 0; i < len(r); i++ {
		if string(r[i]) == "X" && len(r[i:]) > 3 {
			// xStartsFrom = i
			if string(r[i+1]) == "M" && string(r[i+2]) == "A" && string(r[i+3]) == "S" {
				c++
				i += 2
				completeStr += "XMAS"
				continue
			}
		}
		completeStr += "."

	}

	return c, completeStr
}

func topBottom(r string, listBelow []string, s string) (int, string) {
	c := 0
	for i := 0; i < len(r); i++ {
		if string(r[i]) == "X" && len(listBelow) > 2 && len(r[i:]) > 3 {
			if string(listBelow[0][i+1]) == "M" && string(listBelow[1][i+2]) == "A" && string(listBelow[2][i+3]) == "S" {
				s = replaceAtIndex(s, 'X', i)
				c++
			}
		}
	}
	for i := len(r) - 1; i >= 0; i-- {
		if string(r[i]) == "S" && len(listBelow) > 2 && len(r[:i]) >= 3 {
			if string(listBelow[0][i-1]) == "A" && string(listBelow[1][i-2]) == "M" && string(listBelow[2][i-3]) == "X" {
				s = replaceAtIndex(s, 'S', i)
				c++
			}
		}
	}

	for i := 0; i <= len(listBelow) && i < len(r); i++ {
		if string(r[i]) == "X" && len(listBelow) > 2 {
			if string(listBelow[0][i]) == "M" && string(listBelow[1][i]) == "A" && string(listBelow[2][i]) == "S" {
				s = replaceAtIndex(s, 'X', i)
				c++
			}
		}
	}

	return c, s
}

func rightToLeft(r string, s string) (int, string) {
	c := 0
	// xStartsFrom := -1
	for i := len(r) - 1; i >= 0; i-- {
		remainingStr := r[:i]
		if string(r[i]) == "X" && len(remainingStr) >= 3 {
			// xStartsFrom = i
			if string(r[i-1]) == "M" && string(r[i-2]) == "A" && string(r[i-3]) == "S" {
				c++
				s = replaceAtIndex(s, 'X', i)
				s = replaceAtIndex(s, 'M', i-1)
				s = replaceAtIndex(s, 'A', i-2)
				s = replaceAtIndex(s, 'S', i-3)
				i -= 2
				continue
			}
		}
	}

	return c, s
}

func bottomUp(r string, listAbove []string, s string) (int, string) {
	c := 0
	for i := 0; i < len(r); i++ {
		if string(r[i]) == "X" && len(listAbove) > 2 && len(r[i:]) > 3 {
			if string(listAbove[len(listAbove)-1][i+1]) == "M" && string(listAbove[len(listAbove)-2][i+2]) == "A" && string(listAbove[len(listAbove)-3][i+3]) == "S" {
				c++
				s = replaceAtIndex(s, 'X', i)
			}
		}
	}

	for i := 0; i < len(r); i++ {
		if string(r[i]) == "S" && len(listAbove) > 2 && len(r[i:]) > 3 {
			if string(listAbove[len(listAbove)-1][i+1]) == "A" && string(listAbove[len(listAbove)-2][i+2]) == "M" && string(listAbove[len(listAbove)-3][i+3]) == "X" {
				c++
				s = replaceAtIndex(s, 'X', i)
			}
		}
	}

	for i := len(r) - 1; i >= 0; i-- {
		if string(r[i]) == "S" && len(listAbove) > 2 && len(r[:i]) >= 3 {
			if string(listAbove[len(listAbove)-1][i-1]) == "A" && string(listAbove[len(listAbove)-1][i-2]) == "M" && string(listAbove[len(listAbove)-1][i-3]) == "X" {
				c++
				s = replaceAtIndex(s, 'S', i)
			}
		}
	}

	for i := len(r) - 1; i >= 0; i-- {
		if string(r[i]) == "X" && len(listAbove) > 2 && len(r[:i]) >= 3 {
			if string(listAbove[len(listAbove)-1][i-1]) == "M" && string(listAbove[len(listAbove)-2][i-2]) == "A" && string(listAbove[len(listAbove)-3][i-3]) == "S" {
				c++
				s = replaceAtIndex(s, 'S', i)
			}
		}
	}

	for i := len(r) - 1; i >= 0; i-- {
		if string(r[i]) == "X" && len(listAbove) >= 2 {
			if string(listAbove[len(listAbove)-1][i]) == "M" && string(listAbove[len(listAbove)-2][i]) == "A" && string(listAbove[len(listAbove)-3][i]) == "S" {
				s = replaceAtIndex(s, 'X', i)
				c++
			}
		}
	}
	return c, s

}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
