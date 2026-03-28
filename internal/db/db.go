package db

import (
    "database/sql"
    "fmt"
    "os"

    _ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_SSL_MODE"),
    )

    return sql.Open("postgres", dsn)
}

/**
CREATE TABLE queue_entries (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ticket_id   TEXT UNIQUE NOT NULL,
    user_id     TEXT NOT NULL,
    queue_id    TEXT NOT NULL,
    position    INT NOT NULL,
    status      TEXT DEFAULT 'waiting',
    created_at  TIMESTAMP DEFAULT NOW()
);
*/