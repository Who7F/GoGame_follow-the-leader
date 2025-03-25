package entities

import (
	"encoding/json"
	"fmt"
	"os"
)

type NPCData struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	FollowsLast bool `json:"followslast"`
	ImagePath string `json:"imagePath"`
}

func LoadNPCs(jsonFile string) ([]NPCData, error) {
	file, err := os.Open(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open NPC file: %v", err)
	}
	
	defer file.Close()
	
	npcData := []NPCData{}
	decoder := json.NewDecoder(file)
	if err:= decoder.Decode(&npcData); err != nil {
		return nil, fmt.Errorf("failed to decode NPC file: %v", err)
	}
	
	return npcData, nil
}