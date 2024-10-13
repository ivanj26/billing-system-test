package model

import (
	"amartha-billing-app/common/models"
	"time"
)

type LoanSchedule struct {
	models.Base

	LoanId  uint64    `json:"loanId"`
	Payment *Payment  `json:"payment"`
	Amount  float32   `gorm:"default:0" json:"amount"`
	DueDate time.Time `json:"dueDate"`
}
