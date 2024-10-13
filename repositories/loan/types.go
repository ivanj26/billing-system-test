package loan

import (
	model "amartha-billing-app/models"
	"context"

	"github.com/jinzhu/gorm"
)

type repo struct {
	db *gorm.DB
}

type Repository interface {
	FindLatestInprogressLoan(ctx context.Context, borrowerId uint64) (model.Loan, error)
	GetOutstandingByLoanId(ctx context.Context, loanId uint64) (float32, error)
	GetAllPendingLoanPayment(ctx context.Context, loanId uint64) ([]model.LoanSchedule, error)
	IsDelinquent(ctx context.Context, loanId uint64) (bool, error)
}
