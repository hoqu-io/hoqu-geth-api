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
        r.POST("/sell", middleware.SignRequired(), postSellLeadAction)
        r.GET("/:id", getLeadAction)
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
    request := &models.SetStatusRequest{}
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

// swagger:route POST /platform/lead/sell leads sellLead
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
func postSellLeadAction(c *gin.Context) {
    request := &models.IdRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().SellLead(request.Id)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

// swagger:route GET /platform/lead/:id leads getLead
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
    id := c.Param("id")

    lead, err := geth.GetHoquPlatform().GetLead(id)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "Lead": lead,
    })
}
