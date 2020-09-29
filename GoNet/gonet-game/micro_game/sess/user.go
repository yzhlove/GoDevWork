package sess

// User 用户信息
type User struct {
	Id            int32
	Name          string
	Level         uint8
	Score         int32
	LastLoginTime int64
	CreateTime    int64
}
