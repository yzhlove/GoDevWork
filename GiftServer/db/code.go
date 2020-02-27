package db

import (
	"WorkSpace/GoDevWork/GiftServer/misc/helper"
	"WorkSpace/GoDevWork/GiftServer/obj"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

func SetCodeInfo(code *obj.CodeInfo) error {

	c := RedisClient.Get()
	defer c.Close()

	data, err := code.MarshalMsg(nil)
	if err != nil {
		return err
	}
	if err := c.Send("HSET", buildKey("CodeInfoList"), code.Id, data); err != nil {
		return err
	}
	if err := incrMaxId(c); err != nil {
		return err
	}
	_ = c.Flush()
	for i := 0; i < 2; i++ {
		if _, err := c.Receive(); err != nil {
			return err
		}
	}
	return nil
}

func GetCodeInfoList() (map[uint32]*obj.CodeInfo, error) {

	c := RedisClient.Get()
	defer c.Close()

	data, err := redis.ByteSlices(c.Do("HVALS", buildKey("CodeInfoList")))
	if err != nil {
		return nil, err
	}
	ts := time.Now().Unix()
	codes := make(map[uint32]*obj.CodeInfo, len(data))
	ids := make([]uint32, 0, len(codes))
	for _, v := range data {
		code := &obj.CodeInfo{}
		if _, err = code.UnmarshalMsg(v); err != nil {
			return nil, err
		}
		if code.EndTime < ts {
			codes[code.Id] = code
		} else {
			ids = append(ids, code.Id)
		}
	}
	if len(ids) > 0 {
		go DelCodeInfoList(ids)
	}
	return codes, nil
}

func DelCodeInfoList(ids []uint32) error {
	c := RedisClient.Get()
	defer c.Close()

	_, err := c.Do("HMDEL", redis.Args{}.Add(buildKey("CodeInfoList")).AddFlat(ids)...)
	return err
}

func SetCodes(id uint32, codes map[int][]string) error {
	c := RedisClient.Get()
	defer c.Close()
	for key, value := range codes {
		status := make(map[string]int, len(value))
		for _, v := range value {
			status[v] = 0
		}
		if _, err := c.Do("HMSET", redis.Args{}.Add(buildBucketKey("Codes", id, strconv.Itoa(key))).AddFlat(status)...); err != nil {
			return err
		}
	}
	return nil
}

func ExistsCode(id uint32, code string) (bool, error) {
	c := RedisClient.Get()
	defer c.Close()

	i := strconv.Itoa(helper.GetBucketTop(code))
	return redis.Bool(c.Do("HEXISTS", buildBucketKey("Codes", id, i)))
}

func GetCodeTimes(id uint32, code string) (int, error) {
	c := RedisClient.Get()
	defer c.Close()

	i := strconv.Itoa(helper.GetBucketTop(code))
	return redis.Int(c.Do("HGET", buildBucketKey("Codes", id, i), code))
}

func IsUseCode(uid uint64, id uint32, code string) (bool, error) {
	c := RedisClient.Get()
	defer c.Close()

	sid := strconv.FormatUint(uint64(id), 10)
	return redis.Bool(c.Do("SISMEMBER", buildUserKey("UserCodes", uid, sid), code))
}

func SetUseCode(uid uint64, id uint32, code string) error {
	c := RedisClient.Get()
	defer c.Close()

	i := strconv.Itoa(helper.GetBucketTop(code))
	sid := strconv.FormatUint(uint64(id), 10)
	if err := c.Send("HINCRBY", buildBucketKey("Codes", id, i), 1); err != nil {
		return err
	}
	if err := c.Send("SADD", buildUserKey("UserCodes", uid, sid), code); err != nil {
		return err
	}
	_ = c.Flush()
	for i := 0; i < 2; i++ {
		if _, err := c.Receive(); err != nil {
			return err
		}
	}
	return nil
}
