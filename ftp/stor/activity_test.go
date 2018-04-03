package listfiles

import (
	"fmt"
	"testing"

	"io/ioutil"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
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
var path = "temp.txt"
var content = "Cleared problem markers for Projects"

func TestListFiles(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("host", host)
	tc.SetInput("user", user)
	tc.SetInput("password", password)
	tc.SetInput("path", path)
	tc.SetInput("content", content)

	//eval
	act.Eval(tc)
	val := tc.GetOutput("result")

	fmt.Printf("result: %v\n", val)

}
