package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func ReqInput(ctx *gin.Context) {
	rid, _ := uuid.NewV4()
	ctx.Set("rid", rid.String())
	ctx.Set("done", make(chan bool))
}
