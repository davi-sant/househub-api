package config // declarando pacote

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var host,
	port,
	user,
	password,
	dbname,
	sslMode string

var DB *sql.DB
var logError string

func DBConnection() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("No .env file found")
	}

	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname = os.Getenv("DB_NAME")
	sslMode = os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslMode)
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		logError = fmt.Sprintf("error connecting to the database %s: error -> %s:", "DataBase connection: ", err.Error())
		log.Fatal(logError)
	}
	err = db.Ping()
	if err != nil {
		logError = fmt.Sprintf("error the Ping to database:  error -> %s:", err.Error())
		log.Fatal(logError)
	}

	DB = db
	fmt.Println("Database connection established.")
}

const checkExistsTableQuery = `
	select exists(
		select 1
			from information_schema.tables
		where table_schema = 'public' and  table_name = 'registros'
	);
`

func RunMigrations(sqlFile string) {

	var exists bool
	content, err := os.ReadFile(sqlFile)

	if err != nil {
		log.Fatal("Error reading migration file: ", err)
	}

	if err = DB.QueryRow(checkExistsTableQuery).Scan(&exists); err != nil {
		log.Fatal("Error executing consult table: ", err)
	}

	if exists {
		fmt.Println("Migrations table already exists.")
		return
	}

	for query := range strings.SplitSeq(string(content), ";") {
		strings.TrimSpace(query)
		if query == "" || strings.HasPrefix(query, "--") {
			continue
		}
		fmt.Println("Executing migration:", query)

		if _, err := DB.Exec(query); err != nil {
			log.Fatal("Error running migration: ", err)
		}
	}
	fmt.Println("Database migrations completed.")
}
