package main

import (
	"net/http"

	"github.com/Befous/BackendGin/models"
	"github.com/Befous/BackendGin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	routes.UserRoutes(router)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, models.Pesan{Status: false, Message: "Page not found"})
	})

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
