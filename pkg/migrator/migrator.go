package migrator

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"

	"github.com/victor-nach/postr-backend/internal/domain"
)

// Migrate applies all up migrations from the specified migrations path
func Migrate(db *sql.DB, migrationsPath string) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("could not create migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"sqlite3",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %w", err)
	}

	log.Println("Migrations applied successfully")
	return nil
}

// Seed seeds the db
func Seed(db *gorm.DB) error {
	// Read seed users and Insert
	users, err := readUsersFromJSON("seeds/users.json")
	if err != nil {
		return fmt.Errorf("failed to read users from JSON: %w", err)
	}

	if err := db.Create(&users).Error; err != nil {
		return fmt.Errorf("failed to insert users: %w", err)
	}

	// Read seed posts and Insert
	posts, err := readPostsFromJSON("seeds/posts.json")
	if err != nil {
		return fmt.Errorf("failed to read posts from JSON: %w", err)
	}

	if err := db.Create(&posts).Error; err != nil {
		return fmt.Errorf("failed to insert posts: %w", err)
	}

	log.Println("Database seeded successfully with users and posts from JSON files")
	return nil
}

func readUsersFromJSON(filePath string) ([]domain.User, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var users []domain.User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func readPostsFromJSON(filePath string) ([]domain.Post, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var posts []domain.Post
	err = json.Unmarshal(data, &posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
