package main

import "database/sql"

func openDB(connectionString string) (*sql.DB, error) {
	return sql.Open(driverName, connectionString)
}
