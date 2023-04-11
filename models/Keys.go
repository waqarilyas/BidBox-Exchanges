package models

import (
	"context"
	"errors"
	"fmt"
	"html"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/adshao/go-binance/v2"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Key struct {
	Keyid     uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()" json:"key_id"`
	Uid       string    `gorm:"size:255" json:"uid"`
	Service   string    `gorm:"size:255;not null" json:"service"`
	ApiKey    string    `gorm:"not null;unique" json:"api_key"`
	SecretKey string    `gorm:"not null;unique" json:"secret_key"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func Verify(hashedpassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
}

func (u *Key) BeforeSave() error {
	hashapi, err := Hash(u.ApiKey)
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "models/Keys.go",
			"function": "BeforeSave",
		}).Error("Error hashing API Key")
		return err
	}
	hashsecret, err := Hash(u.SecretKey)
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "models/Keys.go",
			"function": "BeforeSave",
		}).Error("Error hashing Secret Key")
		return err
	}
	u.ApiKey = string(hashapi)
	u.SecretKey = string(hashsecret)
	log.WithFields(log.Fields{
		"file":     "models/Keys.go",
		"function": "BeforeSave",
	}).Info("Hashed Keys")
	return nil
}

func (u *Key) Prepare() {
	u.ApiKey = html.EscapeString(strings.TrimSpace(u.ApiKey))
	u.SecretKey = html.EscapeString(strings.TrimSpace(u.SecretKey))
}

var services = [...]string{
	"Binance",
	"Bitget",
	"OKX",
}

func (u *Key) Validate() error {
	if u.Uid == "" {
		log.WithFields(log.Fields{
			"file":     "models/Keys.go",
			"function": "Validate",
		}).Error("Validation Error - user id required")
		return errors.New("uid required")
	}
	if u.Service == "" {
		log.WithFields(log.Fields{
			"file":     "models/Keys.go",
			"function": "Validate",
		}).Error("Validation Error - service required")
		return errors.New("service required")
	}
	if u.SecretKey == "" {
		log.WithFields(log.Fields{
			"file":     "models/Keys.go",
			"function": "Validate",
		}).Error("Validation Error - secret key required")
		return errors.New("secret key required")
	}
	if u.ApiKey == "" {
		log.WithFields(log.Fields{
			"file":     "models/Keys.go",
			"function": "Validate",
		}).Error("Validation Error - api key required")
		return errors.New("api key required")
	}
	found := false
	for _, v := range services {
		if strings.ToLower(u.Service) == v {
			found = true
			fmt.Println(found)
			break
		}
	}
	/*if !found {
		return errors.New("service not found")
	}*/
	client := binance.NewClient(u.ApiKey, u.SecretKey)
	_, err := client.NewListPricesService().Do(context.Background())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"file":     "models/Keys.go",
		"function": "Validate",
	}).Info("Validation passed")
	return nil
}

/*
func (u *Key) TestforBinance() error {

}*/

func (u *Key) SaveKey(db *gorm.DB) (*Key, error) {
	err := db.Debug().Create(&u).Error
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "models/Keys.go",
			"function": "SaveKey",
		}).Error("Database Error - could not save key")
		return &Key{}, err
	}
	log.WithFields(log.Fields{
		"file":     "models/Keys.go",
		"function": "SaveKey",
	}).Info("User saved to database successfully")
	return u, nil
}

func (u *Key) FindAllKeys(db *gorm.DB) (*[]Key, error) {
	Keys := []Key{}
	err := db.Debug().Model(&Key{}).Limit(100).Find(&Keys).Error
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "models/Keys.go",
			"function": "FindAllKeys",
		}).Error("Database Error - could not get keys")
		return &[]Key{}, err
	}
	log.WithFields(log.Fields{
		"file":     "models/Keys.go",
		"function": "FindAllKeys",
	}).Info("Successfully retrieved keys")
	return &Keys, nil
}

func (u *Key) FindKeyById(db *gorm.DB, kid uuid.UUID) (*Key, error) {
	err := db.Debug().Model(Key{}).Where("keyid = ?", kid).Take(&u).Error
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "models/Keys.go",
			"function": "FindKeyById",
		}).Error("Database Error - could not get key ny id")
		return &Key{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		log.WithFields(log.Fields{
			"file":     "models/Keys.go",
			"function": "FindKeyById",
		}).Info("Database - key not found")
		return &Key{}, errors.New("Key not found")
	}
	log.WithFields(log.Fields{
		"file":     "models/Keys.go",
		"function": "FindKeyById",
	}).Info("Successfully got key by id")
	return u, nil
}
