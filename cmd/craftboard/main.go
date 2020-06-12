package main

import (
	"github.com/DenBeke/craftboard"
)

func main() {

	config := craftboard.Config{
		BoardFile: craftboard.GetEnv("BOARD_FILE", "board.json"),
	}

	craftboard.Serve(config)
}
