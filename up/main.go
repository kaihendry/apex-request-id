package main

import (
	"context"
	"net/http"
	"os"

	"github.com/apex/log"
	r "github.com/kaihendry/apex-request-id"
)

func main() {
	ctx := context.Background()
	log.Infof("%#v", ctx)
	h, err := r.New(ctx)
	if err != nil {
		log.WithError(err).Fatal("error setting configuration")
		return
	}

	addr := ":" + os.Getenv("PORT")
	app := h.BasicEngine()

	if err := http.ListenAndServe(addr, app); err != nil {
		log.WithError(err).Fatal("error listening")
	}

}
