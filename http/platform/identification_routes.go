package platform

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/sdk/http/rest"
    "hoqu-geth-api/geth"
    "hoqu-geth-api/geth/models"
    "hoqu-geth-api/sdk/http/middleware"
)

func InitIdentificationRoutes(routerGroup *gin.RouterGroup) {
    r := routerGroup.Group("/identification")
    {
        r.POST("/add", middleware.SignRequired(), postAddIdentificationAction)
        r.POST("/kyc", middleware.SignRequired(), postAddKycReportAction)
        r.GET("/:id", getIdentificationAction)
    }
}

// swagger:route POST /platform/identification/add identifications addIdentification
//
// Add Identification.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: AddSuccessResponse
//   400: RestErrorResponse
//
func postAddIdentificationAction(c *gin.Context) {
    request := &models.AddIdentificationRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, id, err := geth.GetHoquPlatform().AddIdentification(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
        "id": id,
    })
}

// swagger:route POST /platform/identification/kyc identifications addKyc
//
// Add KYC report.
//
// Each Identification has KYC level which reflects his trust level.
// Identification KYC level can be changed only by KYC reports.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: TxSuccessResponse
//   400: RestErrorResponse
//
func postAddKycReportAction(c *gin.Context) {
    request := &models.AddKycReportRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().AddKycReport(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

// swagger:route GET /platform/identification/:id identifications getIdentification
//
// Get Identification by ID.
//
// Produces:
// - application/json
// Responses:
//   200: IdentificationDataResponse
//   400: RestErrorResponse
//
func getIdentificationAction(c *gin.Context) {
    id := c.Param("id")

    identification, err := geth.GetHoQuStorage().GetIdentification(id)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "identification": identification,
    })
}
