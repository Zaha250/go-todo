package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-todo/common/db"
)

func main() {
	viper.SetConfigFile("./.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	db.Init(dbUrl)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbUrl,
		})
	})

	r.Run(":" + port)
}
