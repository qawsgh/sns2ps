// Return entities for various components including competitors, match, and squads
package entities

import (
	"github.com/qawsgh/sns2ps/pkg/competitors"
	"github.com/qawsgh/sns2ps/pkg/match"
	"github.com/qawsgh/sns2ps/pkg/requests"
	"github.com/qawsgh/sns2ps/pkg/squads"
)

// Competitors returns an internal representation of the competitors for a match using
// the Competitors package.
// If useLocal is True, instead of requesting details from the web API, local files will be
// used. This should be used for testing and development only.
func Competitors(
	url string, categories map[string]string, divisions map[string]string, match match.Match, regions map[string]string,
	squads []squads.Squad, username string, password string, useLocal bool) []competitors.Competitor {

	var body []byte
	if useLocal {
		body = requests.FileRequest("sample_content/sg_competitors.json")
	} else {
		body = requests.WebRequest(url, username, password)
	}
	competitors := competitors.GetCompetitors(body, categories, divisions, match, regions, squads)
	return competitors
}

// Match returns an internal representation of a match using the Match package.
// If useLocal is True, instead of requesting details from the web API, local files will be
// used. This should be used for testing and development only.
func Match(url string, username string, password string, useLocal bool) match.Match {
	var body []byte
	if useLocal {
		body = requests.FileRequest("sample_content/sg_match.json")
	} else {
		body = requests.WebRequest(url, username, password)
	}

	m := match.GetMatch(body)
	return m
}

// getSquads returns an internal representation of the squads for a match using the
// Squads package.
// If useLocal is True, instead of requesting details from the web API, local files will be
// used. This should be used for testing and development only.
func Squads(url string, username string, password string, useLocal bool) []squads.Squad {
	var body []byte
	if useLocal {
		body = requests.FileRequest("sample_content/squads.json")
	} else {
		body = requests.WebRequest(url, username, password)
	}

	squads := squads.GetSquads(body)
	return squads
}
