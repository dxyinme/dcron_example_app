package crontasks

import (
	"app/config"
	"app/internal/db"
	"time"

	"github.com/libi/dcron"
	"github.com/libi/dcron/driver"
	redis "github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var (
	cronTasksContainerInstance *CronTasksContainer = nil
)

type CronTasksContainer struct {
	dcronDriverInstance driver.DriverV2
	dcronInstance       *dcron.Dcron
	redisCli            *redis.Client
}

func (ctc *CronTasksContainer) AddTask(job *CronTask) error {
	return ctc.dcronInstance.AddJob(job.Name, job.CronStr, job)
}

func (ctc *CronTasksContainer) RemoveTask(jobName string) {
	ctc.dcronInstance.Remove(jobName)
}

type CronTasksContainerUtil struct{}

func (du CronTasksContainerUtil) Initial() {
	cfg := config.I()
	redisOpts := &redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	}
	redisCli := redis.NewClient(redisOpts)
	dcronDriverInstance := driver.NewRedisDriver(redisCli)
	dcronInstance := dcron.NewDcronWithOption(
		cfg.Dcron.ServiceName,
		dcronDriverInstance,
		dcron.CronOptionSeconds(),
		dcron.WithLogger(new(Logger)),
		dcron.WithHashReplicas(10),
		dcron.WithNodeUpdateDuration(time.Second*10),
		dcron.WithRecoverFunc(func(d *dcron.Dcron) {
			ss := db.SelfStoreUtil{}.I()
			limit := 10
			lastId := uint(0)
			for {
				tasks, err := ss.GetTasksByIDLimit(lastId, limit)
				if err != nil {
					logrus.Error(err)
					continue
				}
				if len(tasks) == 0 {
					break
				}
				for _, task := range tasks {
					cronTask := &CronTask{}
					cronTask.FromDBTask(&task)
					if err = cronTask.Initial(); err != nil {
						logrus.Error(err)
						continue
					}
					if err = d.AddJob(cronTask.Name, cronTask.CronStr, cronTask); err != nil {
						logrus.Error(err)
						continue
					}
				}
				ltasks := len(tasks)
				if ltasks < limit {
					break
				}
				lastId = tasks[ltasks-1].ID + 1
			}
		}),
	)
	cronTasksContainerInstance = &CronTasksContainer{
		dcronDriverInstance: dcronDriverInstance,
		dcronInstance:       dcronInstance,
		redisCli:            redisCli,
	}
	dcronInstance.Start()
}

func (du CronTasksContainerUtil) I() *CronTasksContainer {
	return cronTasksContainerInstance
}
