package http

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/sdk/http/rest"
    "hoqu-geth-api/geth"
    "hoqu-geth-api/sdk/http/middleware"
    "hoqu-geth-api/http/platform"
    models2 "hoqu-geth-api/sdk/models"
)

func initHoquPlatformRoutes(router *gin.Engine) {
    r := router.Group("/platform")
    {
        r.POST("/deploy", middleware.SignRequired(), postDeployHoquPlatformAction)
        r.POST("/events", postPlatformEventsAction)
        platform.InitUserRoutes(r)
        platform.InitIdentificationRoutes(r)
        platform.InitStatsRoutes(r)
        platform.InitCompanyRoutes(r)
        platform.InitNetworkRoutes(r)
        platform.InitTrackerRoutes(r)
        platform.InitOfferRoutes(r)
        platform.InitAdRoutes(r)
        platform.InitLeadRoutes(r)
    }
}

func postDeployHoquPlatformAction(c *gin.Context) {
    addr, tx, err := geth.GetHoquPlatform().Deploy()
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    _, err = geth.GetHoQuConfig().AddOwner(addr.String())
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "address": addr.String(),
        "tx":      tx.Hash().String(),
    })
}

// swagger:route POST /platform/events platform events
//
// Get HOQU platform events.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: ContractEventResponse
//   400: RestErrorResponse
//
func postPlatformEventsAction(c *gin.Context) {
    request := &models2.Events{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    events, err := geth.GetHoquPlatform().Events(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "events": events,
    })
}
