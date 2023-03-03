/*
 * Author : Ismail Ash Shidiq (https://www.eulbyvan.com)
 * Created on : Fri Mar 03 2023 9:41:07 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/eulbyvan/go-user-management/controller"
	"github.com/eulbyvan/go-user-management/repository"
	"github.com/eulbyvan/go-user-management/usecase"
	"github.com/eulbyvan/go-user-management/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Config
	dbHost 		:= utils.DotEnv("DB_HOST")
	dbPort 		:= utils.DotEnv("DB_PORT")
	dbUser 		:= utils.DotEnv("DB_USER")
	dbPassword 	:= utils.DotEnv("DB_PASSWORD")
	dbName 		:= utils.DotEnv("DB_NAME")
	sslMode		:= utils.DotEnv("SSL_MODE")
	serverPort 	:= utils.DotEnv("SERVER_PORT")

	// DB Connection
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, sslMode))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize Gin router
	router := gin.Default()

	// Group for user routes
	userRoutes := router.Group("/users")

	// Initialize dependencies
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userController := controller.NewUserController(userUsecase)

	// Routes for user group
	userRoutes.GET("", userController.GetAll)
	userRoutes.GET("/:id", userController.GetOne)
	userRoutes.POST("", userController.Add)
	userRoutes.PUT("", userController.Edit)
	userRoutes.DELETE("/:id", userController.Remove)

	// Run the server
	if err := router.Run(serverPort); err != nil {
		log.Fatal(err)
	}
}