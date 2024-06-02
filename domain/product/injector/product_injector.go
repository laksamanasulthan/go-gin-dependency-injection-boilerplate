package injector

import (
	"go-gin-dependency-injection-boilerplate/domain/product/controllers"
	"go-gin-dependency-injection-boilerplate/domain/product/repositories"
	"go-gin-dependency-injection-boilerplate/domain/product/services"

	"gorm.io/gorm"
)

type ProductControllers struct {
	ProductController *controllers.ProductController
	//Other Controlller
}

type ProductInjector struct {
	Controllers ProductControllers
}

func NewProductInjector(db *gorm.DB) *ProductInjector {

	productRepo := repositories.NewProductRepositroy(db)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	//Other DI Repo
	//Other DI Service
	//Other DI Controller

	controllers := ProductControllers{
		ProductController: productController,
		//Other Controller
	}

	return &ProductInjector{
		Controllers: controllers,
	}
}
