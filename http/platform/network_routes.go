package platform

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/sdk/http/rest"
    "hoqu-geth-api/geth"
    "hoqu-geth-api/geth/models"
    "hoqu-geth-api/sdk/http/middleware"
)

func InitNetworkRoutes(routerGroup *gin.RouterGroup) {
    r := routerGroup.Group("/network")
    {
        r.POST("/register", middleware.SignRequired(), postRegisterNetworkAction)
        r.POST("/status", middleware.SignRequired(), postSetNetworkStatusAction)
        r.GET("/:id", getNetworkAction)
    }
}

// swagger:route POST /platform/network/register networks registerNetwork
//
// Register Network.
//
// The Network is an essential part of a platform.
// All offers can be accessible only through networks.
// Merchants should join networks as well as Affiliates to communicate with each other.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: AddSuccessResponse
//   400: RestErrorResponse
//
func postRegisterNetworkAction(c *gin.Context) {
    request := &models.RegisterNetworkRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, id, err := geth.GetHoquPlatform().RegisterNetwork(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
        "id": id,
    })
}

// swagger:route POST /platform/network/status networks setNetworkStatus
//
// Set Network Status.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: TxSuccessResponse
//   400: RestErrorResponse
//
func postSetNetworkStatusAction(c *gin.Context) {
    request := &models.SetStatusRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().SetNetworkStatus(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

// swagger:route GET /platform/network/:id networks getNetwork
//
// Get Network by ID.
//
// Produces:
// - application/json
// Responses:
//   200: NetworkDataResponse
//   400: RestErrorResponse
//
func getNetworkAction(c *gin.Context) {
    id := c.Param("id")

    network, err := geth.GetHoQuStorage().GetNetwork(id)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "network": network,
    })
}
