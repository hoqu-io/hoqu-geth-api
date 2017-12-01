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

func InitCompanyRoutes(routerGroup *gin.RouterGroup) {
    r := routerGroup.Group("/company")
    {
        r.POST("/register", middleware.SignRequired(), postRegisterCompanyAction)
        r.POST("/status", middleware.SignRequired(), postSetCompanyStatusAction)
        r.GET("/:id", getCompanyAction)
    }
}

func postRegisterCompanyAction(c *gin.Context) {
    request := &models.RegisterCompanyRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().RegisterCompany(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

func postSetCompanyStatusAction(c *gin.Context) {
    request := &models.SetStatusRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().SetCompanyStatus(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

func getCompanyAction(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        rest.NewResponder(c).Error(fmt.Errorf("wrong Company id provided: %v", err.Error()))
        return
    }

    company, err := geth.GetHoquPlatform().GetCompany(uint32(id))
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "Company": company,
    })
}
