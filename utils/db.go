package utils

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"

	"polynux/yatm/db"
)

var DB *sql.DB
var Q *db.Queries

func Connect() *sql.DB {
  dbUrl := GetEnv("DB_URL")
  dbToken := GetEnv("DB_TOKEN")
  if dbUrl == "" {
    log.Fatal("DB_URL is not set")
    os.Exit(1)
  }
  if dbToken == "" {
    log.Fatal("DB_TOKEN is not set")
    os.Exit(1)
  }

  db, err := sql.Open("libsql", dbUrl + "?authToken=" + dbToken)
  if err != nil {
    log.Fatalf("Error opening database: %v", err)
    os.Exit(1)
  }

  return db
}

func init() {
  LoadEnv()
  DB = Connect()
  Q = db.New(DB)
  CreateTables(context.Background())
}

func LoadSql() string {
  sql, err := os.ReadFile("schema.sql")
  if err != nil {
    log.Fatalf("Error reading schema.sql: %v", err)
    os.Exit(1)
  }
  return string(sql)
}

func CreateTables(ctx context.Context) {
  schema := LoadSql()

  _, err := DB.ExecContext(ctx, schema)
  if err != nil {
    log.Fatalf("Error creating Praticiens table: %v", err)
    os.Exit(1)
  }
}

func LoadEnv() {
  err := godotenv.Load(".env.local")
  if err != nil {
    log.Fatalf("Error loading .env file: %v", err)
    os.Exit(1)
  }
}

func GetEnv(key string) string {
  return os.Getenv(key)
}
