package redispop

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("activity-redis-pop")

const (
	ivHost     = "host"
	ivPort     = "port"
	ivPassword = "password"
	ivKey      = "key"
	ovResult   = "result"
)

// REDISPOP
// inputs : {host,port,user,password,path,content}
// outputs: {result}
type REDISPOPActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new RESTActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &REDISPOPActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *REDISPOPActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Invokes a LISTFILE Operation
func (a *REDISPOPActivity) Eval(context activity.Context) (done bool, err error) {

	host := context.GetInput(ivHost).(string)
	port := context.GetInput(ivPort).(int)
	password := context.GetInput(ivPassword).(string)
	key := context.GetInput(ivKey).(string)

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + strconv.Itoa(port),
		Password: "" + password, // no password set
		DB:       0,             // use default DB
	})

	resp := client.LPop(key)
	content, _ := resp.Result()
	fmt.Println(content)

	context.SetOutput(ovResult, content)
	return true, nil
}
