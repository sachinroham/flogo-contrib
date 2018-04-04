package redissub

import (
	"context"
	"github.com/go-redis/redis"
	syslog "log"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("trigger-flogo-Redis")
var singleton *RedisTrigger

// RedisTrigger AWS Redis trigger struct
type RedisTrigger struct {
	metadata *trigger.Metadata
	config   *trigger.Config
	handlers []*trigger.Handler
}

//NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &RedisFactory{metadata: md}
}

// RedisFactory AWS Redis Trigger factory
type RedisFactory struct {
	metadata *trigger.Metadata
}

//New Creates a new trigger instance for a given id
func (t *RedisFactory) New(config *trigger.Config) trigger.Trigger {
	singleton = &RedisTrigger{metadata: t.metadata, config: config}
	return singleton
}

// Metadata implements trigger.Trigger.Metadata
func (t *RedisTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

func (t *RedisTrigger) Initialize(ctx trigger.InitContext) error {
	t.handlers = ctx.GetHandlers()
	syslog.Println("init")

	return nil
}

func (t *RedisTrigger) Start() error {
	syslog.Println("Start")
	syslog.Println("Processing handlers")
	for _, handler := range t.handlers {

		syslog.Println("Init Hadler", handler)
		t.processMessage(handler)
	}

	return nil

}

func (t *RedisTrigger) processMessage(endpoint *trigger.Handler) {
	syslog.Println("Inside processMessage")

	client := redis.NewClient(&redis.Options{
		Addr:     t.config.GetSetting("host") + ":" + t.config.GetSetting("port"),
		Password: "" + t.config.GetSetting("password"), // no password set
		DB:       0,                                    // use default DB
	})

	pubsub := client.Subscribe("redisChat")
	defer pubsub.Close()

	for {
		// ReceiveTimeout is a low level API. Use ReceiveMessage instead.
		msgi, err := pubsub.ReceiveTimeout(time.Second)
		if err != nil {
		}

		switch msg := msgi.(type) {
		case *redis.Message:
			syslog.Println("received", msg.Payload, "from", msg.Channel)
			data := map[string]interface{}{
				"message": msg.Payload,
				"channel": msg.Channel,
			}
			_, err := endpoint.Handle(context.Background(), data)
			if err != nil {
				syslog.Println("Error running handler: ", err.Error())
			}
		default:
		}
	}

}

// Stop implements util.Managed.Stop
func (t *RedisTrigger) Stop() error {
	syslog.Println("Stop")
	return nil
}
