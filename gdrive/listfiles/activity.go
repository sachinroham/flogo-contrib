package listfiles

import (
	"fmt"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("activity-gdrive-list")


const (
	ivAuthorizatonCode = "authorizatonCode"
	ivPageSize         = "pageSize"

	ovResult = "result"
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


	
	log.Debugf("response Body:", authorizatonCode)

	context.SetOutput(ovResult, authorizatonCode)

	return true, nil
}


