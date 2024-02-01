package controllers

import (
	"net/http"

	"github.com/Befous/BackendGin/middleware"
	"github.com/Befous/BackendGin/models"
	"github.com/Befous/BackendGin/utils"
	"github.com/gin-gonic/gin"
)

func TokenValue(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		var response models.CredentialUser
		// Authorization
		middleware.Authorization(publickey)(c)
		if c.IsAborted() {
			return
		}
		name := c.GetString("name")
		username := c.GetString("username")
		role := c.GetString("role")
		// Cek Username
		if !utils.UsernameExists(mconn, collname, models.User{Username: username}) {
			c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Akun tidak ditemukan"})
			c.Abort()
			return
		}
		// Create Response
		response.Status = true
		response.Message = "Berhasil decode token"
		response.Data.Name = name
		response.Data.Username = username
		response.Data.Role = role

		c.JSON(http.StatusOK, response)
	}
}
