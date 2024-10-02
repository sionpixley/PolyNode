package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"web/internal"
	"web/internal/middleware"
	"web/internal/services"

	"github.com/gorilla/mux"
	"github.com/sionpixley/PolyNode/pkg/polynrc"
)

func main() {
	operatingSystem := runtime.GOOS

	polyNodeConfig := polynrc.LoadPolyNodeConfig()
	err := overwriteGuiConfig(polyNodeConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()
	mapEndpoints(apiRouter)

	guiRouter := router.PathPrefix("/gui").Subrouter()
	guiRouter.PathPrefix("").HandlerFunc(serveGui)

	if operatingSystem == "darwin" || operatingSystem == "linux" {
		err = exec.Command("open", "http://localhost:"+strconv.Itoa(polyNodeConfig.GuiPort)+"/gui").Run()
	} else if operatingSystem == "windows" {
		err = exec.Command("cmd", "/c", "start", "http://localhost:"+strconv.Itoa(polyNodeConfig.GuiPort)+"/gui").Run()
	} else {
		err = errors.New("unsupported operating system")
	}
	if err != nil {
		log.Fatal(err.Error())
	}

	http.ListenAndServe(":"+strconv.Itoa(polyNodeConfig.GuiPort), router)
}

func mapEndpoints(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/add/{version}", services.Add).Methods(http.MethodPost, http.MethodOptions)
	apiRouter.HandleFunc("/list", services.List).Methods(http.MethodGet, http.MethodOptions)
	apiRouter.HandleFunc("/remove/{version}", services.Remove).Methods(http.MethodDelete, http.MethodOptions)
	apiRouter.HandleFunc("/search/{prefix}", services.SearchPrefix).Methods(http.MethodGet, http.MethodOptions)
	apiRouter.HandleFunc("/search", services.Search).Methods(http.MethodGet, http.MethodOptions)
	apiRouter.HandleFunc("/use/{version}", services.Use).Methods(http.MethodPatch, http.MethodOptions)
	apiRouter.HandleFunc("/version", services.Version).Methods(http.MethodGet, http.MethodOptions)

	apiRouter.Use(mux.CORSMethodMiddleware(apiRouter))
	apiRouter.Use(middleware.SetCorsOriginHeader)
	apiRouter.Use(middleware.SetCacheControlHeader)
	apiRouter.Use(middleware.SetContentTypeHeaders)
	apiRouter.Use(middleware.SetFrameHeaders)
}

func overwriteGuiConfig(polyNodeConfig polynrc.PolyNodeConfig) error {
	jsonData, err := json.Marshal(polyNodeConfig)
	if err != nil {
		return err
	}
	return os.WriteFile(internal.PolyNodeHomeDir+"/gui/dist/gui/browser/config/.polynrc", jsonData, 0644)
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
