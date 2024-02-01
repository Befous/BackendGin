package routes

import (
	"github.com/Befous/BackendGin/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/TokenValue", controllers.TokenValue("publickey", "mongoenv", "befous", "user"))
	incomingRoutes.POST("/Registrasi", controllers.Registrasi("publickey", "mongoenv", "befous", "user"))
	incomingRoutes.POST("/Login", controllers.Login("privatekey", "mongoenv", "befous", "user"))
	incomingRoutes.GET("/AmbilSemuaUser", controllers.AmbilSemuaUser("publickey", "mongoenv", "befous", "user"))
	incomingRoutes.PUT("/EditUser", controllers.EditUser("publickey", "mongoenv", "befous", "user"))
	incomingRoutes.DELETE("/HapusUser", controllers.HapusUser("publickey", "mongoenv", "befous", "user"))
}
