package routes

import (
	"github.com/Befous/BackendGin/controllers"
	"github.com/gin-gonic/gin"
)

func GeojsonRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/MembuatGeojsonPoint", controllers.MembuatGeojsonPoint("publickey", "mongoenv", "befous", "geojson"))
	incomingRoutes.GET("/MembuatGeojsonPolyline", controllers.MembuatGeojsonPolyline("publickey", "mongoenv", "befous", "geojson"))
	incomingRoutes.GET("/MembuatGeojsonPolygon", controllers.MembuatGeojsonPolygon("publickey", "mongoenv", "befous", "geojson"))
	incomingRoutes.GET("/AmbilDataGeojson", controllers.AmbilDataGeojson("mongoenv", "befous", "geojson"))
}
