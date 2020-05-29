package main

import "fmt"

//option模式
//选项设计模式

type ModOption func(opt *Option)

type AbstractRedis struct {
	ID        int
	OptParams Option
}

type Option struct {
	Addr     string
	DB       int
	Password string
}

func WithAddr(addr string) ModOption {
	return func(opt *Option) {
		opt.Addr = addr
	}
}

func WithDB(db int) ModOption {
	return func(opt *Option) {
		opt.DB = db
	}
}

func WithPasswd(password string) ModOption {
	return func(opt *Option) {
		opt.Password = password
	}
}

func (redis *AbstractRedis) String() string {
	return fmt.Sprintf("[AbstractRedis] id:%d addr:%s db:%d passwd:%s",
		redis.ID, redis.OptParams.Addr, redis.OptParams.DB, redis.OptParams.Password)
}

func NewAbstractRedis(id int, options ...ModOption) *AbstractRedis {
	opt := Option{
		Addr:     "127.0.0.1",
		DB:       0,
		Password: "1234567",
	}
	for _, fn := range options {
		fn(&opt)
	}
	return &AbstractRedis{ID: id, OptParams: opt}
}

func main() {

	fmt.Println(NewAbstractRedis(0).String())
	fmt.Println(NewAbstractRedis(0, WithAddr("192.168.0.1")).String())
	fmt.Println(NewAbstractRedis(0, WithDB(4)).String())
	fmt.Println(NewAbstractRedis(0, WithDB(1), WithPasswd("1234")).String())
}
