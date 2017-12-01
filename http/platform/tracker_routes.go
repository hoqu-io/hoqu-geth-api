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

func InitTrackerRoutes(routerGroup *gin.RouterGroup) {
    r := routerGroup.Group("/tracker")
    {
        r.POST("/register", middleware.SignRequired(), postRegisterTrackerAction)
        r.POST("/status", middleware.SignRequired(), postSetTrackerStatusAction)
        r.GET("/:id", getTrackerAction)
    }
}

func postRegisterTrackerAction(c *gin.Context) {
    request := &models.RegisterTrackerRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().RegisterTracker(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

func postSetTrackerStatusAction(c *gin.Context) {
    request := &models.SetStatusRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().SetTrackerStatus(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

func getTrackerAction(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        rest.NewResponder(c).Error(fmt.Errorf("wrong Tracker id provided: %v", err.Error()))
        return
    }

    tracker, err := geth.GetHoquPlatform().GetTracker(uint32(id))
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "Tracker": tracker,
    })
}
