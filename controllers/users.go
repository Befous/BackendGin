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

func RegistrasiMongo(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		var user models.Users
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

func LoginMongo(privatekey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		var user models.Users
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

func AmbilSemuaUserMongo(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
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

func EditUserMongo(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		var user models.Users
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

func HapusUserMongo(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mconn := utils.SetConnection(mongoenv, dbname)
		var user models.Users
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

func RegistrasiPostgres(c *gin.Context) {
	var users models.Users
	pconn := utils.SetConnectionPostgres("HOST", "USER", "PASSWORD", "DB_NAME", "PORT_POSTGRES", "SSL")
	defer pconn.Close()
	err := c.BindJSON(&users)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "Error parsing application/json: " + err.Error()})
		return
	}
	if users.Username == "" {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Parameter username tidak boleh kosong"})
		return
	}
	if users.Password == "" {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Parameter password tidak boleh kosong"})
		return
	}
	// Hash password
	hash, hashErr := helpers.HashPassword(users.Password)
	if hashErr != nil {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Gagal hash password: " + hashErr.Error()})
		return
	}
	_, err = pconn.Exec("insert into users(username, password, role) values ($1, $2, $3)", users.Username, hash, users.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Error sql: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Berhasil input data"})
}

func LoginPostgres(c *gin.Context) {
	var users models.Users
	pconn := utils.SetConnectionPostgres("HOST", "USER", "PASSWORD", "DB_NAME", "PORT_POSTGRES", "SSL")
	defer pconn.Close()
	err := c.BindJSON(&users)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "Error parsing application/json: " + err.Error()})
		return
	}
	if users.Username == "" {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Parameter username tidak boleh kosong"})
		return
	}
	if users.Password == "" {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Parameter password tidak boleh kosong"})
		return
	}
	// Cek password
	row := pconn.QueryRow("SELECT * FROM users WHERE username = $1", users.Username)
	var spesifikuser models.Users
	err = row.Scan(&spesifikuser.Username, &spesifikuser.Password, &spesifikuser.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Error sql: " + err.Error()})
		return
	}
	hashChecker := helpers.CheckPasswordHash(users.Password, spesifikuser.Password)
	if !hashChecker {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Password salah"})
		return
	}
	// Encode data login
	tokenstring, tokenerr := helpers.Encode(spesifikuser.Name, spesifikuser.Username, spesifikuser.Role, os.Getenv("privatekey"))
	if tokenerr != nil {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Gagal encode token: " + tokenerr.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Berhasil login", Token: tokenstring})
}

func AmbilSemuaUserPostgres(c *gin.Context) {
	var users []models.Users
	pconn := utils.SetConnectionPostgres("HOST", "USER", "PASSWORD", "DB_NAME", "PORT_POSTGRES", "SSL")
	defer pconn.Close()
	rows, err := pconn.Query("SELECT * FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Error sql: " + err.Error()})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var user models.Users
		err := rows.Scan(&user.Username, &user.Password, &user.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Error scanning rows: " + err.Error()})
			return
		}
		users = append(users, user)
	}
	c.JSON(http.StatusOK, users)
}
