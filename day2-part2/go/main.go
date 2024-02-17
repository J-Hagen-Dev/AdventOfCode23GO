package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
			input := string(line[:])

			if n, err := getPower(input); err != nil {
				fmt.Println(err.Error())
				return
			} else {
				total += n
			}
		}
	}
	fmt.Println("Total:", total)
}

func getPower(input string) (int, error) {
	minR, minG, minB := 0, 0, 0
	suffix := strings.Split(input, ":")[1]

	for _, s := range strings.Split(suffix, ";") {

		for _, v := range strings.Split(s, ",") {

			draw := strings.Split(strings.Trim(v, " "), " ")

			if n, err := strconv.Atoi(draw[0]); err == nil {

				switch draw[1] {

				case "red":
					if n > minR {
						minR = n
					}

				case "blue":
					if n > minB {
						minB = n
					}

				case "green":
					if n > minG {
						minG = n
					}

				}
			} else {
				return 0, err
			}
		}
	}

	return minB * minG * minR, nil
}
