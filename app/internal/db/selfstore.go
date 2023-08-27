package db

import (
	"app/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// store the data of dcron app.
type SelfStore struct {
	db *gorm.DB
}

func newSelfStore() (ss *SelfStore, err error) {
	c := &Database{}
	c.FromConfig(config.I().MySQL)
	ss = new(SelfStore)
	ss.db, err = gorm.Open(mysql.Open(c.DSN()))
	if err != nil {
		return nil, err
	}
	err = ss.db.AutoMigrate(&Task{})
	if err != nil {
		return nil, err
	}
	err = ss.db.AutoMigrate(&Database{})
	if err != nil {
		return nil, err
	}
	return
}

func (ss *SelfStore) UpsertDataBase(db *Database) (err error) {
	return ss.db.Save(db).Error
}

func (ss *SelfStore) GetDataBaseByCustomerName(customerName string) (db Database, err error) {
	err = ss.db.Where("customer_name = ?", customerName).First(&db).Error
	return
}

func (ss *SelfStore) DeleteDataBaseByCustomerName(customerName string) (err error) {
	err = ss.db.Where("customer_name = ?", customerName).Delete(&Database{}).Error
	return
}

var (
	SelfStoreInstance *SelfStore
)

// singleton
// only for initial and get instance
//
type SelfStoreUtil struct{}

func (s SelfStoreUtil) Initial() (err error) {
	SelfStoreInstance, err = newSelfStore()
	return
}

func (s SelfStoreUtil) I() *SelfStore {
	return SelfStoreInstance
}