package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func getDNSFromEnv() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	if sslmode == "" {
		sslmode = "disable"
	}

	if host == "" {
		host = "localhost"
	}

	if port == "" {
		port = "5432"
	}

	fmt.Println("host:", host)
	fmt.Println("user:", user)
	fmt.Println("password:", password)
	fmt.Println("dbname:", dbname)
	fmt.Println("port:", port)
	fmt.Println("sslmode:", sslmode)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host,
		user,
		password,
		dbname,
		port,
		sslmode,
	)

	return dsn
}

// ConnectToDatabase connects to the database and returns a pointer to the database
// object and an error if there is one.
func ConnectToDatabase() *gorm.DB {
	dsn := getDNSFromEnv()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
