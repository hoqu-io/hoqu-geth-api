package tariff_group

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/models"
    "hoqu-geth-api/sdk/http/middleware"
    "hoqu-geth-api/sdk/entity"
)

func InitTariffGroupRoutes(routerGroup *gin.RouterGroup) {
    r := routerGroup.Group("/tariff_group")
    {
        r.POST("/add", middleware.SignRequired(), postAddTariffGroupAction)
        r.POST("/status", middleware.SignRequired(), postSetTariffGroupStatusAction)
        r.GET("/:id", getTariffGroupAction)
    }
}

// swagger:route POST /platform/tariff_group/add tariff_groups addTariffGroup
//
// Add TariffGroup.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: AddSuccessResponse
//   400: RestErrorResponse
//
func postAddTariffGroupAction(ctx *gin.Context) {
    entity.CreateAction(ctx, &models.AddTariffGroupRequest{}, GetService().Create)
}

// swagger:route POST /platform/tariff_group/status tariff_groups setTariffGroupStatus
//
// Set TariffGroup Status.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: TxSuccessResponse
//   400: RestErrorResponse
//
func postSetTariffGroupStatusAction(ctx *gin.Context) {
    entity.SetAction(ctx, &models.SetStatusRequest{}, GetService().SetStatus)
}

// swagger:route GET /platform/tariff_group/:id tariff_groups getTariffGroup
//
// Get TariffGroup by ID.
//
// Produces:
// - application/json
// Responses:
//   200: TariffGroupDataResponse
//   400: RestErrorResponse
//
func getTariffGroupAction(ctx *gin.Context) {
    id := ctx.Param("id")
    entity.GetByIdAction(ctx, id, &models.TariffGroupSuccessData{}, GetService().GetById)
}
