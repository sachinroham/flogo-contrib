package listfiles

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("activity-tibco-rest")

const (
	ivAuthorizatonCode = "authorizatonCode"
	ivPageSize         = "pageSize"

	ovResult = "result"
)

var validMethods = []string{methodGET, methodPOST, methodPUT, methodPATCH, methodDELETE}

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

	log.Debug("response Body:", result)

	context.SetOutput(ovResult, result)
	context.SetOutput(ovStatus, resp.StatusCode)

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved client_secret.json.
	config, err := google.ConfigFromJSON(b, drive.DriveMetadataReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	srv, err := drive.New(getClient(config, authorizatonCode))
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	r, err := srv.Files.List().PageSize(pageSize).
		Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	fmt.Println("Files:")
	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, i := range r.Files {
			fmt.Printf("%s (%s)\n", i.Name, i.Id)
		}
	}

	log.Debug("response Body:", r)

	context.SetOutput(ovResult, r)

	return true, nil
}

////////////////////////////////////////////////////////////////////////////////////////
// Utils

// returns the generated client.
func getClient(config *oauth2.Config, token string) *http.Client {

	tok, err := config.Exchange(oauth2.NoContext, token)
	if err != nil {
		log.Fatalf("Unable to validate token %v", err)
	}
	return config.Client(context.Background(), tok)
}
