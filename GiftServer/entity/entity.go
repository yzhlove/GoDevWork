package entity

import (
	"WorkSpace/GoDevWork/GiftServer/db"
	"WorkSpace/GoDevWork/GiftServer/obj"
)



type entity struct {
	MaxId       uint32
	CodesMap    map[uint32]*obj.CodeInfo
	FixCodesMap map[string]uint32
}

var _entity entity

func Init() (err error) {
	if id, err := db.GetMaxId(); err != nil {
		return err
	} else {
		_entity.MaxId = uint32(id)
	}

	if _entity.CodesMap, err = db.GetCodeInfoList(); err != nil {
		return
	}

	_entity.FixCodesMap = make(map[string]uint32, len(_entity.CodesMap)>>1)
	for _, code := range _entity.CodesMap {
		if code.FixCode != "" {
			_entity.FixCodesMap[code.FixCode] = code.Id
		}
	}

	return
}

func GetMaxId() uint32 {
	return _entity.MaxId
}

func GetCodesMap() map[uint32]*obj.CodeInfo {
	return _entity.CodesMap
}

func GetFixCodeId(code string) (id uint32, ok bool) {
	id, ok = _entity.FixCodesMap[code]
	return
}

func AddCode(id uint32, code *obj.CodeInfo) {
	_entity.CodesMap[id] = code
	_entity.MaxId++
}
