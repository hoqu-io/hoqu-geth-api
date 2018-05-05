package http

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/sdk/http/rest"
    "hoqu-geth-api/geth"
    "hoqu-geth-api/geth/models"
    "hoqu-geth-api/sdk/http/middleware"
)

func initBurnerRoutes(router *gin.Engine) {
    claim := router.Group("/burner")
    {
        claim.POST("/deploy", middleware.SignRequired(), postDeployBurnerAction)
        claim.POST("/burn", middleware.SignRequired(), postBurnFromAction)
    }
}

func postDeployBurnerAction(c *gin.Context) {
    addr, tx, err := geth.GetBurner().Deploy()
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "address": addr.String(),
        "tx":      tx.Hash().String(),
    })
}

func postBurnFromAction(c *gin.Context) {
    request := &models.AddressWithAmount{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetBurner().BurnFrom(request.Address, request.Amount)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}


