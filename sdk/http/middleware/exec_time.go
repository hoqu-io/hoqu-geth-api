package middleware

import (
    "time"

    "github.com/sirupsen/logrus"
    "github.com/gin-gonic/gin"
)

func ExecutionTime() gin.HandlerFunc {
    return func (c *gin.Context) {
        st := gin.Param{
            Key: "startTime",
            Value: time.Now().Format(time.RFC3339Nano),
        }

        done := make(chan bool)

        c.Set("done", done)
        c.Params = append(c.Params, st)
        c.Next()

        go func () {
            for {
                select {
                case <-done:
                    st, err := time.Parse(time.RFC3339, c.Param("startTime"))
                    if err == nil {
                        logrus.Warnf(
                            "%s %s %d Executed withing %f seconds",
                            c.Request.RequestURI,
                            c.Request.Method,
                            c.Writer.Status(),
                            time.Now().Sub(st).Seconds(),
                        )
                    }
                    return
                case <-time.After(3 * time.Second):
                    logrus.Warnf(
                    "%s %s %d Executing too long",
                        c.Request.RequestURI,
                        c.Request.Method,
                        c.Writer.Status(),
                    )
                    return
                }
            }
        }()
    }
}