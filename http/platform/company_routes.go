package platform

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/sdk/http/rest"
    "hoqu-geth-api/geth"
    "hoqu-geth-api/geth/models"
    "hoqu-geth-api/sdk/http/middleware"
)

func InitCompanyRoutes(routerGroup *gin.RouterGroup) {
    r := routerGroup.Group("/company")
    {
        r.POST("/register", middleware.SignRequired(), postRegisterCompanyAction)
        r.POST("/status", middleware.SignRequired(), postSetCompanyStatusAction)
        r.GET("/:id", getCompanyAction)
    }
}

// swagger:route POST /platform/company/register companies registerCompany
//
// Register Company.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: AddSuccessResponse
//   400: RestErrorResponse
//
func postRegisterCompanyAction(c *gin.Context) {
    request := &models.RegisterCompanyRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, id, err := geth.GetHoquPlatform().RegisterCompany(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
        "id": id,
    })
}

// swagger:route POST /platform/company/status companies setCompanyStatus
//
// Set Company Status.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: TxSuccessResponse
//   400: RestErrorResponse
//
func postSetCompanyStatusAction(c *gin.Context) {
    request := &models.SetStatusRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().SetCompanyStatus(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

// swagger:route GET /platform/company/:id companies getCompany
//
// Get Company by ID.
//
// Produces:
// - application/json
// Responses:
//   200: CompanyDataResponse
//   400: RestErrorResponse
//
func getCompanyAction(c *gin.Context) {
    id := c.Param("id")

    company, err := geth.GetHoQuStorage().GetCompany(id)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "Company": company,
    })
}
