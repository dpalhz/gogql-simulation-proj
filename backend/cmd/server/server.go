package main

import (
	"net/http"
	"notes/backend/graphql"
	"notes/backend/graphql/resolver"
	log "notes/backend/internal/logger"
	middlware "notes/backend/middleware"
	"os"

	"notes/backend/utils"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

const defaultPort = "8080"

func main() {
	log.InitializeLogger("Gopher", log.TimeFormatUnix)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Initiate Redis
	utils.InitRedis("localhost:6379", "")

	// Initialize Database
	db, err := utils.InitDB()
	if err != nil {
		log.LogFatalf("failed to connect database: %v", err)
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
	services := utils.InitServices(db)
	resolver := resolver.NewResolver(services)

	// Create a new GraphQL server
	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolver}))

	// Create a handler for the GraphQL server with session middleware
	gqlHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srv.ServeHTTP(w, r)
	})
	gqlHandlerWithMiddleware := middlware.SessionMiddleware(gqlHandler)
	

	// Create a handler for the GraphQL playground
	playgroundHandler := playground.Handler("GraphQL playground", "/query")

	// Setup routes
	app.Post("/query", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(gqlHandlerWithMiddleware)(c.Context())
		return nil
	})

	app.Get("/", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(playgroundHandler)(c.Context())
		return nil
	})

	log.LogInfof("connect to http://localhost:%s/ for GraphQL playground", port)
	log.LogFatal(app.Listen(":" + port))
}
