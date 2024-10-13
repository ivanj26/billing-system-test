package model

import "amartha-billing-app/common/models"

type Borrower struct {
	models.Base

	Loan Loan   `json:"loan"`
	Name string `gorm:"size:255" json:"name"`
}
