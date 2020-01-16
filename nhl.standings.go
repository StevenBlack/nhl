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

func standings() {
	url := urls["standings"]

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

	var stats statistics
	err = json.Unmarshal(data, &stats)
	if err != nil {
		fmt.Println("error:", err)
	}

	// create and load standings
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
