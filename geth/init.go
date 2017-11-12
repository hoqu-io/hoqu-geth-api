package geth

import (
	"github.com/spf13/viper"
	"hoqu-geth-api/sdk/geth"
	"fmt"
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
		return fmt.Errorf(
			"Private Placement token address %v doesn't match with token address %v from config",
			tokenAddr,
			GetToken().Address,
		)
	}

	if err := InitPresale(); err != nil {
		return err
	}

	return nil
}
