package injector

import (
	"go-gin-dependency-injection-boilerplate/domain/product/injector"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Injector struct {
	ProductInjector *injector.ProductInjector
	//Other domain Injector
}

func NewInjector() (*Injector, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_crud?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	productInjector := injector.NewProductInjector(db)
	//Other Domain Injector

	return &Injector{
		ProductInjector: productInjector,
		//Other Domain Injector
	}, nil
}
