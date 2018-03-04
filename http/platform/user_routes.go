package platform

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/sdk/http/rest"
    "hoqu-geth-api/geth"
    "hoqu-geth-api/geth/models"
    "hoqu-geth-api/sdk/http/middleware"
)

func InitUserRoutes(routerGroup *gin.RouterGroup) {
    r := routerGroup.Group("/user")
    {
        r.POST("/register", middleware.SignRequired(), postRegisterUserAction)
        r.POST("/kyc", middleware.SignRequired(), postAddUserKycReportAction)
        r.POST("/address", middleware.SignRequired(), postAddUserAddressAction)
        r.POST("/status", middleware.SignRequired(), postSetUserStatusAction)
        r.GET("/:id", getUserAction)
    }
}

// swagger:route POST /platform/user/register users registerUser
//
// Register User.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: AddSuccessResponse
//   400: RestErrorResponse
//
func postRegisterUserAction(c *gin.Context) {
    request := &models.RegisterUserRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, id, err := geth.GetHoquPlatform().RegisterUser(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
        "id": id,
    })
}

// swagger:route POST /platform/user/kyc users addUserKyc
//
// Add User KYC report.
//
// Each User has KYC level which reflects his trust level.
// User KYC level can be changed only by KYC reports.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: TxSuccessResponse
//   400: RestErrorResponse
//
func postAddUserKycReportAction(c *gin.Context) {
    request := &models.AddUserKycReportRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().AddUserKycReport(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

// swagger:route POST /platform/user/address users addUserAddress
//
// Add User Ethereum Address.
//
// Each User can have several Ethereum addresses.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: TxSuccessResponse
//   400: RestErrorResponse
//
func postAddUserAddressAction(c *gin.Context) {
    request := &models.AddUserAddressRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().AddUserAddress(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

// swagger:route POST /platform/user/status users setUserStatus
//
// Set User Status.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: TxSuccessResponse
//   400: RestErrorResponse
//
func postSetUserStatusAction(c *gin.Context) {
    request := &models.SetStatusRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().SetUserStatus(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

// swagger:route GET /platform/user/:id users getUser
//
// Get User by ID.
//
// Produces:
// - application/json
// Responses:
//   200: UserDataResponse
//   400: RestErrorResponse
//
func getUserAction(c *gin.Context) {
    id := c.Param("id")

    user, err := geth.GetHoquPlatform().GetUser(id)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "user": user,
    })
}
