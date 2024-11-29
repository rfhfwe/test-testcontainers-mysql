package logic

import (
	"context"

	"test_container/models"
	"test_container/repo"
)

type CustomerLogic struct {
	customerRepo *repo.CustomerRepo
}

func NewCustomerLogic(customerRepo *repo.CustomerRepo) *CustomerLogic {
	return &CustomerLogic{
		customerRepo: customerRepo,
	}
}

func (c *CustomerLogic) FindOne(ctx context.Context, id int64) (*models.Customer, error) {
	customer, err := c.customerRepo.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *CustomerLogic) Create(ctx context.Context, customer *models.Customer) error {
	return c.customerRepo.Create(ctx, customer)
}

func (c *CustomerLogic) Update(ctx context.Context, customer *models.Customer) error {
	return c.customerRepo.Update(ctx, customer)
}

func (c *CustomerLogic) Delete(ctx context.Context, id int64) error {
	return c.customerRepo.Delete(ctx, id)
}
