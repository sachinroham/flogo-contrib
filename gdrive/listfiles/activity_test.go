package rest

import (
	"encoding/json"
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
			panic("No Json Metadata found for activity.json path")
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

var authCode ="ya29.GluRBdk5mVFLYH_hSQBo66LDINkzf0-1Jfe6T9fDIxKpvWDezuHfMtaghoQH409ovB6cFLryRV3SRPTp94uQ7psY8k6a9z1tjZXZUDAs2md9Zg1-inhjWUxexOcv"

func TestListFiles(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("authorizatonCode", authCode)
	tc.SetInput("pageSize", 10)

	//eval
	act.Eval(tc)
	val := tc.GetOutput("result")

	fmt.Printf("result: %v\n", val)

}

