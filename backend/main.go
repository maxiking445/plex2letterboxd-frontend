package main

import (
	"fmt"
	"net/http"

	_ "github.com/maxiking445/plex-letterboxd-backend/docs"
	"github.com/maxiking445/plex-letterboxd-backend/internal/handlers"
	_ "github.com/swaggo/files"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title BFF API
// @version 1.0
// @host localhost:8080
// @BasePath /

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/executeScript", handlers.ExecuteScriptHandler)
	mux.HandleFunc("/exports", handlers.GetAllExportsHandler)
	mux.HandleFunc("/exports/{id}", handlers.GetExportHandler)

	mux.HandleFunc("/exports/remove/{id}", handlers.DeleteExportHandler)

	mux.HandleFunc("/settings", handlers.GetSettingsHandler)
	mux.HandleFunc("/settings/save", handlers.SaveSettingsHandler)
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	fmt.Println("Server: http://localhost:8080/swagger/index.html")
	http.ListenAndServe(":8080", cors(mux))
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Vary", "Origin")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Content-Type, Authorization, X-Zip-Password, X-Requested-With")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
