package http

import (
    "fmt"
    "sync"
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
    "hoqu-geth-api/sdk/http/middleware"
    "hoqu-geth-api/sdk/tracing"
)

var router *gin.Engine
var routerOnceInitializer sync.Once

func Run(initRouters func(engine *gin.Engine)) error {
    router = getRouter()
    initRouters(router)

    addr := fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port"))
    return router.Run(addr)
}

func UseTracing(router *gin.Engine) {
    router.Use(middleware.ReqInput, middleware.ExecutionTime, tracing.Middleware, middleware.ReqOutput)
}

func getRouter() *gin.Engine {
    routerOnceInitializer.Do(func() {
        router = gin.New()
    })
    return router
}
