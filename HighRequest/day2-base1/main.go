package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func payloadH(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var content = &PayloadCollection{}
	if err := json.NewDecoder(io.LimitReader(r.Body, 4096)).Decode(&content); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for _, payload := range content.Payloads {
		JobQueue <- Job{Pay: payload}
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()
	http.HandleFunc("/payload", payloadH)
	fmt.Println(http.ListenAndServe(":1234", nil))
}
