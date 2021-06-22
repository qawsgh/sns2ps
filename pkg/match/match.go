// Package match represents the details of a match in Shoot 'n Score It that are
// required to create Competitor list for Practiscore. Many fields are not
// currently used.
package match

import (
	"encoding/json"
	"log"
)

// UnmarshalJSON unmarshals the json content provided for use by other functions
func (f *MatchEntries) UnmarshalJSON(bs []byte) error {
	return json.Unmarshal(bs, &f.MatchEntries)
}

// MatchEntries is a list of MatchEntry items
type MatchEntries struct {
	MatchEntries []MatchEntry
}

// MatchEntry is a representation of a match
type MatchEntry struct {
	Model string `json:"model"`
	Pk    uint64 `json:"pk"`
	Squad Match  `json:"fields"`
}

// The Match contains all fields required to populate a competitor list for import
// into Practiscore, plus some that may be used for validation in the future.
type Match struct {
	Pk             uint64
	MatchName      string `json:"name"`
	Region         string `json:"region"`
	MaxCompetitors uint64 `json:"max_competitors"`
	Comment        string `json:"comment"`
	Registration   string `json:"registration"`
	Prematch       string `json:"prematch"`
	Event          uint64 `json:"event"`
	Firearms       string `json:"firearms"`
}

// GetMatch nmarshals the supplied json, create a list of Match items from this, and
// return the first item in the matches list.
func GetMatch(byteValue []byte) Match {
	allMatches := MatchEntries{}
	err := json.Unmarshal(byteValue, &allMatches)
	if err != nil {
		log.Printf("Failed to unmarshal match")
		log.Println(err)
	}
	matches := GetMatchFromJSON(allMatches)
	// Because we are using top-level json array, we get a list of matches, even
	// though there is only 1 - we return the first.
	return matches[0]
}

// GetMatchFromJSON gets the match from json and create a list of Match items with the
// appropriate values.
func GetMatchFromJSON(matchJSON MatchEntries) []Match {
	var matches []Match
	for i := 0; i < len(matchJSON.MatchEntries); i++ {
		m := matchJSON.MatchEntries[i].Squad
		m.Pk = matchJSON.MatchEntries[i].Pk
		matches = append(matches, m)
	}
	return matches
}
