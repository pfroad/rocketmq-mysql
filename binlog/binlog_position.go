package binlog

import "github.com/jinzhu/gorm"

type BinlogPositionManager struct {
	BinlogFile  string
	NextPostion int64
	DB          *gorm.DB
}
