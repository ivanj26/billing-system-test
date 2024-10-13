package model

import "amartha-billing-app/common/models"

type LoanStatus string

const (
	LOAN_STATUS_INPROGRESS = "inprogress"
	LOAN_STATUS_COMPLETED  = "completed"
)

type Loan struct {
	models.Base

	BorrowerId uint64         `json:"borrowerId"`
	Schedules  []LoanSchedule `json:"schedules"`

	Name   string `gorm:"size:255;not null" json:"name"`
	Status string `gorm:"size:255;not null" json:"status"`
}
