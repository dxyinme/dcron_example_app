package logic

import (
	"app/internal/customerdb"
	"app/internal/db"

	"github.com/sirupsen/logrus"
)

type DBLogic struct{}

func (dbl DBLogic) UpsertDatabase(dbData *db.Database) (err error) {
	ss := db.SelfStoreUtil{}.I()
	err = ss.UpsertDataBase(dbData)
	if err != nil {
		logrus.Error(err)
		return
	}
	return
}

func (dbl DBLogic) UpsertDataBaseToCache(dbData *db.Database) (err error) {
	return customerdb.DBStoresUtil{}.I().Add(dbData)
}

func (dbl DBLogic) Remove(dbName string) (err error) {
	ss := db.SelfStoreUtil{}.I()
	if err = ss.DeleteDataBaseByCustomerName(dbName); err != nil {
		logrus.Error(err.Error())
		return
	}
	return
}

func (dbl DBLogic) Get(dbName string) (dbData db.Database, err error) {
	ss := db.SelfStoreUtil{}.I()
	if dbData, err = ss.GetDataBaseByCustomerName(dbName); err != nil {
		logrus.Error(err)
		return
	}
	return
}
