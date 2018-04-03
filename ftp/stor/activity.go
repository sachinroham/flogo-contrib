package stor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"	
	"strconv"
	
	"github.com/jlaffaye/ftp"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("activity-ftp-stor")

const (
	ivHost     = "host"
	ivPort     = "port"
	ivUser     = "user"
	ivPassword = "password"
	ivPath     = "path"
	ivContent  = "content"
	ovResult   = "result"
)

// STORActivity is an Activity that is used to list filed from Google Drive
// inputs : {host,port,user,password,path,content}
// outputs: {result}
type STORActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new RESTActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &STORActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *STORActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Invokes a LISTFILE Operation
func (a *STORActivity) Eval(context activity.Context) (done bool, err error) {

	host := context.GetInput(ivHost).(string)
	port := context.GetInput(ivPort).(int)
	if port == 0 {
		port = 21
	}
	user := context.GetInput(ivUser).(string)
	password := context.GetInput(ivPassword).(string)
	path := context.GetInput(ivPath).(string)
	content := context.GetInput(ivContent)

	//	log.Debugf("STORActivity Call: [%s] %s\n", host, pageSize)
	//	fmt.Printf("STORActivity Call: [%s] %v\n", authorizatonCode, pageSize)
	
	client, err := ftp.Dial(strings.Join([]string{host,strconv.Itoa(port)},":"))

	if err != nil {
		fmt.Printf("err %v", err)
	}

	if err := client.Login(user, password); err != nil {
		fmt.Printf("err %v", err)
	}

	var stream io.Reader
	if content != nil {
		if str, ok := content.(string); ok {
			stream = bytes.NewBuffer([]byte(str))
		} else {
			b, _ := json.Marshal(content) //todo handle error
			stream = bytes.NewBuffer([]byte(b))
		}
	}

	err = client.Stor(path, stream)
	if err != nil {
		context.SetOutput(ovResult, err)
	} else {
		context.SetOutput(ovResult, "Success")
	}

	return true, nil
}
