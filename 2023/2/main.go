package main

import (
	"fmt"
	"log"
	"minhajthekhan/advent/2023/2/cube"
	"os"
)

func main() {

	b, err := os.ReadFile("2023/2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	games, err := cube.ParseGames(string(b))
	if err != nil {
		log.Fatal(err)
	}

	sum := games.SumOfPossibleGameIDs(map[cube.CubeColor]int{
		cube.CubeColorRed:   12,
		cube.CubeColorGreen: 13,
		cube.CubeColorBlue:  14,
	})

	fmt.Println(sum)

}
