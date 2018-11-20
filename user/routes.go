package user

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/models"
    "hoqu-geth-api/sdk/http/middleware"
    "hoqu-geth-api/sdk/entity"
)

func InitUserRoutes(routerGroup *gin.RouterGroup) {
    r := routerGroup.Group("/user")
    {
        r.POST("/register", middleware.SignRequired(), postRegisterUserAction)
        r.POST("/status", middleware.SignRequired(), postSetUserStatusAction)
        r.POST("/address", middleware.SignRequired(), postAddUserAddressAction)
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
func postRegisterUserAction(ctx *gin.Context) {
    entity.CreateAction(ctx, &models.RegisterUserRequest{}, GetService().Create)
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
func postSetUserStatusAction(ctx *gin.Context) {
    entity.SetAction(ctx, &models.SetStatusRequest{}, GetService().SetStatus)
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
func postAddUserAddressAction(ctx *gin.Context) {
    entity.SetAction(ctx, &models.AddUserAddressRequest{}, GetService().AddAddress)
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
func getUserAction(ctx *gin.Context) {
    id := ctx.Param("id")
    entity.GetByIdAction(ctx, id, &models.UserSuccessData{}, GetService().GetById)
}
