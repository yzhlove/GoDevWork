package manager

import (
	"WorkSpace/GoDevWork/GiftServer/db"
	"WorkSpace/GoDevWork/GiftServer/entity"
	"WorkSpace/GoDevWork/GiftServer/misc/helper"
	"WorkSpace/GoDevWork/GiftServer/obj"
	pb "WorkSpace/GoDevWork/GiftServer/proto"
	"errors"
	"strconv"
)

func CodeVerify(uid uint64, zone uint32, code string) (bool, error) {
	id, ok := helper.Decode(helper.CodeToID(code))
	exists := false
	if !ok {
		//是否是固定code
		if id, ok = entity.GetFixCodeId(code); !ok {
			return false, errors.New("code is invalid :" + code)
		}
		exists = true
	}

	//是否存在
	if !exists {
		ok, err := db.ExistsCode(id, code)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, errors.New("not found code:" + code)
		}
	}

	//获取code的详细信息
	base, ok := entity.GetCodesMap()[id]
	if !ok {
		return false, errors.New("id is invalid:" + strconv.Itoa(int(id)))
	}

	//时间
	if !base.IsExpired() {
		return false, errors.New("code use time err")
	}

	//区
	if !base.IsMatchZone(zone) {
		return false, errors.New("code no match zones")
	}

	//用户是否使用
	ok, err := db.IsUseCode(uid, id, code)
	if err != nil {
		return false, err
	}
	//已经使用
	if ok {
		return false, errors.New("code is use")
	}

	//使用次数
	count, err := db.GetCodeTimes(id, code)
	if err != nil {
		return false, err
	}
	if base.TimesPerCode <= uint16(count) {
		return false, errors.New("times limit")
	}

	//设置code为已使用状态
	if err := db.SetUseCode(uid, id, code); err != nil {
		return false, err
	}
	return true, nil
}

func GeneratePtoCodeMessage(code *obj.CodeInfo) *pb.Manager_CodeInfo {

	return &pb.Manager_CodeInfo{
		Id:   code.Id,
		Used: 0,
		GenInfo: &pb.Manager_GenReq{
			FixCode:      code.FixCode,
			Num:          code.Num,
			StartTime:    code.StartTime,
			EndTime:      code.EndTime,
			TimesPerCode: uint32(code.TimesPerCode),
			TimesPerUser: uint32(code.TimesPerUser),
			ZoneIds:      code.ZoneIds,
			Items:        generatePtoItem(code.Items),
		},
	}
}

func generatePtoItem(items []obj.Item) []*pb.Manager_Item {

	ptoItems := make([]*pb.Manager_Item, 0, len(items))
	for _, item := range items {
		ptoItems = append(ptoItems, &pb.Manager_Item{
			Id:  item.Id,
			Num: item.Num,
		})
	}
	return ptoItems
}
