package controllers

import (
	"net/http"
	"os"

	"github.com/Befous/BackendGin/helpers"
	"github.com/Befous/BackendGin/middleware"
	"github.com/Befous/BackendGin/models"
	"github.com/Befous/BackendGin/utils"
	"github.com/gin-gonic/gin"
)

func Registrasi(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		var user models.User
		err := c.BindJSON(&user)
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
			c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Anda tidak memiliki akses"})
			c.Abort()
			return
		}
		// Cek apakah username telah dipakai
		if user.Username == "" {
			c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Parameter dari function ini adalah username"})
			return
		}
		if utils.UsernameExists(mconn, collname, user) {
			c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Username telah dipakai"})
			return
		}
		// Hash password
		hash, hashErr := helpers.HashPassword(user.Password)
		if hashErr != nil {
			c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Gagal hash password: " + hashErr.Error()})
			return
		}
		user.Password = hash
		// Insert data user
		utils.InsertUser(mconn, collname, user)
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Berhasil input data"})
	}
}

func Login(privatekey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		var user models.User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "Error parsing application/json: " + err.Error()})
			return
		}
		// Cek apakah username ada
		if user.Username == "" {
			c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Parameter dari function ini adalah username"})
			return
		}
		if !utils.UsernameExists(mconn, collname, user) {
			c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Akun tidak ditemukan"})
			return
		}
		// Cek password
		if !utils.IsPasswordValid(mconn, collname, user) {
			c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Password salah"})
			return
		}
		// Encode data login
		datauser := utils.FindUser(mconn, collname, user)
		tokenstring, tokenerr := helpers.Encode(datauser.Name, datauser.Username, datauser.Role, os.Getenv(privatekey))
		if tokenerr != nil {
			c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Gagal encode token: " + tokenerr.Error()})
			return
		}
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Berhasil login", Token: tokenstring})
	}
}

func AmbilSemuaUser(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		// Authorization
		middleware.Authorization(publickey)(c)
		if c.IsAborted() {
			return
		}
		role := c.GetString("role")
		// Cek role
		if role != "owner" {
			c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Anda tidak memiliki akses"})
			c.Abort()
			return
		}
		// Get data user
		datauser, err := utils.GetAllUser(mconn, collname)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "GetAllDoc error: " + err.Error()})
			return
		}
		c.JSON(http.StatusOK, datauser)
	}
}

func EditUser(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		var user models.User
		err := c.BindJSON(&user)
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
			c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Anda tidak memiliki akses"})
			c.Abort()
			return
		}
		// Cek apakah username ada
		if user.Username == "" {
			c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Parameter dari function ini adalah username"})
			return
		}
		if !utils.UsernameExists(mconn, collname, user) {
			c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Akun yang ingin diedit tidak ditemukan"})
			return
		}
		// Hash password jika password di isi
		if user.Password != "" {
			hash, hashErr := helpers.HashPassword(user.Password)
			if hashErr != nil {
				c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Gagal hash password: " + hashErr.Error()})
				return
			}
			user.Password = hash
		} else {
			datauser := utils.FindUser(mconn, collname, user)
			user.Password = datauser.Password
		}
		// Update data user
		utils.UpdateUser(mconn, collname, user)
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Berhasil update " + user.Username + " dari database"})
	}
}

func HapusUser(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		var user models.User
		err := c.BindJSON(&user)
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
			c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Anda tidak memiliki akses"})
			c.Abort()
			return
		}
		// Cek apakah username ada
		if user.Username == "" {
			c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Parameter dari function ini adalah username"})
			return
		}
		if !utils.UsernameExists(mconn, collname, user) {
			c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Akun yang ingin dihapus tidak ditemukan"})
			return
		}
		// Update data user
		utils.DeleteUser(mconn, collname, user)
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Berhasil hapus " + user.Username + " dari database"})
	}
}
