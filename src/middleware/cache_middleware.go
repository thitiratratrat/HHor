package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/thitiratratrat/hhor/src/constant"
	"github.com/thitiratratrat/hhor/src/utils"
)

func VerifyCache(client *redis.Client, nameSpace constant.CacheNameSpace) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		var value []byte
		var err error

		if len(id) > 0 {
			value, err = client.Get(fmt.Sprintf("%s:%s", nameSpace, id)).Bytes()
		} else {
			value, err = client.Get(fmt.Sprintf("%s:%s", nameSpace, context.Param("userid"))).Bytes()
		}

		if err != nil {
			logrus.Info("Data not found in cache")
			context.Next()

			return
		}

		data := utils.ToJson(value)

		logrus.Info("Data found in cache")
		context.AbortWithStatusJSON(http.StatusOK, data)
	}
}
