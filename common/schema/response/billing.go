package response

type GetOutstandingDataResponse struct {
	OutstandingAmount float32 `json:"amount"`
}

type IsDelinquentResponse struct {
	IsDelinquent bool `json:"isDelinquent"`
}

type MakePaymentResponse struct {
	SuccessfulPayments int `json:"successfulPayments"`
}
