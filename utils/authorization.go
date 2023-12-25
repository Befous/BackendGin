package utils

import (
	"net/http"
	"os"

	"github.com/Befous/BackendGin/helpers"
	"github.com/Befous/BackendGin/models"
	"github.com/gin-gonic/gin"
)

func Authorization(publickey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		if header == "" {
			c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Header login tidak ditemukan"})
			c.Abort()
			return
		}
		tokenname := helpers.DecodeGetName(os.Getenv(publickey), header)
		tokenusername := helpers.DecodeGetUsername(os.Getenv(publickey), header)
		tokenrole := helpers.DecodeGetRole(os.Getenv(publickey), header)
		if tokenname == "" || tokenusername == "" || tokenrole == "" {
			c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Hasil decode tidak ditemukan"})
			c.Abort()
			return
		}
		c.Set("name", tokenname)
		c.Set("username", tokenusername)
		c.Set("role", tokenrole)
	}
}
