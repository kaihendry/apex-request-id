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

type key int

const (
	logger key = iota
)

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
			"requestID": ctxObj.AwsRequestID,
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
	app.HandleFunc("/", h.showversion)
	return app
}

func (h handler) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logging := log.WithFields(log.Fields{
			"requestID":  r.Header.Get("X-Request-Id"),
			"method":     r.Method,
			"requestURI": r.RequestURI,
		})

		// Why doesn't this work?
		h.Log = logging

		// Setting context works
		ctx := context.WithValue(r.Context(), logger, logging)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h handler) showversion(w http.ResponseWriter, r *http.Request) {
	// Doesn't work, logging isn't setup
	h.Log.Info("Hello from log handler")

	// Context however, works
	log, ok := r.Context().Value(logger).(*log.Entry)
	if !ok {
		http.Error(w, "Unable to get context logger", http.StatusBadRequest)
		return
	}
	log.Info("Hello from the context logger")

	fmt.Fprintf(w, "%s", os.Getenv("UP_COMMIT"))
}
