package services

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Add(w http.ResponseWriter, r *http.Request) {

}

func Current(w http.ResponseWriter, r *http.Request) {

}

func Install(w http.ResponseWriter, r *http.Request) {

}

func List(w http.ResponseWriter, r *http.Request) {
	l, err := list()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	versions := []string{}
	temp := strings.Split(l, "Node.js - ")
	for i := 0; i < len(temp); i += 1 {
		temp[i] = strings.TrimSpace(temp[i])
		if temp[i] != "" {
			versions = append(versions, temp[i])
		}
	}

	err = json.NewEncoder(w).Encode(versions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func Remove(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	v, exists := params["version"]
	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := remove(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if strings.Contains(output, "error") {
		err = json.NewEncoder(w).Encode(false)
	} else {
		err = json.NewEncoder(w).Encode(true)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func Search(w http.ResponseWriter, r *http.Request) {

}

func Use(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	v, exists := params["version"]
	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := use(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if strings.Contains(output, "error") {
		err = json.NewEncoder(w).Encode(false)
	} else {
		err = json.NewEncoder(w).Encode(true)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func Version(w http.ResponseWriter, r *http.Request) {
	v, err := version()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	v = strings.TrimSpace(v)

	err = json.NewEncoder(w).Encode(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
