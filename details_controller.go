package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func fetchEmployeeDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var responseObject Details
	id := vars["id"]
	fmt.Println("1. Performing Http Get...")
	path := os.Getenv("EMPLOYEE_STORE") + id
	fmt.Println(path)
	resp, err := http.Get(strings.TrimSpace(path))
	if err != nil {
		log.Fatalln(err)
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject)
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}
