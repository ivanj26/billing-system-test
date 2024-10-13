package payment

import (
	"amartha-billing-app/database"
	model "amartha-billing-app/models"
	"context"
	"fmt"
	"time"
)

var instance *repo

func NewRepository() Repository {
	return &repo{
		db: database.GetInstance(),
	}
}

func (r *repo) MarkAsPaid(ctx context.Context, loanSchedules interface{}) (int, error) {
	tx := r.db.Begin()

	if tx.Error != nil {
		return 0, tx.Error
	}

	now := time.Now()
	counter := 0

	switch v := loanSchedules.(type) {
	case []model.LoanSchedule:
		{
			for _, schedule := range v {
				payment := schedule.Payment

				if payment != nil {
					payment.PaidDate = &now
					payment.Status = model.PAYMENT_STATUS_PAID
					payment.Amount = schedule.Amount
				} else {
					payment = &model.Payment{
						LoanScheduleId: schedule.ID,
						Amount:         schedule.Amount,
						Status:         model.PAYMENT_STATUS_PAID,
						PaidDate:       &now,
					}
				}

				err := tx.Save(&payment).Error
				if err != nil {
					tx.Rollback()
					return 0, fmt.Errorf("[PaymentRepo] Failed to mark the payment status as paid. Err=%v", err)
				}

				counter++
			}
		}
	case model.LoanSchedule:
		{
			payment := v.Payment
			payment.PaidDate = &now
			payment.Status = model.PAYMENT_STATUS_PAID
			payment.Amount = v.Amount

			err := tx.Save(&payment).Error
			if err != nil {
				tx.Rollback()
				return 0, fmt.Errorf("[PaymentRepo] Failed to mark the payment status as paid. Err=%v", err)
			}

			counter++
		}
	case model.Payment:
		{
			v.PaidDate = &now
			v.Status = model.PAYMENT_STATUS_PAID

			err := tx.Save(&v).Error
			if err != nil {
				tx.Rollback()
				return 0, fmt.Errorf("[PaymentRepo] Failed to mark the payment status as paid. Err=%v", err)
			}

			counter++
		}
	default:
		{
			// unsupported
			// rollback!

			tx.Rollback()
			return 0, fmt.Errorf("[PaymentRepo] Failed to mark the payment status as paid. Err=%s", "unsupported data type")
		}
	}

	err := tx.Commit().Error
	if err != nil {
		return 0, fmt.Errorf("[PaymentRepo] Failed to mark the payment status as paid. Err=%v", err)
	}

	return counter, nil
}
