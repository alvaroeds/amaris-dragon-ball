package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var (
	postgresClient *Client
	postgresOnce   sync.Once
)

// Client is a client for the PostgreSQL db engine.
type Client struct {
	*sql.DB
}

// NewPostgresClient returns a new client for postgres.
func NewPostgresClient(source string) *Client {
	postgresOnce.Do(func() {
		db, err := sql.Open("postgres", source)
		if err != nil {
			fmt.Printf("❌ Error connecting to PostgreSQL: %v\n", err)
			panic(err)
		}

		// Limit the number of open connections to avoid
		// memory problems in the database.
		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(25)
		db.SetConnMaxLifetime(2 * time.Minute)

		err = db.PingContext(context.Background())
		if err != nil {
			fmt.Printf("❌ Error pinging PostgreSQL database: %v\n", err)
			panic(err)
		}

		fmt.Println("✅ Successfully connected to PostgreSQL")
		postgresClient = &Client{db}
	})
	return postgresClient
}

// Ping tests the PostgreSQL connection
func (c *Client) Ping() error {
	return c.DB.PingContext(context.Background())
}
