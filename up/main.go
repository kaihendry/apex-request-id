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
	h := r.New(ctx)

	app := h.BasicEngine()

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), app); err != nil {
		log.WithError(err).Fatal("error listening")
	}

}
