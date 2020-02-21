package entity

import (
	"WorkSpace/GoDevWork/GiftServer/db"
	"WorkSpace/GoDevWork/GiftServer/obj"
	"strings"
)

type EntityObject struct {
	MaxId       uint32
	CodesMap    map[uint32]*obj.CodeInfo
	FixCodesMap map[string]uint32
}

var Entity *EntityObject

func Init() error {
	Entity = new(EntityObject)

	id, err := db.GetMaxId()
	if err != nil {
		return err
	}
	Entity.MaxId = uint32(id)

	codes, err := db.GetCodeInfoList()
	if err != nil {
		return err
	}

	Entity.CodesMap = codes
	Entity.FixCodesMap = make(map[string]uint32, len(codes)>>1)
	for _, code := range Entity.CodesMap {
		if strings.TrimSpace(code.FixCode) != "" {
			Entity.FixCodesMap[code.FixCode] = code.Id
		}
	}
	return nil
}

func AddCode(id uint32, code *obj.CodeInfo) {
	Entity.CodesMap[id] = code
	Entity.MaxId++
}
