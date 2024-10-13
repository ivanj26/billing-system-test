package billing

import (
	"amartha-billing-app/common/schema/request"
	"amartha-billing-app/common/schema/response"
	"context"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func (c *handler) GetOutstanding(ctx echo.Context) error {
	resp := response.NewResponse(ctx)

	// Validate param body
	params := new(request.GetOutstandingRequest)
	if err := ctx.Bind(params); err != nil {
		resp.SetMessage(err.Error())
		resp.SetStatusCode(http.StatusBadRequest)
		resp.Error()

		return echo.NewHTTPError(http.StatusBadRequest, resp)
	}

	if err := ctx.Validate(params); err != nil {
		resp.SetMessage(err.Error())
		resp.SetStatusCode(http.StatusBadRequest)
		resp.Error()

		return echo.NewHTTPError(http.StatusBadRequest, resp)
	}

	outstanding, err := c.billingService.GetOutstandingByBorrowerId(context.Background(), params.BorrowerId)
	if err != nil {
		resp.SetMessage(fmt.Sprintf("Failed to get outstanding loan from the given borrower id: %d!", params.BorrowerId))
		resp.SetStatusCode(http.StatusInternalServerError)
		resp.Error()

		return echo.NewHTTPError(http.StatusInternalServerError, resp)
	}

	resp.SetData(response.GetOutstandingDataResponse{
		OutstandingAmount: outstanding,
	})
	resp.SetStatusCode(http.StatusOK)
	resp.Success()

	return ctx.JSON(http.StatusOK, resp)
}

func (c *handler) IsDelinquent(ctx echo.Context) error {
	resp := response.NewResponse(ctx)

	// Validate param body
	params := new(request.IsDelinquentRequest)
	if err := ctx.Bind(params); err != nil {
		resp.SetMessage(err.Error())
		resp.SetStatusCode(http.StatusBadRequest)
		resp.Error()

		return echo.NewHTTPError(http.StatusBadRequest, resp)
	}

	if err := ctx.Validate(params); err != nil {
		resp.SetMessage(err.Error())
		resp.SetStatusCode(http.StatusBadRequest)
		resp.Error()

		return echo.NewHTTPError(http.StatusBadRequest, resp)
	}

	isDelinquent, err := c.billingService.IsDelinquent(context.Background(), params.BorrowerId)
	if err != nil {
		resp.SetMessage(fmt.Sprintf("Failed to get loan delinquent status from the given borrower id: %d!", params.BorrowerId))
		resp.SetStatusCode(http.StatusInternalServerError)
		resp.Error()

		return echo.NewHTTPError(http.StatusInternalServerError, resp)
	}

	resp.SetData(response.IsDelinquentResponse{
		IsDelinquent: isDelinquent,
	})
	resp.SetStatusCode(http.StatusOK)
	resp.Success()

	return ctx.JSON(http.StatusOK, resp)
}

func (c *handler) MakePayment(ctx echo.Context) error {
	resp := response.NewResponse(ctx)

	// Validate param body
	params := new(request.MakePaymentRequest)
	if err := ctx.Bind(params); err != nil {
		resp.SetMessage(err.Error())
		resp.SetStatusCode(http.StatusBadRequest)
		resp.Error()

		return echo.NewHTTPError(http.StatusBadRequest, resp)
	}

	if err := ctx.Validate(params); err != nil {
		resp.SetMessage(err.Error())
		resp.SetStatusCode(http.StatusBadRequest)
		resp.Error()

		return echo.NewHTTPError(http.StatusBadRequest, resp)
	}

	successfulPayments, err := c.billingService.MakePayment(context.Background(), params.LoanId, params.PaidAmount)
	if err == gorm.ErrRecordNotFound {
		resp.SetMessage(fmt.Sprintf("[BillingService] Failed to make payment for loanId=%d. Err=%s", params.LoanId, "There is no pending payments at the moment!"))
		resp.SetStatusCode(http.StatusBadRequest)
		resp.Error()

		return echo.NewHTTPError(http.StatusBadRequest, resp)
	}

	if err != nil {
		resp.SetMessage(err.Error())
		resp.SetStatusCode(http.StatusInternalServerError)
		resp.Error()

		return echo.NewHTTPError(http.StatusInternalServerError, resp)
	}

	resp.SetData(response.MakePaymentResponse{
		SuccessfulPayments: successfulPayments,
	})
	resp.SetStatusCode(http.StatusOK)
	resp.Success()

	return ctx.JSON(http.StatusOK, resp)
}
