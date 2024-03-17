package routes

import (
	"github.com/Befous/BackendGin/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/Mongo/TokenValue", controllers.TokenValueMongo("publickey", "mongoenv", "befous", "users"))
	incomingRoutes.POST("/Mongo/Registrasi", controllers.RegistrasiMongo("publickey", "mongoenv", "befous", "users"))
	incomingRoutes.POST("/Mongo/Login", controllers.LoginMongo("privatekey", "mongoenv", "befous", "users"))
	incomingRoutes.GET("/Mongo/AmbilSemuaUser", controllers.AmbilSemuaUserMongo("publickey", "mongoenv", "befous", "users"))
	incomingRoutes.PUT("/Mongo/EditUser", controllers.EditUserMongo("publickey", "mongoenv", "befous", "userss"))
	incomingRoutes.DELETE("/Mongo/HapusUser", controllers.HapusUserMongo("publickey", "mongoenv", "befous", "users"))
	incomingRoutes.POST("/Postgres/SudahLogin", controllers.SudahLogin)
	incomingRoutes.POST("/Postgres/Registrasi", controllers.RegistrasiPostgres)
	incomingRoutes.POST("/Postgres/Login", controllers.LoginPostgres)
	incomingRoutes.GET("/Postgres/AmbilSemuaUser", controllers.AmbilSemuaUserPostgres)
}
