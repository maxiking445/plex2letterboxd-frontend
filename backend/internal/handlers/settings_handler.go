package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/maxiking445/plex-letterboxd-backend/internal/models"
	"gopkg.in/ini.v1"
)

const SettingsPath = "./script/config.ini"

// GetSettingsHandler godoc
//
// @Summary      returns current setings
// @Tags         Settings
// @Accept       json
// @Produce      json
// @Success      200      {string} string  "Sucessfully returned settings"
// @Failure      400      {string} string  "Settings may missing or are incomplete"
// @Failure      500      {string} string  "Unknown error during execution"
// @Router       /settings [GET]
func GetSettingsHandler(w http.ResponseWriter, r *http.Request) {
	cfg, err := ini.Load(SettingsPath)
	if err != nil {
		http.Error(
			w,
			`{"error":"Settings extraction failed: `+err.Error()+`"}`,
			http.StatusInternalServerError,
		)
		return
	}

	config := &models.Settings{}

	err = cfg.Section("auth").MapTo(config)
	if err != nil {
		http.Error(
			w,
			`{"error":"Settings parsing failed: `+err.Error()+`"}`,
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config)
}

// SaveSettingsHandler godoc
// @Summary      saves settings
// @Description Receives baseurl and token and stores them in config.ini
// @Tags         Settings
// @Accept       json
// @Produce      json
// @Param        settings  body      models.Settings  true  "Settings payload"
// @Success      200       {string}   string  "Successfully saved settings"
// @Failure      400       {string}   string  "Settings may be missing or are incomplete"
// @Failure      500       {string}   string  "Unknown error during execution"
// @Router       /settings/save [post]
func SaveSettingsHandler(w http.ResponseWriter, r *http.Request) {
	var input models.Settings

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"error":"Invalid JSON body"}`, http.StatusBadRequest)
		return
	}

	if input.BaseURL == "" || input.Token == "" {
		http.Error(w, `{"error":"baseurl and token are required"}`, http.StatusBadRequest)
		return
	}

	cfg, err := ini.LooseLoad(SettingsPath)
	if err != nil {
		http.Error(
			w,
			`{"error":"Failed to load settings file: `+err.Error()+`"}`,
			http.StatusInternalServerError,
		)
		return
	}
	// write file
	section := cfg.Section("auth")
	section.Key("baseurl").SetValue(input.BaseURL)
	section.Key("token").SetValue(input.Token)

	if err := cfg.SaveTo(SettingsPath); err != nil {
		http.Error(
			w,
			`{"error":"Failed to save settings file: `+err.Error()+`"}`,
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"settings saved"}`))
}
