package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

const author = "Steven Black (https://github.com/StevenBlack/nhl)"
const appVersion = "Version 0.1.3 (Jan 15 2019)"
const description = "NHL plaintext standings and stats"

var urls = map[string]string{
	"standings": "https://statsapi.web.nhl.com/api/v1/standings?expand=standings.record",
	"schedule":  "https://statsapi.web.nhl.com/api/v1/schedule",
}

type teams []struct {
	ID           int      `json:"id"`
	Abbreviation string   `json:"abbreviation"`
	City         string   `json:"city"`
	Alias        string   `json:"alias"`
	Selected     bool     `json:"selected"`
	Aliases      []string `json:"aliases"`
}

type (
	statistics struct {
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
)

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

func main() {

	// Process flags
	version := flag.Bool("v", false, "prints current version")
	description := flag.Bool("d", false, "prints a description of this utility")
	author := flag.Bool("a", false, "prints the author information")
	flag.Parse()
	if *version {
		fmt.Println(appVersion)
		os.Exit(0)
	}

	if *description {
		fmt.Println(description)
		os.Exit(0)
	}

	if *author {
		fmt.Println(author)
		os.Exit(0)
	}

	// Sanitize options
	options := os.Args[1:]
	if len(options) > 0 {
		// lowercase the options
		for n := range options {
			options[n] = strings.ToLower(options[n])
		}
	}

	var mode = reckonMode(options)
	switch mode {
	case "standings":
		standings()
	default:
		schedule()
	}
}

func reckonMode(opt []string) string {
	scoringAliases := [...]string{"assists", "scoring", "goals", "points"}
	if any(opt, scoringAliases) {
		return "scoring"
	}

	scoresAliases := [...]string{"score", "scores"}
	if any(opt, scoresAliases) {
		return "scores"
	}

	scheduleAliases := [...]string{"sched", "schedule", "sked", "games"}
	if any(opt, scheduleAliases) {
		return "schedule"
	}

	return "standings"
}
