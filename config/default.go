package config

import "time"


var defaultConfig = map[string]interface{} {
	"auth.context_key":                      "claims",
	"auth.sign_key":                         "jwt_secret",
	"auth.access_subject":                   "ac",
	"auth.refresh_subject":                  "rt",
	"auth.access_expiration_time":           time.Minute * 15,
	"auth.refresh_expiration_time":          time.Hour * 24 * 7,
}