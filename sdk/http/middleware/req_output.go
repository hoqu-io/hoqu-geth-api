package middleware

import (
	"github.com/gin-gonic/gin"
)

func ReqOutput(ctx *gin.Context) {
	ctx.Next()

	if done, exists := ctx.Get("done"); exists && done != nil {
		close(done.(chan bool))
	}
}
