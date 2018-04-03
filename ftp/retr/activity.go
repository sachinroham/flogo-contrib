package retr

import (

	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/jlaffaye/ftp"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("activity-ftp-stor")

const (
	ivHost       = "host"
	ivPort       = "port"
	ivUser       = "user"
	ivPassword   = "password"
	ivLocalPath  = "localpath"
	ivRemotePath = "remotepath"
	ovResult     = "result"
)

// STORActivity is an Activity is used to store a file on FTP server
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
	localPath := context.GetInput(ivLocalPath).(string)
	if localPath == "" {
		fmt.Printf("localPath %v", localPath)
	}

	remotePath := context.GetInput(ivRemotePath).(string)

	client, err := ftp.Dial(strings.Join([]string{host, strconv.Itoa(port)}, ":"))

	if err != nil {
		fmt.Printf("err %v", err)
	}

	if err := client.Login(user, password); err != nil {
		fmt.Printf("err %v", err)
	}

	resp, err := client.Retr(remotePath)
	if err != nil {
		fmt.Printf("err %v", err)
	}

	buffer := make([]byte, 1024)
	
	if _, err := io.ReadFull(resp, buffer); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("buffer %s", buffer)
	
	context.SetOutput(ovResult, "buffer")
	return true, nil
}
