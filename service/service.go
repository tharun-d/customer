package service

import (
	"context"

	"customer/models"

	"customer/repository"
)

type CustomerService interface {
	GetCustomerData(ctx context.Context, customerDataReq models.GetCustomerReq) (models.Customer, error)
}

type customerService struct {
	customerRepos repository.CustomerRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository) CustomerService {
	return &customerService{customerRepository}
}
