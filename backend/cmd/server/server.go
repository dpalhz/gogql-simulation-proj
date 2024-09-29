package main

import (
	"net/http"
	"notes/backend/graphql"
	"notes/backend/graphql/resolver"
	log "notes/backend/internal/logger"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"notes/backend/models"
	"notes/backend/utils"
)

const defaultPort = "8080"

func main() {
	// Initialize logger first
	log.InitializeLogger("Gopher", log.TimeFormatUnix)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	utils.InitRedis("localhost:6379", "") // Alamat dan password Redis

	// Initialize database connection
	dsn := "host=localhost user=postgres password=postgres dbname=gogql port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.LogFatalf("failed to connect database: %v", err)
	}

	// Menentukan apakah akan melakukan migrasi otomatis
	autoMigrate := true // Ubah ke false jika tidak ingin melakukan migrasi otomatis

	if autoMigrate {
		if err := db.AutoMigrate(&models.User{}, &models.Note{}); err != nil {
			log.LogFatalf("failed to migrate: %v", err)
		}
		log.LogInfo("Database migrated successfully!")
	}

	app := fiber.New()

	// Setup CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, http://localhost:8080",
		AllowCredentials: true,
		AllowMethods:     "POST, OPTIONS",
		AllowHeaders:     "Content-Type, Authorization",
	}))

	// Create a new Resolver with the database connection
	resolver := resolver.NewResolver(db)

	// Create a new GraphQL server
	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolver}))

	// Create a handler for the GraphQL server
	gqlHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srv.ServeHTTP(w, r)
	})

	// Create a handler for the GraphQL playground
	playgroundHandler := playground.Handler("GraphQL playground", "/query")

	// Setup routes
	app.All("/query", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(gqlHandler)(c.Context())
		return nil
	})

	app.All("/", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(playgroundHandler)(c.Context())
		return nil
	})

	log.LogInfof("connect to http://localhost:%s/ for GraphQL playground", port)
	log.LogFatal(app.Listen(":" + port))
}
