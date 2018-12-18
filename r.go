package r

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/gorilla/mux"
)

type handler struct {
	Log *log.Entry
}

// New creates a handler for this application to co-ordinate shared resources
func New(ctx context.Context) (h handler, err error) {
	var logWithRequestID *log.Entry
	ctxObj, ok := lambdacontext.FromContext(ctx)
	if ok {
		logWithRequestID = log.WithFields(log.Fields{
			"RequestID": ctxObj.AwsRequestID,
		})
	} else {
		logWithRequestID = log.WithFields(log.Fields{})
	}
	h = handler{
		Log: logWithRequestID,
	}
	return
}

// Apex lambda stuff
func (h handler) HellofromApex() error {
	h.Log.Info("Hello from Apex")
	return nil
}

// Apex Up stuff
func (h handler) BasicEngine() http.Handler {
	app := mux.NewRouter()
	app.HandleFunc("/", h.showversion).Methods("GET")
	return app
}

func (h handler) showversion(w http.ResponseWriter, r *http.Request) {
	h.Log.Info("Helo from Up")
	fmt.Fprintf(w, "%s", os.Getenv("UP_COMMIT"))
}
