package boot

import (
	"log"
	"os"
	"runtime"

	"github.com/bluefoxcode/rome/services/Caesar/lib/util"
	"github.com/bluefoxcode/rome/services/Caesar/model/race"
	"github.com/jmoiron/sqlx"
	// need postgres drivers
	_ "github.com/lib/pq"
)

// Info contains application settings.
type Info struct {
	Port        string
	DatabaseURL string
}

func init() {
	log.SetFlags(log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// LoadConfig loads the config object from env vars.
func LoadConfig() *Info {
	config := &Info{}

	config.Port = os.Getenv("PORT")
	config.DatabaseURL = os.Getenv("DATABASE_URL")
	return config
}

// RegisterServices sets up services.
func RegisterServices(config *Info) {
	db, err := sqlx.Open("postgres", config.DatabaseURL)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	util.StoreDB(db)

	race.Initialize(db)

}
