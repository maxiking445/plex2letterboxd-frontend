package models

import "time"

type Export struct {
	Name            string    `json:"Name"`
	ExportItemCount int64     `json:"ExportItemCount"`
	Date            time.Time `json:"Date"`
}
