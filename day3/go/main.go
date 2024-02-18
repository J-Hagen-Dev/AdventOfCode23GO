package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Cell struct {
	CellType       string
	Value          rune
	SymbolAdjacent bool
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
	s := ""
	adjacent := false

	for _, row := range cellTable {

		for j, cell := range row {

			if cell.CellType == "n" {
				s += string(cell.Value)

				adjacent = adjacent || cell.SymbolAdjacent
			}

			if s != "" && (cell.CellType != "n" || j == len(row)-1) {

				if n, err := strconv.Atoi(s); err != nil {
					fmt.Println(err.Error())

					return

				} else if adjacent {
					total += n
				}

				s, adjacent = "", false
			}
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
				if cellTable[i-1][j].CellType == "s" {
					cellTable[i][j].SymbolAdjacent = true

					continue
				} else if j > 0 && cellTable[i-1][j-1].CellType == "s" {
					cellTable[i][j].SymbolAdjacent = true

					continue
				} else if j < len(row)-1 && cellTable[i-1][j+1].CellType == "s" {
					cellTable[i][j].SymbolAdjacent = true

					continue
				}
			}

			if i < len(cellTable)-1 {
				if cellTable[i+1][j].CellType == "s" {
					cellTable[i][j].SymbolAdjacent = true

					continue
				} else if j > 0 && cellTable[i+1][j-1].CellType == "s" {
					cellTable[i][j].SymbolAdjacent = true

					continue
				} else if j < len(row)-1 && cellTable[i+1][j+1].CellType == "s" {
					cellTable[i][j].SymbolAdjacent = true

					continue
				}
			}

			if j > 0 && cellTable[i][j-1].CellType == "s" {
				cellTable[i][j].SymbolAdjacent = true

				continue
			}

			if j < len(row)-1 && cellTable[i][j+1].CellType == "s" {
				cellTable[i][j].SymbolAdjacent = true

				continue
			}
		}
	}

	return cellTable
}
