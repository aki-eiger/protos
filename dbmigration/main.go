package main

import (
	"dbmigration/migrations"
	"fmt"

	"github.com/ahakkila/migrator"
)

func main() {
	migrator := migrator.NewMigrator("dbmigrations/migrations")
	migrations.Initialize_001(migrator)
	migrator.IterateVersions()

	fmt.Println("Done!")

}
