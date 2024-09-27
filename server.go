package main

import (
	"log"
	"net/http"
	"notes/graphql"
	"notes/models"
	"notes/utils"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Inisialisasi Redis client
	utils.InitRedis("localhost:6379", "") // Alamat dan password Redis

	// Membuat konfigurasi untuk koneksi database
	dsn := "host=localhost user=postgres password=postgres dbname=gogql port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Menentukan apakah akan melakukan migrasi otomatis
	autoMigrate := true // Ubah ke false jika tidak ingin melakukan migrasi otomatis

	if autoMigrate {
		if err := db.AutoMigrate(&models.User{}, &models.Note{}); err != nil {
			log.Fatalf("failed to migrate: %v", err)
		}
		log.Println("Database migrated successfully!")
	}

	// Atur CORS
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"}, // Ganti dengan domain yang diizinkan
		AllowCredentials: true,
		AllowedMethods:   []string{"POST","OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		Debug:            true,
	})


	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{DB: db}}))

	http.Handle("/graphql", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", corsMiddleware.Handler(utils.SessionMiddleware(srv)))


	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
