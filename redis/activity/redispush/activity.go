package redispush

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("activity-redis-push")

const (
	ivHost     = "host"
	ivPort     = "port"
	ivPassword = "password"
	ivKey      = "key"
	ivMessage  = "message"
	ovResult   = "result"
)

// REDISPUSH
// inputs : {host,port,user,password,path,content}
// outputs: {result}
type REDISPUSHActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new RESTActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &REDISPUSHActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *REDISPUSHActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Invokes a LISTFILE Operation
func (a *REDISPUSHActivity) Eval(context activity.Context) (done bool, err error) {

	host := context.GetInput(ivHost).(string)
	port := context.GetInput(ivPort).(int)
	password := context.GetInput(ivPassword).(string)
	key := context.GetInput(ivKey).(string)
	message := context.GetInput(ivMessage).(string)

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + strconv.Itoa(port),
		Password: "" + password, // no password set
		DB:       0,             // use default DB
	})

	fmt.Println("Data ", host+":"+strconv.Itoa(port))
	resp := client.LPush(key, message)
	fmt.Println(resp.Result())

	context.SetOutput(ovResult, resp)
	return true, nil
}
