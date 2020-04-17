package db

import (
	"WorkSpace/GoDevWork/GiftServer/misc/helper"
	"WorkSpace/GoDevWork/GiftServerTwo/obj"
	"github.com/garyburd/redigo/redis"
)

func Save(base *obj.Code, codeUsers obj.CodeUsers, uid uint64, code string) error {

	c := RedisClient.Get()
	defer c.Close()

	base.Used++
	codeData, err := base.MarshalMsg(nil)
	if err != nil {
		return err
	}
	c.Send("HSET", getStrKey("Codes"), base.Id, codeData)

	top := helper.GetBucketTop(code)
	codeUserData, err := codeUsers.MarshalMsg(nil)
	if err != nil {
		return err
	}
	c.Send("HSET", getSpecKey("UseCodes", uint64(base.Id), top), code, codeUserData)
	c.Send("SADD", getSpecKey("Users", uid, int(base.Id)), code)

	if err = c.Flush(); err != nil {
		return err
	}

	for i := 0; i < 3; i++ {
		if _, err = c.Receive(); err != nil {
			return err
		}
	}
	return nil
}

func SetUseCodes(id uint32, group map[int][]string) error {
	c := RedisClient.Get()
	defer c.Close()

	data, _ := obj.CodeUser{}.MarshalMsg(nil)
	for top, codes := range group {
		if len(codes) > 0 {
			values := make(map[string][]byte, len(codes))
			for _, code := range codes {
				values[code] = data
			}
			if _, err := c.Do("HMSET", redis.Args{}.Add(getSpecKey("UseCodes", uint64(id), top)).AddFlat(values)...); err != nil {
				return err
			}
		}
	}
	return nil
}

func SetCode(code *obj.Code) error {

	c := RedisClient.Get()
	defer c.Close()

	c.Send("SET", getStrKey("AutoID"), code.Id+1)

	data, err := code.MarshalMsg(nil)
	if err != nil {
		return err
	}
	c.Send("HSET", getStrKey("Codes"), code.Id, data)

	if err = c.Flush(); err != nil {
		return err
	}
	for i := 0; i < 2; i++ {
		if _, err = c.Receive(); err != nil {
			return err
		}
	}
	return nil
}
