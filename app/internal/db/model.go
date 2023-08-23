package db

import (
	"app/config"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name    string
	CronStr string
	SQLStr  string
}

type Database struct {
	gorm.Model
	CustomerName string `gorm:"unique"`
	DBType       string
	User         string
	Password     string
	DatabaseName string
	Addr         string
}

func (db Database) DSN() string {
	if db.DBTypeEnum() == DBTypeMYSQL {
		return fmt.Sprintf(MySQLDSNFormat, db.User, db.Password, db.Addr, db.DatabaseName)
	} else {
		panic("not implementation")
	}
}

func (db Database) DBTypeEnum() DBType {
	ret, ok := DBTypeMap[strings.ToLower(db.DBType)]
	if !ok {
		return DBTypeInvalid
	}
	return ret
}

func (db *Database) FromConfig(c config.MySQLType) {
	db.Addr = c.Addr
	db.DBType = "mysql"
	db.Password = c.Password
	db.User = c.User
	db.DatabaseName = "dcronapp"
}
