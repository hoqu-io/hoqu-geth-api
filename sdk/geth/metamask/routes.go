package metamask

import (
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/sdk/http/rest"
    "github.com/ethereum/go-ethereum/common"
)

func AddMetamaskAuthRoutes(gin *gin.Engine) {
    router := gin.Group("/metamask")
    {
        router.POST("challenge", postCreateChallenge)
        router.POST("recover", postRecoverSigner)
    }
}

// swagger:route POST /metamask/challenge metamask createChallenge
//
// Create challenge
//
// Create metamask challenge and return challenge data array
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: ChallengeSuccessResponse
//   400: RestErrorResponse
//
func postCreateChallenge(c *gin.Context) {
    request := &ChallengeRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    challenge, err := GetAuth().CreateChallenge(
        common.HexToAddress(request.Address),
    )
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "result": challenge,
    })
}

// swagger:route POST /metamask/recover metamask recoverSigner
//
// Recover signer
//
// Recover metamask challenge signer from result signature
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: RecoverSuccessResponse
//   400: RestErrorResponse
//
func postRecoverSigner(c *gin.Context) {
    request := &RecoverRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    addr, err := GetAuth().RecoverSignerAddr(
        request.Challenge,
        request.Signature,
    )
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "address": addr,
    })
}
