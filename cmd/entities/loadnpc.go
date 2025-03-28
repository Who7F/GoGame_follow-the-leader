package entities

type NPCData struct {
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
	FollowsLast bool    `json:"followslast"`
	ImagePath   string  `json:"imagePath"`
}

func LoadNPCs(jsonFile string) ([]NPCData, error) {
	return LoadJSON[NPCData](jsonFile, "Rum")
}
