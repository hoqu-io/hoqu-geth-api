package http

import (
    "fmt"
    "sync"
    "os"
    "github.com/sirupsen/logrus"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
)

var router *gin.Engine
var routerOnceInitializer sync.Once

func getRouter() *gin.Engine {
    routerOnceInitializer.Do(func(){
        router = gin.New()
    })
    return router
}

func initHttp(initRouters func(engine *gin.Engine)) {
    router = getRouter()

    initRouters(router)
}

func Run(initRouters func(engine *gin.Engine)) error {
    initHttp(initRouters)

    addr := fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port"))
    return router.Run(addr)
}

func StartTicking(interrupt chan os.Signal) {
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()

    for {
        select {
        case int := <-interrupt:
            if int == os.Interrupt {
                logrus.Info("Shutting down...")
                select {
                case <-time.After(time.Second):
                }

                return
            }
        }
    }
}