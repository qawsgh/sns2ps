// Package competitors represents the details of a competitor in Shoot 'n Score It that
// are required for creation of a competitor import file for Practiscore
package competitors

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/qawsgh/sns2ps/pkg/match"
	"github.com/qawsgh/sns2ps/pkg/squads"
)

// UnmarshalJSON unmarshals the json content provided for use by other functions
func (f *CompetitorEntries) UnmarshalJSON(bs []byte) error {
	return json.Unmarshal(bs, &f.CompetitorEntries)
}

// CompetitorEntries are a list of competitorEntry items
type CompetitorEntries struct {
	CompetitorEntries []CompetitorEntry
}

// A CompetitorEntry represents a competitor in Shoot 'n Score It
type CompetitorEntry struct {
	Model      string     `json:"model"`
	Pk         uint64     `json:"pk"`
	Competitor Competitor `json:"fields"`
}

// The Competitor represents all of the fields set for a competitor in Shoot 'n Score It
// that are needed to import the Competitor into Practiscore.
// Not all fields here are currently used, but may be in the future.
type Competitor struct {
	Pk             uint64
	Squad          string
	Division       string
	Number         int    `json:"number"`
	Shooter        int    `json:"shooter"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Sex            string `json:"sex"`
	Email          string `json:"email"`
	Region         string `json:"region"`
	Club           string `json:"club"`
	EventID        uint64 `json:"event"`
	SquadPk        uint64 `json:"squad"`
	HandgunDiv     string `json:"handgun_div"`
	HandgunPf      string `json:"handgun_pf"`
	RifleDiv       string `json:"rifle_div"`
	RiflePf        string `json:"rifle_pf"`
	MiniRifleDiv   string `json:"mini_rifle_div"`
	PrecRifleDiv   string `json:"prec_rifle_div"`
	PrecRiflePf    string `json:"prec_rifle_pf"`
	ShotgunDiv     string `json:"shotgun_div"`
	AirDiv         string `json:"air_div"`
	PccDiv         string `json:"pcc_div"`
	TournamentDiv  string `json:"tournament_division"`
	Category       string `json:"category"`
	Category2      string
	Classification string `json:"classification"`
}

// Get the competitor category by looking up the value set for the competitor key
// in the categories map.
func getCompetitorCategory(competitor Competitor, categories map[string]string) Competitor {
	competitor.Category2 = ""
	switch competitor.Category {
	case "-":
		competitor.Category = ""
	case "J":
		competitor.Category = "junior"
	case "SJ":
		// BUG(qawsgh): We need to set Super Junior in the Category2 field due to the iOS
		// version of Practiscore not handling this correctly.
		competitor.Category = ""
		competitor.Category2 = "Super Junior"
	case "L":
		competitor.Category = "lady"
	case "S":
		competitor.Category = "senior"
	case "SS":
		competitor.Category = "super senior"
	}

	return competitor
}

// Get the competitor division by looking up the code set for the competitor
// in the divisions map.
func getCompetitorDivision(competitor Competitor, divisions map[string]string,
	match match.Match) Competitor {
	var divCode string
	switch match.Firearms {
	case "ai":
		divCode = competitor.AirDiv
		competitor.AirDiv = divisions[divCode]
		competitor.Division = divisions[divCode]
	case "hg":
		divCode = competitor.HandgunDiv
		competitor.HandgunDiv = divisions[divCode]
		competitor.Division = divisions[divCode]
	case "mr":
		divCode = competitor.MiniRifleDiv
		competitor.MiniRifleDiv = divisions[divCode]
		competitor.Division = divisions[divCode]
	case "sg":
		divCode = competitor.ShotgunDiv
		competitor.ShotgunDiv = divisions[divCode]
		competitor.Division = divisions[divCode]
	}
	return competitor
}

// GetCompetitors parses the supplied json, calls GetCompetitorsFromJSON,
// and returns a list of Competitor entities
func GetCompetitors(byteValue []byte, categories map[string]string, divisions map[string]string,
	match match.Match, regions map[string]string, squads []squads.Squad) []Competitor {
	allCompetitors := CompetitorEntries{}
	err := json.Unmarshal(byteValue, &allCompetitors)
	if err != nil {
		log.Printf("Failed to unmarshal competitors")
		log.Println(err)
	}
	competitors := GetCompetitorsFromJSON(allCompetitors, categories, divisions, match, regions, squads)
	return competitors
}

// GetCompetitorsFromJSON takes the unmarshaled json reppresentation of a competitor,
// and then returns a list of Competitor items including information from other sources
// such as squad, category, region and division.
func GetCompetitorsFromJSON(competitorSON CompetitorEntries, categories map[string]string, divisions map[string]string,
	match match.Match, regions map[string]string, squads []squads.Squad) []Competitor {
	var competitors []Competitor
	for i := 0; i < len(competitorSON.CompetitorEntries); i++ {
		competitor := competitorSON.CompetitorEntries[i].Competitor
		competitor.Pk = competitorSON.CompetitorEntries[i].Pk
		for squad := range squads {
			if squads[squad].Pk == competitor.SquadPk {
				competitor.Squad = strconv.FormatUint(squads[squad].Number, 10)
			}
		}
		competitor = getCompetitorCategory(competitor, categories)
		competitor = getCompetitorDivision(competitor, divisions, match)
		competitors = append(competitors, competitor)
	}
	return competitors
}
