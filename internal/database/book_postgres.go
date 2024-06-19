package database

import (
	"book_shop_api/config"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DatabaseConnection(config *config.Config) (pool *pgxpool.Pool, err error) {
	host := config.Database.Host
	port := config.Database.Port
	dbName := config.Database.Name
	dbUser := config.Database.DbUser
	password := config.Database.Password

	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password)

	//это необходимо для того, что бы контейнер с postgres
	//запустился раньше и наш контейнер с приложением
	//спокойно подключится к базе
	time.Sleep(10 * time.Second)
	pool, err = pgxpool.New(context.Background(), dsn)

	return
}
