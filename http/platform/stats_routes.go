package platform

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/sdk/http/rest"
    "hoqu-geth-api/geth"
)

func InitStatsRoutes(routerGroup *gin.RouterGroup) {
    r := routerGroup.Group("/stats")
    {
        r.GET("/:id", getStatsAction)
    }
}

// swagger:route GET /platform/stats/:id stats getStats
//
// Get Stats by ID.
//
// Produces:
// - application/json
// Responses:
//   200: StatsDataResponse
//   400: RestErrorResponse
//
func getStatsAction(c *gin.Context) {
    id := c.Param("id")

    stats, err := geth.GetHoQuStorage().GetStats(id)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "stats": stats,
    })
}
