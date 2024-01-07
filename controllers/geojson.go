package controllers

import (
	"net/http"

	"github.com/Befous/BackendGin/middleware"
	"github.com/Befous/BackendGin/models"
	"github.com/Befous/BackendGin/utils"
	"github.com/gin-gonic/gin"
)

func MembuatGeojsonPoint(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		var geojsonpoint models.GeoJsonPoint
		err := c.BindJSON(&geojsonpoint)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "Error parsing application/json: " + err.Error()})
			return
		}
		// Authorization
		middleware.Authorization(publickey)(c)
		if c.IsAborted() {
			return
		}
		role := c.GetString("role")
		// Cek role
		if role != "owner" {
			if role != "owner" {
				c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Anda tidak memiliki akses"})
				c.Abort()
				return
			}
		}
		// Insert data user
		utils.PostPoint(mconn, collname, geojsonpoint)
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Berhasil input data"})
	}
}

func MembuatGeojsonPolyline(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		var geojsonpolyline models.GeoJsonLineString
		err := c.BindJSON(&geojsonpolyline)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "Error parsing application/json: " + err.Error()})
			return
		}
		// Authorization
		middleware.Authorization(publickey)(c)
		if c.IsAborted() {
			return
		}
		role := c.GetString("role")
		// Cek role
		if role != "owner" {
			if role != "owner" {
				c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Anda tidak memiliki akses"})
				c.Abort()
				return
			}
		}
		// Insert data user
		utils.PostLinestring(mconn, collname, geojsonpolyline)
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Berhasil input data"})
	}
}

func MembuatGeojsonPolygon(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		var geojsonpolygon models.GeoJsonPolygon
		err := c.BindJSON(&geojsonpolygon)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "Error parsing application/json: " + err.Error()})
			return
		}
		// Authorization
		middleware.Authorization(publickey)(c)
		if c.IsAborted() {
			return
		}
		role := c.GetString("role")
		// Cek role
		if role != "owner" {
			if role != "owner" {
				c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Anda tidak memiliki akses"})
				c.Abort()
				return
			}
		}
		// Insert data user
		utils.PostPolygon(mconn, collname, geojsonpolygon)
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Berhasil input data"})
	}
}

func AmbilDataGeojson(mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		datagedung := utils.GetAllBangunan(mconn, collname)
		c.JSON(http.StatusOK, datagedung)
	}
}
