package entities

import (
	"encoding/json"
	"fmt"
	"os"
)

type RumData struct {
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	GiveDrunk uint    `json:"GiveDrunk"`
	ImagePath string  `json:"imagePath"`
}

func LoadRums(jsonFile string) ([]RumData, error) {
	file, err := os.Open(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open Rum file: %v", err)
	}

	defer file.Close()

	rumData := []RumData{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&rumData); err != nil {
		return nil, fmt.Errorf("failed to decode Rum file: %v", err)
	}

	return rumData, nil
}
