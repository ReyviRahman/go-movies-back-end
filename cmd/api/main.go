package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ReyviRahman/go-movies-back-end/internal/repository"
	"github.com/ReyviRahman/go-movies-back-end/internal/repository/dbrepo"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const port = 8080

type application struct {
	Domain string
	DSN    string
	DB     repository.DatabaseRepo
}

func main() {
	var app application
	app.Domain = "example.com"

	// Load file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Ambil DSN dari env
	app.DSN = os.Getenv("DSN")

	// Lanjutkan proses koneksi DB seperti tadi...
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal("Gagal ping database:", err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	log.Println("Database terhubung ke Supabase!")
	log.Println("Starting on port", port)

	// Jalankan seeder
	// err = app.DB.SeedMovies()
	// if err != nil {
	// 	log.Println("Gagal seeding (mungkin data sudah ada):", err)
	// } else {
	// 	log.Println("Seeding berhasil!")
	// }

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
