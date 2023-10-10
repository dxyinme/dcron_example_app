package db

import (
	"app/config"

	"github.com/sirupsen/logrus"
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

	if config.I().EnableReporter {
		err = ss.db.AutoMigrate(&TaskMetric{})
		if err != nil {
			logrus.Error(err)
		}
	}
	return
}

func (ss *SelfStore) ReportTaskMetric(tm *TaskMetric) {
	_ = ss.db.Save(tm)
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

func (ss *SelfStore) UpsertTask(task *Task) error {
	if tID, err := ss.GetTaskID(task.Name); err != nil {
		return ss.db.Create(task).Error
	} else {
		task.ID = tID
		return ss.db.Save(task).Error
	}
}

func (ss *SelfStore) DeleteTaskByTaskName(taskName string) (err error) {
	return ss.db.Where("name = ?", taskName).Delete(&Task{}).Error
}

func (ss *SelfStore) GetTaskByTaskName(taskName string) (task Task, err error) {
	err = ss.db.Where("name = ?", taskName).First(&task).Error
	return
}

func (ss *SelfStore) GetTaskID(taskName string) (ID uint, err error) {
	t := Task{}
	err = ss.db.Where("name = ?", taskName).Select("id").First(&t).Error
	return t.ID, err
}

func (ss *SelfStore) GetTasksByIDLimit(ID uint, limit int) (tasks []Task, err error) {
	tasks = make([]Task, 0)
	err = ss.db.Where("id >= ?", ID).Order("id").Limit(limit).Find(&tasks).Error
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
