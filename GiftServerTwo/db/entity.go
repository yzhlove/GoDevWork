package db

import (
	"WorkSpace/GoDevWork/GiftServerTwo/obj"
	"github.com/garyburd/redigo/redis"
	"time"
)

func GetAutoId() (uint32, error) {
	c := RedisClient.Get()
	defer c.Close()

	auto, err := redis.Int(c.Do("GET", getStrKey("AutoID")))
	if err != nil {
		if err == redis.ErrNil {
			auto = 0
		}
	}
	return uint32(auto), err
}

func GetCodes() (map[uint32]*obj.Code, error) {

	c := RedisClient.Get()
	defer c.Close()

	result, err := redis.ByteSlices(c.Do("HVALS", getStrKey("Codes")))
	if err != nil {
		return nil, err
	}

	ts := time.Now().Unix()
	codes := make(map[uint32]*obj.Code, len(result))
	expireIds := make([]uint32, 0, len(result))

	for _, data := range result {
		code := &obj.Code{}
		if _, err = code.UnmarshalMsg(data); err != nil {
			return nil, err
		}
		if code.EndTime < ts {
			codes[code.Id] = code
		} else {
			expireIds = append(expireIds, code.Id)
		}
	}
	if len(expireIds) > 0 {
		go DelExpireCodes(expireIds)
	}
	return codes, nil
}

func DelExpireCodes(ids []uint32) (err error) {
	c := RedisClient.Get()
	defer c.Close()

	_, err = c.Do("HMDEL", redis.Args{}.Add(getStrKey("Codes")).AddFlat(ids)...)
	return
}
