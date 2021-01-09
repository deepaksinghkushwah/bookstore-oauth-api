package db

import (
	"bookstore/bookstore-oauth-api/src/domain/access_token"
	"bookstore/bookstore-oauth-api/src/utils/errors"
)

type DBRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func NewRepository() DBRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetByID(string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("Database connection must be implemented")
}
