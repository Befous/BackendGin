package routes

import (
	"github.com/Befous/BackendGin/controllers"
	"github.com/gin-gonic/gin"
)

func GeojsonRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/MembuatGeojsonPoint", controllers.MembuatGeojsonPoint("publickey", "mongoenv", "befous", "geojson"))
	incomingRoutes.POST("/MembuatGeojsonPolyline", controllers.MembuatGeojsonPolyline("publickey", "mongoenv", "befous", "geojson"))
	incomingRoutes.POST("/MembuatGeojsonPolygon", controllers.MembuatGeojsonPolygon("publickey", "mongoenv", "befous", "geojson"))
	incomingRoutes.POST("/PostGeoIntersects", controllers.PostGeoIntersects("publickey", "mongoenv", "befous", "geojson"))
	incomingRoutes.POST("/PostGeoWithin", controllers.PostGeoWithin("publickey", "mongoenv", "befous", "geojson"))
	incomingRoutes.POST("/PostNear", controllers.PostNear("publickey", "mongoenv", "befous", "geojson"))
	incomingRoutes.POST("/PostNearSphere", controllers.PostNearSphere("publickey", "mongoenv", "befous", "geojson"))
	incomingRoutes.POST("/PostBox", controllers.PostBox("publickey", "mongoenv", "befous", "geojson"))
	incomingRoutes.GET("/AmbilDataGeojson", controllers.AmbilDataGeojson("mongoenv", "befous", "geojson"))
}
