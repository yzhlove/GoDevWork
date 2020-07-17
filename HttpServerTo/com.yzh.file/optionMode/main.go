package main

import "fmt"

type BuildRedis struct {
	ID     int
	Addr   string
	DB     int
	Passwd string
}

func NewBuilderRedis(id int) *BuildRedis {
	return &BuildRedis{
		ID:     id,
		DB:     0,
		Addr:   "127.0.0.1",
		Passwd: "12345678",
	}
}

func (builder *BuildRedis) WithAddd(addr string) *BuildRedis {
	builder.Addr = addr
	return builder
}

func (builder *BuildRedis) WithDB(db int) *BuildRedis {
	builder.DB = db
	return builder
}

func (builder *BuildRedis) WithPasswd(passwd string) *BuildRedis {
	builder.Passwd = passwd
	return builder
}

func (builder *BuildRedis) String() string {
	return fmt.Sprintf("[BuildRedis] id:%d addr:%s db:%d passwd:%s",
		builder.ID, builder.Addr, builder.DB, builder.Passwd)
}

func main() {
	fmt.Println(NewBuilderRedis(1).WithAddd("192.168.0.1").String())
	fmt.Println(NewBuilderRedis(1).WithDB(12).String())
	fmt.Println(NewBuilderRedis(1).WithDB(12).WithPasswd("12345673fasd").String())
}
