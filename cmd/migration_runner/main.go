package main

import (
	"log"
	"flag"
	"os"

	"go.uber.org/zap"

	"github.com/victor-nach/postr-backend/internal/config"
	"github.com/victor-nach/postr-backend/pkg/migrator"
	"github.com/victor-nach/postr-backend/internal/infrastructure/db"
	"github.com/victor-nach/postr-backend/pkg/logger"

)

func main() {
	seed := flag.Bool("seed", false, "seed the database after migrations")
	flag.Parse()

	appEnv, ok := os.LookupEnv(config.EnvAppEnv)
	if !ok {
		appEnv = config.DefaultAppEnv
	}

	logr, err := logger.NewLogger(appEnv)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer logr.Sync()

	db, sqlDB, err := db.New()
	if err != nil {
		logr.Fatal("failed to connect to database", zap.Error(err))
	}
	defer sqlDB.Close()

	if err := migrator.Migrate(sqlDB, "file://migrations"); err != nil {
		logr.Fatal("migration failed", zap.Error(err))
	}

	if *seed {
		if err := migrator.Seed(db); err != nil {
			logr.Fatal("seeding failed", zap.Error(err))
		}
	}

	logr.Info("Migration runner finished successfully")
	os.Exit(0)
}
