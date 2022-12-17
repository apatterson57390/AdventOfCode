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

func rps_battle(opponentMoveKey string, outcomeMoveKey string) {

	opponentKey := map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
	}

	opponentMove := opponentKey[opponentMoveKey]

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

	switch outcomeMoveKey {
	// player loose
	case "X":
		playerMove := weakness[opponentMove]
		opponentScore = opponentScore + scoreValues["winScore"] + scoreValues[opponentMove]
		playerScore = playerScore + scoreValues["looseScore"] + scoreValues[playerMove]
	// player and opponent draw
	case "Y":
		playerMove := opponentMove
		opponentScore = opponentScore + scoreValues["drawScore"] + scoreValues[opponentMove]
		playerScore = playerScore + scoreValues["drawScore"] + scoreValues[playerMove]
	// player win
	// for this challenge, we can infer the winner the the existing map by saying
	// that the value of the thing that it itself defeats is what defeats it
	case "Z":
		playerMove := weakness[weakness[opponentMove]]
		opponentScore = opponentScore + scoreValues["looseScore"] + scoreValues[opponentMove]
		playerScore = playerScore + scoreValues["winScore"] + scoreValues[playerMove]
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
