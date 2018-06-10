package platform

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/sdk/http/rest"
    "hoqu-geth-api/geth"
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

// swagger:route POST /platform/offer/add offers addOffer
//
// Add Offer.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: AddSuccessResponse
//   400: RestErrorResponse
//
func postAddOfferAction(c *gin.Context) {
    request := &models.AddOfferRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, id, err := geth.GetHoquPlatform().AddOffer(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
        "id": id,
    })
}

// swagger:route POST /platform/offer/status offers setOfferStatus
//
// Set Offer Status.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: TxSuccessResponse
//   400: RestErrorResponse
//
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

// swagger:route GET /platform/offer/:id offers getOffer
//
// Get Offer by ID.
//
// Produces:
// - application/json
// Responses:
//   200: OfferDataResponse
//   400: RestErrorResponse
//
func getOfferAction(c *gin.Context) {
    id := c.Param("id")

    offer, err := geth.GetHoQuStorage().GetOffer(id)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "offer": offer,
    })
}
