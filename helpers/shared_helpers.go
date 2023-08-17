package helpers

import (
	"github.com/jinzhu/gorm"
	"github.com/kryptomind/crons/models"
)

var IS_TESTNET bool = true
var PRODUCT_TYPE string = "UMCBL"
var MARGIN_COIN string = "USDT"

func InitSharedData(db *gorm.DB) {
	var set models.Settings
	settings, err := set.GetSettings(db)
	if err != nil {

		return
	}
	IS_TESTNET = settings.IsTestnet

	if IS_TESTNET {
		PRODUCT_TYPE = "SUMCBL"
		MARGIN_COIN = "SUSDT"
	} else {
		PRODUCT_TYPE = "UMCBL"
		MARGIN_COIN = "USDT"
	}

}
