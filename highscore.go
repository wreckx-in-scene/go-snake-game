package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveHighScore(score int) {
	data, err := json.Marshal(score)
	if err != nil {
		fmt.Println("Error saving your score : ", err)
		return
	}

	//writing in the json file
	os.WriteFile("scores.json", data, 0644)
}

func LoadHighScore() int {
	data, err := os.ReadFile("scores.json")
	if err != nil {
		fmt.Println("Error loading your score : ", err)
		return 0
	}

	var score int
	json.Unmarshal(data, &score)
	return score
}
