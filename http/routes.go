package http

import (
    sdkHttp "hoqu-geth-api/sdk/http"
    "github.com/gin-gonic/gin"
    "sync"
    "hoqu-geth-api/sdk/http/middleware"
)

var routerOnceInitializer sync.Once

func Run() error {
    return sdkHttp.Run(initRoutes)
}

func initRoutes(router *gin.Engine) {
    routerOnceInitializer.Do(func() {
        router.Use(middleware.ExecutionTime())

        initWalletRoutes(router)
        initTokenRoutes(router)
        initPresaleRoutes(router)
        initSaleRoutes(router)
        initBountyRoutes(router)
        initHoQuConfigRoutes(router)
        initHoquPlatformRoutes(router)
    })
}
