package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var (
	playerScore   int = 0
	opponentScore int = 0
)

func rps_battle(opponentMoveKey string, playerMoveKey string) {

	opponentKey := map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
	}

	playerKey := map[string]string{
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}

	opponentMove := opponentKey[opponentMoveKey]
	playerMove := playerKey[playerMoveKey]

	weakness := map[string]string{
		"rock":     "scissors",
		"paper":    "rock",
		"scissors": "paper",
	}

	scoreValues := map[string]int{
		"rock":       1,
		"paper":      2,
		"scissors":   3,
		"winScore":   6,
		"drawScore":  3,
		"looseScore": 0,
	}

	// a draw
	if opponentMove == playerMove {
		opponentScore = opponentScore + scoreValues["drawScore"] + scoreValues[opponentMove]
		playerScore = playerScore + scoreValues["drawScore"] + scoreValues[playerMove]
	}

	// if the opponents move defeats the players move
	if weakness[opponentMove] == playerMove {
		// should we increment a value and not return? probably
		opponentScore = opponentScore + scoreValues[opponentMove] + scoreValues["winScore"]
		playerScore = playerScore + scoreValues[playerMove]
	}

	// if the player move defeats opponent move
	if weakness[playerMove] == opponentMove {
		playerScore = playerScore + scoreValues[playerMove] + scoreValues["winScore"]
		opponentScore = opponentScore + scoreValues[opponentMove]
	}

}

func main() {
	input, err := ioutil.ReadFile("puzzleInput.txt")

	if err != nil {
		log.Fatal("There was an error opening the file:", err)
	}

	rounds := strings.Split(string(input), "\n")

	// for each round of rock,paper,scissors - initiate the rock paper scissors battle!
	for i := 0; i < len(rounds); i++ {
		moves := strings.Split(rounds[i], " ")
		rps_battle(moves[0], moves[1])
	}

	fmt.Println("The players score is: ", playerScore)
	fmt.Println("The opponents score is: ", opponentScore)
}
