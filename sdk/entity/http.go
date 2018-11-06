package entity

import (
    "context"
    "github.com/gin-gonic/gin"
    "hoqu-geth-api/sdk/tracing"
    "hoqu-geth-api/sdk/http/rest"
    "hoqu-geth-api/models"
)

type ServiceHandler func(ctx context.Context, input interface{}, output interface{}) error

func CreateAction(ctx *gin.Context, input interface{}, handler ServiceHandler) {
    output := &models.CreateResponseData{}
    HttpBindAction(ctx, input, output, handler)
}

func SetAction(ctx *gin.Context, input interface{}, handler ServiceHandler) {
    output := &models.TxSuccessData{}
    HttpBindAction(ctx, input, output, handler)
}

func GetByIdAction(ctx *gin.Context, id string, output interface{}, handler ServiceHandler) {
    input := &models.IdRequest{
        Id: id,
    }
    HttpAction(ctx, input, output, handler)
}

func HttpBindAction(ctx *gin.Context, input interface{}, output interface{}, handler ServiceHandler) {
    err := ctx.BindJSON(input)
    if err != nil {
        rest.NewResponder(ctx).ErrorValidation(err.Error())
        return
    }

    HttpAction(ctx, input, output, handler)
}

func HttpAction(ctx *gin.Context, input interface{}, output interface{}, handler ServiceHandler) {
    err := handler(tracing.GetServerSpan(ctx), input, output)
    if err != nil {
        rest.NewResponder(ctx).Error(err.Error())
        return
    }

    rest.NewResponder(ctx).Success(output)
}