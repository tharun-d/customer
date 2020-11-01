package repository

import (
	"customer/models"
	"database/sql"
)

type CustomerRepository interface {
	GetCustomerData(req models.GetCustomerReq) (models.Customer, error)
}

type mariaCustomerRepository struct {
	Conn *sql.DB
}

// NewMariaCustomerRepository will create an object that represent the customer.Repository interface
func NewMariaCustomerRepository(Conn *sql.DB) CustomerRepository {

	return &mariaCustomerRepository{Conn}
}
