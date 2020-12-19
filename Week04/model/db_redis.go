package model

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var redisDB *redis.Client

const redisPrefix = "ginbro:"

func CreateRedis(redisAddr, redisPassword string, idx int) {
	//initializing redis client
	redisDB = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword, // no password set
		DB:       idx,           // use default DB
	})
	if pong, err := redisDB.Ping().Result(); err != nil || pong != "PONG" {
		logrus.WithError(err).Fatal("could not connect to the redis server")
	}

}
