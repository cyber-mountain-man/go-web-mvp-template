package main

import (
	"log"
	"net/http"
	"os"

	"golangApp-postgres/internal/db"
	"golangApp-postgres/internal/router"

	"github.com/joho/godotenv"
)

func main() {
	// 📦 Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found. Using system env variables...")
	}

	// 🔌 Connect to the database
	if err := db.InitDB(); err != nil {
		log.Fatalf("❌ Failed to initialize database: %v", err)
	}
	defer db.DB.Close()

	// 🌐 Set up the router
	r := router.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server running at http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
