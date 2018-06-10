package platform

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/sdk/http/rest"
    "hoqu-geth-api/geth"
    "hoqu-geth-api/geth/models"
    "hoqu-geth-api/sdk/http/middleware"
)

func InitTrackerRoutes(routerGroup *gin.RouterGroup) {
    r := routerGroup.Group("/tracker")
    {
        r.POST("/register", middleware.SignRequired(), postRegisterTrackerAction)
        r.POST("/status", middleware.SignRequired(), postSetTrackerStatusAction)
        r.GET("/:id", getTrackerAction)
    }
}

// swagger:route POST /platform/tracker/register trackers registerTracker
//
// Register Tracker.
//
// The Tracker is an application which receives visitor's requests to show the appropriate widget.
// All widgets can be created in terms of appropriate Ad (affiliate promotion campaign).
// If visitor interacts with a widget the Tracker creates the Lead and push it to Blockchain.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: AddSuccessResponse
//   400: RestErrorResponse
//
func postRegisterTrackerAction(c *gin.Context) {
    request := &models.RegisterTrackerRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, id, err := geth.GetHoquPlatform().RegisterTracker(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
        "id": id,
    })
}

// swagger:route POST /platform/tracker/status trackers setTrackerStatus
//
// Set Tracker Status.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: TxSuccessResponse
//   400: RestErrorResponse
//
func postSetTrackerStatusAction(c *gin.Context) {
    request := &models.SetStatusRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().SetTrackerStatus(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

// swagger:route GET /platform/tracker/:id trackers getTracker
//
// Get Tracker by ID.
//
// Produces:
// - application/json
// Responses:
//   200: TrackerDataResponse
//   400: RestErrorResponse
//
func getTrackerAction(c *gin.Context) {
    id := c.Param("id")

    tracker, err := geth.GetHoQuStorage().GetTracker(id)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tracker": tracker,
    })
}
