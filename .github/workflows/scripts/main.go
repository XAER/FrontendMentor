package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type ChallengeJSON struct {
	ChallengesDirs []string `json:"challengesDirectories"`
}

func arrayContains(array []string, element string) bool {
	for _, a := range array {
		if a == element {
			return true
		}
	}
	return false
}

func main() {
	// Get a list of the directories in the current directory
	dirs, err := os.ReadDir("../../../")
	if err != nil {
		panic(err)
	}

	for _, dir := range dirs {
		fmt.Println(dir.Name())
	}

	// Getting the challenges.json file and unmarshaling it
	challengesJSONFile, err := os.ReadFile("../../../challenges.json")
	if err != nil {
		panic(err)
	}

	var challengesJSON ChallengeJSON
	err = json.Unmarshal(challengesJSONFile, &challengesJSON)

	if err != nil {
		panic(err)
	}

	// Loop through the directories
	for _, dir := range dirs {
		// Check if this is a directory
		if dir.IsDir() {
			if !arrayContains(challengesJSON.ChallengesDirs, dir.Name()) && !strings.HasPrefix(dir.Name(), ".") {
				challengesJSON.ChallengesDirs = append(challengesJSON.ChallengesDirs, dir.Name())
			}
		}
	}

	// Marshal the challengesJSON struct
	challengesJSONBytes, err := json.Marshal(challengesJSON)
	if err != nil {
		panic(err)
	}

	// Write the challenges.json file
	err = os.WriteFile("challenges.json", challengesJSONBytes, 0644)

	if err != nil {
		panic(err)
	}
}
