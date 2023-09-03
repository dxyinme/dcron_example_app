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

func (webTask *Task) FromDBTask(task *db.Task) {
	webTask.CronStr = task.CronStr
	webTask.Name = task.Name
	webTask.DBName = task.DBCustomerName
	webTask.SQLStr = task.SQLStr
}
