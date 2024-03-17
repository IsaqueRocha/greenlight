package main

import (
	"net/http"
)

const errHealthMessage = "The server encountered a problem and could not process your request"

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, errHealthMessage, http.StatusInternalServerError)
	}
}
