package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Supabase struct {
	db *sql.DB
}

type SupabaseConnection struct {
	user     string
	password string
	host     string
	port     string
	dbname   string
}

func New() *Supabase {
	return &Supabase{
		db: &sql.DB{},
	}
}

func (sb *Supabase) Connect() *sql.DB {
	godotenv.Load(".env")

	conn := SupabaseConnection{
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		dbname:   os.Getenv("DB_NAME"),
	}

	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", conn.user, conn.password, conn.host, conn.port, conn.dbname))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connect to db successfully!")

	return db
}
