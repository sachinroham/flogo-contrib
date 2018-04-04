package redissub

import (
	syslog "log"

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
	handlers := t.handlers

	syslog.Println("Processing handlers")
	for _, handler := range handlers {

		syslog.Println("Init Hadler",handler)
		
	}

	return nil
	
	return nil
}

// Stop implements util.Managed.Stop
func (t *RedisTrigger) Stop() error {
	syslog.Println("Stop")
	return nil
}

//func Invoke() (interface{}, error) {
//
//	log.Info("Starting AWS Redis Trigger")
//	syslog.Println("Starting AWS Redis Trigger")
//
//	return nil,nil
//}

//func run(t *RedisTrigger) error {
//	syslog.Println("run")
//	return nil
//}
