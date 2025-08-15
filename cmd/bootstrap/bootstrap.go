package bootstrap

import (
	"log"

	"github.com/alvaroeds/amaris-dragon-ball/internal/config"
	"github.com/alvaroeds/amaris-dragon-ball/internal/health"
	"github.com/alvaroeds/amaris-dragon-ball/internal/infrastructure/db/cache/redis"
	"github.com/alvaroeds/amaris-dragon-ball/internal/infrastructure/db/postgres"
	"github.com/alvaroeds/amaris-dragon-ball/internal/infrastructure/server/http"
	"github.com/alvaroeds/amaris-dragon-ball/pkg/character/application"
	character_external "github.com/alvaroeds/amaris-dragon-ball/pkg/character/infrastructure/external"
	character_handler "github.com/alvaroeds/amaris-dragon-ball/pkg/character/infrastructure/handler"
	character_repository "github.com/alvaroeds/amaris-dragon-ball/pkg/character/infrastructure/repository"
)

// Run bootstraps and starts the Dragon Ball API application
func Run() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Database connections
	postgresClient := postgres.NewPostgresClient(cfg.PostgresConnectionString())
	redisClient := redis.NewRedisClient(cfg.RedisAddress(), cfg.RedisPassword, cfg.RedisDB)

	// External API client
	dragonBallAPI := character_external.NewDragonBallAPIClient(cfg.DragonBallAPIURL, cfg.DragonBallAPITimeout)

	// Character handler
	characterRepository := character_repository.NewPostgresRepository(postgresClient.DB)
	cacheRepository := character_repository.NewCacheRepository(redisClient)
	characterService := application.NewService(characterRepository, dragonBallAPI, cacheRepository)
	characterHandler := character_handler.NewHandler(characterService)

	// Health handler
	healthHandler := health.NewHandler(postgresClient, redisClient)

	// Setup and start server
	router := http.Routes(characterHandler, healthHandler)
	server := http.NewServer(cfg.ServerPort, router)
	server.Start()
}
