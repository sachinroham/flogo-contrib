package retr

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json yes Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

var host = "sroham-t430"
var user = "admin"
var password = "admin"
var localpath = "temp.txt"
var remotepath = "temp.txt"


func TestListFiles(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("host", host)
	tc.SetInput("port", 21)
	tc.SetInput("user", user)
	tc.SetInput("password", password)
	tc.SetInput("localpath", localpath)
	tc.SetInput("remotepath", remotepath)
	fmt.Println(strings.Join([]string{host, strconv.Itoa(21)}, ":"))
	//eval
	act.Eval(tc)
	val := tc.GetOutput("result")

	fmt.Printf("result: %v\n", val)

}
