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

func InitLeadRoutes(routerGroup *gin.RouterGroup) {
    r := routerGroup.Group("/lead")
    {
        r.POST("/add", middleware.SignRequired(), postAddLeadAction)
        r.POST("/intermediary", middleware.SignRequired(), postAddLeadIntermediaryAction)
        r.POST("/status", middleware.SignRequired(), postSetLeadStatusAction)
        r.POST("/sell", middleware.SignRequired(), postSellLeadAction)
        r.GET("/:id", getLeadAction)
    }
}

func postAddLeadAction(c *gin.Context) {
    request := &models.AddLeadRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().AddLead(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

func postAddLeadIntermediaryAction(c *gin.Context) {
    request := &models.AddLeadIntermediaryRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().AddLeadIntermediary(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

func postSetLeadStatusAction(c *gin.Context) {
    request := &models.SetStatusRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().SetLeadStatus(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

func postSellLeadAction(c *gin.Context) {
    request := &models.IdRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().SellLead(request.Id)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

func getLeadAction(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        rest.NewResponder(c).Error(fmt.Errorf("wrong Lead id provided: %v", err.Error()))
        return
    }

    lead, err := geth.GetHoquPlatform().GetLead(uint32(id))
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "Lead": lead,
    })
}
