package manager

import (
	"WorkSpace/GoDevWork/GiftServer/config"
	"WorkSpace/GoDevWork/GiftServer/db"
	"WorkSpace/GoDevWork/GiftServer/entity"
	"WorkSpace/GoDevWork/GiftServer/misc/helper"
	"WorkSpace/GoDevWork/GiftServer/obj"
	pb "WorkSpace/GoDevWork/GiftServer/proto"
	log "github.com/sirupsen/logrus"
	"sync"
)

func GenerateCodeInfo(req *pb.Manager_GenReq) (*pb.Manager_CodeInfo, error) {

	id := entity.GetMaxId()
	generate, err := helper.GenerateCode(id, req.Num)
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go setCodes(id, &wg, generate)
	}
	wg.Wait()

	code := generateCode(id, req)
	if err := db.SetCodeInfo(code); err != nil {
		return nil, err
	}

	entity.AddCode(id, code)

	return &pb.Manager_CodeInfo{
		Id:      id,
		Used:    0,
		GenInfo: req,
	}, nil
}

func setCodes(id uint32, wg *sync.WaitGroup, generate chan int64) {
	defer wg.Done()
	max, index := 100, 0
	data := make(map[int][]string, config.BucketMax)
	for number := range generate {
		if index >= max {
			if err := db.SetCodes(id, data); err != nil {
				log.Error(err)
				return
			}
			data = make(map[int][]string, config.BucketMax)
			index = 0
		}
		code := helper.IdToCode(helper.Encode(id, number))
		i := helper.GetBucketTop(code)
		data[i] = append(data[i], code)
		index++
	}
	if index > 0 {
		if err := db.SetCodes(id, data); err != nil {
			log.Error(err)
			return
		}
	}
}

func generateCode(id uint32, req *pb.Manager_GenReq) *obj.CodeInfo {
	return &obj.CodeInfo{
		Id:           id,
		FixCode:      req.FixCode,
		Num:          req.Num,
		StartTime:    req.StartTime,
		EndTime:      req.EndTime,
		TimesPerCode: uint16(req.TimesPerCode),
		TimesPerUser: uint16(req.TimesPerUser),
		ZoneIds:      req.ZoneIds,
		Items:        generateItem(req.Items),
	}
}

func generateItem(ptoItems []*pb.Manager_Item) []obj.Item {
	items := make([]obj.Item, len(ptoItems))
	for _, item := range ptoItems {
		items = append(items, obj.Item{Id: item.Id, Num: item.Num,})
	}
	return items
}
