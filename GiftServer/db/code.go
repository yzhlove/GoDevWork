package db

import (
	"WorkSpace/GoDevWork/GiftServer/obj"
	"github.com/garyburd/redigo/redis"
	"time"
)

func SetCodeInfoList(infos []obj.CodeInfo) error {
	c := RedisClient.Get()
	defer c.Close()
	codes := make(map[uint32][]byte, 4)
	for _, code := range infos {
		if data, err := code.MarshalMsg(nil); err != nil {
			return err
		} else {
			codes[code.Id] = data
		}
	}
	_, err := c.Do("HMSET", redis.Args{}.Add("CodeInfoList").AddFlat(codes)...)
	return err
}

func GetCodeInfoList() (map[uint32]obj.CodeInfo, error) {

	c := RedisClient.Get()
	defer c.Close()

	data, err := redis.ByteSlices(c.Do("HVALS", "CodeInfoList"))
	if err != nil {
		return nil, err
	}
	ts := time.Now().Unix()
	codes := make(map[uint32]obj.CodeInfo, len(data))
	ids := make([]uint32, 0, len(codes))
	for _, v := range data {
		var code obj.CodeInfo
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

	_, err := c.Do("HMDEL", redis.Args{}.Add("CodeInfoList").AddFlat(ids)...)
	return err
}

func GetCodeInfo(id uint32) (code obj.CodeInfo, err error) {
	c := RedisClient.Get()
	defer c.Close()

	data, err := redis.Bytes(c.Do("HGET", "CodeInfoList", id))
	if err != nil {
		return obj.CodeInfo{}, err
	}
	_, err = code.UnmarshalMsg(data)
	return
}

