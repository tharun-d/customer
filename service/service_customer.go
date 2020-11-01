package service

import (
	"context"
	"customer/models"
)


func (c *customerService) GetCustomerData(_ context.Context, customerDataReq models.GetCustomerReq) (models.Customer, error) {
	cus, err := c.customerRepos.GetCustomerData(customerDataReq)
	if err != nil {
		return models.Customer{}, err
	}

	return cus, nil
}
