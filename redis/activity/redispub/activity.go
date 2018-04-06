package redispub

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("activity-redis-pub")

const (
	ivHost     = "host"
	ivPort     = "port"
	ivPassword = "password"
	ivChannel  = "channel"
	ivMessage  = "message"
	ovResult   = "result"
)

// REDISPUBActivity is an Activity is used to store a file on FTP server
// inputs : {host,port,user,password,path,content}
// outputs: {result}
type REDISPUBActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new RESTActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &REDISPUBActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *REDISPUBActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Invokes a LISTFILE Operation
func (a *REDISPUBActivity) Eval(context activity.Context) (done bool, err error) {

	host := context.GetInput(ivHost).(string)
	port := context.GetInput(ivPort).(int)
	password := context.GetInput(ivPassword).(string)
	channel := context.GetInput(ivChannel).(string)
	message := context.GetInput(ivMessage).(string)

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + strconv.Itoa(port),
		Password: "" + password, // no password set
		DB:       0,             // use default DB
	})

	fmt.Println("Data ", host+":"+strconv.Itoa(port))
	resp, err := client.Publish(channel, message).Result()
	if err != nil {
		panic(err)
	}

	context.SetOutput(ovResult, resp)
	return true, nil
}
