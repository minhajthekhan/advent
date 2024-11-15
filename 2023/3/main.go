package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	b, err := os.ReadFile("2023/3/input.txt")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(b), "\n")
	for _, row := range rows {
		fmt.Println(row)
	}
}
