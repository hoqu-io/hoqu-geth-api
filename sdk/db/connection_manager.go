package db

import (
    "hoqu-geth-api/sdk/app"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/jinzhu/gorm"
)

type ConnectionManager struct {
    dsn string
    driver string
}

func NewConnectionManager(driver, dsn string) *ConnectionManager {
    return &ConnectionManager{
        dsn: dsn,
        driver: driver,
    }
}

func (st *ConnectionManager) NewConnection() (*gorm.DB, error) {
    db, err := gorm.Open(st.driver, st.dsn)

    if app.Env() == app.TESTNET_ENV {
        db.LogMode(true)
    }

    return db, err
}
