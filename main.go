package main

import (
	"conway-game-of-life/life"
	"conway-game-of-life/render"
	"flag"
	"fmt"
	"time"
)

func main() {
	file := flag.String("file", "life_106.txt", "Path to input Life 1.06 file")
	gens := flag.Int("gens", 10, "Number of generations to simulate")
	cleanTerminal := flag.Bool("clear", false, "Clear terminal between generations")
	rate := flag.Int("rate", 500, "Update rate between generations in milliseconds")
	flag.Parse()

	board, err := life.LoadLife106(*file)
	if err != nil {
		panic(err)
	}
	renderer := render.GetRenderer()

	for i := 0; i <= *gens; i++ {
		if *cleanTerminal {
			fmt.Print("\033[2J\033[H")
		}
		fmt.Printf("Generation %d:\n", i)
		renderer(board)
		board = life.NextGen(board)
		time.Sleep(time.Duration(*rate) * time.Millisecond)
	}
}
