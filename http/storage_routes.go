package http

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/sdk/http/rest"
    "hoqu-geth-api/geth"
    "hoqu-geth-api/sdk/http/middleware"
)

func InitHoQuStorageRoutes(router *gin.Engine) {
    config := router.Group("/storage", middleware.SignRequired())
    {
        config.POST("/deploy", postDeployHoQuStorageAction)
    }
}

func postDeployHoQuStorageAction(c *gin.Context) {
    addr, tx, err := geth.GetHoQuStorage().Deploy()
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    _, err = geth.GetHoQuConfig().AddOwner(addr.String())
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "address": addr.String(),
        "tx":      tx.Hash().String(),
    })
}