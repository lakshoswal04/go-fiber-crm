package lead

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/lakshoswal04/go-fiber-crm/database"
)


// Lead model
type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	return c.JSON(leads)
}


func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.First(&lead, id)
	return c.JSON(lead)
}


func CreateLead(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	db.Create(&lead)
	return c.JSON(lead)
}


func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		return c.Status(404).SendString("No lead found with given ID")
	}

	db.Delete(&lead)
	return c.SendString("Lead successfully deleted")
}
