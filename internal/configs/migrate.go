package configs

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// RunMigrations applies all pending .sql migrations from the given directory.
// Migrations are tracked in a schema_migrations table by filename.
// Each file is expected to follow the goose format with `-- +goose Up` and
// `-- +goose Down` sections; only the Up section is executed.
func RunMigrations(db *sql.DB, dir string) error {
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS schema_migrations (
		filename text PRIMARY KEY,
		applied_at timestamptz NOT NULL DEFAULT NOW()
	)`); err != nil {
		return fmt.Errorf("create schema_migrations: %w", err)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("read migrations dir %q: %w", dir, err)
	}

	var files []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".sql") {
			files = append(files, e.Name())
		}
	}
	sort.Strings(files)

	for _, name := range files {
		var exists bool
		if err := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM schema_migrations WHERE filename=$1)`, name).Scan(&exists); err != nil {
			return fmt.Errorf("check migration %s: %w", name, err)
		}
		if exists {
			continue
		}

		raw, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			return fmt.Errorf("read %s: %w", name, err)
		}
		up := extractGooseUp(string(raw))
		if strings.TrimSpace(up) == "" {
			log.Printf("migration %s has empty Up section; skipping", name)
		} else {
			if _, err := db.Exec(up); err != nil {
				return fmt.Errorf("apply %s: %w", name, err)
			}
		}
		if _, err := db.Exec(`INSERT INTO schema_migrations(filename) VALUES ($1)`, name); err != nil {
			return fmt.Errorf("record %s: %w", name, err)
		}
		log.Printf("applied migration %s", name)
	}
	return nil
}

// extractGooseUp returns the SQL between `-- +goose Up` and `-- +goose Down`,
// stripping `-- +goose StatementBegin/End` markers.
func extractGooseUp(s string) string {
	upIdx := strings.Index(s, "-- +goose Up")
	if upIdx < 0 {
		return s
	}
	rest := s[upIdx+len("-- +goose Up"):]
	if downIdx := strings.Index(rest, "-- +goose Down"); downIdx >= 0 {
		rest = rest[:downIdx]
	}
	rest = strings.ReplaceAll(rest, "-- +goose StatementBegin", "")
	rest = strings.ReplaceAll(rest, "-- +goose StatementEnd", "")
	return rest
}

// SeedDefaultUsers ensures `admin` (superuser) and `user` (regular) exist
// with the password "password" stored as a bcrypt hash. Idempotent.
func SeedDefaultUsers(db *sql.DB) error {
	hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("bcrypt hash: %w", err)
	}
	users := []struct {
		username, role string
	}{
		{"admin", "admin"},
		{"user", "user"},
	}
	for _, u := range users {
		if _, err := db.Exec(`INSERT INTO users (username, password_hash, role)
			VALUES ($1, $2, $3) ON CONFLICT (username) DO NOTHING`,
			u.username, string(hash), u.role); err != nil {
			return fmt.Errorf("seed user %s: %w", u.username, err)
		}
	}
	return nil
}
