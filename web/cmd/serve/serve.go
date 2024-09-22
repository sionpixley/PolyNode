package main

import (
	"errors"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"web/internal"
	"web/internal/middleware"
	"web/internal/services"

	"github.com/gorilla/mux"
)

func main() {
	operatingSystem := runtime.GOOS
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()
	mapEndpoints(apiRouter)

	guiRouter := router.PathPrefix("/gui").Subrouter()
	guiRouter.PathPrefix("").HandlerFunc(serveGui)

	var err error
	if operatingSystem == "darwin" || operatingSystem == "linux" {
		err = exec.Command("open", "http://localhost:2334/gui").Run()
	} else if operatingSystem == "windows" {
		err = exec.Command("start", "http://localhost:2334/gui").Run()
	} else {
		err = errors.New("unsupported operating system")
	}
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":2334", router)
}

func mapEndpoints(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/list", services.List).Methods(http.MethodGet)
	apiRouter.HandleFunc("/remove/{version}", services.Remove).Methods(http.MethodDelete)
	apiRouter.HandleFunc("/search", services.Search).Methods(http.MethodGet)
	apiRouter.HandleFunc("/use/{version}", services.Use).Methods(http.MethodPatch)
	apiRouter.HandleFunc("/version", services.Version).Methods(http.MethodGet)

	apiRouter.Use(middleware.SetCacheControlHeader)
	apiRouter.Use(middleware.SetContentTypeHeaders)
	apiRouter.Use(middleware.SetFrameHeaders)
}

func serveGui(w http.ResponseWriter, r *http.Request) {
	trimmedPath := strings.TrimPrefix(r.URL.Path, "/gui")
	if trimmedPath == "" || trimmedPath == "/" {
		trimmedPath = "index.html"
	}
	filePath := internal.PolyNodeHomeDir + "/gui/dist/gui/browser" + trimmedPath
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.ServeFile(w, r, internal.PolyNodeHomeDir+"/gui/dist/gui/browser/index.html")
	} else {
		http.ServeFile(w, r, filePath)
	}
}
