package obj

//go:generate msgp -io=false -tests=false
type CodeUser struct {
	UID  uint64
	Zone uint32
}

//msgp:ignore RespCodeUser
type RespCodeUser struct {
	CodeUser
	Code string
}

type CodeUsers struct {
	Users []*CodeUser
}

func (c *CodeUsers) Use(uid uint64) bool {
	for _, user := range c.Users {
		if user.UID == uid {
			return true
		}
	}
	return false
}

func (c *CodeUsers) GetTimes() int {
	return len(c.Users)
}
