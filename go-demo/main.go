package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://postgres:postgres@localhost:5432/satgo?sslmode=disable"
	}

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to PostgreSQL database")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS forms (
		id VARCHAR(255) PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		content TEXT NOT NULL,
		created_at VARCHAR(255) NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created forms table")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS messages (
		id VARCHAR(255) PRIMARY KEY,
		form_id VARCHAR(255) NOT NULL,
		content TEXT NOT NULL,
		created_at VARCHAR(255) NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created messages table")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS invoices (
		payment_hash VARCHAR(255) PRIMARY KEY,
		payment_request TEXT NOT NULL,
		site_id VARCHAR(255) NOT NULL,
		message TEXT NOT NULL,
		amount_sats BIGINT NOT NULL,
		created_at VARCHAR(255) NOT NULL,
		expires_at VARCHAR(255) NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created invoices table")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS webhook_events (
		id VARCHAR(255) PRIMARY KEY,
		event_type VARCHAR(255) NOT NULL,
		status VARCHAR(50) NOT NULL,
		created_at VARCHAR(255) NOT NULL,
		processed_at VARCHAR(255)
	)`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created webhook_events table")

	fmt.Println("All tables created successfully with PostgreSQL (lib/pq)")
}