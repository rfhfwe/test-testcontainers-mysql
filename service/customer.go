package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"test_container/logic"
	"test_container/models"
)

type CustomerService struct {
	customerLogic *logic.CustomerLogic
}

func NewCustomerService(customerLogic *logic.CustomerLogic) *CustomerService {
	return &CustomerService{
		customerLogic: customerLogic,
	}
}

func (c *CustomerService) Create(ctx *gin.Context) error {
	return c.customerLogic.Create(ctx, &models.Customer{
		Name:    "",
		Age:     22,
		Address: "abc",
		Email:   "2653@qq.com",
	})
}

func (c *CustomerService) FindOne(ctx *gin.Context) error {
	one, err := c.customerLogic.FindOne(ctx, 1)
	if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": one,
	})
	return nil
}
