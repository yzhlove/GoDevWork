package db

import "github.com/garyburd/redigo/redis"

const minAutoIncrement = 64

func GetMaxId() (max int, err error) {
	c := RedisClient.Get()
	defer c.Close()

	var ok bool
	if ok, err = redis.Bool(c.Do("EXISTS", buildKey("AutoIncrement"))); err != nil {
		return
	}
	if ok {
		max, err = redis.Int(c.Do("GET", buildKey("AutoIncrement")))
	} else {
		max = minAutoIncrement
		_, err = c.Do("SET", buildKey("AutoIncrement"), minAutoIncrement)
	}
	return
}

func incrMaxId(c redis.Conn) error {
	return c.Send("INCR", buildKey("AutoIncrement"))
}
