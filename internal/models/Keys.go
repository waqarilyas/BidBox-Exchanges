package models

import (
	"errors"
	"html"
	"strings"

	"github.com/kryptomind/bidboxapi/KeyService/helpers"
	//	log "github.com/sirupsen/logrus"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)
type UpdateLeverage struct {
	Category     string `json:"category"`
	Symbol       string `json:"symbol"`
	BuyLeverage  string `json:"buyLeverage"`
	SellLeverage string `json:"sellLeverage"`
}
type UpdateLeverageBitget struct {
	Symbol       string `json:"symbol"`
	MarginCoin   string `json:"marginCoin"`
	Leverage     string `json:"leverage"`
}
type UpdateLeverageBinance struct {
	Symbol       string `json:"symbol"`
	Timestamp    int64 `json:"timestamp"`
	Leverage     int `json:"leverage"`
}
type Key struct {
	Keyid      uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()" json:"key_id"`
	Uid        string    `gorm:"null;size:255" json:"uid"`
	Service    string    `gorm:"size:255;not null" json:"service"`
	ApiKey     string    `gorm:"not null;unique" json:"api_key"`
	SecretKey  string    `gorm:"not null;unique" json:"secret_key"`
	Passphrase string    `gorm:"" json:"passphrase"`
	UserEmail  string    `gorm:"" json:"user_email"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func Verify(hashedpassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
}

func (u *Key) BeforeSave() error {
	hashapi, err := helpers.EncryptStrings(u.ApiKey)
	if err != nil {
		return err
	}

	hashsecret, err := helpers.EncryptStrings(u.SecretKey)
	if err != nil {
		return err
	}

	if u.Passphrase != "" {
		hashpassphrase, err := helpers.EncryptStrings(u.Passphrase)
		if err != nil {
			return err
		}
		u.Passphrase = string(hashpassphrase)
	}
	u.ApiKey = string(hashapi)
	u.SecretKey = string(hashsecret)
	return nil
}

func (u *Key) Prepare() {
	u.ApiKey = html.EscapeString(strings.TrimSpace(u.ApiKey))
	u.SecretKey = html.EscapeString(strings.TrimSpace(u.SecretKey))
	u.Passphrase = html.EscapeString(strings.TrimSpace(u.Passphrase))
}

var services = []string{
	"Binance",
	"Bitget",
	"OKEX",
}

func (u *Key) Validate() error {

	if u.Service == "bitget" && u.Passphrase == "" {
		return errors.New("passphrase is required")
	}
	if u.Service == "" {
		return errors.New("service required")
	}
	if u.SecretKey == "" {
		return errors.New("secret_key required")
	}
	if u.ApiKey == "" {
		return errors.New("api_key required")
	}
	if u.UserEmail == "" {
		return errors.New("user_email is required")
	}

	// found := false
	// for _, v := range services {
	// 	if strings.ToLower(u.Service) == v {
	// 		found = true
	// 		fmt.Println(found)
	// 		break
	// 	}
	// }
	// if !found {
	// 	return errors.New("service not found")
	// }
	// client := binance.NewClient(u.ApiKey, u.SecretKey)
	// _, err := client.NewListPricesService().Do(context.Background())
	// if err != nil {
	// 	return err
	// }
	return nil
}

/*
func (u *Key) TestforBinance() error {

}*/

func (u *Key) SaveKey(db *gorm.DB) (*Key, error) {
	err := db.Debug().Create(&u).Error
	if err != nil {
		return &Key{}, err
	}
	return u, nil
}

func (u *Key) FindAllKeys(db *gorm.DB) (*[]Key, error) {
	Keys := []Key{}
	err := db.Debug().Model(&Key{}).Limit(100).Find(&Keys).Error
	if err != nil {
		return &[]Key{}, err
	}
	return &Keys, nil
}

func (u *Key) FindKeyByEmailAndService(db *gorm.DB, service string, email string) (*Key, error) {
	err := db.Debug().Model(Key{}).Where("service = ? AND user_email= ?", service, email).Take(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *Key) FindKeyById(db *gorm.DB, kid uuid.UUID) (*Key, error) {
	err := db.Debug().Model(Key{}).Where("keyid = ?", kid).Take(&u).Error
	if err != nil {
		return &Key{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Key{}, errors.New("Key not found")
	}
	return u, nil
}

func (u *Key) FindKeyByEmail(db *gorm.DB, email string) ([]Key, error) {
	Keys := []Key{}
	err := db.Debug().Model(&Key{}).Where("user_email= ?", email).Find(&Keys).Error
	if err != nil {
		return []Key{}, err
	}
	return Keys, nil
}