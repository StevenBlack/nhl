package main

import (
	"fmt"
	"time"

	"github.com/thedevsaddam/gojsonq"
)

func main() {
	fmt.Println("NHL Standings Interpreter")

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

	jq := gojsonq.New().File("./data.json")

	// res := jq.From("records.[0].teamRecords").Get()
	// fmt.Println(res)

	res2 := jq.From("records.[].teamRecords.[3].team.name").Get()
	fmt.Println(res2)

}
