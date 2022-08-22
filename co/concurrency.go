package concurrency

import (
	"encoding/json"
	"fmt"
	E "github.com/MehrbanooEbrahimzade/FootballTeams/Entitis"
	T "github.com/MehrbanooEbrahimzade/FootballTeams/Tools"
	"io/ioutil"
	"math"
	"net/http"
	"sync"
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
var allPlayers E.Players
var selectedTeams []E.Team
var wg sync.WaitGroup
var mutex sync.Mutex

func Concurrency() {
	teams := make(chan E.Team, 200)

	getTeams(teams)
	SeparateSelectedTeams(teams)

	for _, tm := range selectedTeams {
		wg.Add(1)
		go getPlayers(tm)
	}
	wg.Wait()

	for _, tm := range selectedTeams {
		wg.Add(1)
		go addListOfTeam(tm)
	}
	wg.Wait()

	fmt.Println("len of players in co is ", len(allPlayers))
	fmt.Println("len of selected team in co is ", len(selectedTeams))
}
func addListOfTeam(team E.Team) {
	defer wg.Done()
	for _, tp := range team.Players {
		for i, ap := range allPlayers {
			wg.Add(1)
			go func(aplr E.Player, tplr E.Player, index int) {
				defer wg.Done()
				if aplr.Name == tplr.Name {
					aplr.ListOfTeams = append(aplr.ListOfTeams, team.Name)
					allPlayers[index] = aplr
				}
			}(ap, tp, i)
		}
	}

}
func getTeams(c chan E.Team) {
	for i := 1; i < math.MaxInt8; i++ {
		wg.Add(1)
		go func(indexer int) {
			defer wg.Done()
			response, err := http.Get(fmt.Sprintf("https://api-origin.onefootball.com/score-one-proxy/api/teams/en/%d.json", indexer))
			if err != nil {
				return
			}
			responseData, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return
			}
			var footballData E.FootballData
			json.Unmarshal(responseData, &footballData)
			c <- footballData.Data.Team
		}(i)
	}
	wg.Wait()
	close(c)
}
func SeparateSelectedTeams(c chan E.Team) {
	for team := range c {
		wg.Add(1)
		go func(tm E.Team) {
			defer wg.Done()
			if T.Contains(names, tm.Name) {
				selectedTeams = append(selectedTeams, tm)
			}
		}(team)
	}
	wg.Wait()
}
func getPlayers(tm E.Team) {
	defer wg.Done()
	for _, plr := range tm.Players {
		wg.Add(1)
		go func(p E.Player) {
			defer wg.Done()
			mutex.Lock()
			if !allPlayers.Contains(p) {
				allPlayers = append(allPlayers, p)
			}
			mutex.Unlock()
		}(plr)
	}
}
