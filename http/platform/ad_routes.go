package platform

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/sdk/http/rest"
    "hoqu-geth-api/geth"
    "hoqu-geth-api/geth/models"
    "hoqu-geth-api/sdk/http/middleware"
)

func InitAdRoutes(routerGroup *gin.RouterGroup) {
    r := routerGroup.Group("/ad")
    {
        r.POST("/add", middleware.SignRequired(), postAddAdAction)
        r.POST("/status", middleware.SignRequired(), postSetAdStatusAction)
        r.GET("/:id", getAdAction)
    }
}

// swagger:route POST /platform/ad/add ads addAd
//
// Add Ad.
//
// The Ad is a campaign of the affiliate to promote particular offer.
// The Ad is created by the affiliate. All widgets placed on affiliate website should be linked to particular Ad.
// The affiliate can specify money receiver address (beneficiary) per each Ad.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: AddSuccessResponse
//   400: RestErrorResponse
//
func postAddAdAction(c *gin.Context) {
    request := &models.AddAdRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, id, err := geth.GetHoquPlatform().AddAd(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
        "id": id,
    })
}

// swagger:route POST /platform/ad/status ads setAdStatus
//
// Set Ad Status.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: TxSuccessResponse
//   400: RestErrorResponse
//
func postSetAdStatusAction(c *gin.Context) {
    request := &models.SetStatusRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().SetAdStatus(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

// swagger:route GET /platform/ad/:id ads getAd
//
// Get Ad by ID.
//
// Produces:
// - application/json
// Responses:
//   200: AdDataResponse
//   400: RestErrorResponse
//
func getAdAction(c *gin.Context) {
    id := c.Param("id")

    ad, err := geth.GetHoquPlatform().GetAd(id)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "Ad": ad,
    })
}
