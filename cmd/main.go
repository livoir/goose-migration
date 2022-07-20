package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	db2 "github.com/livoir/goose-migration/db/sqlc"
	"github.com/livoir/goose-migration/internal/core/services"
	"github.com/livoir/goose-migration/internal/handlers"
	"github.com/livoir/goose-migration/internal/repositories"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goose_migration")
	if err != nil {
		log.Println(err.Error())
		return
	}
	queries := db2.New(db)

	userRepository := repositories.NewUserRepositoryImpl(queries)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewAPIHandler(userService)

	router := gin.Default()
	router.GET("/users", userHandler.GetAll)
	router.GET("/users/:id", userHandler.GetById)

	router.Run("localhost:8080")
}
