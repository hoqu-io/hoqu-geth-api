package http

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/geth"
    "hoqu-geth-api/sdk/http/rest"
    "hoqu-geth-api/geth/models"
)

func initTokenRoutes(router *gin.Engine) {
    hqx := router.Group("/hqx")
    {
        hqx.GET("/balance/:address", getHqxBalanceAction)
        hqx.POST("/balances", getHqxBalancesAction)
    }
}

func getHqxBalanceAction(c *gin.Context) {
    addr := c.Param("address")
    bal, err := geth.GetToken().Balance(addr)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "address": addr,
        "balance": bal.String(),
    })
}

func getHqxBalancesAction(c *gin.Context) {
    request := &models.Addresses{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    bals := map[string]string{}
    for _, addr := range request.Addresses {
        bal, err := geth.GetToken().Balance(addr)
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
