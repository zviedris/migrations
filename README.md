# migrations
Golang library that can be used to run DB migrations on app startup

In project where you add the package - you need to use for example (directory you set where the migration files is)
//go:embed directory
var folder embed.FS

And then pass folder and named directory as a parameters to the RunAutoMigrate function

For example you can call auto migrations with
err, _, _ := migrations.RunAutoMigrate(migrateDBInstance.DBX().DB, folder, "directory")
		if err != nil {
			log.Error(err)
		}