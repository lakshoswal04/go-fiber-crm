package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"

	"github.com/lakshoswal04/go-fiber-crm/database"
	"github.com/lakshoswal04/go-fiber-crm/lead"
)

// Setup all routes
func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.CreateLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

// Initialize database connection
func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Dialector{
		DriverName: "sqlite",
		DSN:        "leads.db",
	}, &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	initDatabase()
	setupRoutes(app)

	app.Listen(":3000")
}
