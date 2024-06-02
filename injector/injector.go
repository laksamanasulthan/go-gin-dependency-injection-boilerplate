package injector

import (
	"fmt"
	product_injector "go-gin-dependency-injection-boilerplate/domain/product/injector"
	product_models "go-gin-dependency-injection-boilerplate/domain/product/models"
	user_models "go-gin-dependency-injection-boilerplate/domain/user/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Injector struct {
	ProductInjector *product_injector.ProductInjector
	//Other domain Injector
}

func NewInjector() (*Injector, error) {

	godotenv.Load()

	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_database := os.Getenv("DB_DATABASE")

	// dsn := "root:@tcp(127.0.0.1:3306)/go_boilerplate?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db_username, db_password, db_host, db_port, db_database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	//Auto-Migration
	err = db.AutoMigrate(
		&product_models.Product{},
		&user_models.User{},
	)

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
