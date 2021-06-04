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
	squads []squads.Squad, username string, password string, useLocal bool) (*[]competitors.Competitor, *requests.HTTPError) {

	var body []byte
	var err error
	if useLocal {
		body = requests.FileRequest("../sample_content/sg_competitors.json")
	} else {
		body, err = requests.WebRequest(url, username, password)
		if err != nil {
			re := err.(*requests.HTTPError)
			return nil, re
		}
	}
	competitors := competitors.GetCompetitors(body, categories, divisions, match, regions, squads)
	return &competitors, nil
}

// Match returns an internal representation of a match using the Match package.
// If useLocal is True, instead of requesting details from the web API, local files will be
// used. This should be used for testing and development only.
func Match(url string, username string, password string, useLocal bool) (*match.Match, *requests.HTTPError) {
	var body []byte
	var err error
	if useLocal {
		body = requests.FileRequest("../sample_content/sg_match.json")
	} else {
		body, err = requests.WebRequest(url, username, password)
		if err != nil {
			re := err.(*requests.HTTPError)
			return nil, re
		}
	}

	m := match.GetMatch(body)
	return &m, nil
}

// getSquads returns an internal representation of the squads for a match using the
// Squads package.
// If useLocal is True, instead of requesting details from the web API, local files will be
// used. This should be used for testing and development only.
func Squads(url string, username string, password string, useLocal bool) (*[]squads.Squad, *requests.HTTPError) {
	var body []byte
	var err error
	if useLocal {
		body = requests.FileRequest("../sample_content/squads.json")
	} else {
		body, err = requests.WebRequest(url, username, password)
		if err != nil {
			re := err.(*requests.HTTPError)
			return nil, re
		}
	}

	squads := squads.GetSquads(body)
	return &squads, nil
}
