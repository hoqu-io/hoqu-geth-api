package http

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/sdk/http/rest"
    "hoqu-geth-api/geth"
    "hoqu-geth-api/models"
    "hoqu-geth-api/sdk/http/middleware"
    models2 "hoqu-geth-api/sdk/models"
)

func InitPresaleRoutes(router *gin.Engine) {
    presale := router.Group("/presale")
    {
        presale.POST("/deploy", middleware.SignRequired(), postDeployAction)
        presale.GET("/summary", getSummaryAction)
        presale.GET("/balance/:address", getClaimableTokensBalanceAction)
        presale.POST("/balances", postClaimableTokensBalancesAction)
        presale.GET("/approved/:address", getApprovedAction)
        presale.POST("/approved", postApprovedManyAction)
        presale.POST("/transactions", postPresaleTransactionsAction)
        presale.POST("/add", middleware.SignRequired(), postAddAction)
        presale.POST("/topup", middleware.SignRequired(), postTopUpAction)
        presale.POST("/approve", middleware.SignRequired(), postApproveAction)
    }
}

func postDeployAction(c *gin.Context) {
    request := &models.PresaleDeployParams{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    addr, tx, err := geth.GetPresale().Deploy(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "address": addr.String(),
        "tx":      tx.Hash().String(),
    })
}

func getSummaryAction(c *gin.Context) {
    sum, err := geth.GetPresale().Summary()
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "summary": sum,
    })
}

func getClaimableTokensBalanceAction(c *gin.Context) {
    addr := c.Param("address")
    bal, err := geth.GetPresale().Balance(addr)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "balance": bal.String(),
    })
}

func postClaimableTokensBalancesAction(c *gin.Context) {
    request := &models.Addresses{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    bals := map[string]string{}
    for _, addr := range request.Addresses {
        bal, err := geth.GetPresale().Balance(addr)
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

func getApprovedAction(c *gin.Context) {
    addr := c.Param("address")
    approved, err := geth.GetPresale().Approved(addr)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "approved": approved,
    })
}

func postApprovedManyAction(c *gin.Context) {
    request := &models.Addresses{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    approvedMap := map[string]bool{}
    for _, addr := range request.Addresses {
        approved, err := geth.GetPresale().Approved(addr)
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

func postPresaleTransactionsAction(c *gin.Context) {
    request := &models2.Events{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    events, err := geth.GetPresale().Events(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "transactions": events,
    })
}

func postAddAction(c *gin.Context) {
    request := &models.AddressWithAmount{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetPresale().Add(request.Address, request.Amount)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

func postTopUpAction(c *gin.Context) {
    request := &models.AddressWithAmount{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetPresale().TopUp(request.Address, request.Amount)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

func postApproveAction(c *gin.Context) {
    request := &models.Address{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetPresale().Approve(request.Address)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}
