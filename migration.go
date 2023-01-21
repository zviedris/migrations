package main

import (
	"database/sql"
	"embed"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

// embed sql directory where migrations is stored
// and can be accessed to run them
//

// function to run migration from go - when app is started
// migration is set just to postgreSql database
func RunAutoMigrate(db *sql.DB, migrationPath string) (err error, step string, dbVersion uint) {
	// embed sql directory where migrations is stored
	// and can be accessed to run them
	var fs embed.FS

	d, err := iofs.New(fs, migrationPath)
	if err != nil {
		return err, "open sql path", 0
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err, "initiate driver", 0
	}

	m, err := migrate.NewWithInstance("iofs", d, "postgres", driver)
	if err != nil {
		return err, "new instance", 0
	}

	defer m.Close()
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err, "migration failed", 0
	}
	dbversion, _, err := m.Version()

	return nil, "", dbversion

}
