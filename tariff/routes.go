package tariff

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/models"
    "hoqu-geth-api/sdk/http/middleware"
    "hoqu-geth-api/sdk/entity"
)

func InitTariffRoutes(routerGroup *gin.RouterGroup) {
    r := routerGroup.Group("/tariff")
    {
        r.POST("/add", middleware.SignRequired(), postAddTariffAction)
        r.POST("/status", middleware.SignRequired(), postSetTariffStatusAction)
        r.GET("/:id", getTariffAction)
    }
}

// swagger:route POST /platform/tariff/add tariffs addTariff
//
// Add Tariff.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: AddSuccessResponse
//   400: RestErrorResponse
//
func postAddTariffAction(ctx *gin.Context) {
    entity.CreateAction(ctx, &models.AddTariffRequest{}, GetService().Create)
}

// swagger:route POST /platform/tariff/status tariffs setTariffStatus
//
// Set Tariff Status.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: TxSuccessResponse
//   400: RestErrorResponse
//
func postSetTariffStatusAction(ctx *gin.Context) {
    entity.SetAction(ctx, &models.SetStatusRequest{}, GetService().SetStatus)
}

// swagger:route GET /platform/tariff/:id tariffs getTariff
//
// Get Tariff by ID.
//
// Produces:
// - application/json
// Responses:
//   200: TariffDataResponse
//   400: RestErrorResponse
//
func getTariffAction(ctx *gin.Context) {
    id := ctx.Param("id")
    entity.GetByIdAction(ctx, id, &models.TariffSuccessData{}, GetService().GetById)
}
