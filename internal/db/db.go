package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(databaseURL string) (*sql.DB, error) {
	// Implement your database connection logic here
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}
	db.SetMaxOpenConns(25)//why set max open connections to 25? This is a common practice to limit the number of concurrent connections to the database, which can help prevent resource exhaustion and improve performance. The optimal number of connections may vary depending on the specific application and database configuration, so it may be adjusted based on performance testing and monitoring.
    db.SetConnMaxIdleTime(25)//why set max idle time to 25? Setting the maximum idle time for connections helps manage resources by closing idle connections after a certain period. This can prevent the database from holding onto unused connections, which can lead to resource exhaustion. The value of 25 seconds is a reasonable default, but it may be adjusted based on the application's usage patterns and performance requirements.
	db.SetConnMaxLifetime(5*time.Minute)//why set max lifetime to 5 minutes? Setting the maximum lifetime for connections ensures that connections are periodically refreshed, which can help prevent issues related to long-lived connections, such as stale connections or resource leaks. A value of 5 minutes is a common choice, but it may be adjusted based on the application's needs and database behavior.
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)//
	defer cancel()
	
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("db.PingContext: %w", err)
	}
	return db, nil
}