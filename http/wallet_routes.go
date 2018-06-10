package http

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/sdk/http/rest"
    sdkGeth "hoqu-geth-api/sdk/geth"
    "hoqu-geth-api/geth/models"
    "github.com/ethereum/go-ethereum/common"
)

func initWalletRoutes(router *gin.Engine) {
    eth := router.Group("/eth")
    {
        eth.GET("/balance/:address", getEthBalanceAction)
        eth.POST("/balances", getEthBalancesAction)
        eth.GET("/block", getLatestBlockHeaderAction)
        eth.GET("/block/:hash", getBlockHeaderByHashAction)
    }
}

// swagger:route GET /eth/balance/:address ethereum getEthBalance
//
// Get ETH balance
//
// Get particular Ethereum address balance.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: GetBalanceSuccessResponse
//   400: RestErrorResponse
//
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

// swagger:route POST /eth/balances ethereum getEthBalances
//
// Get ETH balances
//
// Get balances for list of Ethereum addresses.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: GetBalancesSuccessResponse
//   400: RestErrorResponse
//
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

// swagger:route GET /eth/block ethereum getEthLastBlock
//
// Get last block
//
// Get last Ethereum block data.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: GetBlockSuccessResponse
//   400: RestErrorResponse
//
func getBlockHeaderByHashAction(c *gin.Context) {
    hash := c.Param("hash")
    b, err := sdkGeth.GetWallet().GetBlockHeaderByHash(common.HexToHash(hash))
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "number": b.Number.String(),
        "hash": b.Hash().String(),
        "timestamp": b.Time.String(),
        "raw": b,
    })
}

// swagger:route GET /eth/block/:hash ethereum getEthBlockByHash
//
// Get block by hash
//
// Get Ethereum block data by its hash.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: GetBlockSuccessResponse
//   400: RestErrorResponse
//
func getLatestBlockHeaderAction(c *gin.Context) {
    b, err := sdkGeth.GetWallet().GetBlockHeaderByNumber(nil)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "number": b.Number.String(),
        "hash": b.Hash().String(),
        "timestamp": b.Time.String(),
        "raw": b,
    })
}
