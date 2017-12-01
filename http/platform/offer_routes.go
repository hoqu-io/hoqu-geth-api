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

func InitOfferRoutes(routerGroup *gin.RouterGroup) {
    r := routerGroup.Group("/offer")
    {
        r.POST("/add", middleware.SignRequired(), postAddOfferAction)
        r.POST("/status", middleware.SignRequired(), postSetOfferStatusAction)
        r.GET("/:id", getOfferAction)
    }
}

func postAddOfferAction(c *gin.Context) {
    request := &models.AddOfferRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().AddOffer(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

func postSetOfferStatusAction(c *gin.Context) {
    request := &models.SetStatusRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().SetOfferStatus(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

func getOfferAction(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        rest.NewResponder(c).Error(fmt.Errorf("wrong Offer id provided: %v", err.Error()))
        return
    }

    offer, err := geth.GetHoquPlatform().GetOffer(uint32(id))
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "Offer": offer,
    })
}
