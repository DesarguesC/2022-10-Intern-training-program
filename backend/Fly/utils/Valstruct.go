package utils

import (
	"github.com/go-playground/validator/v10"
	"sync"
)

type CustomValidator struct {
	once     sync.Once
	validate *validator.Validate
}

type PermitPerson struct {
	ID     int    `json:"ID" validate:"required,min=0,max=10"`
	PassWd string `json:"PassWd" validate:"required,excludesall=!@#$%^&*()_-{}"`
}
