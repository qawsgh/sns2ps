// The practiscorecsv package uses Competitor and Match information to create a
// csv file in an appropriate format for importing competitor registration into
// Practiscore.
package practiscorecsv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/qawsgh/sns2ps/pkg/competitors"
	"github.com/qawsgh/sns2ps/pkg/match"
)

// WriteCSV opens a csv file, writes the headings that Practiscore uses to identify
// content in the row, and then for each competitor, creates a row representing
// their registration details.
// The match name from Shoot 'n Score It is used as the CSV filename.
func WriteCSV(competitors []competitors.Competitor, match match.Match) {
	var data [][]string
	filename := strings.ReplaceAll(match.MatchName+".csv", " ", "_")
	fmt.Printf("\nCreating CSV named \"%v\"\n", filename)

	var headings = []string{"number", "first name", "last name", "email", "phone",
		"squad", "age", "category", "gender", "division", "power factor", "class",
		"special", "team", "region"}
	data = append(data, headings)

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot open file", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// BUG(qawsgh): Category2 is required due to Practiscore iOS not handling
	// 'super junior' correctly in the Category field.
	for c := range competitors {
		comp_number := strconv.Itoa(competitors[c].Number)
		var csvContent = []string{comp_number, competitors[c].FirstName, competitors[c].LastName,
			competitors[c].Email, "", competitors[c].Squad, competitors[c].Category, competitors[c].Category2,
			competitors[c].Sex, competitors[c].Division, "", competitors[c].Classification, "", "",
			competitors[c].Region}
		data = append(data, csvContent)
	}

	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			log.Fatal("Cannot write to file", err)
		}
	}
	fmt.Printf("Finished creating competitor csv - you can now import this to Practiscore\n")
}
