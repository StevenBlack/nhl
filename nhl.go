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
