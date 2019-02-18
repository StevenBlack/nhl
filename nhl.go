package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	// "github.com/thedevsaddam/gojsonq"
)

type Stats struct {
	Copyright string `json:"copyright"`
	Records   []struct {
		StandingsType string `json:"standingsType"`
		League        struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"league"`
		Division struct {
			ID           int    `json:"id"`
			Name         string `json:"name"`
			NameShort    string `json:"nameShort"`
			Link         string `json:"link"`
			Abbreviation string `json:"abbreviation"`
		} `json:"division"`
		Conference struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"conference"`
		TeamRecords []struct {
			Team struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"team"`
			LeagueRecord struct {
				Wins   int    `json:"wins"`
				Losses int    `json:"losses"`
				Ot     int    `json:"ot"`
				Type   string `json:"type"`
			} `json:"leagueRecord"`
			GoalsAgainst   int    `json:"goalsAgainst"`
			GoalsScored    int    `json:"goalsScored"`
			Points         int    `json:"points"`
			DivisionRank   string `json:"divisionRank"`
			ConferenceRank string `json:"conferenceRank"`
			LeagueRank     string `json:"leagueRank"`
			WildCardRank   string `json:"wildCardRank"`
			Row            int    `json:"row"`
			GamesPlayed    int    `json:"gamesPlayed"`
			Streak         struct {
				StreakType   string `json:"streakType"`
				StreakNumber int    `json:"streakNumber"`
				StreakCode   string `json:"streakCode"`
			} `json:"streak"`
			Records struct {
				DivisionRecords []struct {
					Wins   int    `json:"wins"`
					Losses int    `json:"losses"`
					Ot     int    `json:"ot"`
					Type   string `json:"type"`
				} `json:"divisionRecords"`
				OverallRecords []struct {
					Wins   int    `json:"wins"`
					Losses int    `json:"losses"`
					Ot     int    `json:"ot,omitempty"`
					Type   string `json:"type"`
				} `json:"overallRecords"`
				ConferenceRecords []struct {
					Wins   int    `json:"wins"`
					Losses int    `json:"losses"`
					Ot     int    `json:"ot"`
					Type   string `json:"type"`
				} `json:"conferenceRecords"`
			} `json:"records"`
			LastUpdated time.Time `json:"lastUpdated"`
		} `json:"teamRecords"`
	} `json:"records"`
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
}

func main() {
	fmt.Println("NHL Standings Interpreter")

	data, err := ioutil.ReadFile("./data.json")
	if err != nil {
		fmt.Print(err)
	}

	var stats Stats
	err = json.Unmarshal(data, &stats)
	if err != nil {
		fmt.Println("error:", err)
	}

	// create amd load standings
	standings := make([]Team, 0)

	for _, r := range stats.Records {
		fmt.Println(r.Conference.Name, r.Division.Name)
		for _, tr := range r.TeamRecords {
			team := Team{
				Conference: r.Conference.Name,
				Division:   r.Division.Name,
				Team:       tr.Team.Name,
				W:          tr.LeagueRecord.Wins,
				L:          tr.LeagueRecord.Losses,
				Wl:         (tr.LeagueRecord.Wins - tr.LeagueRecord.Losses)}
			standings = append(standings, team)
			fmt.Printf("%-25v", tr.Team.Name)
			fmt.Println(tr.LeagueRecord.Wins - tr.LeagueRecord.Losses)
		}
		fmt.Println(" ")
	}

	fmt.Print(standings)
}
