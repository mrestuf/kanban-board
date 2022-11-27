package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mrestuf/kanban-board/config"
	"github.com/mrestuf/kanban-board/httpserver"
	"github.com/mrestuf/kanban-board/httpserver/controllers"
	"github.com/mrestuf/kanban-board/httpserver/services"
	"github.com/mrestuf/kanban-board/httpserver/repositories/gorm"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("cannot load env")
		return
	}
}

func main() {
	db, err := config.ConnectPostgresGORM()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	config.GenerateJwtSignature()

	userRepo := gorm.NewUserRepo(db)
	userSvc := services.NewUserSvc(userRepo)
	userHandler := controllers.NewUserController(userSvc)

	app := httpserver.NewRouter(router, userHandler)
	PORT := os.Getenv("PORT")
	app.Start(":" + PORT)
}
