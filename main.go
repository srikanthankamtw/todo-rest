package main

import (
	"github.com/gin-gonic/gin"
	"github.com/srikanthankamtw/todo-rest/config"
	"github.com/srikanthankamtw/todo-rest/migration"
	"github.com/srikanthankamtw/todo-rest/route"
	"log"
)

func init() {
	db := config.Init()
	migration.Migrate(db)
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := route.SetupRoutes()
	
	if err := router.Run(":8080"); err != nil {
		log.Panicf("error: %s", err)
	}
}
