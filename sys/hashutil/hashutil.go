package hashutil

import (
	"github.com/nridwan/config/configutil"
	"github.com/speps/go-hashids/v2"
)

var hasher *hashids.HashID = nil

func LoadConfiguration() {
	hashData := hashids.NewData()
	hashData.Salt = configutil.Getenv("HASH_SECRET", "")
	hashData.MinLength = 30
	hasher, _ = hashids.NewWithData(hashData)
}

func EncodeSingle(id int64) (string, error) {
	return hasher.EncodeInt64([]int64{id})
}

func DecodeSingle(hash string) (int64, error) {
	data, err := hasher.DecodeInt64WithError(hash)
	if err != nil {
		return 0, err
	}
	return data[0], err
}
