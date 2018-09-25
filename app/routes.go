// Package classification HOQU blockchain API by GoLang.
//
// Provides all projects the opportunity to access Ethereum blockchain through RESTful API.
//
//     Schemes: http, https
//     BasePath: /
//     Version: 1.1.7
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

package app

import (
    sdkHttp "hoqu-geth-api/sdk/http"
    "github.com/gin-gonic/gin"
    "sync"
    "hoqu-geth-api/sdk/geth/metamask"
    "hoqu-geth-api/http"
)

var routerOnceInitializer sync.Once

func Run() error {
    return sdkHttp.Run(initRoutes)
}

func initRoutes(router *gin.Engine) {
    routerOnceInitializer.Do(func() {
        sdkHttp.UseTracing(router)

        metamask.AddMetamaskAuthRoutes(router)

        http.InitWalletRoutes(router)
        http.InitTokenRoutes(router)
        http.InitPresaleRoutes(router)
        http.InitSaleRoutes(router)
        http.InitBountyRoutes(router)
        http.InitClaimRoutes(router)
        http.InitBurnerRoutes(router)
        http.InitHoQuConfigRoutes(router)
        http.InitHoQuStorageRoutes(router)
        http.InitHoquPlatformRoutes(router)
        http.InitHoQuRaterRoutes(router)
        http.InitHoQuTransactorRoutes(router)
    })
}
