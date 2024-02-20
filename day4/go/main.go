package main

import (
	"bufio"
	"fmt"
	"os"
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

	for {
		line, _, err := r.ReadLine()

		if err != nil {
			break
		}

		if len(line) > 0 {
			input := string(line[:])
		}
	}

	fmt.Println("Total:", total)
}
