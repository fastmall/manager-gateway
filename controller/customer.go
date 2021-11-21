package controller

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"fmt"
	"github.com/fastmall/customer/api"
	"github.com/fastmall/manager-gateway/dubbo"
	"github.com/gin-gonic/gin"
)

var customerService = dubbo.CustomerService

func Register(c *gin.Context) {

	createCustomerResponse, err := customerService.CreateCustomer(context.Background(), &api.CreateCustomerRequest{
		Username: "zhurungen",
		Password: "test01",
	})

	if err != nil {
		logger.Error(err)
		return
	}

	msg := fmt.Sprintf("customerId:%d, token:%s", createCustomerResponse.CustomerId, createCustomerResponse.Token)

	logger.Info(msg)
	c.String(200, msg)
}
