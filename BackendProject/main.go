package main

import (
	"github.com/gin-gonic/gin"

	"github.com/backend/db"
	"github.com/backend/routes"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
