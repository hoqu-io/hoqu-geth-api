package offer

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/models"
    "hoqu-geth-api/sdk/http/middleware"
    "hoqu-geth-api/sdk/entity"
)

func InitOfferRoutes(routerGroup *gin.RouterGroup) {
    r := routerGroup.Group("/offer")
    {
        r.POST("/add", middleware.SignRequired(), postAddOfferAction)
        r.POST("/status", middleware.SignRequired(), postSetOfferStatusAction)
        r.POST("/name", middleware.SignRequired(), postSetOfferNameAction)
        r.POST("/payer_address", middleware.SignRequired(), postSetOfferPayerAddressAction)
        r.POST("/tariff_group", middleware.SignRequired(), postAddOfferTariffGroupAction)
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
func postAddOfferAction(ctx *gin.Context) {
    entity.CreateAction(ctx, &models.AddOfferRequest{}, GetService().Create)
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
func postSetOfferStatusAction(ctx *gin.Context) {
    entity.SetAction(ctx, &models.SetStatusRequest{}, GetService().SetStatus)
}

// swagger:route POST /platform/offer/name offers setOfferName
//
// Set Offer Name.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: TxSuccessResponse
//   400: RestErrorResponse
//
func postSetOfferNameAction(ctx *gin.Context) {
    entity.SetAction(ctx, &models.SetNameRequest{}, GetService().SetName)
}

// swagger:route POST /platform/offer/payer_address offers setOfferPayerAddress
//
// Set Offer Payer Address.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: TxSuccessResponse
//   400: RestErrorResponse
//
func postSetOfferPayerAddressAction(ctx *gin.Context) {
    entity.SetAction(ctx, &models.SetOfferPayerAddressRequest{}, GetService().SetPayerAddress)
}

// swagger:route POST /platform/offer/tariff_group offers addOfferTariffGroup
//
// Add Offer tariff.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: TxSuccessResponse
//   400: RestErrorResponse
//
func postAddOfferTariffGroupAction(ctx *gin.Context) {
    entity.SetAction(ctx, &models.IdWithChildIdRequest{}, GetService().AddChildById)
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
func getOfferAction(ctx *gin.Context) {
    id := ctx.Param("id")
    entity.GetByIdAction(ctx, id, &models.OfferSuccessData{}, GetService().GetById)
}
