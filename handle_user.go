package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func HandleUserNew(w http.ResponseWriter, r *http.Request, params 	httprouter.Params) {
	RenderTemplate(w, r, "users/new", nil)
}