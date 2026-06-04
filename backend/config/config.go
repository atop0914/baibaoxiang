package config

import "os"

type Config struct {
	Port     string
	DBPath   string
	Debug    bool
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "data/baibaoxiang.db"
	}
	debug := os.Getenv("DEBUG") == "true"
	return &Config{
		Port:   port,
		DBPath: dbPath,
		Debug:  debug,
	}
}
