package race

import (
	"database/sql"
	"fmt"

	"github.com/bluefoxcode/rome/services/Nero/lib/util"
)

// getCount is used to see if there are any rows in the table
func getCount(db Connection) (int, bool, error) {
	var result int
	err := db.Get(&result, fmt.Sprintf(`
	SELECT COUNT(*)
	FROM %v`, table))
	return result, err == sql.ErrNoRows, err
}

// creatTable creates the table.
func createTable(db Connection) (err error) {
	_, err = db.Exec(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %v 
			(
			raceID text,
			racename text, 
			firstname text,
            lastname text,
            party text,
            candidateID text,
            voteCount int,
            source text
			)`, table))

	return err
}

// populateDB prepopulates the DB if it was empty
func populateDB(db Connection) {

	defaultItems := []Item{
		Item{
			"0",
			"President of the United States",
			"Donald",
			"Trump",
			"GOP",
			"9677",
			0,
			"TEST",
		},
		Item{"0",
			"President of the United States",
			"Hillary",
			"Clinton",
			"Dem",
			"9678",
			0,
			"TEST",
		},
	}

	for _, item := range defaultItems {
		_, err := db.Exec(fmt.Sprintf(`
		INSERT INTO %v
		(raceID, racename, firstname, lastname, party, candidateID,voteCount, source)
		VALUES
		($1,$2, $3, $4, $5, $6, $7, $8)
		`, table),
			item.RaceID, item.RaceName, item.First, item.Last, item.Party, item.CandidateID, item.VoteCount, item.Source)
		util.CheckErr(err)
	}

}
