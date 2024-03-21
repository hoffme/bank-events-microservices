package controllers

import (
	"context"
	"errors"

	appErrors "github.com/hoffme/backend-transactions/internal/shared/errors"
	"github.com/hoffme/backend-transactions/internal/shared/logger"

	"github.com/hoffme/backend-transactions/internal/application/api/generated"
)

func (c Controllers) NewError(ctx context.Context, err error) *generated.ErrorStatusCode {
	logger.Debug("%s", err.Error())

	status := 500
	code := "internal Error"
	description := generated.OptString{}

	var appErr appErrors.Error
	if errors.As(err, &appErr) {
		status = appErr.Status()
		code = appErr.Code()
		description = generated.NewOptString(appErr.Description())
	}

	return &generated.ErrorStatusCode{
		StatusCode: status,
		Response: generated.Error{
			Status:      status,
			Code:        code,
			Description: description,
		},
	}
}
