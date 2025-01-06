package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	todo "go-todo/internal/modules/todo/handlers"
)

func main() {
	viper.SetConfigFile("./.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	//dbUrl := viper.Get("DB_URL").(string)

	router := gin.Default()
	app := router.Group("/api")
	//db.Init(dbUrl)
	todo.RegisterHandlers(app)

	router.Run(":" + port)
}
