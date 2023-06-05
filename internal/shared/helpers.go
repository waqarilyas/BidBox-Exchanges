package shared

import (
	"errors"

	"github.com/kryptomind/bidboxapi/KeyService/helpers"
	"github.com/kryptomind/bidboxapi/KeyService/models"
)

func DecryptUserKeys(keys *models.Key) (*DecryptedKeys, error) {

	api_key, err := helpers.DecryptStrings(keys.ApiKey)
	if err != nil {
		return nil, errors.New("unable to decrypt api key")
	}

	api_secret, err := helpers.DecryptStrings(keys.SecretKey)
	if err != nil {
		return nil, errors.New("unable to decrypt api secret")

	}

	dKeys := &DecryptedKeys{
		ApiKey: api_key,
		Secret: api_secret,
	}

	if keys.Passphrase != "" {
		passphrase, err := helpers.DecryptStrings(keys.Passphrase)
		if err != nil {
			return nil, errors.New("unable to decrypt passphrase")
		}

		dKeys.Passphrase = passphrase

	}

	return dKeys, nil
}
