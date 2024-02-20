package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ScratchCard struct {
	ID             int
	NumbersRolled  []int
	WinningNumbers []int
	Score          int
}

func main() {
	total := 0

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	r := bufio.NewReader(file)

	var cards []ScratchCard

	for {
		line, _, err := r.ReadLine()

		if err != nil || len(line) == 0 {
			break
		}

		if card, err := createCard(string(line[:])); err != nil {
			fmt.Println(err.Error())
			return
		} else {
			cards = append(cards, card)
		}
	}

	for _, card := range cards {
		total += card.Score
	}

	fmt.Println("Total:", total)
}

func arrayContains[T comparable](arr []T, v T) bool {
	for _, p := range arr {
		if v == p {
			return true
		}
	}
	return false
}

func calcScore(card *ScratchCard) {
	for _, n := range card.NumbersRolled {
		if arrayContains[int](card.WinningNumbers, n) {
			switch {
			case card.Score == 0:
				card.Score++
			case card.Score > 0:
				card.Score *= 2
			}
		}
	}
}

func createCard(input string) (ScratchCard, error) {
	var card ScratchCard
	cardLeading := strings.Split(input, ":")[0]
	cardBody := strings.Split(input, ":")[1]

	if n, err := strconv.Atoi(strings.Split(cardLeading, " ")[len(strings.Split(cardLeading, " "))-1]); err != nil {
		return card, err
	} else {
		card.ID = n
	}

	for i, s := range strings.Split(cardBody, "|") {
		numStrings := strings.Split(s, " ")

		for _, v := range numStrings {
			if v == "" {
				continue
			}
			if n, err := strconv.Atoi(v); err != nil {
				return card, err
			} else if i == 0 {
				card.NumbersRolled = append(card.NumbersRolled, n)
			} else {
				card.WinningNumbers = append(card.WinningNumbers, n)
			}
		}
	}

	calcScore(&card)

	return card, nil
}
