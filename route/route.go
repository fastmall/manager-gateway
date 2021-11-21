package route

import (
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/fastmall/manager-gateway/controller"
	"github.com/gin-gonic/gin"
)

func SetCustomerRoute(r *gin.Engine) {
	g := r.Group("/customer")
	g.POST("/register", controller.Register)
	g.GET("/register", controller.Register)
}

func SetGoodsRoute(r *gin.Engine) {
	g := r.Group("/goods")
	g.GET("/detail", controller.GetGoodsDetail)
}

func SetOrderRoute(r *gin.Engine) {
	g := r.Group("/order")
	g.GET("/detail", controller.GetOrderDetail)
}

func SetCartRoute(r *gin.Engine) {
	g := r.Group("/cart")
	g.GET("/detail", controller.GetCartDetail)
}
