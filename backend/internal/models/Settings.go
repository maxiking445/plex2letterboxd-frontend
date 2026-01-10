package models

type Settings struct {
	BaseURL string `ini:"baseurl" json:"baseurl"`
	Token   string `ini:"token"   json:"token"`
}
