package platform

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/sdk/http/rest"
    "hoqu-geth-api/geth"
    "strconv"
    "fmt"
    "hoqu-geth-api/geth/models"
    "hoqu-geth-api/sdk/http/middleware"
)

func InitUserRoutes(routerGroup *gin.RouterGroup) {
    r := routerGroup.Group("/user")
    {
        r.POST("/register", middleware.SignRequired(), postRegisterUserAction)
        r.POST("/kyc", middleware.SignRequired(), postAddUserKycReportAction)
        r.POST("/address", middleware.SignRequired(), postAddUserAddressAction)
        r.POST("/status", middleware.SignRequired(), postSetUserStatusAction)
        r.GET("/:id", getUserAction)
    }
}

func postRegisterUserAction(c *gin.Context) {
    request := &models.RegisterUserRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().RegisterUser(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

func postAddUserKycReportAction(c *gin.Context) {
    request := &models.AddUserKycReportRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().AddUserKycReport(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

func postAddUserAddressAction(c *gin.Context) {
    request := &models.AddUserAddressRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().AddUserAddress(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

func postSetUserStatusAction(c *gin.Context) {
    request := &models.SetStatusRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().SetUserStatus(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

func getUserAction(c *gin.Context) {
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
