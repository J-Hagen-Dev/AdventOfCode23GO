package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

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
			runes, s := []rune(string(line[:])), ""

			for _, r := range runes {
				if unicode.IsDigit(r) {
					s = string(r)
					break
				}
			}

			for j := len(runes) - 1; j >= 0; j-- {
				if unicode.IsDigit(runes[j]) {
					s = s + string(runes[j])
					break
				}
			}

			if c, err := strconv.Atoi(s); err == nil {
				total += c
			}
		}
	}

	fmt.Printf("Total: %d", total)
}
