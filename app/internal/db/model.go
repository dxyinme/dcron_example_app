package db

import (
	"app/config"
	"app/internal/common"
	"fmt"
	"strings"
)

type Task struct {
	ID             uint   `gorm:"primarykey"`
	Name           string `gorm:"index;type:VARCHAR(128)"`
	CronStr        string `gorm:"type:VARCHAR(64)"`
	SQLStr         string
	DBCustomerName string `gorm:"type:VARCHAR(128)"`
}

type Database struct {
	CustomerName string `gorm:"primaryKey;type:VARCHAR(128)"`
	DBType       string `gorm:"type:VARCHAR(64)"`
	User         string `gorm:"type:VARCHAR(64)"`
	Password     string `gorm:"type:VARCHAR(64)"`
	DatabaseName string `gorm:"type:VARCHAR(64)"`
	Addr         string `gorm:"type:VARCHAR(128)"`
}

func (db Database) DSN() string {
	if db.DBTypeEnum() == common.DBTypeMYSQL {
		return fmt.Sprintf(common.MySQLDSNFormat, db.User, db.Password, db.Addr, db.DatabaseName)
	} else {
		panic("not implementation")
	}
}

func (db Database) DBTypeEnum() common.DBType {
	ret, ok := common.DBTypeMap[strings.ToLower(db.DBType)]
	if !ok {
		return common.DBTypeInvalid
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
