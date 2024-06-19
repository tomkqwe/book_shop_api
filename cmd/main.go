package main

import (
	"book_shop_api/config"
	"book_shop_api/internal/database"
	"book_shop_api/internal/repository/pgrepo"
	"book_shop_api/internal/service"
	"book_shop_api/internal/service/middleware"
	"book_shop_api/internal/transport/http"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Start application...")

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		fmt.Println("Переменная окружения CONFIG_PATH не установлена")
		return
	}

	config, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("error with load config: %v", err)
	}
	err = config.Validate()
	if err != nil {
		log.Fatalf("error with validate config: %v", err)
	}

	pool, err := database.DatabaseConnection(config)
	if err != nil {
		log.Fatal("Can't connect to database... ", err)
	}
	defer pool.Close()

	db, err := sql.Open("postgres", pool.Config().ConnString())
	if err != nil {
		log.Fatalf("Could not open database connection: %v", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not create postgres driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		//"file://../migrations", // Локальный путь к папке с миграциями
		"file:///app/migrations", // путь к папке с миграциями в docker
		"postgres", driver,
	)
	if err != nil {
		log.Fatalf("Could not create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Could not apply migrations: %v", err)
	}

	// Создание обработчиков и маршрутизации
	r := gin.Default()
	bookRepo := pgrepo.NewBookRepo(pool)
	userRepo := pgrepo.NewUserRepo(pool)
	cartRepo := pgrepo.NewCartRepo(pool)
	cartItemRepo := pgrepo.NewCartItemRepo(pool)
	orderRepo := pgrepo.NewOrderRepo(pool)
	orderItemRepo := pgrepo.NewOrderItemRepo(pool)

	bookSrv := service.NewBookService(bookRepo)
	userSrv := service.NewUserService(userRepo)
	cartSrv := service.NewCartService(cartRepo)
	cartItemSrv := service.NewCartItemService(cartItemRepo)
	orderSrv := service.NewOrderService(orderRepo)
	orderItemSrv := service.NewOrderItemService(orderItemRepo)

	bookHandler := http.NewBookHandler(bookSrv)
	userHandler := http.NewUserHandler(userSrv)
	cartHandler := http.NewCartHandler(cartSrv)
	cartItemHandler := http.NewCartItemHandler(cartItemSrv)
	orderHandler := http.NewOrderHandler(orderSrv)
	orderItemHandler := http.NewOrderItemHandler(orderItemSrv)
	registerHandler := http.NewRegisterHandler(userSrv)

	// Настройка маршрутов
	r.GET("/books/:id", bookHandler.GetBook)
	r.GET("/books", bookHandler.GetBooks)
	r.POST("/books", bookHandler.CreateBook)
	r.PUT("/books/:id", bookHandler.UpdateBook)
	r.DELETE("/books/:id", bookHandler.DeleteBook)

	r.GET("/users/:id", userHandler.GetUserById)
	r.GET("/users", userHandler.GetUsers)
	r.POST("/users", userHandler.PostUser)
	r.PUT("/users/:id", userHandler.PutUser)
	r.DELETE("/users/:id", userHandler.DeleteUserById)

	r.GET("/carts/:id", cartHandler.GetCart)
	r.GET("/carts", cartHandler.GetCarts)
	r.PUT("/carts/:id", cartHandler.UpdateCart)
	r.POST("/carts", cartHandler.CreateCart)
	r.DELETE("/carts/:id", cartHandler.DeleteCart)

	r.GET("/cartItems/:id", cartItemHandler.GetCartItemById)
	r.GET("/cartItems", cartItemHandler.GetCartItems)
	r.PUT("/cartItems/:id", cartItemHandler.PutCartItemById)
	r.POST("/cartItems", cartItemHandler.PostCartItem)
	r.DELETE("/cartItems/:id", cartItemHandler.DeleteCartItemById)

	r.GET("/orders/:id", orderHandler.GetOrderById)
	r.GET("/orders", orderHandler.GetOrders)
	r.PUT("/orders/:id", orderHandler.PutOrderById)
	r.POST("/orders", orderHandler.PostOrder)
	r.DELETE("/orders/:id", orderHandler.DeleteOrderById)

	r.GET("/orderItems/:id", orderItemHandler.GetOrderItem)
	r.GET("/orderItems", orderItemHandler.GetOrderItems)
	r.PUT("/orderItems/:id", orderItemHandler.UpdateOrderItem)
	r.POST("/orderItems", orderItemHandler.CreateOrderItem)
	r.DELETE("/orderItems/:id", orderItemHandler.DeleteOrderItem)

	r.POST("/register", registerHandler.Register)
	r.POST("/login", registerHandler.Login)

	authorized := r.Group("/say")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/hello", func(ctx *gin.Context) {
			Id := ctx.MustGet("id").(uuid.UUID)
			email := ctx.MustGet("email").(string)
			ctx.JSON(200, gin.H{
				"message": fmt.Sprintf("Hello, %s !", email),
				"id":      Id,
			})
		})
	}

	// Запуск сервера
	r.Run(fmt.Sprintf(":%d", config.HttpServer.Port))
}
