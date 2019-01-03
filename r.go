package r

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/apex/log"
	jsonhandler "github.com/apex/log/handlers/json"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/gorilla/mux"
)

func init() {
	log.SetHandler(jsonhandler.Default)
}

// handler is the share state between the functions
type handler struct {
	Log *log.Entry
}

// New creates a handler for this application to co-ordinate shared resources
func New(ctx context.Context) (h handler) {
	var logWithRequestID *log.Entry
	ctxObj, ok := lambdacontext.FromContext(ctx)
	if ok {
		logWithRequestID = log.WithFields(log.Fields{
			"RequestID": ctxObj.AwsRequestID,
		})
	} else {
		// I want this to be replaced by loggingMiddleware
		logWithRequestID = log.WithFields(log.Fields{
			"foo": "bar",
		})
	}
	h = handler{
		Log: logWithRequestID,
	}
	return
}

// Apex lambda stuff
func (h handler) HellofromApex() error {
	h.Log.Info("Hello from Apex!")
	return nil
}

// Apex Up stuff
func (h handler) BasicEngine() http.Handler {
	app := mux.NewRouter()
	app.Use(h.loggingMiddleware)
	app.HandleFunc("/", h.showversion).Methods("GET")
	return app
}

func (h handler) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Log.WithFields(log.Fields{
			"requestid": r.Header.Get("X-Request-Id"),
			"method":    r.Method,
		})
		next.ServeHTTP(w, r)
	})
}

func (h handler) showversion(w http.ResponseWriter, r *http.Request) {
	// Real test is here
	h.Log.Info("Hello from Up")
	fmt.Fprintf(w, "%s", os.Getenv("UP_COMMIT"))
}
