package payment

import (
	"context"

	"github.com/jinzhu/gorm"
)

type repo struct {
	db *gorm.DB
}

type Repository interface {
	MarkAsPaid(ctx context.Context, payments interface{}) (int, error)
}
