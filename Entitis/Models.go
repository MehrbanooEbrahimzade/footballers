package Entitis

type Player struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Age         string   `json:"age"`
	ListOfTeams []string `json:"their teams"`
}

type Team struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Players []Player `json:"players"`
}
type Data struct {
	Team Team `json:"team"`
}
type FootballData struct {
	Data Data `json:"data"`
}

type Players []Player

func (ps Players) Contains(p Player) bool {
	for _, a := range ps {
		if a.ID == p.ID {
			return true
		}
	}
	return false
}
