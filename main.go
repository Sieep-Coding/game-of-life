package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"

	"golang.org/x/term"
)

const universeHeight = 40

var universeWidth int

type Cell bool

type Universe [][]Cell

var dir = [][]int{
	{-1, -1}, {0, -1}, {1, -1},
	{1, 0}, {1, 1},
	{0, 1}, {-1, 1}, {-1, 0},
}

func newUniverse(density float64) Universe {
	universe := make(Universe, universeHeight)
	for i := range universe {
		universe[i] = make([]Cell, universeWidth)
		for j := range universe[i] {
			universe[i][j] = Cell(rand.Float64() < density)
		}
	}
	return universe
}

func countNeighbors(u Universe, x, y int) int {
	count := 0
	for _, d := range dir {
		row := (y + d[0] + universeHeight) % universeHeight
		col := (x + d[1] + universeWidth) % universeWidth
		if u[row][col] {
			count++
		}
	}
	return count
}

func (u Universe) next(survivalRule, birthRule func(int) bool) Universe {
	newUniverse := make(Universe, universeHeight)
	for i := range newUniverse {
		newUniverse[i] = make([]Cell, universeWidth)
		for j := range newUniverse[i] {
			neighbors := countNeighbors(u, j, i)
			cell := u[i][j]
			if cell {
				newUniverse[i][j] = Cell(survivalRule(neighbors))
			} else {
				newUniverse[i][j] = Cell(birthRule(neighbors))
			}
		}
	}
	return newUniverse
}

func printUniverse(u *Universe) {
	for i := range *u {
		for j := range (*u)[i] {
			if (*u)[i][j] {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func getTerminalSize() (int, int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0, 0
	}
	return width, height
}

func applyPatterns(u *Universe) {
	patterns := []struct {
		x, y    int
		pattern [][]Cell
	}{
		{x: universeWidth / 4, y: universeHeight / 4, pattern: gliderPattern},
		{x: 3 * universeWidth / 4, y: universeHeight / 4, pattern: pentominoPattern},
		{x: universeWidth / 4, y: 3 * universeHeight / 4, pattern: spaceshipPattern},
		{x: 3 * universeWidth / 4, y: 3 * universeHeight / 4, pattern: gliderGunPattern},
	}

	for _, p := range patterns {
		for i := 0; i < len(p.pattern); i++ {
			for j := 0; j < len(p.pattern[i]); j++ {
				row := (p.y + i) % universeHeight
				col := (p.x + j) % universeWidth
				(*u)[row][col] = p.pattern[i][j]
			}
		}
	}
}

var gliderPattern = [][]Cell{
	{false, true, false},
	{false, false, true},
	{true, true, true},
}

var pentominoPattern = [][]Cell{
	{false, true, true},
	{true, true, false},
	{false, true, false},
}

var spaceshipPattern = [][]Cell{
	{false, false, true, true, false},
	{true, true, false, true, true},
	{true, true, true, true, false},
	{false, true, true, false, false},
}

var gliderGunPattern = [][]Cell{
	{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, true, false, false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false, false, false, true, true, false, false, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, true, true},
	{false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, true, true},
	{true, true, false, false, false, false, false, false, false, false, true, false, false, false, false, false, true, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
	{true, true, false, false, false, false, false, false, false, false, true, false, false, false, true, false, true, true, false, false, false, false, true, false, true, false, false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, true, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
}

func main() {
	rand.Seed(time.Now().UnixNano())

	width, _ := getTerminalSize()
	universeWidth = width

	density := 0.3
	universe := newUniverse(density)

	survivalRule := func(neighbors int) bool {
		return neighbors == 2 || neighbors == 3
	}
	birthRule := func(neighbors int) bool {
		return neighbors == 3
	}

	applyPatterns(&universe)

	for {
		clearScreen()

		printUniverse(&universe)

		time.Sleep(100 * time.Millisecond)

		universe = universe.next(survivalRule, birthRule)
	}
}
