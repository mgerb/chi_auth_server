package db

import (
	"database/sql"
	"log"

	"github.com/mgerb/chi_auth_server/config"

	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

var Conn *sql.DB

func Connect() {

	dbConfig := "postgres://" +
		config.Config.Database.User + ":" +
		config.Config.Database.Password + "@" +
		config.Config.Database.Address + "/" +
		config.Config.Database.Name +
		"?sslmode=disable"

	var err error
	Conn, err = sql.Open("postgres", dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(Conn, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)

	err = m.Up()

	if err != nil {
		log.Println(err)
	}

	version, dirty, err := m.Version()

	if err != nil {
		log.Println(err)
	} else {
		log.Print("Migration Version: ", version)
		if dirty {
			log.Println("Migration is dirty.")
		}
	}
}
