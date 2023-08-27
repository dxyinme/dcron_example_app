package types

import "app/internal/db"

func (webDB *DB) FromDBDatabase(database *db.Database) {
	webDB.CustomerName = database.CustomerName
	webDB.DatabaseName = database.DatabaseName
	webDB.Addr = database.Addr
	webDB.DBType = database.DBType
	webDB.Password = database.Password
	webDB.User = database.User
}
