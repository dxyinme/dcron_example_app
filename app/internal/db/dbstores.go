package db

import (
	"errors"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBStores struct {
	customerName2DB *sync.Map
}

func (dbs *DBStores) Add(db *Database) (err error) {
	var dbconn *gorm.DB = nil
	switch db.DBTypeEnum() {
	case DBTypeMYSQL:
		{
			dbconn, err = gorm.Open(mysql.Open(db.DSN()))
		}
	case DBTypePOSTGRESQL:
		{
			panic("not implementation")
		}
	default:
		{
			err = errors.New("invalid dbType")
		}
	}
	if err != nil {
		return
	}
	dbs.customerName2DB.Store(db.CustomerName, dbconn)
	return
}

func (dbs *DBStores) Load(CustomerName string) (*gorm.DB, bool) {
	dbInterface, ok := dbs.customerName2DB.Load(CustomerName)
	if ok {
		return dbInterface.(*gorm.DB), ok
	}
	return nil, ok
}

func (dbs *DBStores) Remove(CustomerName string) (isExist bool) {
	_, isExist = dbs.customerName2DB.LoadAndDelete(CustomerName)
	return
}

func newDBStores() (dbs *DBStores, err error) {
	dbs = &DBStores{
		customerName2DB: new(sync.Map),
	}

	return
}

var (
	dbStoresInstance *DBStores
)

type DBStoresUtil struct{}

func (db DBStoresUtil) Initial() (err error) {
	dbStoresInstance, err = newDBStores()
	return
}

func (db DBStoresUtil) I() *DBStores {
	return dbStoresInstance
}
