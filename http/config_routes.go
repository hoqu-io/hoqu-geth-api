package http

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/sdk/http/rest"
    "hoqu-geth-api/geth"
    "hoqu-geth-api/geth/models"
    "hoqu-geth-api/sdk/http/middleware"
)

func initHoQuConfigRoutes(router *gin.Engine) {
    config := router.Group("/config", middleware.SignRequired())
    {
        config.POST("/deploy", postDeployHoQuConfigAction)
        config.POST("/owners/add", postAddOwnersAction)
    }
}

func postDeployHoQuConfigAction(c *gin.Context) {
    request := &models.ConfigDeployParams{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    addr, tx, err := geth.GetHoQuConfig().Deploy(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "address": addr.String(),
        "tx":      tx.Hash().String(),
    })
}

func postAddOwnersAction(c *gin.Context) {
    request := &models.Addresses{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    txs := map[string]string{}
    for _, addr := range request.Addresses {
        tx, err := geth.GetHoQuConfig().AddOwner(addr)
        if err != nil {
            rest.NewResponder(c).Error(err.Error())
            return
        }
        txs[addr] = tx.String()
    }

    rest.NewResponder(c).Success(gin.H{
        "txs": txs,
    })
}