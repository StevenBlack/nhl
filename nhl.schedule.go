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

	client := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

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

	// create and load schedule
	fmt.Println(sched)

}

// sort by wins and losses, last 10 games, league wide
type By10Wl []Team

func (c By10Wl) Len() int      { return len(c) }
func (c By10Wl) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c By10Wl) Less(i, j int) bool {
	if c[i].WL10 == c[j].WL10 {
		if c[i].GP == c[j].GP {
			return c[i].GD > c[j].GD
		}
		return c[i].GP < c[j].GP

	}
	return c[i].WL10 > c[j].WL10
}

// sort by wins and losses, league wide
type ByWl []Team

func (c ByWl) Len() int      { return len(c) }
func (c ByWl) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c ByWl) Less(i, j int) bool {
	if c[i].Wl == c[j].Wl {
		if c[i].GP == c[j].GP {
			return c[i].GD > c[j].GD
		}
		return c[i].GP < c[j].GP

	}
	return c[i].Wl > c[j].Wl
}

// sort by division
type ByDivision []Team

func (c ByDivision) Len() int      { return len(c) }
func (c ByDivision) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c ByDivision) Less(i, j int) bool {
	if c[i].Conference == c[j].Conference {
		if c[i].Division == c[j].Division {
			if c[i].Wl == c[j].Wl {
				if c[i].GP == c[j].GP {
					return c[i].GD > c[j].GD
				}
				return c[i].GP < c[j].GP
			}
			return c[i].Wl > c[j].Wl
		}
		return c[i].Division < c[j].Division
	}
	return c[i].Conference < c[j].Conference
}

// sort by conference
type ByConference []Team

func (c ByConference) Len() int      { return len(c) }
func (c ByConference) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c ByConference) Less(i, j int) bool {
	if c[i].Conference == c[j].Conference {
		if c[i].Wl == c[j].Wl {
			if c[i].GP == c[j].GP {
				return c[i].GD > c[j].GD
			}
			return c[i].GP < c[j].GP
		}
		return c[i].Wl > c[j].Wl
	}
	return c[i].Conference < c[j].Conference
}
