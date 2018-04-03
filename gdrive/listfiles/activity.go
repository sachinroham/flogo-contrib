package listfiles

import (
	"bytes"
	"fmt"
	"github.com/jlaffaye/ftp"
	"io"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("activity-gdrive-list")

const (
	ivAuthorizatonCode = "authorizatonCode"
	ivPageSize         = "pageSize"
	ivFlowInfo         = "flowInfo"
	ivAddToFlow        = "addToFlow"
	ovResult           = "result"
)

// LISTFILEActivity is an Activity that is used to list filed from Google Drive
// inputs : {authorizatonCode,pageSize}
// outputs: {result}
type LISTFILEActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new RESTActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &LISTFILEActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *LISTFILEActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Invokes a LISTFILE Operation
func (a *LISTFILEActivity) Eval(context activity.Context) (done bool, err error) {

	authorizatonCode := strings.ToUpper(context.GetInput(ivAuthorizatonCode).(string))
	pageSize := context.GetInput(ivPageSize).(int)

	log.Debugf("LISTFILE Call: [%s] %s\n", authorizatonCode, pageSize)
	fmt.Printf("LISTFILE Call: [%s] %v\n", authorizatonCode, pageSize)

	client, err := ftp.Dial("localhost:21")
	if err != nil {
		fmt.Printf("err %v", err)
	}

	if err := client.Login("admin", "admin"); err != nil {
		fmt.Printf("err %v", err)
	}

	entries, _ := client.List("*")

	for _, entry := range entries {
		name := entry.Name
		fmt.Printf("name %v\n", name)
	}

	var content io.Reader
	content = bytes.NewBuffer([]byte(authorizatonCode))
	err = client.Stor("temp", content)
	if err != nil {
		fmt.Printf("err %v", err)
	}
	log.Debugf("response Body:", authorizatonCode)

	context.SetOutput(ovResult, entries)

	return true, nil
}
