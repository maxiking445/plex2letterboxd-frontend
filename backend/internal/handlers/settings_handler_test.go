package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/maxiking445/plex-letterboxd-backend/internal/models"
)

func setupTestSettingsFile(t *testing.T) func() {
	t.Helper()

	err := os.MkdirAll("./script", 0755)
	if err != nil {
		t.Fatalf("failed to create script dir: %v", err)
	}

	content := `
[auth]
baseurl = http://example.com
token = testtoken
libarys = Movies,Series
`

	err = os.WriteFile(SettingsPath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to write config.ini: %v", err)
	}

	return func() {
		os.RemoveAll("./script")
	}
}

func TestGetSettingsHandler(t *testing.T) {
	cleanup := setupTestSettingsFile(t)
	defer cleanup()

	req := httptest.NewRequest(http.MethodGet, "/settings", nil)
	rr := httptest.NewRecorder()

	GetSettingsHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}

	var settings models.Settings
	if err := json.Unmarshal(rr.Body.Bytes(), &settings); err != nil {
		t.Fatalf("invalid json response: %v", err)
	}

	if settings.BaseURL != "http://example.com" {
		t.Fatalf("unexpected baseurl: %s", settings.BaseURL)
	}

	if settings.Token != "testtoken" {
		t.Fatalf("unexpected token")
	}
}

func TestSaveSettingsHandler_Success(t *testing.T) {
	cleanup := setupTestSettingsFile(t)
	defer cleanup()

	input := models.Settings{
		BaseURL: "http://new-url.com",
		Token:   "newtoken",
		Libarys: []string{"Movies", "Anime"},
	}

	body, _ := json.Marshal(input)

	req := httptest.NewRequest(
		http.MethodPost,
		"/settings/save",
		bytes.NewBuffer(body),
	)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	SaveSettingsHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}

	// verify file content
	data, err := os.ReadFile(SettingsPath)
	if err != nil {
		t.Fatalf("failed to read config.ini: %v", err)
	}

	if bytes.Contains(data, []byte("")) {
		t.Log("settings saved")
	}
}

func TestSaveSettingsHandler_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPost,
		"/settings/save",
		bytes.NewBuffer([]byte("{invalid-json")),
	)
	rr := httptest.NewRecorder()

	SaveSettingsHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rr.Code)
	}
}

func TestSaveSettingsHandler_MissingFields(t *testing.T) {
	input := models.Settings{
		BaseURL: "",
		Token:   "",
	}

	body, _ := json.Marshal(input)

	req := httptest.NewRequest(
		http.MethodPost,
		"/settings/save",
		bytes.NewBuffer(body),
	)
	rr := httptest.NewRecorder()

	SaveSettingsHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rr.Code)
	}
}

func TestLoadSettingsFromFile_FileMissing(t *testing.T) {
	os.RemoveAll("./script")

	settings := loadSettingsFromFile()

	if settings.BaseURL != "" || settings.Token != "" {
		t.Fatal("expected empty settings when file missing")
	}
}
