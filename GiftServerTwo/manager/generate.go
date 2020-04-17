package manager

import (
	"WorkSpace/GoDevWork/GiftServerTwo/config"
	"WorkSpace/GoDevWork/GiftServerTwo/db"
	"WorkSpace/GoDevWork/GiftServerTwo/misc/helper"
	"WorkSpace/GoDevWork/GiftServerTwo/obj"
	"WorkSpace/GoDevWork/GiftServerTwo/pb"
	"context"
	log "github.com/sirupsen/logrus"
	"strings"
	"sync"
)

type pbGenReq = *pb.Manager_GenReq
type pbItems = []*pb.Manager_Item

func GenerateCodes(id uint32, req pbGenReq) (*obj.Code, error) {

	//如果不是固定code
	if fixed := strings.TrimSpace(req.FixCode); len(fixed) == 0 {
		ctx, cancel := context.WithCancel(context.Background())
		codeChan := helper.GenerateCode(ctx, req.Num)
		var wg sync.WaitGroup
		for i, j := 0, config.BucketMax>>1; i < j; i++ {
			wg.Add(1)
			go generateCode(cancel, id, &wg, codeChan)
		}
		wg.Wait()
		cancel()
	}

	code := generateCodeObject(id, req)
	if err := db.SetCode(code); err != nil {
		return nil, err
	}

	return code, nil
}

func generateCode(cancel context.CancelFunc, id uint32, wg *sync.WaitGroup, codeChan chan int64) {
	defer wg.Done()
	max, count := 100, 0
	group := make(map[int][]string, config.BucketMax)
	for number := range codeChan {
		if count >= max {
			if err := db.SetUseCodes(id, group); err != nil {
				log.Error(err)
				cancel()
				return
			}
			//clear
			for i, codes := range group {
				codes[i] = codes[i][:0]
			}
			count = 0
		}
		newCode := helper.ToEncodeStr(helper.ToEncodeNumber(id, number))
		top := helper.GetBucketTop(newCode)
		group[top] = append(group[top], newCode)
		count++
	}
	if count > 0 {
		if err := db.SetUseCodes(id, group); err != nil {
			log.Error(err)
			return
		}
	}
}

func generateCodeObject(id uint32, req pbGenReq) *obj.Code {

	return &obj.Code{
		Id:           id,
		FixCode:      req.FixCode,
		Num:          req.Num,
		StartTime:    req.StartTime,
		EndTime:      req.EndTime,
		TimesPerCode: uint16(req.TimesPerCode),
		TimesPerUser: uint16(req.TimesPerUser),
		ZoneIds:      req.ZoneIds,
		Items:        generateItemObject(req.Items),
		Used:         0,
	}
}

func generateItemObject(items pbItems) []obj.Item {

	newItems := make([]obj.Item, 0, len(items))
	for _, it := range items {
		newItems = append(newItems, obj.Item{Id: it.Id, Num: it.Num})
	}
	return newItems
}
