package cube

import (
	"strconv"
	"strings"
)

type CubeColor int

const (
	CubeColorUndefined = iota
	CubeColorRed
	CubeColorBlue
	CubeColorGreen
)

type Bag struct {
	cubeShows []map[CubeColor]int
}

type Record struct {
	cubesRecorded map[CubeColor][]int
}

func (r *Record) Add(cubeColor CubeColor, cubeCount int) {
	if _, ok := r.cubesRecorded[cubeColor]; !ok {
		r.cubesRecorded[cubeColor] = make([]int, 0)
	}

	c := r.cubesRecorded[cubeColor]
	c = append(c, cubeCount)
	r.cubesRecorded[cubeColor] = c
}

func (r *Record) hasMoreCubeThan(color CubeColor, count int) bool {
	max := -1
	for _, v := range r.cubesRecorded[color] {
		if max < v {
			max = v
		}
	}

	return count < max
}

func (b Bag) Shows() []map[CubeColor]int {
	return b.cubeShows
}

type Game struct {
	id  int
	bag Bag
}

type GameSlice []Game

func (g GameSlice) SumOfPossibleGameIDs(possibility map[CubeColor]int) int {
	sum := 0
	for _, game := range g {
		if id, ok := game.isPossible(possibility); ok {
			sum += id
		}
	}

	return sum
}

func (g Game) isPossible(possibility map[CubeColor]int) (int, bool) {
	record := Record{cubesRecorded: make(map[CubeColor][]int)}
	for _, show := range g.bag.Shows() {
		for cubeColor, cubeCount := range show {
			record.Add(cubeColor, cubeCount)
		}
	}

	for color, count := range possibility {
		if record.hasMoreCubeThan(color, count) {
			return 0, false
		}
	}
	return g.id, true
}

func ParseGames(input string) (GameSlice, error) {
	games := make([]Game, 0)
	for _, gameString := range strings.Split(input, "\n") {
		if gameString == "" {
			continue
		}
		x := strings.Split(gameString, ":")   // ["Game 1", ...]
		gameID := strings.Split(x[0], " ")[1] // ["Game", "1"]
		id, err := strconv.Atoi(gameID)
		if err != nil {
			return nil, err
		}

		bag, err := parseBag(x[1])
		if err != nil {
			return nil, err
		}
		games = append(games, Game{id: id, bag: bag})
	}
	return games, nil
}

func parseBag(input string) (Bag, error) {
	tupleString := strings.Split(input, ";") //

	bag := Bag{cubeShows: make([]map[CubeColor]int, 0)}
	for _, tuple := range tupleString {
		cubeShow, err := parseCubeShow(tuple)
		if err != nil {
			return Bag{}, err
		}
		bag.cubeShows = append(bag.cubeShows, cubeShow)
	}
	return bag, nil
}

func parseCubeShow(tuple string) (map[CubeColor]int, error) {
	bagShow := make(map[CubeColor]int)
	for _, colorAndCount := range strings.Split(tuple, ",") {
		x := strings.Split(strings.TrimSpace(colorAndCount), " ")
		cubeCountString, cubeColorString := strings.TrimSpace(x[0]), strings.TrimSpace(x[1])
		cubeColor, cubeCount, err := parseColorAndCount(cubeColorString, cubeCountString)
		if err != nil {
			return nil, err
		}
		bagShow[cubeColor] = cubeCount
	}
	return bagShow, nil
}

func parseColorAndCount(color, count string) (CubeColor, int, error) {
	var cubeColor CubeColor
	switch color {
	case "red":
		cubeColor = CubeColorRed
	case "blue":
		cubeColor = CubeColorBlue
	case "green":
		cubeColor = CubeColorGreen
	}

	cubeCount, err := strconv.Atoi(count)
	if err != nil {
		return -1, -1, err
	}
	return cubeColor, cubeCount, nil
}
