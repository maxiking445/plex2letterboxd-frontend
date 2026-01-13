package handlers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	shortuuid "github.com/lithammer/shortuuid/v4"
	"github.com/maxiking445/plex-letterboxd-backend/internal/models"
	_ "github.com/maxiking445/plex-letterboxd-backend/internal/models"
)

const (
	DataPath    = "/data/"
	Time_Format = "020106_1504"
)

func NewShortUUID() string {
	id := shortuuid.New()
	if len(id) < 6 {
		return id
	}
	return id[:4]
}

// ExecuteScriptHandler godoc
//
// @Summary      executes script
// @Description  executes plex2letterbox python script with current settings
// @Tags         Data
// @Accept       json
// @Produce      json
// @Success      200      {string} string  "Sucessfully executed export"
// @Failure      400      {string} string  "Settings may missing or are incomplete"
// @Failure      500      {string} string  "Unknown error during execution"
// @Router       /executeScript [GET]
func ExecuteScriptHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	timestamp := time.Now().Format(Time_Format)

	csvFilename := fmt.Sprintf("Export_"+NewShortUUID()+"_%s.csv", timestamp)

	venvPython := "./env/bin/python"

	libarys := loadSettingsFromFile().Libarys

	cmd := exec.Command(venvPython, "-m", "plex2letterboxd",
		"-i", "config.ini",
		"-o", ".."+DataPath+csvFilename,
		"-s")

	cmd.Args = append(cmd.Args, libarys...)
	cmd.Dir = "./script"

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Python Error: %v\nOutput: %s", err, output)
		http.Error(w, `{"error": "Export failed: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	log.Printf("Python Output: %s", output)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Export Sucessfull! CSV: letterboxd.csv",
	})
}

// GetAllExportsHandler godoc
//
// @Summary      Returns all existing exports
// @Description  Returns all existing exports under /data
// @Tags         Data
// @Produce      json
// @Success      200 {array} models.Export "Successfully retrieved exports"
// @Failure      500 {string} string "Unknown error during execution"
// @Router       /exports [GET]
func GetAllExportsHandler(w http.ResponseWriter, r *http.Request) {
	exports := []models.Export{}

	files, err := os.ReadDir("./data")
	if err != nil {
		http.Error(w, "files could not be read", http.StatusBadRequest)
		return
	}

	for _, file := range files {
		var fileName = file.Name()

		re := regexp.MustCompile(`^Export_(.+)_(\d{6})_(\d{4})\.csv$`)
		matches := re.FindStringSubmatch(fileName)

		timestampStr := parseToTime(matches[2] + "_" + matches[3])

		exports = append(exports, models.Export{
			Name:            fileName,
			Date:            timestampStr,
			ExportItemCount: countEntrys(fileName),
		})
	}

	sort.Slice(exports, func(i, j int) bool {
		return exports[i].Date.After(exports[j].Date)
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exports)
}

func parseToTime(timestampStr string) time.Time {
	parsedTime, err := time.Parse(Time_Format, timestampStr)
	if err != nil {
		log.Printf("Time convertion failed!: %s", err)
		return time.Now()
	}
	return parsedTime
}

func countEntrys(fileName string) int64 {

	file, err := os.Open("." + DataPath + fileName)
	if err != nil {
		log.Printf("Python Output: %s", err)

		return 0
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}

	if len(records) <= 1 {
		return 0
	}

	records = records[1:]

	count := 0

	for _, record := range records {
		if len(record) < 2 {
			continue
		}
		count++
	}

	return int64(count)
}

// GetExportHandler godoc
//
// @Summary      returns CSV of export
// @Description  reads CSV
// @Tags         Data
// @Param        id   path      string  true  "Export ID (Filename)"
// @Produce      application/octet-stream
// @Success      200  {file}  file  "CSV File"
// @Failure      404  {string} string  "File not found"
// @Router       /exports/{id} [GET]
func GetExportHandler(w http.ResponseWriter, r *http.Request) {

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 || pathParts[1] != "exports" {
		http.Error(w, `{"error": "invalid path"}`, http.StatusBadRequest)
		return
	}

	id := pathParts[2]
	if id == "" {
		http.Error(w, `{"error": "ID required"}`, http.StatusBadRequest)
		return
	}

	filePath := filepath.Join("." + DataPath + id)

	data, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, `{"error": "File not found"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", id))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// DeleteExportHandler godoc
//
// @Summary      Deletes a specific export
// @Description  Deletes the CSV export file with the given ID
// @Tags         Data
// @Produce      json
// @Param        id   path      string  true  "Export ID (Filename)"
// @Success      200  {string}  string  "Export successfully deleted"
// @Failure      404  {string}  string  "File not found"
// @Failure      500  {string}  string  "Internal server error"
// @Router       /exports/remove/{id} [delete]
func DeleteExportHandler(w http.ResponseWriter, r *http.Request) {

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 || pathParts[1] != "exports" {
		http.Error(w, `{"error": "invalid path"}`, http.StatusBadRequest)
		return
	}

	id := pathParts[3]
	if id == "" {
		http.Error(w, `{"error": "ID required"}`, http.StatusBadRequest)
		return
	}
	filePath := filepath.Join("." + DataPath + id)
	err := os.Remove(filePath)
	if err != nil {
		log.Printf("Failed to delete file %s: %v", filePath, err)
		http.Error(w, `{"error": "Could not remove file"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Deletion Sucessfull!",
	})
}
