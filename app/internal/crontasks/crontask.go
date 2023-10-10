package crontasks

import (
	"app/internal/customerdb"
	"app/internal/db"
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CronTask struct {
	Name           string
	CronStr        string
	DBCustomerName string
	SQLStr         string
	dbconn         *gorm.DB
}

func (cronTask *CronTask) Run() {
	logrus.Debugf(
		"name=%s|dbname=%s|sql-str=`%s`",
		cronTask.Name, cronTask.DBCustomerName, cronTask.SQLStr)

	dbi := db.SelfStoreUtil{}.I()
	go dbi.ReportTaskMetric(&db.TaskMetric{
		TaskName: cronTask.Name,
		NodeID:   CronTasksContainerUtil{}.I().dcronInstance.NodeID(),
	})
	if err := cronTask.dbconn.Exec(cronTask.SQLStr).Error; err != nil {
		logrus.Error(err)
	}
}

func (cronTask *CronTask) FromDBTask(dbTask *db.Task) {
	cronTask.Name = dbTask.Name
	cronTask.CronStr = dbTask.CronStr
	cronTask.DBCustomerName = dbTask.DBCustomerName
	cronTask.SQLStr = dbTask.SQLStr
}

func (cronTask *CronTask) Initial() error {
	var existed bool
	dbs := customerdb.DBStoresUtil{}.I()

	cronTask.dbconn, existed = dbs.Load(cronTask.DBCustomerName)
	if !existed {
		return errors.New("can not load dbconn")
	}
	return nil
}
