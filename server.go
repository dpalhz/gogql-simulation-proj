package main

import (
	"log"
	"net/http"
	"os"

	"notes/graph" // Ganti dengan path yang sesuai untuk package graph Anda
	"notes/models"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gorm.io/driver/postgres" // Impor driver PostgreSQL
	"gorm.io/gorm"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

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

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
