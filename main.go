// main.go
package main

import (
    "fmt"
	"goapp/api"
	"goapp/db"
	"goapp/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
    "github.com/tealeg/xlsx"
)

var DB *gorm.DB
func GenerateLeaveReport(r *fiber.Ctx) error {
    // Retrieve leave data from the database
    var leaves []models.Leave
    db.DB.Find(&leaves)

    // Create a new Excel file
    file := xlsx.NewFile()
    sheet, err := file.AddSheet("Leave Report")
    if err != nil {
        return err
    }

    // Add headers to the Excel file
    headers := sheet.AddRow()
    headers.AddCell().Value = "ID"
    headers.AddCell().Value = "User ID"
    headers.AddCell().Value = "Start Date"
    headers.AddCell().Value = "End Date"
    headers.AddCell().Value = "Status"
    headers.AddCell().Value = "Description"

    // Add leave data to the Excel file
    for _, leave := range leaves {
        row := sheet.AddRow()
        row.AddCell().Value = fmt.Sprintf("%d", leave.ID)
        row.AddCell().Value = fmt.Sprintf("%d", leave.UserID)
        row.AddCell().Value = leave.StartDate.Format("2022-07-23")
        row.AddCell().Value = leave.EndDate.Format("2022-08-23")
        row.AddCell().Value = leave.Status
        row.AddCell().Value = leave.Description
    }

    // Save the Excel file
    err = file.Save("leave_report.xlsx")
    if err != nil {
        return err
    }

    return r.SendString("Leave report generated successfully")
}

func main() {
    app := fiber.New()

    // Initialize database
    Db:= db.InitDB()
	Db.AutoMigrate(&models.Leave{})


    // Leave API routes
    app.Post("/api/apply-leave", api.ApplyLeave)
	app.Get("/auth", api.AuthenticateUser)
    app.Get("/auth/callback", api.HandleOAuthCallback)
    app.Get("/main",GenerateLeaveReport)

    // Start the server
    app.Listen(":8080")
}
