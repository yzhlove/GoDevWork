package db

import (
	"WorkSpace/GoDevWork/GiftServerTwo/misc/helper"
	"WorkSpace/GoDevWork/GiftServerTwo/obj"
	"github.com/garyburd/redigo/redis"
)

func GetUseCode(id uint32, code string) (codeUsers obj.CodeUsers, err error) {

	c := RedisClient.Get()
	defer c.Close()
	top := helper.GetBucketTop(code)
	var data []byte
	if data, err = redis.Bytes(c.Do("HGET", getSpecKey("UseCodes", uint64(id), top), code)); err != nil {
		if err == redis.ErrNil {
			return codeUsers, nil
		}
	}
	_, err = codeUsers.UnmarshalMsg(data)
	return
}

func GetRedemptionTimes(uid uint64, id uint32) (int, error) {
	c := RedisClient.Get()
	defer c.Close()

	return redis.Int(c.Do("SCARD", getSpecKey("Users", uid, int(id))))
}

func GetUseCodes(id uint32, top int) (respCodeUsers []obj.RespCodeUser, err error) {

	c := RedisClient.Get()
	defer c.Close()

	values, err := redis.StringMap(c.Do("HGETALL", getSpecKey("UseCodes", uint64(id), top)))
	if err != nil {
		return
	}

	var codeUsers = &obj.CodeUsers{}
	for code, data := range values {
		if _, err = codeUsers.UnmarshalMsg([]byte(data)); err != nil {
			return
		}
		for _, codeUser := range codeUsers.Users {
			respCodeUsers = append(respCodeUsers, obj.RespCodeUser{
				Code:     code,
				CodeUser: obj.CodeUser{UID: codeUser.UID, Zone: codeUser.Zone},
			})
		}
	}
	return
}
