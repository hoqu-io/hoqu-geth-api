package middleware

import (
    "time"

    "github.com/sirupsen/logrus"
    "github.com/gin-gonic/gin"
    "github.com/spf13/cast"
)

func ExecutionTime(ctx *gin.Context) {
    done, has := ctx.Get("done")
    if !has {
        return
    }

    ctx.Set("startTime", time.Now().Format(time.RFC3339Nano))
    go tickExecTime(ctx, done.(chan bool))
    ctx.Next()
}

func tickExecTime(ctx *gin.Context, done chan bool) {
    for {
        select {
        case <-done:
            startTime, _ := ctx.Get("startTime")
            st, err := time.Parse(time.RFC3339, cast.ToString(startTime))
            if err == nil {
                logrus.Infof(
                    "[%s] %s %s %d Executed withing %f seconds",
                    time.Now().Format("02.01.2006 15:04:05"),
                    ctx.Request.RequestURI,
                    ctx.Request.Method,
                    ctx.Writer.Status(),
                    time.Now().Sub(st).Seconds(),
                )
            }
            return
        case <-time.After(300 * time.Second):
            logrus.Warnf(
                "[%s] %s %s %d Executing too long",
                time.Now().Format("02.01.2006 15:04:05"),
                ctx.Request.RequestURI,
                ctx.Request.Method,
                ctx.Writer.Status(),
            )
            return
        }
    }
}