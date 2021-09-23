package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/manulife-ca/aff-trivia-gophers/sharedtypes"
)


func UpdatePlayers(winners []string, resetScores bool) {
	file, _ := ioutil.ReadFile("teams/endzymex.json")
	playersData := sharedtypes.Players{}
	json.Unmarshal([]byte(file), &playersData)

	if resetScores {
		for i := 0; i < len(playersData.Players); i++ {
			playersData.Players[i].Score = 0
		}
	} else {
		for ii := 0; ii < len(winners); ii++ {
			for i := 0; i < len(playersData.Players); i++ {
				if playersData.Players[i].Name == winners[ii] {
					playersData.Players[i].Score = playersData.Players[i].Score + 1
				}
			}
		}
	}

	newPlayersJson, err := json.Marshal(playersData)
	if err != nil {
		panic(err)
	}
	writeErr := ioutil.WriteFile("teams/endzymex.json", newPlayersJson, 0777)

	if writeErr != nil {
		panic(writeErr)
	}
}

func UpdateTeam(teamName string) {

}

func GetScores(w http.ResponseWriter, r *http.Request) {

	jsonFile, err := os.Open("teams/endzymex.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)

	var players sharedtypes.Players

	json.Unmarshal(byteValue, &players)
	json.NewEncoder(w).Encode(players)
}

func GetScore(teamName string) sharedtypes.Players {

	jsonFile, err := os.Open("teams/"+teamName+".json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)

	var players sharedtypes.Players

	json.Unmarshal(byteValue, &players)
	return players
}
