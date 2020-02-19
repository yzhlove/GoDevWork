package obj

//go:generate msgp -io=false -tests=false
type Item struct {
	Id  uint32
	Num int32
}

type CodeInfo struct {
	Id           uint32   //批次
	FixCode      string   //固定兑换码
	Num          uint32   //生成的数量
	StartTime    int64    //开始时间
	EndTime      int64    //结束时间
	TimesPerCode uint16   //兑换码可用次数
	TimesPerUser uint16   //同一批次兑换码单人可使用次数
	ZoneIds      []uint32 //可用区
	Items        []Item   //奖励
}
