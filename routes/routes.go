package routes

import (
	product_route "go-gin-dependency-injection-boilerplate/domain/product/route"
	"go-gin-dependency-injection-boilerplate/injector"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, injector *injector.Injector) {
	//``
	product_route.RegisterProductRoutes(router, injector.ProductInjector.Controllers)

}
