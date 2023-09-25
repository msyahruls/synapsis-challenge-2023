package main

import (
	"fmt"
	"os"
	"synapsis-challange/config"
	"synapsis-challange/controller"
	"synapsis-challange/helper"
	"synapsis-challange/repository"
	"synapsis-challange/service"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	time.Local = time.UTC
	db := config.NewDB()
	validate := validator.New()

	logRepository := repository.NewLogRepository(db)
	logService := service.NewLogService(logRepository)
	logController := controller.NewLogController(logService)

	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository, validate)
	authController := controller.NewAuthController(authService, logService)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository, validate)
	categoryController := controller.NewCategoryController(categoryService, logService)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository, validate)
	productController := controller.NewProductController(productService, logService)

	cartRepository := repository.NewCartRepository(db)
	cartService := service.NewCartService(cartRepository, validate)
	cartController := controller.NewCartController(cartService, logService)

	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepository, validate)
	transactionController := controller.NewTransactionController(transactionService, cartService, logService)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "*",
		AllowCredentials: true,
	}))

	authController.NewAuthRouter(app)
	logController.NewLogRouter(app)
	categoryController.NewCategoryRouter(app)
	productController.NewProductRouter(app)
	cartController.NewCartRouter(app)
	transactionController.NewTransactionRouter(app)

	host := fmt.Sprintf("%s:%s", os.Getenv("SERVER_URI"), os.Getenv("SERVER_PORT"))
	err := app.Listen(host)
	helper.PanicIfError(err)

}
