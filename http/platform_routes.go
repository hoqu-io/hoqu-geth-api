package http

import (
    "github.com/gin-gonic/gin"
    "hoqu-api/sdk/http/rest"
    "hoqu-api/geth"
    "hoqu-api/sdk/http/middleware"
    "strconv"
    "fmt"
)

func initHoquPlatformRoutes(router *gin.Engine) {
    platform := router.Group("/platform")
    {
        platform.POST("/deploy", middleware.SignRequired(), postDeployHoquPlatformAction)
        platform.GET("/user/:id", getPlatformUserAction)
    }
}

func postDeployHoquPlatformAction(c *gin.Context) {
    addr, tx, err := geth.GetHoquPlatform().Deploy()
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "address": addr.String(),
        "tx":      tx.Hash().String(),
    })
}

func getPlatformUserAction(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        rest.NewResponder(c).Error(fmt.Errorf("wrong user id provided: %v", err.Error()))
        return
    }

    user, err := geth.GetHoquPlatform().GetUser(uint32(id))
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "user": user,
    })
}
