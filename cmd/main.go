// Create a competitor registration csv file representing a match on
// Shoot 'n Score It formatted for importing into practiscore

// This currently only supports IPSC matches, specifically
//   - action air
//   - mini rifle
//   - shotgun

package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"

	"github.com/qawsgh/sns2ps/pkg/categories"
	"github.com/qawsgh/sns2ps/pkg/divisions"
	"github.com/qawsgh/sns2ps/pkg/entities"
	"github.com/qawsgh/sns2ps/pkg/practiscorecsv"
	"github.com/qawsgh/sns2ps/pkg/regions"
	"github.com/qawsgh/sns2ps/pkg/requests"
	"golang.org/x/term"
)

var Version = "development"

// The arguments function ensures that all required arguments are set as flags to the
// command, or requests these interactively from the user.
func arguments(
	matchIDIn string, usernameIn string, passwordIn string, useLocalFiles bool, versionIn bool) (
	matchID string, username string, password string) {

	matchID = matchIDIn
	username = usernameIn
	password = passwordIn
	// useLocal = useLocalFilesIn
	displayVersion := versionIn

	if displayVersion {
		fmt.Println("Version: ", Version)
		os.Exit(0)
	}

	if !useLocalFiles {
		if matchID == "" {
			fmt.Print("Enter the Match ID: ")
			fmt.Scanln(&matchID)
		}
		if username == "" {
			fmt.Print("Enter your Shoot 'n Score It login (normally e-mail address): ")
			fmt.Scanln(&username)
		}
		if password == "" {
			fmt.Print("Enter your Shoot 'n Score It password (nothing will show when you type here): ")
			enteredPassword, _ := term.ReadPassword(int(syscall.Stdin))
			password = string(enteredPassword)
		}
	}
	fmt.Printf("\n\n")
	return matchID, username, password
}

func main() {
	var matchID string
	var username string
	var password string
	var useLocal bool
	var version bool

	flag.StringVar(&matchID, "m", "", "ID of the match in Shoot 'n Score It (shorthand)")
	flag.StringVar(&matchID, "matchid", "", "ID of the match in Shoot 'n Score It")
	flag.StringVar(&username, "u", "", "Your Shoot 'n Score It login (e-mail address normally) (shorthand)")
	flag.StringVar(&username, "user", "", "Your Shoot 'n Score It login (e-mail address normally)")
	flag.StringVar(&password, "p", "", "Your Shoot 'n Score It password (shorthand)")
	flag.StringVar(&password, "password", "", "Your Shoot 'n Score It password")
	flag.BoolVar(&useLocal, "uselocal", false, "Use local json files instead of requesting from Shoot 'n Score It")
	flag.BoolVar(&version, "v", false, "Display version information and quit (shorthand)")
	flag.BoolVar(&version, "version", false, "Display version information and quit")

	flag.Parse()

	matchID, username, password = arguments(matchID, username, password, useLocal, version)

	matchURL := "https://shootnscoreit.com/api/ipsc/match/" + matchID + "/"
	squadsURL := matchURL + "squads/"
	competitorsURL := matchURL + "competitors/"

	// Get categories, divisions and regions from respective packages
	categories := categories.Categories()
	divisions := divisions.Divisions()
	regions := regions.Regions()

	match, err := entities.Match(matchURL, username, password, useLocal)
	if err != nil {
		os.Exit(2)
	}

	fmt.Printf("Generating competitor list for \"%v\"\n", match.MatchName)

	squads, err := entities.Squads(squadsURL, username, password, useLocal)
	if err != nil {
		re := err.(*requests.HTTPError)
		if re.StatusCode == 404 {
			fmt.Println("Could not get squad info for this match - please check that squads are defined in Shoot 'n Score It")
		} else {
			fmt.Println("Unknown error trying to get squad info")
		}
		os.Exit(2)
	}
	fmt.Printf("Found %d squads\n", len(*squads))

	competitors, err := entities.Competitors(competitorsURL, categories, divisions, *match, regions, *squads, username, password, useLocal)
	if err != nil {
		re := err.(*requests.HTTPError)
		if re.StatusCode == 404 {
			fmt.Println("Could not get competitor info for this match")
		} else {
			fmt.Println("Unknown error trying to get competitor info")
		}
		os.Exit(2)
	}
	fmt.Printf("Found %d competitors\n", len(*competitors))

	// Write competitor information to Practiscore CSV file
	csvContent := practiscorecsv.CSVContent(*competitors)
	practiscorecsv.WriteCSV(csvContent, *match)
}
