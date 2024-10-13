package model

import (
	"amartha-billing-app/common/models"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type PaymentStatus string

const (
	PAYMENT_STATUS_PENDING PaymentStatus = "pending"
	PAYMENT_STATUS_PAID    PaymentStatus = "paid"
)

type Payment struct {
	models.Base

	LoanScheduleId uint64        `json:"loanScheduleId"`
	Status         PaymentStatus `gorm:"size:255" json:"status"`
	Amount         float32       `gorm:"default:0" json:"amount"`
	PaidDate       *time.Time    `json:"paidDate"`
}

func (p Payment) AfterSave(tx *gorm.DB) (err error) {
	err = tx.
		Exec(queryUpdateLoanStatusToCompletedHook, p.LoanScheduleId).
		Error

	if err != nil {
		log.Println("[Payment - AfterSave] Error updating the loan status to completed! Err=", err.Error())
	}

	return
}
