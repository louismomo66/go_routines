package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func initDB() *sql.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("can't connect to database")
	}
	if err := verifyTables(conn); err != nil {
		log.Printf("Database verification failed: %v", err)
		// Attempt to initialize schema
		if err := initializeSchema(conn); err != nil {
			log.Panic("Could not initialize database schema:", err)
		}
	}
	return conn
}
func initializeSchema(db *sql.DB) error {
	// Read schema file
	schema, err := os.ReadFile("db.sql")
	if err != nil {
		return fmt.Errorf("could not read schema file: %v", err)
	}

	// Execute schema
	_, err = db.Exec(string(schema))
	if err != nil {
		return fmt.Errorf("could not initialize schema: %v", err)
	}

	log.Println("Schema initialized successfully")
	return nil
}
func verifyTables(db *sql.DB) error {
	tables := []string{"users", "plans", "user_plans"}

	for _, table := range tables {
		var exists bool
		query := `
            SELECT EXISTS (
                SELECT FROM information_schema.tables 
                WHERE table_schema = 'public' 
                AND table_name = $1
            );`

		err := db.QueryRow(query, table).Scan(&exists)
		if err != nil {
			return fmt.Errorf("error checking table %s: %v", table, err)
		}

		if !exists {
			return fmt.Errorf("table %s does not exist", table)
		}
	}

	return nil
}
func connectToDB() *sql.DB {
	counts := 0
	dsn := os.Getenv("DSN")
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("postgres not yet ready...")
		} else {
			log.Print("connected to database!")
			return connection
		}
		if counts > 10 {
			return nil
		}
		log.Print("Backing off for 1 second")
		time.Sleep(1 * time.Second)
		counts++
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
