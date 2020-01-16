package main

import (
	"time"
)

type schedule struct {
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
