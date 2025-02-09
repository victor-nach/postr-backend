package migrator

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
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

func Seed(db *gorm.DB) error {
	users := []domain.User{
		{
			ID:       uuid.NewString(),
			Firstname: "John",
			Lastname:  "Doe",
			Email:     "john@example.com",
			Street:    "123 Elm Street",
			City:      "New York",
			State:     "NY",
			Zipcode:   "10001",
			CreatedAt: time.Now(),
		},
		{
			ID:       uuid.NewString(),
			Firstname: "Jane",
			Lastname:  "Smith",
			Email:     "jane@example.com",
			Street:    "456 Oak Avenue",
			City:      "Los Angeles",
			State:     "CA",
			Zipcode:   "90001",
			CreatedAt: time.Now(),
		},
	}

	if err := db.Create(&users).Error; err != nil {
		return fmt.Errorf("failed to insert users: %w", err)
	}

	posts := []domain.Post{
		{ID:uuid.NewString(), UserID: users[0].ID, Title: "Post 1", Body: "Content of post 1", CreatedAt: time.Now()},
		{ID:uuid.NewString(), UserID: users[0].ID, Title: "Post 2", Body: "Content of post 2", CreatedAt: time.Now()},
		{ID:uuid.NewString(), UserID: users[1].ID, Title: "Post 3", Body: "Content of post 3", CreatedAt: time.Now()},
	}

	if err := db.Create(&posts).Error; err != nil {
		return fmt.Errorf("failed to insert posts: %w", err)
	}

	log.Println("Database seeded successfully with users and posts")
	return nil
}
