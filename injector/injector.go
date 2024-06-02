package injector

import (
	product_injector "go-gin-dependency-injection-boilerplate/domain/product/injector"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Injector struct {
	ProductInjector *product_injector.ProductInjector
	//Other domain Injector
}

func NewInjector() (*Injector, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/go_boilerplate?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	productInjector := product_injector.NewProductInjector(db)
	//Other Domain Injector

	return &Injector{
		ProductInjector: productInjector,
		//Other Domain Injector
	}, nil
}
