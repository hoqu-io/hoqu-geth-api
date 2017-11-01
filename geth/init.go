package geth

import (
	"github.com/spf13/viper"
	"errors"
	"hoqu-geth-api/sdk/geth"
)

func InitGeth() error {
	err := geth.InitWallet(
		viper.GetString("geth.endpoint"),
		viper.GetString("geth.main.key_file"),
		viper.GetString("geth.main.pass"),
	)
	if err != nil {
		return err
	}

	if err := InitToken(); err != nil {
		return err
	}

	if err := InitPrivatePlacement(); err != nil {
		return err
	}

	if tokenAddr, _ := GetPrivatePlacement().TokenAddr(); tokenAddr != GetToken().Address {
		return errors.New("Private Placement token address doesn't match with token address from config")
	}

	if err := InitPresale(); err != nil {
		return err
	}

	return nil
}
