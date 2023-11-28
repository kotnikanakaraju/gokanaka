package models
 
import (
	"gorm.io/gorm"
	"time"
)
 
type Leave struct {
    gorm.Model
    UserID      uint
    StartDate   time.Time
    EndDate     time.Time
    Status      string
    Description string
}
 