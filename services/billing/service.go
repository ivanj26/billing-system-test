package billing

import (
	"amartha-billing-app/repositories/loan"
	"amartha-billing-app/repositories/payment"
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
)

var instance BillingService

func NewService() BillingService {
	if instance != nil {
		return instance
	}

	return &service{
		loanRepo:    loan.NewRepository(),
		paymentRepo: payment.NewRepository(),
	}
}

func (s *service) GetOutstandingByBorrowerId(ctx context.Context, borrowerId uint64) (float32, error) {
	loan, err := s.loanRepo.FindLatestInprogressLoan(ctx, borrowerId)
	if err != nil {
		return 0, fmt.Errorf("[BillingService] Failed to get outstanding from borrowerId=%d. Err=%s", borrowerId, err.Error())
	}

	outstanding, err := s.loanRepo.GetOutstandingByLoanId(ctx, loan.ID)
	if err != nil {
		return 0, fmt.Errorf("[BillingService] Failed to get outstanding from loanId=%d. Err=%s", loan.ID, err.Error())
	}

	return outstanding, nil
}

func (s *service) IsDelinquent(ctx context.Context, borrowerId uint64) (bool, error) {
	loan, err := s.loanRepo.FindLatestInprogressLoan(ctx, borrowerId)
	if err != nil {
		return false, fmt.Errorf("[BillingService] Failed to get delinquent status of the borrowerId=%d. Err=%s", borrowerId, err.Error())
	}

	isDelinquent, err := s.loanRepo.IsDelinquent(ctx, loan.ID)
	if err != nil {
		return false, fmt.Errorf("[BillingService] Failed to get delinquent status of the loanId=%d. Err=%s", loan.ID, err.Error())
	}

	return isDelinquent, nil
}

func (s *service) MakePayment(ctx context.Context, loanId uint64, paidAmount float32) (int, error) {
	list, err := s.loanRepo.GetAllPendingLoanPayment(ctx, loanId)
	if err != nil {
		return 0, fmt.Errorf("[BillingService] Failed to make payment for loanId=%d. Err=%s", loanId, err.Error())
	}

	if len(list) == 0 {
		return 0, gorm.ErrRecordNotFound
	}

	loanAmount := list[0].Amount
	if uint(paidAmount)%uint(loanAmount) != 0 {
		return 0, fmt.Errorf("[BillingService] Failed to make payment for loanId=%d. Err=%s", loanId, fmt.Sprintf("Payments must be made in multiples of %.2f!", loanAmount))
	}

	nbOfPayments := int(paidAmount / loanAmount)
	if nbOfPayments > len(list) {
		nbOfPayments = len(list)
	}

	pendingPayments := list[0:nbOfPayments]
	rowAffected, err := s.paymentRepo.MarkAsPaid(ctx, pendingPayments)
	if err != nil {
		return 0, fmt.Errorf("[BillingService] Failed to make payment for loanId=%d. Err=%s", loanId, err.Error())
	}

	return rowAffected, nil
}
