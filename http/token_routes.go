package http

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/geth"
    "hoqu-geth-api/sdk/http/rest"
    "hoqu-geth-api/models"
    "hoqu-geth-api/sdk/http/middleware"
    sdkModels "hoqu-geth-api/sdk/models"
)

func InitTokenRoutes(router *gin.Engine) {
    hqx := router.Group("/hqx")
    {
        hqx.POST("/deploy", middleware.SignRequired(), postDeployTokenAction)
        hqx.GET("/balance/:address", getHqxBalanceAction)
        hqx.POST("/balances", getHqxBalancesAction)
        hqx.POST("/allowance", getHqxAllowanceAction)
        hqx.POST("/transactions", postTokenTransactionsAction)
        hqx.POST("/holders", postTokenHoldersAction)
    }
}

func postDeployTokenAction(c *gin.Context) {
    request := &models.TokenDeployParams{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    addr, tx, err := geth.GetToken().Deploy(request.TotalSupply)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "address": addr.String(),
        "tx":      tx.Hash().String(),
    })
}

// swagger:route GET /hqx/balance/:address platform getHqxBalance
//
// Get HQX balance
//
// Get particular Ethereum address HQX balance.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: GetBalanceSuccessResponse
//   400: RestErrorResponse
//
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

// swagger:route GET /hqx/allowance platform getHqxAllowance
//
// Get HQX allowance
//
// Get allowance given by owner of HQX tokens to spender.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: GetAllowanceSuccessResponse
//   400: RestErrorResponse
//
func getHqxAllowanceAction(c *gin.Context) {
    request := &models.AllowanceRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    allowance, err := geth.GetToken().Allowance(request.Owner, request.Spender)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "allowance": allowance.String(),
    })
}

// swagger:route POST /hqx/balances platform getHqxBalances
//
// Get HQX balances
//
// Get HQX balances for list of Ethereum addresses.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: GetBalancesSuccessResponse
//   400: RestErrorResponse
//
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

func postTokenTransactionsAction(c *gin.Context) {
    request := &sdkModels.Events{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    events, err := geth.GetToken().Events(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "transactions": events,
    })
}

func postTokenHoldersAction(c *gin.Context) {
    request := &sdkModels.Events{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    holders, err := geth.GetToken().Holders(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "count": len(holders),
        "holders": holders,
    })
}