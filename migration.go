package Migrations

import (
	"database/sql"
	"embed"

	"github.com/golang-migrate/migrate/v4"
	database "github.com/golang-migrate/migrate/v4/database"
	mysql "github.com/golang-migrate/migrate/v4/database/mysql"
	postgres "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

// function to run migration from go - when app is started
// migration is set just to postgreSql database
func RunAutoMigrate(db *sql.DB, fs embed.FS, migrationPath string, dbType string, dbName string) (err error, step string, dbVersion uint) {

	d, err := iofs.New(fs, migrationPath)
	if err != nil {
		return err, "open sql path", 0
	}

	var driver database.Driver

	//check which driver to use
	if dbType == "mysql" {
		driver, err = mysql.WithInstance(db, &mysql.Config{})
		if err != nil {
			return err, "initiate driver", 0
		}
	} else {
		driver, err = postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			return err, "initiate driver", 0
		}
	}

	m, err := migrate.NewWithInstance("iofs", d, dbName, driver)
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
