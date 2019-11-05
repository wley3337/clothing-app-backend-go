package utils

import (
	"encoding/json"
	"net/http"
)

//Message this is my first GO app so I'm not sure what to write here
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

//Respond this is my response comment as it is still my first GO app
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
