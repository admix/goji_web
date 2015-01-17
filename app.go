package main

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
)

func home(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Testing app, params: %s", c.URLParams["user"])
}

func main() {
	goji.Get("/home/:user", home)
	goji.Serve()
}