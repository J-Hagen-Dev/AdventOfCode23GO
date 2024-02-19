package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Coord struct {
	Y int
	X int
}

type Cell struct {
	CellType string
	Value    rune
	Coord    Coord
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

	var runeTable [][]rune

	for {
		line, _, err := r.ReadLine()

		if err != nil {
			break
		}

		if len(line) > 0 {
			runeTable = append(runeTable, []rune(string(line[:])))
		}
	}

	cellTable := createCellTable(runeTable)

	gears := map[Coord][]int{}

	for _, row := range cellTable {
		coords := []Coord{{Y: -1, X: -1}}
		s := ""

		for j, cell := range row {

			if cell.CellType == "n" {
				s += string(cell.Value)

				if cell.Coord.X != -1 {
					coords = append(coords, cell.Coord)
				}
			}

			if s != "" && (cell.CellType != "n" || j == len(row)-1) {
				if n, err := strconv.Atoi(s); err != nil {
					fmt.Println(err.Error())

					return

				} else {
					for _, c := range deduplicateList[Coord](coords) {
						if c.X != -1 {
							gears[c] = append(gears[c], n)
						}
					}
				}
				coords = []Coord{{Y: -1, X: -1}}
				s = ""
			}
		}
	}

	for key := range gears {
		if len(gears[key]) == 2 {
			num := 1
			for _, n := range gears[key] {
				num *= n
			}
			total += num
		}
	}

	fmt.Println("Total:", total)
}

func createCellTable(runeTable [][]rune) [][]Cell {
	var cellTable [][]Cell

	for i, row := range runeTable {
		cellTable = append(cellTable, []Cell{})

		for _, r := range row {
			var cell Cell

			if r == 46 {
				cell.CellType = "p"

			} else if r >= 48 && r <= 57 {
				cell.CellType = "n"
				cell.Value = r
				cell.Coord = Coord{X: -1, Y: -1}

			} else if r == 42 {
				cell.CellType = "g"

			} else {
				cell.CellType = "s"
			}

			cellTable[i] = append(cellTable[i], cell)
		}
	}

	for i, row := range cellTable {

		for j, cell := range row {

			if cell.CellType != "n" {
				continue
			}
			if i > 0 {
				if cellTable[i-1][j].CellType == "g" {
					cellTable[i][j].Coord = Coord{Y: i - 1, X: j}

					continue
				} else if j > 0 && cellTable[i-1][j-1].CellType == "g" {
					cellTable[i][j].Coord = Coord{Y: i - 1, X: j - 1}

					continue
				} else if j < len(row)-1 && cellTable[i-1][j+1].CellType == "g" {
					cellTable[i][j].Coord = Coord{Y: i - 1, X: j + 1}

					continue
				}
			}

			if i < len(cellTable)-1 {
				if cellTable[i+1][j].CellType == "g" {
					cellTable[i][j].Coord = Coord{Y: i + 1, X: j}

					continue
				} else if j > 0 && cellTable[i+1][j-1].CellType == "g" {
					cellTable[i][j].Coord = Coord{Y: i + 1, X: j - 1}

					continue
				} else if j < len(row)-1 && cellTable[i+1][j+1].CellType == "g" {
					cellTable[i][j].Coord = Coord{Y: i + 1, X: j + 1}

					continue
				}
			}

			if j > 0 && cellTable[i][j-1].CellType == "g" {
				cellTable[i][j].Coord = Coord{Y: i, X: j - 1}

				continue
			}

			if j < len(row)-1 && cellTable[i][j+1].CellType == "g" {
				cellTable[i][j].Coord = Coord{Y: i, X: j + 1}

				continue
			}
		}
	}

	return cellTable
}

func deduplicateList[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
