package db

import (
	"WorkSpace/GoDevWork/GiftServerTwo/config"
	"github.com/garyburd/redigo/redis"
	log "github.com/sirupsen/logrus"
	"runtime"
	"time"
)

var (
	RedisClient *redis.Pool
)

func Init() error {
	RedisClient = &redis.Pool{
		MaxIdle:     10 * runtime.NumCPU(),
		MaxActive:   50 * runtime.NumCPU(),
		IdleTimeout: 60 * time.Second,
		Wait:        true,
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial(
				"tcp",
				config.RedisHost,
				redis.DialDatabase(config.RedisDb),
				redis.DialConnectTimeout(2*time.Second),
			)
		},
	}
	for {
		conn := RedisClient.Get()
		if err := conn.Err(); err != nil {
			log.Info(err)
			conn.Close()
			time.Sleep(2 * time.Second)
			continue
		}
		for {
			if _, err := conn.Do("PING"); err != nil {
				log.Info("wait for redis online")
				time.Sleep(2 * time.Second)
			} else {
				break
			}
		}
		conn.Close()
	}
}
