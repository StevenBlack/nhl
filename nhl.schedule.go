package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type teams []struct {
	ID           int      `json:"id"`
	Abbreviation string   `json:"abbreviation"`
	City         string   `json:"city"`
	Alias        string   `json:"alias"`
	Selected     bool     `json:"selected"`
	Aliases      []string `json:"aliases"`
}

type Team struct {
	Conference string
	Division   string
	Team       string
	W          int
	L          int
	Wl         int
	W10        int
	L10        int
	WL10       int
	GP         int
	GD         int
	WCFlag     int
}

type gameschedule struct {
	Copyright    string `json:"copyright"`
	TotalItems   int    `json:"totalItems"`
	TotalEvents  int    `json:"totalEvents"`
	TotalGames   int    `json:"totalGames"`
	TotalMatches int    `json:"totalMatches"`
	Wait         int    `json:"wait"`
	Dates        []struct {
		Date         string `json:"date"`
		TotalItems   int    `json:"totalItems"`
		TotalEvents  int    `json:"totalEvents"`
		TotalGames   int    `json:"totalGames"`
		TotalMatches int    `json:"totalMatches"`
		Games        []struct {
			GamePk   int       `json:"gamePk"`
			Link     string    `json:"link"`
			GameType string    `json:"gameType"`
			Season   string    `json:"season"`
			GameDate time.Time `json:"gameDate"`
			Status   struct {
				AbstractGameState string `json:"abstractGameState"`
				CodedGameState    string `json:"codedGameState"`
				DetailedState     string `json:"detailedState"`
				StatusCode        string `json:"statusCode"`
				StartTimeTBD      bool   `json:"startTimeTBD"`
			} `json:"status"`
			Teams struct {
				Away struct {
					LeagueRecord struct {
						Wins   int    `json:"wins"`
						Losses int    `json:"losses"`
						Ot     int    `json:"ot"`
						Type   string `json:"type"`
					} `json:"leagueRecord"`
					Score int `json:"score"`
					Team  struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
						Link string `json:"link"`
					} `json:"team"`
				} `json:"away"`
				Home struct {
					LeagueRecord struct {
						Wins   int    `json:"wins"`
						Losses int    `json:"losses"`
						Ot     int    `json:"ot"`
						Type   string `json:"type"`
					} `json:"leagueRecord"`
					Score int `json:"score"`
					Team  struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
						Link string `json:"link"`
					} `json:"team"`
				} `json:"home"`
			} `json:"teams"`
			Content struct {
				Link string `json:"link"`
			} `json:"content"`
			Venue struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"venue,omitempty"`
		} `json:"games"`
		Events  []interface{} `json:"events"`
		Matches []interface{} `json:"matches"`
	} `json:"dates"`
}

func schedule() {
	url := urls["schedule"]

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Print(err)
	}

	req.Header.Set("User-Agent", "nhl-stats-api")

	res, getErr := client.Do(req)
	if getErr != nil {
		fmt.Print(getErr)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err)
	}

	var sched gameschedule
	err = json.Unmarshal(data, &sched)
	if err != nil {
		fmt.Println("error:", err)
	}

	for _, r := range sched.Dates {
		fmt.Println(r.Date)
		prefix := "  "
		for _, g := range r.Games {
			t := g.Teams
			a := t.Away.Team
			h := t.Home.Team
			switch g.Status.DetailedState {
			case "Final":
				fmt.Println(prefix, lastWord(a.Name), t.Away.Score, lastWord(h.Name), t.Home.Score)
			case "In Progress":
				fmt.Println(prefix, lastWord(a.Name), t.Away.Score, lastWord(h.Name), t.Home.Score)
			default:
				fmt.Println(prefix, lastWord(t.Away.Team.Name), "vs", lastWord(t.Home.Team.Name))
			}
		}
	}
	fmt.Println()

}
