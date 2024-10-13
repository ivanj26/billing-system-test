package billing

import (
	"amartha-billing-app/repositories/loan"
	"amartha-billing-app/repositories/payment"
	"context"
)

type service struct {
	loanRepo    loan.Repository
	paymentRepo payment.Repository
}

type BillingService interface {
	GetOutstandingByBorrowerId(ctx context.Context, borrowerId uint64) (float32, error)
	MakePayment(ctx context.Context, loanId uint64, paidAmount float32) (int, error)
	IsDelinquent(ctx context.Context, borrowerId uint64) (bool, error)
}
