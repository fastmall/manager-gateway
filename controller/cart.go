package controller

import (
	"github.com/fastmall/cart/api"
	"github.com/fastmall/manager-gateway/dubbo"
	"github.com/gin-gonic/gin"
	"strconv"
)

var cartService = dubbo.CartService

func GetCartDetail(c *gin.Context) {
	customerIdStr := c.Query("customerId")
	customerId, err := strconv.ParseInt(customerIdStr, 10, 64)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	cart, err := cartService.GetCustomerCart(c, &api.GetCustomerCartRequest{CustomerId: customerId})
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.JSON(200, cart)
}
