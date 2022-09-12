package dbutils

import (
	f "azure_pg_kv/errorhandling"
	"database/sql"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
	"github.com/microsoft/go-mssqldb/azuread"
)

var db *sql.DB

type Database struct {
	connection *sql.DB
}

func NewConnection() *Database {
	// connString := fmt.Sprintf("Server=%s; Authentication=Active Directory Managed Identity; Database=%s",
	//	"membranedevdb.database.windows.net", "membranedevdb")
	//connString := fmt.Sprintf("sqlserver://%s/?database=%s&fedauth=ActiveDirectoryMSI&user_id=%s", "membranedevdb.database.windows.net", "membranedevdb", "id-db-membranedev")
	connString := fmt.Sprintf("sqlserver://%s/?database=%s&fedauth=ActiveDirectoryMSI", "membranedevdb.database.windows.net", "membranedevdb")

	conn, err := sql.Open(azuread.DriverName, connString)
	f.FailIf(err, "SQL Server connection failure")
	return &Database{connection: conn}
}

func (d *Database) Ping() error {
	return d.connection.Ping()
}
