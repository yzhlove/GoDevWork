package manager

import (
	"WorkSpace/GoDevWork/GiftServer/db"
	"WorkSpace/GoDevWork/GiftServer/misc/helper"
	"WorkSpace/GoDevWork/GiftServer/obj"
	pb "WorkSpace/GoDevWork/GiftServer/proto"
	log "github.com/sirupsen/logrus"
	"time"
)

func CodeVerify(uid uint64, zone uint32, code string) bool {

	id, ok := helper.Decode(helper.CodeToID(code))
	if !ok {
		//检查是否是FixCode
	}
	info, err := db.GetCodeInfo(id)
	if err != nil {
		log.Error(err)
		return false
	}
	//检查开放时间
	if ts := time.Now().Unix(); ts < info.StartTime || ts >= info.EndTime {
		return false
	}
	//检查区支持
	status := true
	for _, z := range info.ZoneIds {
		if z == zone {
			status = false
			break
		}
	}
	if status {
		return false
	}
	//检测次数使用

	return true
}

func GetCodeInfoList(zone uint32) ([]*pb.Manager_CodeInfo, error) {

	codes, err := db.GetCodeInfoList()
	if err != nil {
		return nil, err
	}

	ptoCodes := make([]*pb.Manager_CodeInfo, 0, len(codes))
	for _, code := range codes {
		filter := true
		if len(code.ZoneIds) > 0 {
			for _, id := range code.ZoneIds {
				if id == zone {
					break
				}
			}
			filter = false
		}
		if filter {
			ptoCodes = append(ptoCodes, generatePtoCodeMessage(code))
		}
	}
	return ptoCodes, nil
}

func generatePtoCodeMessage(code obj.CodeInfo) *pb.Manager_CodeInfo {

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
