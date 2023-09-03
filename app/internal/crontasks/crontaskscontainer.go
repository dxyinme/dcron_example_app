package crontasks

import (
	"app/config"
	"time"

	redis "github.com/go-redis/redis/v8"
	"github.com/libi/dcron"
	"github.com/libi/dcron/driver"
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
		}),
	)

	dcronInstance.Start()

	cronTasksContainerInstance = &CronTasksContainer{
		dcronDriverInstance: dcronDriverInstance,
		dcronInstance:       dcronInstance,
		redisCli:            redisCli,
	}
}

func (du CronTasksContainerUtil) I() *CronTasksContainer {
	return cronTasksContainerInstance
}
