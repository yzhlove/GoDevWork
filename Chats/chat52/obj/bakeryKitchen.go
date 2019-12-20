package obj

//go:generate msgp -io=false -tests=false
//厨房
type KitchenController struct {
	Id       uint32 //厨房id
	Ovens    []*Oven
	LastTime int64 //上次结算时间
}

type OvenCat struct {
	CatId    uint32
	StarTime int64
}

type OvenFinishCake struct {
	Id  uint32
	Num uint32
}

//烤箱
type Oven struct {
	Id                 uint32            //烤箱ID
	CakeId             uint32            //正在生产的蛋糕Id
	TotalSet           uint32            //准备生产的set数量
	NowSetPoint        uint32            //当前set完成的点
	FinishSet          uint32            //上次结算已经完成的set
	LastSettlementTime int64             //上次结算的时间(完整的set)
	FinishCakes        []*OvenFinishCake //已经生产完成的蛋糕，以set为单位 key:cakeId value:num
	Cats               []OvenCat         //当前设备上阵的猫娘
	NowSetStartTime    int64             //当前set 开始生产时间 为0 表示烤箱处于停止生产状态
	NowSetEndTime      int64             //当前set 完成时间
}
