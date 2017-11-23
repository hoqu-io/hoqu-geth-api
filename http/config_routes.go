package http

import (
    "github.com/gin-gonic/gin"
    "hoqu-api/sdk/http/rest"
    "hoqu-api/geth"
    "hoqu-api/geth/models"
    "hoqu-api/sdk/http/middleware"
)

func initHoQuConfigRoutes(router *gin.Engine) {
    config := router.Group("/config")
    {
        config.POST("/deploy", middleware.SignRequired(), postDeployHoQuConfigAction)
        config.GET("/owner", getHoQuConfigSystemOwnerAction)
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

func getHoQuConfigSystemOwnerAction(c *gin.Context) {
    addr, err := geth.GetHoQuConfig().SystemOwner()
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "address": addr.String(),
    })
}
