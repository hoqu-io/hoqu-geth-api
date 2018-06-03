package platform

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/sdk/http/rest"
    "hoqu-geth-api/geth"
    "hoqu-geth-api/geth/models"
    "hoqu-geth-api/sdk/http/middleware"
    "github.com/satori/go.uuid"
)

func InitAdRoutes(routerGroup *gin.RouterGroup) {
    r := routerGroup.Group("/ad")
    {
        r.POST("/add", middleware.SignRequired(), postAddAdAction)
        r.POST("/status", middleware.SignRequired(), postSetAdStatusAction)
        r.GET("/:id", getAdAction)
    }
}

// swagger:route POST /platform/ad/add ads addAd
//
// Add Ad.
//
// The Ad is a campaign of the affiliate to promote particular offer.
// The Ad is created by the affiliate. All widgets placed on affiliate website should be linked to particular Ad.
// The affiliate can specify money receiver address (beneficiary) per each Ad.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: AddSuccessResponse
//   400: RestErrorResponse
//
func postAddAdAction(c *gin.Context) {
    request := &models.AddAdRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    offer, err := geth.GetHoQuStorage().GetOffer(request.OfferId)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    id, err := uuid.NewV2('A')
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    deployParams := &models.DeployAdContractRequest{
        OwnerId: request.OwnerId,
        OfferId: request.OfferId,
        AdId: id.String(),
        BeneficiaryAddress: request.BeneficiaryAddress,
        PayerAddress: offer.PayerAddress,
    }
    contractAddress, tx, err := geth.DeployAdCampaign(deployParams)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    adToStorageReq := &models.AddAdToStorageRequest{
        OwnerId: request.OwnerId,
        OfferId: request.OfferId,
        AdId: id.String(),
        ContractAddress: contractAddress.String(),
    }

    _, err = geth.GetHoquPlatform().AddAd(adToStorageReq)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    _, err = geth.GetHoQuConfig().AddOwner(adToStorageReq.ContractAddress)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.Hash().String(),
        "id": id,
        "contractAddress": contractAddress.String(),
    })
}

// swagger:route POST /platform/ad/status ads setAdStatus
//
// Set Ad Status.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: TxSuccessResponse
//   400: RestErrorResponse
//
func postSetAdStatusAction(c *gin.Context) {
    request := &models.SetStatusRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    tx, err := geth.GetHoquPlatform().SetAdStatus(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "tx": tx.String(),
    })
}

// swagger:route GET /platform/ad/:id ads getAd
//
// Get Ad by ID.
//
// Produces:
// - application/json
// Responses:
//   200: AdDataResponse
//   400: RestErrorResponse
//
func getAdAction(c *gin.Context) {
    id := c.Param("id")

    adstorage, err := geth.GetHoQuStorage().GetAd(id)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    adcontract, err := geth.GetAdCampaign(adstorage.ContractAddress)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    adId, err := adcontract.AdId()
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    offId, err := adcontract.OfferId()
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    affId, err := adcontract.AffiliateId()
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    baddr, err := adcontract.BeneficiaryAddress()
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    paddr, err := adcontract.PayerAddress()
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    ad := &models.AdData{
        AdId: adId.String(),
        CreatedAt: adstorage.CreatedAt,
        OwnerId: affId.String(),
        OfferId: offId.String(),
        ContractAddress: adstorage.ContractAddress,
        BeneficiaryAddress: baddr.String(),
        PayerAddress: paddr.String(),
        Status: adstorage.Status,
    }

    rest.NewResponder(c).Success(gin.H{
        "ad": ad,
    })
}
