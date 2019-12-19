package handler

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"zvon/logger"
	"zvon/state"
)

const name = "WEBHOOK"

// Web hook that gets notified of issue change.
func IssuesWebhook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Decode JSON.
	decoder := json.NewDecoder(r.Body)
	data := make(map[string]interface{})
	err := decoder.Decode(&data)
	if err != nil {
		logger.Log(err.Error(), name, logger.LevelInfo)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Get object attributes.
	attributes, ok := data["object_attributes"].(map[string]interface{})
	if !ok {
		logger.Log("Could not get object attributes.", name, logger.LevelInfo)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the action that was executed on the object.
	action, ok := attributes["action"].(string)
	if !ok {
		logger.Log("Could not get action.", name, logger.LevelInfo)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// If action was close, then the issue was closed and needs banging is set to true.
	if action == "close" {
		state.Current.SetNeedsBanging(true)
	}

	// Write response to client.
	w.WriteHeader(http.StatusAccepted)
}
