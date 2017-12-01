package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "hoqu-geth-api/sdk/http/rest"
    "github.com/spf13/viper"
    "fmt"
)

const SIGN_HEADER = "X-Sign"

func SignRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        signHeader := c.GetHeader(SIGN_HEADER)
        if signHeader == "" {
            rest.NewResponder(c).ErrorAuthorization()
            return
        }

        if signHeader != viper.GetString("security.sign") {
            rest.NewResponder(c).ErrorWithCode(
                http.StatusUnauthorized,
                rest.ErrorAuthorization,
                fmt.Sprintf(
                    "Incorrect signature, provided %v, expected %v",
                    signHeader,
                    viper.GetString("security.sign"),
                ),
            )
            return
        }
    }
}
