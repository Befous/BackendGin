package routes

import (
	"github.com/Befous/BackendGin/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/TokenValue", controllers.TokenValue("publickey", "mongoenv", "befous", "users"))
	incomingRoutes.GET("/Registrasi", controllers.Registrasi("publickey", "mongoenv", "befous", "users"))
	incomingRoutes.GET("/Login", controllers.Login("privatekey", "mongoenv", "befous", "users"))
	incomingRoutes.GET("/AmbilSemuaUser", controllers.AmbilSemuaUser("publickey", "mongoenv", "befous", "users"))
	incomingRoutes.GET("/EditUser", controllers.EditUser("publickey", "mongoenv", "befous", "users"))
	incomingRoutes.GET("/HapusUser", controllers.HapusUser("publickey", "mongoenv", "befous", "users"))
}
