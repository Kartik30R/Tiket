package main

import (
	"log"

	"github.com/Kartik30R/Tiket.git/config"
	"github.com/Kartik30R/Tiket.git/db"
	"github.com/Kartik30R/Tiket.git/handlers"
	"github.com/Kartik30R/Tiket.git/middleware"
	"github.com/Kartik30R/Tiket.git/repositories"
	"github.com/Kartik30R/Tiket.git/services"
	"github.com/gin-gonic/gin"
)

func main() {
   cfg, err := config.NewEnvConfig()

   db:= db.Init(cfg,db.DBMigrator)
   if err!=nil{
      	log.Fatal("config cant be loaded", err)

   }

router := gin.Default()
   authRepo := repositories.NewAuthRepository(db)
   eventRepo := repositories.NewEventRepository(db)
   ticketRepo := repositories.NewTicketRepository(db)

   authService := services.NewAuthService(authRepo)

   server := router.Group("/api")
      handlers.NewAuthHandler(server.Group("/auth"),authService)
server.Use(middleware.AuthProtected(db)) 
   
   handlers.NewEventHandler(server.Group("/event"),eventRepo)
      handlers.NewTicketHandler(server.Group("/ticket"),ticketRepo)


router.Run(":" + cfg.ServerPort)
}