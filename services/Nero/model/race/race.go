package race

import (
	"database/sql"
	"fmt"
)

var table = "race"

// Item defines the model.
type Item struct {
	RaceID      string `db:"raceid" json:"raceID"`
	RaceName    string `db:"racename" json:"raceName"`
	First       string `db:"firstname" json:"firstName"`
	Last        string `db:"lastname" json:"lastName"`
	Party       string `db:"party" json:"party"`
	CandidateID string `db:"candidateid" json:"candidateID"`
	VoteCount   int32  `db:"votecount" json:"voteCount"`
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
    SELECT *
    FROM %v`, table))
	return result, err == sql.ErrNoRows, err
}

// Wipeout will delete all race items
func Wipeout(db Connection) (err error) {
	_, err = db.Exec(fmt.Sprintf(`DELETE FROM %v `, table))

	return err
}
