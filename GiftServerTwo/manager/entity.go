package manager

import (
	"WorkSpace/GoDevWork/GiftServerTwo/db"
	"WorkSpace/GoDevWork/GiftServerTwo/obj"
	"WorkSpace/GoDevWork/GiftServerTwo/pb"
	"strings"
)

func EntityInit() (e *obj.Entity, err error) {
	e = &obj.Entity{}
	if e.AutoId, err = db.GetAutoId(); err != nil {
		return
	}
	if e.AutoId == 0 {
		e.AutoId = 1
	}
	if e.Infos, err = db.GetCodes(); err != nil {
		return
	}
	e.Fixed = make(map[string]uint32, len(e.Infos))
	for id, code := range e.Infos {
		if fixed := strings.TrimSpace(code.FixCode); len(fixed) > 0 {
			e.Fixed[fixed] = id
		}
	}
	return
}

func GeneratePtoCodeInfo(c *obj.Code) *pb.Manager_CodeInfo {

	return &pb.Manager_CodeInfo{
		Id:   c.Id,
		Used: c.Used,
		GenInfo: &pb.Manager_GenReq{
			FixCode:      c.FixCode,
			Num:          c.Num,
			StartTime:    c.StartTime,
			EndTime:      c.EndTime,
			TimesPerCode: uint32(c.TimesPerCode),
			TimesPerUser: uint32(c.TimesPerUser),
			ZoneIds:      c.ZoneIds,
			Items:        generatePtoItems(c.Items),
		},
	}
}

func generatePtoItems(its []obj.Item) []*pb.Manager_Item {

	ps := make([]*pb.Manager_Item, 0, len(its))
	for _, it := range its {
		ps = append(ps, &pb.Manager_Item{Id: it.Id, Num: it.Num})
	}
	return ps
}


