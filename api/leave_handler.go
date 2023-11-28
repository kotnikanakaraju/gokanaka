// api/leave_handler.go
package api

import (
	"fmt"
    "net/http"
    "time"
    "goapp/db"
    "github.com/gofiber/fiber/v2"
    "goapp/models"
    "goapp/reporting"
	"gopkg.in/gomail.v2"
)

func ApplyLeave(r *fiber.Ctx) error {
    // Parse request
    var input struct {
        UserID      uint   `json:"user_id"`
        StartDate   string `json:"start_date"`
        EndDate     string `json:"end_date"`
        Description string `json:"description"`
    }

    if err := r.BodyParser(&input); err != nil {
        return r.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
    }

    // Validate input
    if input.UserID == 0 || input.StartDate == "" || input.EndDate == "" {
        return r.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    // Convert date strings to time.Time
    startDate, err := time.Parse("2006-01-02", input.StartDate)
    if err != nil {
        return r.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid start date"})
    }

    endDate, err := time.Parse("2006-01-02", input.EndDate)
    if err != nil {
        return r.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid end date"})
    }

    // Apply leave
    leave := models.Leave{
        UserID:      input.UserID,
        StartDate:   startDate,
        EndDate:     endDate,
        Description: input.Description,
        Status:      "Pending", // You may set an initial status
    }
    if err := db.DB.Create(&leave).Error; err != nil {
        return r.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to apply leave"})
    }


    // Notify manager
    notifyManager(leave.ID)
    leaves := []models.Leave{leave}
    reporting.GenerateLeaveReport(leaves)

    return r.Status(http.StatusOK).JSON(fiber.Map{"message": "Leave applied successfully"})
}


func notifyManager(leaveID uint) {
    // Retrieve manager information based on your application logic
    managerEmail := "manager@example.com"

    
    subject := "Leave Application"
    message := fmt.Sprintf("Leave request ID %d requires your approval.", leaveID)

    
    sendNotification(managerEmail, subject, message)
}

func sendNotification(to, subject, message string) {
    m := gomail.NewMessage()
    m.SetHeader("From", "kanakarajukotni@gmail.com") // Replace with your email
    m.SetHeader("kanakarajukotni@outlook.com", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text/plain", message)

    d := gomail.NewDialer("smtp.example.com", 587, "kanakarajukotni@gmail.com", "kanaka")

    if err := d.DialAndSend(m); err != nil {
        fmt.Println("Error sending email:", err)
    }

    fmt.Println("Email sent successfully")
    fmt.Printf("Sending notification to %s\nSubject: %s\nMessage: %s\n", to, subject, message)
}