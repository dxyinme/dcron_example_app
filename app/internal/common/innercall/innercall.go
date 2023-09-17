package innercall

import (
	"app/config"
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type InnerCall struct {
	channelName string
	redisCli    *redis.Client

	pubsub *redis.PubSub
	input  chan []byte
	output chan []byte
}

func (ic *InnerCall) Start() {
	ic.pubsub = ic.redisCli.Subscribe(context.Background(), ic.channelName)
	ic.input = make(chan []byte, 100)
	ic.output = make(chan []byte, 100)
	go ic.receiveLoop()
	go ic.publishLoop()
}

func (ic *InnerCall) Send(op *Operation) {
	b, err := json.Marshal(op)
	if err != nil {
		logrus.Error(err)
		return
	}
	ic.output <- b
}

func (ic *InnerCall) Receive() (op *Operation) {
	b := <-ic.input
	op = &Operation{}
	if err := json.Unmarshal(b, op); err != nil {
		logrus.Error(err)
		return nil
	}
	return
}

func (ic *InnerCall) Close() {
	ic.pubsub.Close()
}

func (ic *InnerCall) receiveLoop() {
	for {
		msg := <-ic.pubsub.Channel()
		ic.input <- []byte(msg.Payload)
	}
}

func (ic *InnerCall) publishLoop() {
	for {
		msg := <-ic.output
		if err := ic.redisCli.Publish(context.Background(), ic.channelName, msg).Err(); err != nil {
			logrus.Error(err)
		}
	}
}

var (
	innerCallInstance *InnerCall = nil
)

type InnerCallUtil struct{}

func (iu InnerCallUtil) Initial() {
	cfg := config.I()
	innerCallInstance = &InnerCall{
		channelName: cfg.InnerCall.Channel,
		redisCli: redis.NewClient(&redis.Options{
			Addr:     cfg.InnerCall.Redis.Addr,
			DB:       cfg.InnerCall.Redis.DB,
			Password: cfg.Redis.Password,
		}),
	}
	innerCallInstance.Start()
}

func (iu InnerCallUtil) I() *InnerCall {
	return innerCallInstance
}
