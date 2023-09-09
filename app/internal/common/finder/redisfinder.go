package finder

import (
	"context"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type RedisFinder struct {
	selfAddr string

	opts     *redis.Options
	redisCli *redis.Client
}

func (rf *RedisFinder) SetUp(selfAddr string, opts any) {
	rf.setUp(selfAddr, opts.(*redis.Options))
}

func (rf *RedisFinder) Start() (err error) {
	go rf.keep()
	<-time.After(2 * time.Second)
	return
}

func (rf *RedisFinder) GetAllAddrs() (addrs []string, err error) {
	result := rf.redisCli.Keys(context.Background(), IFinderPrefix+"*")
	if err = result.Err(); err != nil {
		return
	}
	addrs = result.Val()
	l := len(addrs)
	for i := 0; i < l; i++ {
		addrs[i] = strings.TrimPrefix(addrs[i], IFinderPrefix)
	}
	return
}

func (rf RedisFinder) keep() {
	tick := time.NewTicker(time.Second)
	key := IFinderPrefix + rf.selfAddr
	for range tick.C {
		if err := rf.redisCli.Set(context.Background(), key, []byte("1"), 5*time.Second).Err(); err != nil {
			logrus.Error(err)
		}
	}
}

func (rf *RedisFinder) setUp(selfAddr string, opts *redis.Options) {
	rf.selfAddr = selfAddr
	rf.opts = opts
	rf.redisCli = redis.NewClient(opts)
}
