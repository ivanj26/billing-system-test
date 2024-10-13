package loan

import (
	"amartha-billing-app/database"
	model "amartha-billing-app/models"
	"context"
	"database/sql"
	"errors"

	"github.com/jinzhu/gorm"
)

var instance *repo

func NewRepository() Repository {
	return &repo{
		db: database.GetInstance(),
	}
}

func (r *repo) FindLatestInprogressLoan(ctx context.Context, borrowerId uint64) (model.Loan, error) {
	var loan model.Loan

	err := r.db.
		Where("borrower_id = ?", borrowerId).
		Where("status = ?", model.LOAN_STATUS_INPROGRESS).
		Last(&loan).
		Error

	if err != nil {
		return loan, errors.New("[LoanRepo] Unable to get latest inprogress loan. Err=" + err.Error())
	}

	return loan, nil
}

func (r *repo) GetOutstandingByLoanId(ctx context.Context, loanId uint64) (float32, error) {
	var result struct {
		Amount float32 `gorm:"column:outstanding_amount"`
	}

	err := r.db.
		Raw(QUERY_GET_OUTSTANDING_BY_LOAN_ID, 1).
		Scan(&result).
		Error

	if err != nil {
		return 0, errors.New("[LoanRepo] Unable to calculate the outstanding. Err=" + err.Error())
	}

	return result.Amount, nil
}

// GetAllPendingLoanPayment - Get all pending payments from the current time
func (r *repo) GetAllPendingLoanPayment(ctx context.Context, loanId uint64) ([]model.LoanSchedule, error) {
	var result []model.LoanSchedule

	err := r.db.
		Model(&model.LoanSchedule{}).
		Preload("Payment").
		Joins("LEFT JOIN payments AS p ON p.loan_schedule_id = loan_schedules.id").
		Where("loan_schedules.loan_id = ?", loanId).
		Where("loan_schedules.due_date <= CURRENT_TIMESTAMP").
		Where("(p.status = 'pending' OR p.id IS NULL)").
		Order("loan_schedules.due_date ASC").
		Find(&result).
		Error

	if err != nil {
		return result, errors.New("[LoanRepo] Unable to find pending loan payments. Err=" + err.Error())
	}

	return result, nil
}

func (r *repo) IsDelinquent(ctx context.Context, loanId uint64) (bool, error) {
	var isDelinquent struct {
		Value sql.NullBool `gorm:"column:1"`
	}

	err := r.db.
		Raw(QUERY_IS_DELINQUENT_BY_LOAN_ID, loanId).
		Scan(&isDelinquent).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if err != nil {
		return false, errors.New("[LoanRepo] Unable to check delinquent status. Err=" + err.Error())
	}

	if isDelinquent.Value.Valid {
		return isDelinquent.Value.Valid, nil
	} else {
		return false, nil
	}
}
