package http

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/sdk/http/rest"
    "hoqu-geth-api/geth"
    "hoqu-geth-api/geth/models"
    "hoqu-geth-api/sdk/http/middleware"
    "math/big"
)

func initSaleRoutes(router *gin.Engine) {
    sale := router.Group("/sale")
    {
        sale.POST("/deploy", middleware.SignRequired(), postDeploySaleAction)
        sale.GET("/summary", getSaleSummaryAction)
        sale.GET("/balance/:address", getSaleTokensBalanceAction)
        sale.POST("/balances", postSaleTokensBalancesAction)
        sale.GET("/approved/:address", getSaleApprovedAction)
        sale.POST("/approved", postSaleApprovedManyAction)
        sale.POST("/transactions", postSaleTransactionsAction)
        sale.POST("/add", middleware.SignRequired(), postSaleAddAction)
        sale.POST("/topup", middleware.SignRequired(), postSaleTopUpAction)
        sale.POST("/approve", middleware.SignRequired(), postSaleApproveAction)
        sale.POST("/claim", middleware.SignRequired(), postSaleClaimAllAction)
    }
}

func postDeploySaleAction(c *gin.Context) {
    request := &models.SaleDeployParams{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    addr, tx, err := geth.GetSale().Deploy(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "address": addr.String(),
        "tx":      tx.Hash().String(),
    })
}

func getSaleSummaryAction(c *gin.Context) {
    sum, err := geth.GetSale().Summary()
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "summary": sum,
    })
}

func getSaleTokensBalanceAction(c *gin.Context) {
    addr := c.Param("address")
    bal, err := geth.GetSale().Balance(addr)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    bbal, err := geth.GetBounty().Balance(addr)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "balance": big.NewInt(0).Add(bal, bbal).String(),
    })
}

func postSaleTokensBalancesAction(c *gin.Context) {
    request := &models.Addresses{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    bals := map[string]string{}
    for _, addr := range request.Addresses {
        bal, err := geth.GetSale().Balance(addr)
        if err != nil {
            rest.NewResponder(c).Error(err.Error())
            return
        }

        bbal, err := geth.GetBounty().Balance(addr)
        if err != nil {
            rest.NewResponder(c).Error(err.Error())
            return
        }

        bals[addr] = big.NewInt(0).Add(bal, bbal).String()
    }

    rest.NewResponder(c).Success(gin.H{
        "balances": bals,
    })
}

func getSaleApprovedAction(c *gin.Context) {
    addr := c.Param("address")
    approved, err := geth.GetSale().Approved(addr)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "approved": approved,
    })
}

func postSaleApprovedManyAction(c *gin.Context) {
    request := &models.Addresses{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    approvedMap := map[string]bool{}
    for _, addr := range request.Addresses {
        approved, err := geth.GetSale().Approved(addr)
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

func postSaleTransactionsAction(c *gin.Context) {
    request := &models.Addresses{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    events, err := geth.GetSale().Events(request.Addresses)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    bEvents, err := geth.GetBounty().Events(request.Addresses)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    events = append(events, bEvents...)

    rest.NewResponder(c).Success(gin.H{
        "transactions": events,
    })
}

func postSaleAddAction(c *gin.Context) {
    request := &models.AddressWithAmount{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetSale().Add(request.Address, request.Amount)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

func postSaleTopUpAction(c *gin.Context) {
    request := &models.AddressWithAmount{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetSale().TopUp(request.Address, request.Amount)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

func postSaleApproveAction(c *gin.Context) {
    request := &models.Address{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetSale().Approve(request.Address)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

func postSaleClaimAllAction(c *gin.Context) {
    tx, err := geth.GetSale().ClaimAll()
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}
