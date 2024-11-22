package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func RunHTTPServer(serviceName string, wrapper func(router *gin.Engine)) {
	addr := viper.Sub(serviceName).GetString("http-addr")
	if addr == "" {
		// TODO: Warning log
	}
	RunHTTPServerOnAddr(addr, wrapper)
}

// wrapper的作用是将gin.Engine传入，然后在wrapper中定义路由，这样就可以在不同的服务中定义不同的路由，而不需要修改RunHTTPServer函数
func RunHTTPServerOnAddr(addr string, wrapper func(router *gin.Engine)) {
	apiRouter := gin.New()
	wrapper(apiRouter)
	apiRouter.Group("/api")
	apiRouter.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})
	if err := apiRouter.Run(addr); err != nil {
		panic(err)
	}
}
