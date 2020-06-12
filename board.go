package craftboard

import (
	"encoding/json"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/pkg/errors"
)

// Board placeholder type def
type Board map[string]interface{}

// GetBoardFromFile reads the board from the json file
func GetBoardFromFile(file string) (Board, error) {

	jsonFile, err := os.Open(file)
	if err != nil {
		log.Warnf("couldn't open JSON board file: %v", err)
		return nil, errors.Wrap(err, "couldn't open JSON board file")
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Warnf("couldn't read JSON board file: %v", err)
		return nil, errors.Wrap(err, "couldn't open read board file")
	}

	board := Board{}
	json.Unmarshal(byteValue, &board)

	return board, nil
}

// SaveBoardToFile writes the board to the json file
func SaveBoardToFile(file string, board Board) error {

	data, err := json.MarshalIndent(board, "", "    ")
	if err != nil {
		log.Warnf("couldn't marshall JSON board: %v", err)
		return errors.Wrap(err, "couldn't marshall board")
	}

	err = ioutil.WriteFile(file, data, 0644)
	if err != nil {
		log.Warnf("couldn't write marshalled JSON board to file: %v", err)
		return errors.Wrap(err, "couldn't write marshalled board to file")
	}

	return nil
}
