package gondola_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAdjacentInRow(t *testing.T) {
	testcases := []struct {
		name     string
		row      string
		expected int
	}{
		{
			name:     "single non numeric",
			row:      "*",
			expected: 0,
		},
		{
			name:     "double non numeric",
			row:      "**",
			expected: 0,
		},
		{
			name:     "starts with non numeric",
			row:      "**12",
			expected: 12,
		},
		{
			name:     "ends with non numeric",
			row:      "12***",
			expected: 12,
		},
		{
			name:     "non numeric between numerics",
			row:      "1*1",
			expected: 2,
		},
		{
			name:     "pattern non numeric between numerics",
			row:      "1*1*1*1",
			expected: 4,
		},
		{
			name:     "repeated non numeric between numerics",
			row:      "5*****5",
			expected: 10,
		},
		{
			name:     "repeated non numeric between numerics with starting and ending non numerics",
			row:      "*****5*****5****5*****",
			expected: 15,
		},
		{
			name:     "from the code example - row 1",
			row:      "467..114..",
			expected: 0,
		},
		{
			name:     "from the code example - row 5",
			row:      "617*......",
			expected: 617,
		},
		{
			name:     "from the code example - row last",
			row:      ".664.598..",
			expected: 0,
		},
		{
			name:     "from the code example - row second last",
			row:      "...$.*....",
			expected: 0,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, getAdjacentInRow(tc.row))
		})
	}
}

func getAdjacentInRow(row string) int {
	lastNonNumeric := -1
	sum := 0
	for i, el := range row {
		if _, err := strconv.Atoi(string(el)); err != nil && string(el) != "." {
			var num int
			if i <= 0 {
				num = 0
				continue
			}
			lastNonNumeric = i
			x := getNumericalSubstring(row, i-1)
			if x == "" {
				continue
			}
			num, _ = strconv.Atoi(x)
			sum += num
		}
	}
	if lastNonNumeric < 0 {
		return sum
	}
	if lastNonNumeric >= len(row)-1 {
		return sum
	}
	num, _ := strconv.Atoi(getNumericalSubstring(row, lastNonNumeric+1))
	return sum + num
}

func TestGetNumericalSubstring(t *testing.T) {
	testcases := []struct {
		name       string
		input      string
		entrypoint int
		substr     string
	}{
		{
			name:       "no numbers",
			input:      "*",
			entrypoint: 0,
			substr:     "",
		},
		{
			name:       "no numbers, multiple non numeric",
			input:      "**",
			entrypoint: 0,
			substr:     "",
		},
		{
			name:       "endings with non numeric",
			input:      "12*",
			entrypoint: 1,
			substr:     "12",
		},
		{
			name:       "starting end ending with non numeric",
			input:      "*3*",
			entrypoint: 1,
			substr:     "3",
		},
		{
			input:      "1*123*1",
			entrypoint: 3,
			substr:     "123",
		},
		{
			input:      "1*123",
			entrypoint: 3,
			substr:     "123",
		},
		{
			input:      "123",
			entrypoint: 0,
			substr:     "123",
		},
		{
			input:      "1",
			entrypoint: 0,
			substr:     "1",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.substr, getNumericalSubstring(tc.input, tc.entrypoint))
		})
	}

}

func getNumericalSubstring(input string, entrypoint int) string {
	offset := getOffset(input, entrypoint)
	limit := getLimit(input, entrypoint)
	return input[offset:limit]
}

func getLimit(input string, index int) int {
	// cut out constraints first.

	// if the current value is not a number, return index itself
	if _, err := strconv.Atoi(string(input[index])); err != nil {
		return index
	}

	for {
		value := input[index]
		_, err := strconv.Atoi(string(value))
		if err != nil {
			break
		}
		index++
		if index >= len(input) {
			break
		}
	}
	return index
}

func getOffset(input string, index int) int {
	if _, err := strconv.Atoi(string(input[index])); err != nil {
		return index
	}

	for {
		value := string(input[index])
		_, err := strconv.Atoi(value)
		if err != nil {
			return index + 1
		}
		index--
		if index < 0 {
			break
		}
	}
	if index < 0 {
		index = 0
	}
	return index
}
