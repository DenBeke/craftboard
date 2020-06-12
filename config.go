package craftboard

import "os"

// Config for CraftBoard
type Config struct {
	// BoardFile is the default file to which the board will be saved/read.
	BoardFile string
}

// GetEnv reads environment variable. If empty it falls back to the fallback value
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
