package billing

import (
	"amartha-billing-app/common"
	billingService "amartha-billing-app/services/billing"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	billingService billingService.BillingService
}

type IHandler interface {
	common.BaseHandler

	GetOutstanding(ctx echo.Context) error
	IsDelinquent(ctx echo.Context) error
	MakePayment(ctx echo.Context) error
}

func New() IHandler {
	return &handler{
		billingService: billingService.NewService(),
	}
}

func (h *handler) Routes() []common.Route {
	return []common.Route{
		{
			Method:  http.MethodGet,
			Path:    "outstanding",
			Handler: h.GetOutstanding,
		},

		{
			Method:  http.MethodPost,
			Path:    "is-delinquent",
			Handler: h.IsDelinquent,
		},

		{
			Method:  http.MethodPost,
			Path:    "payment",
			Handler: h.MakePayment,
		},
	}
}
