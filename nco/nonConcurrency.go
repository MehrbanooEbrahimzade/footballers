package nonConcurrency

import (
	"encoding/json"
	"fmt"
	E "github.com/MehrbanooEbrahimzade/FootballTeams/Entitis"
	T "github.com/MehrbanooEbrahimzade/FootballTeams/Tools"
	"io/ioutil"
	"net/http"
)

type teamNames []string

var names = teamNames{
	"Germany",
	"England",
	"France",
	"Spain",
	"Manchester United",
	"Arsenal",
	"Chelsea",
	"Barcelona",
	"Real Madrid",
	"Bayern Munich",
}
var selectedTeams []E.Team
var allPlayers E.Players
var allTeams []E.Team

func NonConcurrency() {
	getTeams()
	getPlayers()
	addListOfTeam()
	fmt.Println("len of players in Nco is ", len(allPlayers))
	fmt.Println("len of selected team in Nco is ", len(selectedTeams))
}
func getTeams() {
	for i := 1; len(selectedTeams) != len(names); i++ {

		response, err := http.Get(fmt.Sprintf("https://api-origin.onefootball.com/score-one-proxy/api/teams/en/%d.json", i))
		if err != nil {
			continue
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			continue
		}
		var footballData E.FootballData

		json.Unmarshal(responseData, &footballData)
		allTeams = append(allTeams, footballData.Data.Team)
		if T.Contains(names, footballData.Data.Team.Name) {
			selectedTeams = append(selectedTeams, footballData.Data.Team)
		}
	}
}
func getPlayers() {
	for _, Team := range selectedTeams {
		for _, playerOfSelectedTeam := range Team.Players {
			if !allPlayers.Contains(playerOfSelectedTeam) {
				allPlayers = append(allPlayers, playerOfSelectedTeam)
			}
		}
	}
}
func addListOfTeam() {
	for _, team := range selectedTeams {
		for _, tp := range team.Players {
			for i, ap := range allPlayers {
				if ap.Name == tp.Name {
					ap.ListOfTeams = append(ap.ListOfTeams, team.Name)
					allPlayers[i] = ap
				}
			}
		}
	}
}
