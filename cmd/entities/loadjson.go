package entities

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadJSON[T any](jsonFile string, placeholder string) ([]T, error) {
	file, err := os.Open(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s file: %v", placeholder, err)
	}

	defer file.Close()

	var jsonData []T
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&jsonData); err != nil {
		return nil, fmt.Errorf("failed to decode %s file: %v", placeholder, err)
	}

	return jsonData, nil
}
