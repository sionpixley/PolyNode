//go:build !windows

package main

import (
	"net/http"
	"os"
	"os/exec"
	"strings"
	"web/internal"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()
	mapEndpoints(apiRouter)

	guiRouter := router.PathPrefix("/gui").Subrouter()
	guiRouter.PathPrefix("").HandlerFunc(serveGui)

	err := exec.Command("open", "http://localhost:2334/gui").Run()
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":2334", router)
}

func mapEndpoints(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/list", internal.List).Methods(http.MethodGet)
}

func serveGui(w http.ResponseWriter, r *http.Request) {
	trimmedPath := strings.TrimPrefix(r.URL.Path, "/")
	filePath := internal.PolyNodeHomeDir + "/gui/" + trimmedPath
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.ServeFile(w, r, internal.PolyNodeHomeDir+"/gui/dist/gui/browser/index.html")
	} else {
		http.ServeFile(w, r, filePath)
	}
}
