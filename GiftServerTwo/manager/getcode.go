package manager

import (
	"WorkSpace/GoDevWork/GiftServerTwo/config"
	"WorkSpace/GoDevWork/GiftServerTwo/db"
	"WorkSpace/GoDevWork/GiftServerTwo/misc/helper"
	"WorkSpace/GoDevWork/GiftServerTwo/pb"
)

type pbCodeStatus = *pb.Manager_CodeStatus

func GetCodes(id uint32, fixedCode string) (respCodes []pbCodeStatus, err error) {

	if fixedCode != "" {
		values, err := db.GetUseCodes(id, helper.GetBucketTop(fixedCode))
		if err != nil {
			return respCodes, err
		}
		for _, value := range values {
			respCodes = append(respCodes, &pb.Manager_CodeStatus{
				Code: value.Code, UserId: int64(value.UID), ZoneId: value.Zone,
			})
		}
		return respCodes, nil
	}

	for i := 0; i < config.BucketMax; i++ {
		values, err := db.GetUseCodes(id, helper.GetBucketTop(fixedCode))
		if err != nil {
			return respCodes, err
		}
		for _, value := range values {
			respCodes = append(respCodes, &pb.Manager_CodeStatus{
				Code: value.Code, UserId: int64(value.UID), ZoneId: value.Zone,
			})
		}

	}
	return respCodes, nil
}
