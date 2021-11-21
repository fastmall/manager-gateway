package controller

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/fastmall/manager-gateway/dubbo"
	"github.com/fastmall/order/api"
	"github.com/gin-gonic/gin"
	"strconv"
)

var orderService = dubbo.OrderService

func GetOrderDetail(c *gin.Context) {
	orderIdStr := c.Query("orderId")
	orderId, err := strconv.ParseInt(orderIdStr, 10, 64)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	orderDetail, err := orderService.GetOrderDetail(c, &api.GetOrderDetailRequest{OrderId: orderId})
	if err != nil {
		logger.Error(err)
		c.String(500, err.Error())
		return
	}
	c.JSON(200, orderDetail)
}
