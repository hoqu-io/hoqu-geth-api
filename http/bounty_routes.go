package http

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/sdk/http/rest"
    "hoqu-geth-api/geth"
    "hoqu-geth-api/models"
    "hoqu-geth-api/sdk/http/middleware"
)

func InitBountyRoutes(router *gin.Engine) {
    bounty := router.Group("/bounty")
    {
        bounty.POST("/deploy", middleware.SignRequired(), postDeployBountyAction)
        bounty.GET("/summary", getBountySummaryAction)
        bounty.GET("/balance/:address", getBountyTokensBalanceAction)
        bounty.POST("/balances", postBountyTokensBalancesAction)
        bounty.GET("/approved/:address", getBountyApprovedAction)
        bounty.POST("/approved", postBountyApprovedManyAction)
        bounty.POST("/add", middleware.SignRequired(), postBountyAddAction)
    }
}

func postDeployBountyAction(c *gin.Context) {
    request := &models.BountyDeployParams{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    addr, tx, err := geth.GetBounty().Deploy(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "address": addr.String(),
        "tx":      tx.Hash().String(),
    })
}

func getBountySummaryAction(c *gin.Context) {
    sum, err := geth.GetBounty().Summary()
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "summary": sum,
    })
}

func getBountyTokensBalanceAction(c *gin.Context) {
    addr := c.Param("address")
    bal, err := geth.GetBounty().Balance(addr)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "balance": bal.String(),
    })
}

func postBountyTokensBalancesAction(c *gin.Context) {
    request := &models.Addresses{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    bals := map[string]string{}
    for _, addr := range request.Addresses {
        bal, err := geth.GetBounty().Balance(addr)
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

func getBountyApprovedAction(c *gin.Context) {
    addr := c.Param("address")
    approved, err := geth.GetBounty().Approved(addr)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "approved": approved,
    })
}

func postBountyApprovedManyAction(c *gin.Context) {
    request := &models.Addresses{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    approvedMap := map[string]bool{}
    for _, addr := range request.Addresses {
        approved, err := geth.GetBounty().Approved(addr)
        if err != nil {
            rest.NewResponder(c).Error(err.Error())
            return
        }
        approvedMap[addr] = approved
    }

    rest.NewResponder(c).Success(gin.H{
        "approved": approvedMap,
    })
}

func postBountyAddAction(c *gin.Context) {
    request := &models.AddressWithAmount{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetBounty().AddByBounty(request.Address, request.Amount)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

