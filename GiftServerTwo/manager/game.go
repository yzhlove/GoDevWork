package manager

import (
	"WorkSpace/GoDevWork/GiftServerTwo/db"
	"WorkSpace/GoDevWork/GiftServerTwo/obj"
	"errors"
)

func CodeVerify(base *obj.Code, uid uint64, z uint32, code string) (bool, error) {

	//过期时间
	if !base.Expired() {
		return false, errors.New("time is expire")
	}

	//支持区域
	if !base.ZoneCheck(z) {
		return false, errors.New("zone is no match")
	}

	codeUsers, err := db.GetUseCode(base.Id, code)
	if err != nil {
		return false, err
	}

	//是否已经兑换过
	if codeUsers.Use(uid) {
		return false, errors.New("code already used")
	}

	//code可兑换次数
	if uint16(codeUsers.GetTimes()) >= base.TimesPerCode {
		return false, errors.New("usage limit reached")
	}

	//用户可以兑换同一批次兑换码的数量上限
	times, err := db.GetRedemptionTimes(uid, base.Id)
	if err != nil {
		return false, err
	}
	if uint16(times) >= base.TimesPerUser {
		return false, errors.New("redemption times to limit")
	}

	//入库
	codeUsers.Users = append(codeUsers.Users, &obj.CodeUser{UID: uid, Zone: z})
	if err = db.Save(base, codeUsers, uid, code); err != nil {
		base.Used--
		return false, err
	}

	return true, nil
}
