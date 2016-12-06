package hero

import (
	"database/sql"
	"fmt"

	"github.com/bluefoxcode/rome/services/Nero/lib/util"
)

var table = "hero"

// Item defines the model.
type Item struct {
	ID          uint32 `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
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
    SELECT id, name, description
    FROM %v`, table))
	return result, err == sql.ErrNoRows, err
}

// ByID gets single item by ID
func ByID(db Connection, ID string) (Item, bool, error) {
	result := Item{}
	err := db.Get(&result, fmt.Sprintf(`
    SELECT id, name, description
    FROM %v
	WHERE id = $1`, table), ID)

	return result, err == sql.ErrNoRows, err
}

// Create adds an item
func Create(db Connection, name string, description string) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf(`
	INSERT INTO %v
	(name, description)
	VALUES
	($1, $2)
	`, table),
		name, description)
	return result, err
}

// Initialize sets up the database and prepopulates it.
func Initialize(db Connection) {
	var err error
	err = createTable(db)
	util.CheckErr(err)

	count, _, err := getCount(db)

	if count < 1 {
		populateDB(db)
	}

}
