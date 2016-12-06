package race

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bluefoxcode/rome/services/Caesar/lib/util"
)

var table = "race"

// Item defines the model.
type Item struct {
	RaceID      string `db:"raceid" json:"raceID"`
	RaceName    string `db:"racename" json:"raceName"`
	First       string `db:"firstname" json:"firstName"`
	Last        string `db:"lastname" json:"lastName"`
	Party       string `db:"party" json:"party"`
	CandidateID string `db:"candidateID" json:"candidateID"`
	VoteCount   int32  `db:"voteCount" json:"voteCount"`
	Source      string `db:"source" json:"source"`
}

// Connection is an interface for making queries.
type Connection interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

// *******************
// External functions
// *******************

// List gets all items.
func List(db Connection) ([]Item, bool, error) {
	var result []Item
	err := db.Select(&result, fmt.Sprintf(`
    SELECT raceID, raceName
    FROM %v`, table))
	return result, err == sql.ErrNoRows, err
}

// Create adds an item
func Create(db Connection, item Item) (sql.Result, error) {
	i, _, err := countOne(db, item)
	var result sql.Result
	if err != nil {
		log.Println("ERROR>>>>>>>>>>>>>>>>>", err)
	}

	if i < 1 {
		result, err = db.Exec(fmt.Sprintf(`
		INSERT INTO %v
		(raceid, racename, firstname, lastname, party, candidateID,voteCount, source)
		VALUES
		($1,$2, $3, $4, $5, $6, $7, $8)
		`, table),
			item.RaceID, item.RaceName, item.First, item.Last, item.Party, item.CandidateID, item.VoteCount, item.Source)
	} else {
		result, err = db.Exec(fmt.Sprintf(`
		UPDATE %v
		SET (raceid, racename, firstname, lastname, party, candidateID,voteCount, source) =
		($1,$2, $3, $4, $5, $6, $7, $8)
		WHERE raceid = $1
		AND candidateID = $6
		AND source = $8`, table),
			item.RaceID, item.RaceName, item.First, item.Last, item.Party, item.CandidateID, item.VoteCount, item.Source)
	}

	return result, err
}

// Initialize sets up the database and prepopulates it.
func Initialize(db Connection) {
	var err error
	err = createTable(db)
	util.CheckErr(err)

	// count, _, err := getCount(db)

	// if count < 1 {
	// 	populateDB(db)
	// }

}

func countOne(db Connection, item Item) (int, bool, error) {
	var result int
	err := db.Get(&result, fmt.Sprintf(`
		Select COUNT(*)
		FROM %v
		WHERE raceid = $1 
		AND candidateid = $2
		AND source = $3
		`, table), item.RaceID, item.CandidateID, item.Source)

	return result, err == sql.ErrNoRows, err
}
