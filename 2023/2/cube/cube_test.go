package cube

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCubeCount(t *testing.T) {
	testcase := []struct {
		bag                Bag
		game               Game
		bagContainsInTotal map[CubeColor]int
		gamePossible       bool
	}{
		{

			game: Game{
				bag: Bag{cubeShows: []map[CubeColor]int{{CubeColorRed: 10}}},
			},
			bagContainsInTotal: map[CubeColor]int{CubeColorRed: 11},
			gamePossible:       true,
		},
		{
			game: Game{
				bag: Bag{cubeShows: []map[CubeColor]int{{CubeColorRed: 10}}},
			},
			bagContainsInTotal: map[CubeColor]int{CubeColorRed: 5},
			gamePossible:       false,
		},
	}

	for _, tc := range testcase {
		_, gamePossible := tc.game.isPossible(tc.bagContainsInTotal)
		assert.Equal(t, tc.gamePossible, gamePossible)
	}
}

func TestCubeConondrom(t *testing.T) {
	testcase := []struct {
		input       string
		possibility map[CubeColor]int
		expected    bool
	}{
		{
			input: `
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`,
			possibility: map[CubeColor]int{CubeColorBlue: 14, CubeColorGreen: 13, CubeColorRed: 12},
			expected:    true,
		},
	}

	for _, tc := range testcase {
		games, err := ParseGames(tc.input)
		assert.NoError(t, err)
		assert.Len(t, games, 5)

		assert.Equal(t, 8, games.SumOfPossibleGameIDs(tc.possibility))
	}
}

func TestParseInput(t *testing.T) {
	input := `
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	games, err := ParseGames(input)
	if err != nil {
		t.Fatalf("error is not nil")
	}
	assert.Len(t, games, 5)
	assert.Len(t, games[0].bag.cubeShows, 3)
	assert.Len(t, games[1].bag.cubeShows, 3)
	assert.Len(t, games[2].bag.cubeShows, 3)
	assert.Len(t, games[3].bag.cubeShows, 3)
	assert.Len(t, games[4].bag.cubeShows, 2)
}
