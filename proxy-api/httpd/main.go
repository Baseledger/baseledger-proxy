package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/unibrightio/proxy-api/httpd/handler"
)

func main() {
	setupViper()
	r := gin.Default()
	r.GET("/ping", handler.PingGet())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// discuss if we should use config struct or this is enough
func setupViper() {
	viper.AddConfigPath("../")
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("viper read config error")
	}
}
