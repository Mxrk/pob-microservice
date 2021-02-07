package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"pob_api/pob"
)

func main(){
	http.HandleFunc("/", handlePob)

	err := http.ListenAndServe(":" + "8080", nil)
	if err != nil {
		panic(err)
	}
}

func handlePob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		jsonError(w, "please use post", http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		jsonError(w, "can't read body", http.StatusBadRequest)
		return
	}

	type b struct {
		Code string `json:"code"`
	}
	js := b{}

	err = json.Unmarshal(body, &js)
	if err != nil {
		log.Println(err)
		jsonError(w, "can't read body", http.StatusBadRequest)
		return
	}

	show := pob.GetPob(js.Code)

	json, err := json.Marshal(show)
	if err != nil {
		log.Println(err)
		jsonError(w, "", http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func jsonError(w http.ResponseWriter, errorInterface interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	type localError struct {
		MSG  interface{} `json:"error"`
		Code int         `json:"code"`
	}

	var e localError
	e.MSG = errorInterface
	e.Code = code

	err := json.NewEncoder(w).Encode(e)
	if err != nil {
		log.Println(err)
	}
}