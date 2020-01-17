package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const author = "Steven Black (https://github.com/StevenBlack/nhl)"
const appVersion = "Version 0.1.3 (Jan 15 2019)"
const description = "NHL plaintext standings and stats"

var urls = map[string]string{
	"standings": "https://statsapi.web.nhl.com/api/v1/standings?expand=standings.record",
	"schedule":  "https://statsapi.web.nhl.com/api/v1/schedule",
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
