package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)



func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go movies up and running",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)

	if err != nil {
		fmt.Println(err)
	}

	app.Writer(w, r, http.StatusOK).Write(out)
}
