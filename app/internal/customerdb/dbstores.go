package customerdb

import (
	"app/internal/common"
	"app/internal/db"
	"errors"
	"sync"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBStores struct {
	customerName2DB *sync.Map
}

func (dbs *DBStores) Add(db *db.Database) (err error) {
	var dbconn *gorm.DB = nil
	switch db.DBTypeEnum() {
	case common.DBTypeMYSQL:
		{
			dbconn, err = gorm.Open(mysql.Open(db.DSN()))
		}
	case common.DBTypePOSTGRESQL:
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
	for count := 0; count < 2; count++ {
		dbInterface, ok := dbs.customerName2DB.Load(CustomerName)
		if ok {
			return dbInterface.(*gorm.DB), ok
		} else {
			// try to load from db
			logrus.Infof("Load DB count %d", count)
			ss := db.SelfStoreUtil{}.I()
			db, err := ss.GetDataBaseByCustomerName(CustomerName)
			if err != nil {
				logrus.Error(err)
				return nil, false
			}
			if err = dbs.Add(&db); err != nil {
				logrus.Error(err)
				return nil, false
			}
		}
	}
	return nil, false
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
