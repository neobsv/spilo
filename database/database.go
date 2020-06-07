package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"./docs"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //Import sqllite gorm
)

// Postgres settings
const (
	masterhost = "postgres0"
	masterport = 5432
	host       = "haproxy"
	port       = 5000
	user       = "postgres"
	password   = "password"
	dbname     = "postgres"
)

func initDocsDatabase() {

	time.Sleep(100 * time.Second)

	var err error
	docs.DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable binary_parameters=yes", host, port, user, password, dbname))
	if err != nil {
		fmt.Println("Error: ", err.Error())
		panic("Could not connect to DB")
	}

	docs.MasterDB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable binary_parameters=yes", masterhost, masterport, user, password, dbname))
	if err != nil {
		fmt.Println("Error: ", err.Error())
		panic("Could not connect to DB")
	}

	//Migrate the json schema for the docs
	docs.MasterDB.AutoMigrate(&docs.Docs{})

	fmt.Println("Database migrated")
	fmt.Println("Connected to database")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/docs", docs.GetDocs)
	app.Get("/api/v1/doc/:id", docs.GetDoc)
	app.Post("/api/v1/doc", docs.PostDoc)
	app.Delete("/api/v1/doc/:id", docs.DeleteDoc)

	app.Get("/api/v1/stest", docs.TestDoc)
	app.Get("/api/v1/server", docs.TestServer)
	app.Get("/api/v1/json", docs.TestJSON)
}

func main() {
	app := fiber.New()
	port := os.Getenv("WEB_PORT")
	fmt.Println("SRV PORT:", port)

	initDocsDatabase()
	defer docs.MasterDB.Close()
	defer docs.DB.Close()

	setupRoutes(app)

	log.Fatal(app.Listen(port))
}
