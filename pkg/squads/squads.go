// Squads represent the details of a squad in Shoot 'n Score It that are required
// to add squad numbers to Competitor structs.
// A competitor entity in Shoot 'n Score It will contain a key that maps to a
// squad entity.
// We need the squad number for the csv, not this key.
package squads

import (
	"encoding/json"
	"log"
)

// Unmarshal the json content provided for use by other functions
func (f *SquadEntries) UnmarshalJSON(bs []byte) error {
	return json.Unmarshal(bs, &f.SquadEntries)
}

type SquadEntries struct {
	SquadEntries []SquadEntry
}

type SquadEntry struct {
	Model string `json:"model"`
	Pk    uint64 `json:"pk"`
	Squad Squad  `json:"fields"`
}

// A Squad contains all items needed to populate a Competitor and some that
// are not yet used.
// Only Pk and Number are used in the creation of the csv.
type Squad struct {
	Pk             uint64
	Number         uint64 `json:"number"`
	MaxCompetitors uint64 `json:"max_competitors"`
	Comment        string `json:"comment"`
	Registration   string `json:"registration"`
	Prematch       string `json:"prematch"`
	Event          uint64 `json:"event"`
}

// GetSquads unmarshals the supplied json content to create a list of
// Squad items that is then returned
func GetSquads(byteValue []byte) []Squad {
	allSquads := SquadEntries{}
	err := json.Unmarshal(byteValue, &allSquads)
	if err != nil {
		log.Printf("Failed to unmarshal squads")
	}
	squads := GetSquadsFromJSON(allSquads)
	return squads
}

// Get the squads from json and create a list of Squad items with the appropriate
// values.
func GetSquadsFromJSON(squadJSON SquadEntries) []Squad {
	var squads []Squad
	for i := 0; i < len(squadJSON.SquadEntries); i++ {
		squad := squadJSON.SquadEntries[i].Squad
		squad.Pk = squadJSON.SquadEntries[i].Pk
		squads = append(squads, squad)
	}
	return squads
}
