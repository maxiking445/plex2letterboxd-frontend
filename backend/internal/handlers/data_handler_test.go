package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/maxiking445/plex-letterboxd-backend/internal/models"
)

func setupTestDataDir(t *testing.T) func() {
	t.Helper()

	err := os.MkdirAll("./data", 0755)
	if err != nil {
		t.Fatalf("failed to create data dir: %v", err)
	}

	// Test CSV
	content := "title,year\nMovie 1,2020\nMovie 2,2021\n"
	err = os.WriteFile("./data/Export_test_010101_1200.csv", []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to write test csv: %v", err)
	}

	return func() {
		os.RemoveAll("./data")
	}
}

func TestGetAllExportsHandler(t *testing.T) {
	cleanup := setupTestDataDir(t)
	defer cleanup()

	req := httptest.NewRequest(http.MethodGet, "/exports", nil)
	rr := httptest.NewRecorder()

	GetAllExportsHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}

	var exports []models.Export
	err := json.Unmarshal(rr.Body.Bytes(), &exports)
	if err != nil {
		t.Fatalf("invalid json response: %v", err)
	}

	if len(exports) != 1 {
		t.Fatalf("expected 1 export, got %d", len(exports))
	}
}

func TestGetExportHandler(t *testing.T) {
	cleanup := setupTestDataDir(t)
	defer cleanup()

	req := httptest.NewRequest(
		http.MethodGet,
		"/exports/Export_test_010101_1200.csv",
		nil,
	)
	rr := httptest.NewRecorder()

	GetExportHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}

	if rr.Header().Get("Content-Type") != "text/csv" {
		t.Fatalf("expected csv content-type")
	}
}

func TestGetExportHandler_NotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/exports/notfound.csv", nil)
	rr := httptest.NewRecorder()

	GetExportHandler(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", rr.Code)
	}
}

func TestDeleteExportHandler(t *testing.T) {
	cleanup := setupTestDataDir(t)
	defer cleanup()

	req := httptest.NewRequest(
		http.MethodDelete,
		"/exports/remove/Export_test_010101_1200.csv",
		nil,
	)
	rr := httptest.NewRecorder()

	DeleteExportHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}

	if _, err := os.Stat(filepath.Join("./data", "Export_test_010101_1200.csv")); !os.IsNotExist(err) {
		t.Fatalf("file was not deleted")
	}
}

func TestDeleteExportHandler_NotFound(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodDelete,
		"/exports/remove/does-not-exist.csv",
		nil,
	)
	rr := httptest.NewRecorder()

	DeleteExportHandler(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", rr.Code)
	}
}

func TestExecuteScriptHandler_Skipped(t *testing.T) {
	t.Skip("skipped because it executes external python script")

	req := httptest.NewRequest(http.MethodGet, "/executeScript", nil)
	rr := httptest.NewRecorder()

	ExecuteScriptHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
}
