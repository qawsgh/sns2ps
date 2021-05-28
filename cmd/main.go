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
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"syscall"

	"github.com/qawsgh/sns2ps/pkg/categories"
	"github.com/qawsgh/sns2ps/pkg/competitors"
	"github.com/qawsgh/sns2ps/pkg/divisions"
	"github.com/qawsgh/sns2ps/pkg/match"
	"github.com/qawsgh/sns2ps/pkg/practiscorecsv"
	"github.com/qawsgh/sns2ps/pkg/regions"
	"github.com/qawsgh/sns2ps/pkg/squads"
	"golang.org/x/term"
)

var Version = "development"

// fileRequest opens a file and reads the content, returning a byteValue.
// It is intended to be called by any function that needs to get json content for
// testing instead of using the web.
func fileRequest(filename string) []byte {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)

	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

// webRequest will request content from the specified Shoot 'n Score It API endpoint.
// Endpoints are provided as the URL string and should be the complete URL including
// the match ID, and type of request (competitor, match, squad).
// username and password values are used to authenticate to Shoot 'n Score It.
func webRequest(url string, username string, password string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.SetBasicAuth(username, password)
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		fmt.Printf("Could not connect to %v - check your Internet connection and make sure you can access the site in a browser\n", url)
		os.Exit(2)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Could not connect to %v - check your Internet connection and make sure you can access the site in a browser\n", url)
		os.Exit(2)
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == 401 {
			fmt.Println("Unable to get details for this match - please check your username and password and try again.")
		} else if resp.StatusCode == 404 {
			fmt.Println("Unable to get details for this match - please check your matchID and try again.")
		} else {
			fmt.Println("Unknown error: ", resp.StatusCode)
		}
		os.Exit(2)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}
	return body
}

// getMatch returns an internal representation of a match using the Match package.
// If useLocal is True, instead of requesting details from the web API, local files will be
// used. This should be used for testing and development only.
func getMatch(url string, username string, password string, useLocal bool) match.Match {
	var body []byte
	if useLocal {
		body = fileRequest("sample_content/sg_match.json")
	} else {
		body = webRequest(url, username, password)
	}

	m := match.GetMatch(body)
	return m
}

// getSquads returns an internal representation of the squads for a match using the
// Squads package.
// If useLocal is True, instead of requesting details from the web API, local files will be
// used. This should be used for testing and development only.
func getSquads(url string, username string, password string, useLocal bool) []squads.Squad {
	var body []byte
	if useLocal {
		body = fileRequest("sample_content/squads.json")
	} else {
		body = webRequest(url, username, password)
	}

	squads := squads.GetSquads(body)
	return squads
}

// getCompetitors returns an internal representation of the competitors for a match using
// the Competitors package.
// If useLocal is True, instead of requesting details from the web API, local files will be
// used. This should be used for testing and development only.
func getCompetitors(
	url string, categories map[string]string, divisions map[string]string, match match.Match, regions map[string]string,
	squads []squads.Squad, username string, password string, useLocal bool) []competitors.Competitor {

	var body []byte
	if useLocal {
		body = fileRequest("sample_content/sg_competitors.json")
	} else {
		body = webRequest(url, username, password)
	}
	competitors := competitors.GetCompetitors(body, categories, divisions, match, regions, squads)
	return competitors
}

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

	baseURL := "https://shootnscoreit.com/api/ipsc/match/"
	matchURL := baseURL + matchID + "/"
	squadsURL := baseURL + matchID + "/squads/"
	competitorsURL := baseURL + matchID + "/competitors/"

	// Get categories, divisions and regions from respective packages
	categories := categories.Categories()
	divisions := divisions.Divisions()
	regions := regions.Regions()

	match := getMatch(matchURL, username, password, useLocal)
	fmt.Printf("Generating competitor list for \"%v\"\n", match.MatchName)

	squads := getSquads(squadsURL, username, password, useLocal)
	fmt.Printf("Found %d squads\n", len(squads))

	competitors := getCompetitors(competitorsURL, categories, divisions, match, regions, squads, username, password, useLocal)
	fmt.Printf("Found %d competitors\n", len(competitors))

	// Write competitor information to Practiscore CSV file
	practiscorecsv.WriteCSV(competitors, match)
}
