package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init(dbPath string) error {
	// Ensure the directory exists
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create db directory: %w", err)
	}

	var err error
	DB, err = sql.Open("sqlite3", dbPath+"?_journal_mode=WAL&_busy_timeout=5000")
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	// Create tables
	if err = createTables(); err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	log.Printf("Database initialized at %s", dbPath)
	return nil
}

func createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS tool_usage (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		tool_name TEXT NOT NULL,
		input_summary TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := DB.Exec(query)
	return err
}

func RecordUsage(toolName, inputSummary string) {
	if DB == nil {
		return
	}
	_, err := DB.Exec(
		"INSERT INTO tool_usage (tool_name, input_summary) VALUES (?, ?)",
		toolName, inputSummary,
	)
	if err != nil {
		log.Printf("Failed to record usage: %v", err)
	}
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
