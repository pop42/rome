package hero

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
			id serial,
			name text, 
			description text
			)`, table))

	return err
}

// populateDB prepopulates the DB if it was empty
func populateDB(db Connection) {

	defaultItems := []Item{
		Item{
			Name:        "Superman",
			Description: "Man of Steel",
		},
		Item{
			Name:        "Batman",
			Description: "Dark Knight",
		},
	}

	for _, item := range defaultItems {
		_, err := db.Exec(fmt.Sprintf(`
		INSERT INTO %v
		(name, description)
		VALUES
		($1,$2)
		`, table),
			item.Name, item.Description)
		util.CheckErr(err)
	}

}
