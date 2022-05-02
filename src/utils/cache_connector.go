package utils

import (
	"os"

	"github.com/go-redis/redis"
)

type CacheConnector interface {
	Open()
	Close()
	GetClient() *redis.Client
}

func CacheConnectorHandler() CacheConnector {
	return &cacheConnector{}
}

type cacheConnector struct {
	client *redis.Client
}

func (cacheConnector *cacheConnector) Open() {
	cacheConnector.client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("CACHE_HOST"),
		Password: os.Getenv("CACHE_PASS"),
		DB:       0,
	})
}

func (cacheConnector *cacheConnector) Close() {
	cacheConnector.client.Close()
}

func (cacheConnector *cacheConnector) GetClient() *redis.Client {
	return cacheConnector.client
}
