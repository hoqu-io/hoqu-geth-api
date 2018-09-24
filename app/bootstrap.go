package app

import (
    sdkApp "hoqu-geth-api/sdk/app"
    "hoqu-geth-api/geth"
    "github.com/sirupsen/logrus"
    "hoqu-geth-api/sdk/tracing"
    "hoqu-geth-api/offer"
    "hoqu-geth-api/tariff"
    "hoqu-geth-api/ad_campaign"
    "hoqu-geth-api/lead"
    "hoqu-geth-api/user"
    "hoqu-geth-api/tariff_group"
)

func Bootstrap() {
    sdkApp.InitConfig()
    sdkApp.InitLogger()
    if err := geth.InitGeth(); err != nil {
        logrus.Fatal(err)
    }

    if err := tracing.Init(); err != nil {
        logrus.Fatal(err)
    }

    if err := user.InitService(); err != nil {
        logrus.Fatal(err)
    }

    if err := offer.InitService(); err != nil {
        logrus.Fatal(err)
    }

    if err := ad_campaign.InitService(); err != nil {
        logrus.Fatal(err)
    }

    if err := lead.InitService(); err != nil {
        logrus.Fatal(err)
    }

    if err := tariff_group.InitService(); err != nil {
        logrus.Fatal(err)
    }

    if err := tariff.InitService(); err != nil {
        logrus.Fatal(err)
    }
}