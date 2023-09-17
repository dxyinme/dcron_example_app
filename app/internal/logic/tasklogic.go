package logic

import (
	"app/internal/crontasks"
	"app/internal/db"

	"github.com/sirupsen/logrus"
)

type TaskLogic struct{}

func (tl TaskLogic) UpsertCronTask(task *db.Task) (err error) {
	ss := db.SelfStoreUtil{}.I()
	if err = ss.UpsertTask(task); err != nil {
		logrus.Error(err)
	}
	return
}

func (tl TaskLogic) DeleteCronTask(taskName string) (err error) {
	ss := db.SelfStoreUtil{}.I()
	if err = ss.DeleteTaskByTaskName(taskName); err != nil {
		logrus.Error(err)
	}
	return
}

func (tl TaskLogic) UpsertCronTaskToDcron(task *db.Task) (err error) {
	ctc := crontasks.CronTasksContainerUtil{}.I()
	cronTask := &crontasks.CronTask{}
	cronTask.FromDBTask(task)
	if err = cronTask.Initial(); err != nil {
		return err
	}
	err = ctc.AddTask(cronTask)
	return
}

func (tl TaskLogic) RemoveCronTaskFromDcron(taskName string) {
	ctc := crontasks.CronTasksContainerUtil{}.I()
	ctc.RemoveTask(taskName)
}

func (tl TaskLogic) GetCronTask(taskName string) (task db.Task, err error) {
	ss := db.SelfStoreUtil{}.I()
	if task, err = ss.GetTaskByTaskName(taskName); err != nil {
		logrus.Error(err)
		return
	}
	return
}
