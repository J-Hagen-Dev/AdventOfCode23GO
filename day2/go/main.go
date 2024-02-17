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

			if g, err := getDrawFromString(input); err != nil {
				fmt.Println(err.Error())
				return
			} else if g.Possible {
				total += g.ID
			}
		}
	}
	fmt.Println("Total:", total)
}

type Game struct {
	ID       int
	Drws     []Draw
	Possible bool
}

type Draw struct {
	R int
	G int
	B int
}

func getDrawFromString(input string) (Game, error) {
	var g Game
	g.Possible = true

	leading := strings.Split(input, ":")[0]
	suffix := strings.Split(input, ":")[1]

	if n, err := strconv.Atoi(strings.Split(leading, " ")[1]); err != nil {
		return g, err
	} else {
		g.ID = n
	}

	for _, s := range strings.Split(suffix, ";") {
		var d Draw

		for _, v := range strings.Split(s, ",") {

			draw := strings.Split(strings.Trim(v, " "), " ")

			if n, err := strconv.Atoi(draw[0]); err == nil {

				switch draw[1] {
				case "red":
					d.R = n
				case "blue":
					d.B = n
				case "green":
					d.G = n
				}
			}
		}

		if d.B > 14 || d.R > 12 || d.G > 13 {
			g.Possible = false
		}

		g.Drws = append(g.Drws, d)
	}

	return g, nil
}
