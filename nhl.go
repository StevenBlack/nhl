package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"time"
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

type ByWl []Team

func (c ByWl) Len() int           { return len(c) }
func (c ByWl) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByWl) Less(i, j int) bool { return c[i].Wl > c[j].Wl }

type ByDivision []Team

func (c ByDivision) Len() int      { return len(c) }
func (c ByDivision) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c ByDivision) Less(i, j int) bool {
	if c[i].Conference == c[j].Conference {
		if c[i].Division == c[j].Division {
			return c[i].Wl > c[j].Wl
		} else {
			return c[i].Division < c[j].Division
		}
	} else {
		return c[i].Conference < c[j].Conference
	}
}

type ByConference []Team

func (c ByConference) Len() int      { return len(c) }
func (c ByConference) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c ByConference) Less(i, j int) bool {
	if c[i].Conference == c[j].Conference {
		return c[i].Wl > c[j].Wl
	} else {
		return c[i].Conference < c[j].Conference
	}
}

func main() {
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
		// fmt.Println(r.Conference.Name, r.Division.Name)
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
				Wl:         (lr.Wins - lr.Losses),
				WL10:       wl10}
			standings = append(standings, team)
		}
	}

	sort.Sort(ByDivision(standings))

	fmt.Println("NHL Division Standings")
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
		fmt.Printf("%2d", ln)
		fmt.Print(" ")
		fmt.Printf("%-25v", s.Team)
		fmt.Print(" ")
		fmt.Printf("%4d", s.Wl)
		fmt.Print(" ")
		fmt.Printf("%4d", s.WL10)
		fmt.Println()
	}

	sort.Sort(ByConference(standings))

	fmt.Println()
	fmt.Println("NHL Conference Standings")
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
		fmt.Printf("%2d", ln)
		fmt.Print(" ")
		fmt.Printf("%-25v", s.Team)
		fmt.Print(" ")
		fmt.Printf("%4d", s.Wl)
		fmt.Print(" ")
		fmt.Printf("%4d", s.WL10)
		fmt.Println()
	}

	sort.Sort(ByWl(standings))

	fmt.Println()
	fmt.Println("NHL League Standings")

	ln = 0
	for _, s := range standings {
		ln = ln + 1
		fmt.Printf("%2d", ln)
		fmt.Print(" ")
		fmt.Printf("%-25v", s.Team)
		fmt.Print(" ")
		fmt.Printf("%4d", s.Wl)
		fmt.Print(" ")
		fmt.Printf("%4d", s.WL10)
		fmt.Println()
	}

	// fmt.Print(standings)

}
