package ad_campaign

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/models"
    "hoqu-geth-api/sdk/http/middleware"
    "hoqu-geth-api/sdk/entity"
)

func InitAdCampaignRoutes(routerGroup *gin.RouterGroup) {
    r := routerGroup.Group("/ad")
    {
        r.POST("/add", middleware.SignRequired(), postAddAdCampaignAction)
        r.POST("/status", middleware.SignRequired(), postSetAdCampaignStatusAction)
        r.GET("/:id", getAdCampaignAction)
    }
}

// swagger:route POST /platform/ad/add ad_campaigns addAdCampaign
//
// Add AdCampaign.
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
//   200: AddAdCampaignSuccessResponse
//   400: RestErrorResponse
//
func postAddAdCampaignAction(ctx *gin.Context) {
    entity.HttpBindAction(ctx, &models.AddAdCampaignRequest{}, &models.CreateAdCampaignResponseData{}, GetService().Create)
}

// swagger:route POST /platform/ad/status ad_campaigns setAdCampaignStatus
//
// Set AdCampaign Status.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: TxSuccessResponse
//   400: RestErrorResponse
//
func postSetAdCampaignStatusAction(ctx *gin.Context) {
    entity.SetAction(ctx, &models.SetStatusRequest{}, GetService().SetStatus)
}

// swagger:route GET /platform/ad/:id ad_campaigns getAdCampaign
//
// Get AdCampaign by ID.
//
// Produces:
// - application/json
// Responses:
//   200: AdCampaignSuccessData
//   400: RestErrorResponse
//
func getAdCampaignAction(ctx *gin.Context) {
    id := ctx.Param("id")
    entity.GetByIdAction(ctx, id, &models.AdCampaignSuccessData{}, GetService().GetById)
}
