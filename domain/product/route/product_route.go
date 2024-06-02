package product_route

import (
	"go-gin-dependency-injection-boilerplate/domain/product/injector"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.Engine, controllers injector.ProductControllers) {

	productRoutes := router.Group("/product")
	{
		productRoutes.GET("", controllers.ProductController.Index)
	}

	// ptherDomainRoutes := router.Group("/other-domain-routes")
	// {
	// 	otherDomainRoutes.GET("", controllers.OtherController.Index)
	// }

}
