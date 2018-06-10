// Package classification HOQU blockchain API by GoLang.
//
// Provides all projects the opportunity to access Ethereum blockchain through RESTful API.
//
//     Schemes: http, https
//     BasePath: /
//     Version: 0.4.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Denis Degterin <d.degterin@gmail.com> https://github.com/hoqu-io/hoqu-geth-api
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - hash
//     - jwt
//
//     SecurityDefinitions:
//     hash:
//          type: apiKey
//          name: X-Sign
//          in: header
//     jwt:
//         type: apiKey
//         name: X-Authorization
//         in: header
//
// swagger:meta
package http

import (
    sdkHttp "hoqu-geth-api/sdk/http"
    "github.com/gin-gonic/gin"
    "sync"
    "hoqu-geth-api/sdk/http/middleware"
    "hoqu-geth-api/sdk/geth/metamask"
)

var routerOnceInitializer sync.Once

func Run() error {
    return sdkHttp.Run(initRoutes)
}

func initRoutes(router *gin.Engine) {
    routerOnceInitializer.Do(func() {
        router.Use(middleware.ExecutionTime())

        metamask.AddMetamaskAuthRoutes(router)

        initWalletRoutes(router)
        initTokenRoutes(router)
        initPresaleRoutes(router)
        initSaleRoutes(router)
        initBountyRoutes(router)
        initClaimRoutes(router)
        initBurnerRoutes(router)
        initHoQuConfigRoutes(router)
        initHoQuStorageRoutes(router)
        initHoquPlatformRoutes(router)
        initHoQuRaterRoutes(router)
    })
}
