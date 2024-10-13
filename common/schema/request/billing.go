package request

type GetOutstandingRequest struct {
	BorrowerId uint64 `query:"borrower_id" validate:"required,gte=0"`
}

type IsDelinquentRequest struct {
	BorrowerId uint64 `json:"borrower_id" validate:"required,gte=0"`
}

type MakePaymentRequest struct {
	LoanId     uint64  `json:"loan_id" validate:"required,gte=0"`
	PaidAmount float32 `json:"paid_amount" validate:"required,gte=0"` // paid amount by customer
}
