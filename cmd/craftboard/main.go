package main

import (
	"io/ioutil"
	"os"

	"github.com/DenBeke/craftboard"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const sampleBoardFile = "./board_sample.json"

func copyBoardIfNotExists(file string) error {
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		// file exists, so nothing to copy
		log.Printf("Found a board file in %q.", file)
		return nil
	}

	log.Printf("%q doesn't exist. Copying the sample board file now...")

	input, err := ioutil.ReadFile(sampleBoardFile)
	if err != nil {
		return errors.Wrap(err, "couldn't read sample board file")
	}

	err = ioutil.WriteFile(file, input, 0644)
	if err != nil {
		return errors.Wrapf(err, "couldn't copy sample board to destination: %q", file)
	}

	log.Printf("Successfully copied sample board file")

	return nil
}

func main() {

	config := craftboard.Config{
		BoardFile: craftboard.GetEnv("BOARD_FILE", "board.json"),
	}

	err := copyBoardIfNotExists(config.BoardFile)
	if err != nil {
		log.Fatalln("couldn't create board file")
	}

	craftboard.Serve(config)
}
