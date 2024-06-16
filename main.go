package main

import (
	"fmt"
	"net/http"

	"github.com/clemensjur/sudoku/db"
	"github.com/clemensjur/sudoku/models"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	db := db.Db("dev.db")

	r.GET("/devices", func(c *gin.Context) {
		var devices []models.Device

		tx := db.Find(&devices)

		if tx.Error != nil {
			c.JSON(http.StatusOK, gin.H{"status": "no value"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "ok", "devices": devices})
		}
	})

	r.GET("/device/:id", func(c *gin.Context) {
		var device models.Device

		id := c.Params.ByName("id")
		tx := db.First(&device, "id = ?", id)
		if tx.Error != nil {
			c.JSON(http.StatusOK, gin.H{"status": "no value"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "ok", "device": device})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	/*
		authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
			"foo":  "bar", // user:foo password:bar
			"manu": "123", // user:manu password:123
		}))
	*/

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	/*
		authorized.POST("admin", func(c *gin.Context) {
			user := c.MustGet(gin.AuthUserKey).(string)

			// Parse JSON
			var json struct {
				Value string `json:"value" binding:"required"`
			}

			if c.Bind(&json) == nil {
				db[user] = json.Value
				c.JSON(http.StatusOK, gin.H{"status": "ok"})
			}
		})
	*/

	return r
}

func main() {

	db := db.Db("dev.db")

	err := db.AutoMigrate(&models.Device{})
	if err != nil {
		fmt.Println(err)
	}

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
