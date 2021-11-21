package dubbo

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	_ "github.com/dubbogo/triple/pkg/triple"
	cartApi "github.com/fastmall/cart/api"
	customerApi "github.com/fastmall/customer/api"
	goodsApi "github.com/fastmall/goods/api"
	orderApi "github.com/fastmall/order/api"
)

var CustomerService = &customerApi.CustomerServiceClientImpl{}
var GoodsService = &goodsApi.GoodsServiceClientImpl{}
var CartService = &cartApi.CartServiceClientImpl{}
var OrderService = &orderApi.OrderServiceClientImpl{}

func StartDubboConsumer() {
	config.SetConsumerService(CustomerService)
	config.SetConsumerService(CartService)
	config.SetConsumerService(OrderService)
	config.SetConsumerService(GoodsService)
	err := config.Load(config.WithPath("conf/dubbo.yaml"))
	if err != nil {
		logger.Fatal(err)
		return
	}
}
