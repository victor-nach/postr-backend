package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/victor-nach/postr-backend/internal/config"
	"github.com/victor-nach/postr-backend/internal/handlers"
	"github.com/victor-nach/postr-backend/internal/infrastructure/db"
	"github.com/victor-nach/postr-backend/internal/infrastructure/repositories"
	"github.com/victor-nach/postr-backend/internal/services/postsservice"
	"github.com/victor-nach/postr-backend/internal/services/usersservice"
	"github.com/victor-nach/postr-backend/pkg/logger"
)

func main() {
	appEnv, ok := os.LookupEnv(config.EnvAppEnv)
	if !ok {
		appEnv = config.DefaultAppEnv
	}

	logr, err := logger.NewLogger(appEnv)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer logr.Sync()

	//  Load config
	cfg, err := config.Load(logr)
	if err != nil {
		logr.Fatal("failed to load configuration", zap.Error(err))
	}

	// Initialize DB
	gormDB, sqlDB, err := db.New()
	if err != nil {
		logr.Fatal("failed to connect to database", zap.Error(err))
	}
	defer sqlDB.Close()

	// Initialize repos
	userRepo := repositories.NewUserRepository(gormDB)
	postRepo := repositories.NewPostRepository(gormDB)

	// Initialize services
	userSvc := usersservice.New(userRepo, logr)
	postSvc := postsservice.New(postRepo, userRepo, logr)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userSvc, logr)
	postHandler := handlers.NewPostHandler(postSvc,  logr)

	RunServer(cfg.Port, userHandler, postHandler, logr)
}

// RunServer creates and mounts the router, starts the server in a goroutine,
// and listens for OS signals to gracefully shutdown
func RunServer(port string, userHandler *handlers.UserHandler, postHandler *handlers.PostHandler, logr *zap.Logger) {
	router := createRouter(userHandler, postHandler)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Run the server in a goroutine
	go func() {
		logr.Info("Starting server", zap.String("address", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logr.Fatal("failed to start server", zap.Error(err))
		}
	}()

	// Create a channel to listen for OS quit signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	logr.Info("Shutting down server...")

	// Attempt graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logr.Fatal("Server forced to shutdown", zap.Error(err))
	}

	logr.Info("Server exiting")
}

func createRouter(userHandler *handlers.UserHandler, postHandler *handlers.PostHandler) http.Handler {
	router := gin.Default()

	router.Use(cors.Default())

	router.POST("/users", userHandler.CreateUser)
	router.GET("/users", userHandler.ListUsers)
	router.GET("/users/count", userHandler.CountUsers)
	router.GET("/users/:id", userHandler.GetUserByID)

	router.POST("/posts", postHandler.CreatePost)
	router.DELETE("/posts/:id", postHandler.DeletePost)
	router.GET("/posts/:userId", postHandler.ListPostsByUserID)

	return router
}
