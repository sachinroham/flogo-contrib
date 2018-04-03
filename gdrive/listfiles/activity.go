package listfiles

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"

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
func (a *LISTFILEActivity) Eval(context1 activity.Context) (done bool, err error) {

	authorizatonCode := strings.ToUpper(context1.GetInput(ivAuthorizatonCode).(string))
	pageSize := context1.GetInput(ivPageSize).(int)

	log.Debugf("LISTFILE Call: [%s] %s\n", authorizatonCode, pageSize)
	fmt.Printf("LISTFILE Call: [%s] %v\n", authorizatonCode, pageSize)

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Debugf("Unable to read client secret file: %v\n", err)
		fmt.Printf("Unable to read client secret file: %v\n", err)
	}

//	fmt.Printf("b: %v \n", b)
	// If modifying these scopes, delete your previously saved client_secret.json.
	config, err := google.ConfigFromJSON(b, drive.DriveMetadataReadonlyScope)
	if err != nil {
		log.Debugf("Unable to parse client secret file to config: %v\n", err)
		fmt.Printf("Unable to parse client secret file to config: %v\n", err)
	}
	fmt.Printf("config: %v \n", config)

	srv, err := drive.New(getClient(config, authorizatonCode))
	if err != nil {
		log.Debugf("Unable to retrieve Drive client: %v\n", err)
		fmt.Printf("Unable to retrieve Drive client: %v\n", err)
	}
	fmt.Printf("srv: %v \n", srv)

	r, err := srv.Files.List().PageSize(10).
		Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		log.Debugf("Unable to retrieve files: %v\n", err)
	}
	fmt.Printf("Files: %v \n", r)
	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, i := range r.Files {
			fmt.Printf("%s (%s)\n", i.Name, i.Id)
		}
	}

	log.Debugf("response Body:", r)

	context1.SetOutput(ovResult, r)

	return true, nil
}

////////////////////////////////////////////////////////////////////////////////////////
// Utils

// returns the generated client.
func getClient(config *oauth2.Config, token string) *http.Client {
	

	fmt.Printf("token %s\n", token)
	tok, err := config.Exchange(oauth2.NoContext, token)
	if err != nil {
		log.Debugf("Unable to validate token %v\n", err)
		fmt.Printf("Unable to validate token %v\n", err)
	}
	
	return config.Client(context.Background(), tok)
}
