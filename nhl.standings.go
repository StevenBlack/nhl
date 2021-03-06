package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

type (
	nhlStandings struct {
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

func standings() {
	url := urls["standings"]

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

	var stats nhlStandings
	err = json.Unmarshal(data, &stats)
	if err != nil {
		fmt.Println("error:", err)
	}

	// create and load standings
	standings := make([]Team, 0)

	for _, r := range stats.Records {
		lnsr := 0
		for _, tr := range r.TeamRecords {
			lnsr = lnsr + 1
			lr := tr.LeagueRecord
			overallr := tr.Records.OverallRecords
			wl10 := 0
			for _, v := range overallr {
				if v.Type == "lastTen" {
					wl10 = v.Wins - v.Losses
				}
			}
			team := Team{
				Conference: r.Conference.Name,
				Division:   r.Division.Name,
				Team:       tr.Team.Name,
				W:          lr.Wins,
				L:          lr.Losses,
				Wl:         lr.Wins - lr.Losses,
				WL10:       wl10,
				GP:         tr.GamesPlayed,
				GD:         tr.GoalsScored - tr.GoalsAgainst}
			standings = append(standings, team)
		}
	}

	sort.Sort(ByDivision(standings))

	section("NHL Division Standings")
	conf := ""
	div := ""
	ln := 0
	for _, s := range standings {
		if conf != s.Conference {
			conf = s.Conference
			ln = 0
			fmt.Println()
			fmt.Println(s.Conference, "Conference")
		}
		if div != s.Division {
			div = s.Division
			ln = 0
			fmt.Println()
			fmt.Println(s.Division, "Division")
		}
		ln = ln + 1
		teamline(ln, s)
	}

	section("NHL Wildcard Standings")
	conf = ""
	div = ""
	ln = 0

	for i, s := range standings {
		if conf != s.Conference {
			conf = s.Conference
			ln = 0
			fmt.Println()
			fmt.Println(s.Conference, "Conference")
		}
		if div != s.Division {
			div = s.Division
			ln = 0
			fmt.Println()
			fmt.Println(s.Division, "Division")
		}
		ln = ln + 1
		if ln < 4 {
			standings[i].WCFlag = 1
			teamline(ln, s)
		}
	}
	sort.Sort(ByConference(standings))
	for _, s := range standings {
		if conf != s.Conference {
			conf = s.Conference
			ln = 6
			fmt.Println()
			fmt.Println(s.Conference, "Conference Wildcards")
		}
		if s.WCFlag > 0 {
			continue
		}
		ln = ln + 1
		teamline(ln, s)
	}

	sort.Sort(ByConference(standings))

	section("NHL Conference Standings")
	conf = ""
	ln = 0
	for _, s := range standings {
		if conf != s.Conference {
			conf = s.Conference
			ln = 0
			fmt.Println()
			fmt.Println(s.Conference, "Conference")
		}
		ln = ln + 1
		teamline(ln, s)
	}

	sort.Sort(ByWl(standings))

	section("NHL League Standings")

	ln = 0
	for _, s := range standings {
		ln = ln + 1
		teamline(ln, s)
	}

	sort.Sort(By10Wl(standings))

	section("NHL Hot or Not, last 10")

	ln = 0
	for _, s := range standings {
		ln = ln + 1
		teamline(ln, s)
	}

}
func teamline(ln int, s Team) {
	fmt.Printf("%2d", ln)
	fmt.Print(" ")
	fmt.Printf("%-21v", s.Team)
	fmt.Print(" ")
	fmt.Printf("%4d", s.Wl)
	fmt.Print(" ")
	fmt.Printf("%4d", s.WL10)
	fmt.Print(" ")
	fmt.Printf("%4d", s.GP)
	fmt.Print(" ")
	fmt.Printf("%4d", s.GD)
	fmt.Println()
}

func section(title string) {
	fmt.Println()
	fmt.Println(strings.Repeat("=", 44))
	fmt.Println(title)
	fmt.Println(strings.Repeat("=", 44))
	fmt.Println("	                  +/-  +/10  GP   GD")
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
