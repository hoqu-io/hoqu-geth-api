package http

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/sdk/http/rest"
    "hoqu-geth-api/geth"
    "hoqu-geth-api/geth/models"
    "hoqu-geth-api/sdk/http/middleware"
    "github.com/ethereum/go-ethereum/common"
)

func initClaimRoutes(router *gin.Engine) {
    claim := router.Group("/claim")
    {
        claim.POST("/deploy", middleware.SignRequired(), postDeployClaimAction)
        claim.POST("/batch", middleware.SignRequired(), postClaimBatchAction)
        claim.POST("/one", middleware.SignRequired(), postClaimOneAction)
        claim.POST("/transactions", postClaimTransactionsAction)
    }
}

func postDeployClaimAction(c *gin.Context) {
    request := &models.ClaimDeployParams{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    addr, tx, err := geth.GetClaim().Deploy(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "address": addr.String(),
        "tx":      tx.Hash().String(),
    })
}

func postClaimBatchAction(c *gin.Context) {
    request := &models.AddressWithAmount{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    id := geth.GetClaim().BatchId
    tx, err := geth.GetClaim().ClaimAddress(request.Address, request.Amount)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    if tx.String() == (common.Hash{}).String() {
        rest.NewResponder(c).Success(gin.H{
            "id": id.String(),
        })
    } else {
        rest.NewResponder(c).Success(gin.H{
            "id": id.String(),
            "tx": tx.String(),
        })
    }
}

func postClaimOneAction(c *gin.Context) {
    request := &models.AddressWithAmount{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetClaim().ClaimOne(request.Address, request.Amount)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

func postClaimTransactionsAction(c *gin.Context) {
    request := &models.Addresses{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    events, err := geth.GetClaim().Events(request.Addresses)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "transactions": events,
    })
}
