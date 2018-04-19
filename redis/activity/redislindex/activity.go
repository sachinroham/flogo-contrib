package redislindex

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("activity-redis-lindex")

const (
	ivHost     = "host"
	ivPort     = "port"
	ivPassword = "password"
	ivKey      = "key"
	ivIndex    = "index"
	ovResult   = "result"
)

// REDISLINDEX
// inputs : {host,port,user,password,path,content}
// outputs: {result}
type REDISLINDEXActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new RESTActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &REDISLINDEXActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *REDISLINDEXActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Invokes a LISTFILE Operation
func (a *REDISLINDEXActivity) Eval(context activity.Context) (done bool, err error) {

	host := context.GetInput(ivHost).(string)
	port := context.GetInput(ivPort).(int)
	password := context.GetInput(ivPassword).(string)
	key := context.GetInput(ivKey).(string)
	index := context.GetInput(ivIndex).(int)

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + strconv.Itoa(port),
		Password: "" + password, // no password set
		DB:       0,             // use default DB
	})

	resp := client.LIndex(key, int64(index))
	content, _ := resp.Result()
	fmt.Println(content)

	context.SetOutput(ovResult, content)
	return true, nil
}
