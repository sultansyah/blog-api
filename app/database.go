package app

import (
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/blog-api")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

// migrate create -ext sql -dir db/migrations create_users_table
// migrate create -ext sql -dir db/migrations create_categories_table
// migrate create -ext sql -dir db/migrations create_blogs_table
// migrate create -ext sql -dir db/migrations create_comments_table

// up
// migrate -database "mysql://root@tcp(localhost:3306)/blog-api" -path db/migrations up

// down
// migrate -database "mysql://root@tcp(localhost:3306)/blog-api" -path db/migrations down
