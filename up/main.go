package main

import (
	"context"
	"net/http"
	"os"

	"github.com/apex/log"
	jsonloghandler "github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/text"
	r "github.com/kaihendry/apex-request-id"
)

func init() {
	if os.Getenv("UP_STAGE") != "" {
		log.SetHandler(jsonloghandler.Default)
	} else {
		log.SetHandler(text.Default)
	}
}

func main() {
	ctx := context.Background()
	log.Infof("%#v", ctx)
	h := r.New(ctx)

	app := h.BasicEngine()

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), app); err != nil {
		log.WithError(err).Fatal("error listening")
	}

}
