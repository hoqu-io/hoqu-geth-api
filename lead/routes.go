package lead

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/models"
    "hoqu-geth-api/sdk/http/middleware"
    "hoqu-geth-api/sdk/entity"
)

func InitLeadRoutes(routerGroup *gin.RouterGroup) {
    r := routerGroup.Group("/lead")
    {
        r.POST("/add", middleware.SignRequired(), postAddLeadAction)
        r.POST("/status", middleware.SignRequired(), postSetLeadStatusAction)
        r.POST("/price", middleware.SignRequired(), postSetLeadPriceAction)
        r.POST("/url", middleware.SignRequired(), postSetLeadDataUrlAction)
        r.POST("/intermediary", middleware.SignRequired(), postAddLeadIntermediaryAction)
        r.POST("/transact", middleware.SignRequired(), postTransactAction)
        r.POST("/get", getLeadAction)
    }
}

// swagger:route POST /platform/lead/add leads addLead
//
// Add Lead.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: AddSuccessResponse
//   400: RestErrorResponse
//
func postAddLeadAction(ctx *gin.Context) {
    entity.CreateAction(ctx, &models.AddLeadRequest{}, GetService().Create)
}

// swagger:route POST /platform/lead/status leads setLeadStatus
//
// Set Lead Status.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: TxSuccessResponse
//   400: RestErrorResponse
//
func postSetLeadStatusAction(ctx *gin.Context) {
    entity.SetAction(ctx, &models.SetLeadStatusRequest{}, GetService().SetStatus)
}

// swagger:route POST /platform/lead/price leads setLeadPrice
//
// Set Lead Price.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: TxSuccessResponse
//   400: RestErrorResponse
//
func postSetLeadPriceAction(ctx *gin.Context) {
    entity.SetAction(ctx, &models.SetLeadPriceRequest{}, GetService().SetPrice)
}

// swagger:route POST /platform/lead/url leads setLeadDataUrl
//
// Set Lead Data URL.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: TxSuccessResponse
//   400: RestErrorResponse
//
func postSetLeadDataUrlAction(ctx *gin.Context) {
    entity.SetAction(ctx, &models.SetLeadDataUrlRequest{}, GetService().SetUrl)
}

// swagger:route POST /platform/lead/intermediary leads addLeadIntermediary
//
// Add Lead Intermediary.
//
// Each intermediary is a money receiver when selling the lead.
// Warning! Lead owner shouldn't be in this list. Owner receives the rest of amount.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: TxSuccessResponse
//   400: RestErrorResponse
//
func postAddLeadIntermediaryAction(ctx *gin.Context) {
    entity.SetAction(ctx, &models.AddLeadIntermediaryRequest{}, GetService().AddIntermediary)
}

// swagger:route POST /platform/lead/transact leads transactLead
//
// Sell Lead.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: TxSuccessResponse
//   400: RestErrorResponse
//
func postTransactAction(ctx *gin.Context) {
    entity.SetAction(ctx, &models.TransactLeadRequest{}, GetService().Transact)
}

// swagger:route POST /platform/lead/get leads getLead
//
// Get Lead by ID.
//
// Produces:
// - application/json
// Responses:
//   200: LeadDataResponse
//   400: RestErrorResponse
//
func getLeadAction(ctx *gin.Context) {
    entity.HttpBindAction(ctx, &models.TransactLeadRequest{}, &models.LeadSuccessData{}, GetService().GetById)
}
