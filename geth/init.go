package geth

import (
    "github.com/spf13/viper"
    "hoqu-geth-api/sdk/geth"
    "hoqu-geth-api/sdk/geth/metamask"
    "sync"
)

var gOnce sync.Once

func InitGeth() (err error) {
    gOnce.Do(func() {
        err = geth.InitWallet(
            viper.GetString("geth.endpoint"),
            viper.GetString("geth.main.key_file"),
            viper.GetString("geth.main.pass"),
        )
        if err != nil {
            return
        }

        metamask.InitAuth(viper.GetString("geth.meta_auth.banner"))

        if err = InitToken(); err != nil {
            return
        }

        if err = InitPrivatePlacement(); err != nil {
            return
        }

        if err = InitPresale(); err != nil {
            return
        }

        if err = InitSale(); err != nil {
            return
        }

        if err = InitBounty(); err != nil {
            return
        }

        if err = InitClaim(); err != nil {
            return
        }

        if err = InitBurner(); err != nil {
            return
        }

        if err = initHoQuConfig(); err != nil {
            return
        }

        if err = initHoQuStorage(); err != nil {
            return
        }

        if err = initHoQuRater(); err != nil {
            return
        }

        if err = initHoquPlatform(); err != nil {
            return
        }

        if err = initHoQuTransactor(); err != nil {
            return
        }
    })
    return
}
