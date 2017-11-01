package http

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/sdk/http/rest"
    sdkGeth "hoqu-geth-api/sdk/geth"
    "hoqu-geth-api/geth/models"
)

func initWalletRoutes(router *gin.Engine)  {
    eth := router.Group("/eth")
    {
        eth.GET("/balance/:address", getEthBalanceAction)
        eth.POST("/balances", getEthBalancesAction)
    }
}

func getEthBalanceAction(c *gin.Context) {
    addr := c.Param("address")
    bal, err := sdkGeth.GetWallet().BalanceAt(addr)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "address": addr,
        "balance": bal.String(),
    })
}

func getEthBalancesAction(c *gin.Context) {
    request := &models.Addresses{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    bals := map[string]string{}
    for _, addr := range request.Addresses {
        bal, err := sdkGeth.GetWallet().BalanceAt(addr)
        if err != nil {
            rest.NewResponder(c).Error(err.Error())
            return
        }
        bals[addr] = bal.String()
    }

    rest.NewResponder(c).Success(gin.H{
        "balances": bals,
    })
}