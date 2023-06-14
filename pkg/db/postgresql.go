package db

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"database/sql"

	_ "github.com/lib/pq"
	"github.com/sulis96/quinzia-golang-instrumentations/config"
)

type (

	// IDatabase is an interface that has database related functions
	IDatabase interface {
		Close() error
		Ping() error
	}

	// Database is a database instance
	Database struct {
		TZ       string
		Database *sql.DB
	}
)

func NewDatabase(dc *config.DbConfig) (*Database, error) {
	tz := "Asia/Jakarta"
	dbConfig := fmt.Sprintf(
		"host=%s port=%s user=%s password='%s' dbname=%s search_path=%s sslmode=disable",
		dc.Host,
		dc.Port,
		dc.Username,
		dc.Password,
		dc.DbName,
		dc.Schema,
	)
	db, err := sql.Open(dc.Driver, dbConfig)
	if err != nil {
		panic("missing db")
	}

	maxConn := os.Getenv("DB_MAX_CONNECTION")
	c, _ := strconv.Atoi(maxConn)
	db.SetMaxOpenConns(c)

	maxIdle := os.Getenv("DB_MAX_IDLE_CONNECTION")
	i, _ := strconv.Atoi(maxIdle)
	db.SetMaxIdleConns(i)

	lifeTime := os.Getenv("DB_MAX_LIFETIME_CONNECTION")
	lt, _ := time.ParseDuration(lifeTime)
	db.SetConnMaxLifetime(lt)

	if db.Ping() != nil {
		panic("missing db")
	}

	return &Database{
		TZ:       tz,
		Database: db,
	}, nil
}

// Close closes the database connection
func (d *Database) Close() error {
	return d.Database.Close()
}

// Ping is used to ping database connection
func (d *Database) Ping() error {
	return d.Database.Ping()
}
