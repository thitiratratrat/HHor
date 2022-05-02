package utils

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/thitiratratrat/hhor/src/constant"
)

func SaveToCache(cacheClient *redis.Client, nameSpace constant.CacheNameSpace, key string, value interface{}) {
	valueMarshal, marshalErr := json.Marshal(value)

	if marshalErr == nil {
		cacheErr := cacheClient.Set(fmt.Sprintf("%s:%s", nameSpace, key), valueMarshal, 60*time.Second).Err()

		if cacheErr != nil {
			logrus.Error("Unable to set cache", cacheErr)
		}
	}
}

func DeleteCache(cacheClient *redis.Client, nameSpace constant.CacheNameSpace, key string) {
	pipe := cacheClient.Pipeline()

	pipe.Del(fmt.Sprintf("%s:%s", constant.Dorm, key))
	pipe.Exec()
}
