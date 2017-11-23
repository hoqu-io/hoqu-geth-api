package app

import (
    "os"
)

const (
    ENVNAME     string = "APPLICATION_ENV"
    PRONET_ENV         = "pro_net"
    TESTNET_ENV        = "test_net"
    DEFAULT_ENV string = TESTNET_ENV
)

var env string

func init() {
    env = getAppEnv(DEFAULT_ENV)
}

func Env() string {
    return env
}

func getAppEnv(def string) (env string) {
    env = os.Getenv(ENVNAME)

    if env = os.Getenv(ENVNAME); env == "" {
        env = def
    }

    return
}
