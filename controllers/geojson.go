package controllers

import (
	"net/http"
	"strconv"

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
		var geospatial models.Geospatial
		err := c.BindJSON(&geospatial)
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

		geointersects, err := utils.GeoIntersects(mconn, collname, geospatial)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "GetGeoIntersectsDoc: " + err.Error()})
			return
		}
		result := utils.GeojsonNameString(geointersects)
		if result == "" {
			c.JSON(http.StatusOK, models.Pesan{Status: true, Empty: true, Message: "Tidak ada geojson yang bersinggungan dengan koordinat anda"})
			return
		}
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Geojson yang bersinggungan dengan koordinat anda adalah: " + result})
	}
}

func PostGeoWithin(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		var geospatial models.Geospatial
		err := c.BindJSON(&geospatial)
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

		geowithin, err := utils.GeoWithin(mconn, collname, geospatial)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "GetGeoWithinDoc: " + err.Error()})
			return
		}
		result := utils.GeojsonNameString(geowithin)
		if result == "" {
			c.JSON(http.StatusOK, models.Pesan{Status: true, Empty: true, Message: "Tidak ada geojson yang berada dengan koordinat anda"})
			return
		}
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Geojson yang berada di dalam koordinat anda adalah: " + result})
	}
}

func PostNear(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection2dsphere(mongoenv, dbname, collname)
		var geospatial models.Geospatial
		err := c.BindJSON(&geospatial)
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

		near, err := utils.Near(mconn, collname, geospatial)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "GetNearDoc: " + err.Error()})
			return
		}
		result := utils.GeojsonNameString(near)
		if result == "" {
			c.JSON(http.StatusOK, models.Pesan{Status: true, Empty: true, Message: "Tidak terdapat geojson yang berdekatan pada koordinat anda dengan max distance " + strconv.FormatFloat(geospatial.Max, 'f', -1, 64) + " meter dan min distance " + strconv.FormatFloat(geospatial.Min, 'f', -1, 64) + " meter"})
			return
		}
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Geojson yang berdekatan pada koordinat anda dengan max distance " + strconv.FormatFloat(geospatial.Max, 'f', -1, 64) + " meter dan min distance " + strconv.FormatFloat(geospatial.Min, 'f', -1, 64) + " meter adalah: " + result})
	}
}

func PostNearSphere(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection2dsphere(mongoenv, dbname, collname)
		var geospatial models.Geospatial
		err := c.BindJSON(&geospatial)
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

		nearsphere, err := utils.NearSphere(mconn, collname, geospatial)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "GetNearSphereDoc: " + err.Error()})
			return
		}
		result := utils.GeojsonNameString(nearsphere)
		if result == "" {
			c.JSON(http.StatusOK, models.Pesan{Status: true, Empty: true, Message: "Tidak ada geojson yang berdekatan pada koordinat anda dengan max distance " + strconv.FormatFloat(geospatial.Max, 'f', -1, 64) + " meter dan min distance " + strconv.FormatFloat(geospatial.Min, 'f', -1, 64) + " meter"})
			return
		}
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Geojson yang berdekatan pada koordinat anda dengan max distance " + strconv.FormatFloat(geospatial.Max, 'f', -1, 64) + " meter dan min distance " + strconv.FormatFloat(geospatial.Min, 'f', -1, 64) + " meter adalah: " + result})
	}
}

func PostBox(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		var geospatial models.Geospatial
		err := c.BindJSON(&geospatial)
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

		box, err := utils.Box(mconn, collname, geospatial)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "GetBoxDoc: " + err.Error()})
			return
		}
		result := utils.GeojsonNameString(box)
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: result})
	}
}

func PostCenter(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		var geospatial models.Geospatial
		err := c.BindJSON(&geospatial)
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

		center, err := utils.Center(mconn, collname, geospatial)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "GetCenterDoc: " + err.Error()})
			return
		}
		result := utils.GeojsonNameString(center)
		if result == "" {
			c.JSON(http.StatusOK, models.Pesan{Status: true, Empty: true, Message: "Tidak ada geojson yang berada di dalam lingkaran dengan radius " + strconv.FormatFloat(geospatial.Radius, 'f', -1, 64)})
			return
		}
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Geojson yang berada di dalam lingkaran dengan radius " + strconv.FormatFloat(geospatial.Radius, 'f', -1, 64) + " adalah: " + result})
	}
}

func PostCenterSphere(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		var geospatial models.Geospatial
		err := c.BindJSON(&geospatial)
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

		centersphere, err := utils.CenterSphere(mconn, collname, geospatial)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "GetCenterSphereDoc: " + err.Error()})
			return
		}
		result := utils.GeojsonNameString(centersphere)
		if result == "" {
			c.JSON(http.StatusOK, models.Pesan{Status: true, Empty: true, Message: "Tidak ada geojson yang berada di dalam lingkaran dengan radius " + strconv.FormatFloat(geospatial.Radius, 'f', -1, 64)})
			return
		}
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Geojson yang berada di dalam lingkaran dengan radius " + strconv.FormatFloat(geospatial.Radius, 'f', -1, 64) + " adalah: " + result})
	}
}

func AmbilDataGeojson(mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		datagedung := utils.GetAllBangunan(mconn, collname)
		c.JSON(http.StatusOK, datagedung)
	}
}
