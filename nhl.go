package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/markbates/pkger"
	"github.com/thoas/go-funk"
)

const author = "Steven Black (https://github.com/StevenBlack/nhl)"
const appVersion = "Version 0.1.4 (Jan 18 2019)"
const description = "NHL plaintext standings and stats"

var urls = map[string]string{
	"standings": "https://statsapi.web.nhl.com/api/v1/standings?expand=standings.record",
	"schedule":  "https://statsapi.web.nhl.com/api/v1/schedule",
	"scoring":   "https://statsapi.web.nhl.com/api/v1/statTypes",
}

var client = http.Client{
	Timeout: time.Second * 4, // Maximum of 4 secs
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
	// unique options, please
	options = funk.UniqString(options)

	buf := bytes.NewBuffer(nil)
	f, err := pkger.Open("/metadata/teamdata.json")
	if err != nil {
		fmt.Print(err)
		return
	}
	io.Copy(buf, f)
	f.Close()

	var tm teams
	err = json.Unmarshal(buf.Bytes(), &tm)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	var mode = reckonMode(options)

	switch mode {
	case "standings":
		standings()
	case "scoring":
		var teamIds = reckonTeams(options, tm)
		scoring(teamIds)
	default:
		var teamIds = reckonTeams(options, tm)
		schedule(teamIds)
	}
}

func reckonTeams(opt []string, tm teams) []int {
	// analyze options string and return a slice of team ids
	// first, Selected set to false
	for _, t := range tm {
		t.Selected = false
	}
	// scour for teams
	for _, o := range opt {
		for ti, t := range tm {
			if !t.Selected {
				for _, a := range t.Aliases {
					if o == a {
						tm[ti].Selected = true
					}
				}
			}
		}
	}

	// prepare to return selected team ids
	var teamIds []int
	for _, t := range tm {
		if t.Selected {
			teamIds = append(teamIds, t.ID)
		}
	}
	return teamIds
}

func reckonMode(opt []string) string {
	scoringAliases := [...]string{"assists", "scoring", "goals", "points", "rookies", "forwards", "f", "defencemen", "defence", "d"}
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

type teams []struct {
	ID           int      `json:"id"`
	Abbreviation string   `json:"abbreviation"`
	City         string   `json:"city"`
	Alias        string   `json:"alias"`
	Selected     bool     `json:"selected"`
	Aliases      []string `json:"aliases"`
}
