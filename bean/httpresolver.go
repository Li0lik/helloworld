package bean

import (
	"encoding/json"
	"net/http"
	"os"
)

type HttpResolver struct {
}

func NewHttpResolver() *HttpResolver {
	o := HttpResolver{}

	return &o
}
func (o *HttpResolver) Resolver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	resp := make(map[string]string)
	resp["Status"] = "ok"
	resp["Hostname"] = hostname

	if r.Method == "POST" {
		resp["Method"] = "POST"
	}
	if r.Method == "GET" {
		resp["Method"] = "GET"
	}
	if r.Method == "DELETE" {
		resp["Method"] = "DELETE"
	}
	if r.Method == "PUT" {
		resp["Method"] = "PUT"
	}
	rawdata, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "[ERROR] Something was wrong "+err.Error(), http.StatusInternalServerError)
		w.Write(make([]byte, 0))
		return
	}

	w.Write(rawdata)
}
