package crontasks

import (
	"app/config"

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

func (ctc *CronTasksContainer) AddTask() {

}

func (ctc *CronTasksContainer) RemoveTask() {

}

type CronTasksContainerUtil struct{}

func (du CronTasksContainerUtil) Initial() {
	cfg := config.I()
	redisOpts := &redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
	}
	redisCli := redis.NewClient(redisOpts)
	dcronDriverInstance := driver.NewRedisZSetDriver(redisCli)
	dcronInstance := dcron.NewDcronWithOption(
		cfg.Dcron.ServiceName,
		dcronDriverInstance,
		dcron.CronOptionSeconds(),
		dcron.WithLogger(new(Logger)),
	)

	cronTasksContainerInstance = &CronTasksContainer{
		dcronDriverInstance: dcronDriverInstance,
		dcronInstance:       dcronInstance,
		redisCli:            redisCli,
	}
}

func (du CronTasksContainerUtil) I() *CronTasksContainer {
	return cronTasksContainerInstance
}
