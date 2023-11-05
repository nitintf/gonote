package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

// oncer for the config loading
var loadOnce sync.Once

func get(name, def string) string {
	// Load .env file only once
	loadOnce.Do(func() {
		godotenv.Load(".env")
	})

	if value := os.Getenv(name); value != "" {
		return value
	}

	return def
}

var (
	DatabaseURL = get("DATABASE_URL", "file:gonote.db?cache=shared")
	Environment = get("GO_ENV", "development")
	Port        = get("PORT", ":8181")

	GlovesExtensionsToWatch = []string{".go", ".html", ".css", ".js"}
	GlovesExcludePaths      = []string{""}
)
