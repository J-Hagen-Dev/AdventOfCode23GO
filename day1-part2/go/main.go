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

			for i, r := range runes {

				if unicode.IsDigit(r) {
					s = string(r)
					break
				}

				if n := checkAlphabeticalDigits(runes[:i+1]); n > 0 {
					s = strconv.Itoa(n)
					break
				}
			}

			for j := len(runes) - 1; j >= 0; j-- {
				//fmt.Println("Backward char", string(runes[j]), "\tindex", j)
				//fmt.Println("Backward slice", string(runes[j:]), "\tindex", j)
				if unicode.IsDigit(runes[j]) {
					s = s + string(runes[j])
					break
				}

				if n := checkAlphabeticalDigits(runes[j:]); n > 0 {
					s = s + strconv.Itoa(n)
					break
				}
			}

			if c, err := strconv.Atoi(s); err == nil {
				total += c
			}
		}
	}

	fmt.Printf("Total: %d\n", total)
}

func checkAlphabeticalDigits(runes []rune) int {
	rLen := len(runes)

	charNums := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for i := 3; i <= 5; i++ {
		if rLen < i {
			break
		}

		for key := range charNums {
			if key == string(runes[rLen-i:]) || key == string(runes[:i]) {
				return charNums[key]
			}
		}
	}

	return -1
}
