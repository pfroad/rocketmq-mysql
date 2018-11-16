package schema

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"github.com/pfroad/rocketmq-mysql/config"
)

var GormDB *gorm.DB

func InitGormDB() error {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName[:len(defaultTableName)-1]
	}

	db := config.GetConfig().DB
	glog.Infof("Init db %s \n", db)

	var err error
	GormDB, err = gorm.Open("mysql", db)
	if err != nil {
		glog.Errorf("Cannot connect to BD %s: %v\n", db, err)
		return err
	}

	GormDB.DB().SetMaxIdleConns(2)
	GormDB.DB().SetMaxOpenConns(2)

	return nil
}
