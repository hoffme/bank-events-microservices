package controllers

import (
	"github.com/hoffme/backend-transactions/internal/application/api/generated"
	"github.com/hoffme/backend-transactions/internal/domain"
)

var _ generated.Handler = Controllers{}

type Controllers struct {
	Deps domain.Dependencies
}
