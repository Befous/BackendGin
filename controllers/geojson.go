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
			if role != "dosen" {
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
			if role != "dosen" {
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
			if role != "dosen" {
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

func PostGeoIntersects(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		var coordinate models.Point
		err := c.BindJSON(&coordinate)
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
			if role != "dosen" {
				c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Anda tidak memiliki akses"})
				c.Abort()
				return
			}
		}

		geointersects := utils.GeoIntersects(mconn, collname, coordinate)
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: geointersects})
	}
}

func PostGeoWithin(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		var coordinate models.Polygon
		err := c.BindJSON(&coordinate)
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
			if role != "dosen" {
				c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Anda tidak memiliki akses"})
				c.Abort()
				return
			}
		}

		geowithin := utils.GeoWithin(mconn, collname, coordinate)
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: geowithin})
	}
}

func PostNear(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection2dsphere(mongoenv, dbname, collname)
		var coordinate models.Point
		err := c.BindJSON(&coordinate)
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
			if role != "dosen" {
				c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Anda tidak memiliki akses"})
				c.Abort()
				return
			}
		}

		near := utils.Near(mconn, collname, coordinate)
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: near})
	}
}

func PostNearSphere(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection2dsphere(mongoenv, dbname, collname)
		var coordinate models.Point
		err := c.BindJSON(&coordinate)
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
			if role != "dosen" {
				c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Anda tidak memiliki akses"})
				c.Abort()
				return
			}
		}

		nearsphere := utils.NearSphere(mconn, collname, coordinate)
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: nearsphere})
	}
}

func PostBox(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		var coordinate models.Polyline
		err := c.BindJSON(&coordinate)
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
			if role != "dosen" {
				c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Anda tidak memiliki akses"})
				c.Abort()
				return
			}
		}

		box := utils.Box(mconn, collname, coordinate)
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: box})
	}
}

func PostCenter(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		var coordinate models.Point
		err := c.BindJSON(&coordinate)
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
			if role != "dosen" {
				c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Anda tidak memiliki akses"})
				c.Abort()
				return
			}
		}

		box := utils.Center(mconn, collname, coordinate)
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: box})
	}
}

func PostCenterSphere(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		var coordinate models.Point
		err := c.BindJSON(&coordinate)
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
			if role != "dosen" {
				c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Anda tidak memiliki akses"})
				c.Abort()
				return
			}
		}

		box := utils.CenterSphere(mconn, collname, coordinate)
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: box})
	}
}

func AmbilDataGeojson(mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		datagedung := utils.GetAllBangunan(mconn, collname)
		c.JSON(http.StatusOK, datagedung)
	}
}
