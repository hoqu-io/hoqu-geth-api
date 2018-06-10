package platform

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/sdk/http/rest"
    "hoqu-geth-api/geth"
    "hoqu-geth-api/geth/models"
    "hoqu-geth-api/sdk/http/middleware"
)

func InitLeadRoutes(routerGroup *gin.RouterGroup) {
    r := routerGroup.Group("/lead")
    {
        r.POST("/add", middleware.SignRequired(), postAddLeadAction)
        r.POST("/intermediary", middleware.SignRequired(), postAddLeadIntermediaryAction)
        r.POST("/status", middleware.SignRequired(), postSetLeadStatusAction)
        r.POST("/transact", middleware.SignRequired(), postTransactLeadAction)
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
func postAddLeadAction(c *gin.Context) {
    request := &models.AddLeadRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, id, err := geth.GetHoquPlatform().AddLead(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
        "id": id,
    })
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
func postSetLeadStatusAction(c *gin.Context) {
    request := &models.SetLeadStatusRequest{}
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
func postTransactLeadAction(c *gin.Context) {
    request := &models.TransactLeadRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().TransactLead(request.Id, request.AdId)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
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
func getLeadAction(c *gin.Context) {
    request := &models.TransactLeadRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    adContract, err := geth.GetAdCampaign(request.AdId)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    lead, err := adContract.GetLead(request.Id)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "lead": lead,
    })
}
