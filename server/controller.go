package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../logic"
)

// JSONData Struct to maping json received.
type JSONData struct {
	Array interface{} `json:"array"`
}

// GetFlattenArray function to receive the request and read the data, then
// we call to logic for generating flatten array and response to client.
func GetFlattenArray(w http.ResponseWriter, r *http.Request) {
	var jsonData JSONData

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &jsonData)

	flatten, depth, err := logic.Flatten(jsonData.Array)

	if err != nil {
		log.Printf("Error flatting array: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%v, deep: %v", flatten, depth)
}
