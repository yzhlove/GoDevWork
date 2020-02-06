package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/cacheClient/client"
	"flag"
	"fmt"
	"time"
)

type statistic struct {
	count int
	tm    time.Duration
}

type result struct {
	getCount    int
	setCount    int
	missCount   int
	statBuckets []statistic
}

func (r *result) addStatistic(bucket int, stat statistic) {
	if bucket > len(r.statBuckets) {
		newStatBuckets := make([]statistic, bucket+1)
		copy(newStatBuckets, r.statBuckets)
		r.statBuckets = newStatBuckets
	}
	s := r.statBuckets[bucket]
	s.count += stat.count
	s.tm += stat.tm
	r.statBuckets[bucket] = s
}

func (r *result) addDuration(d time.Duration, typ string) {
	bucket := int(d / time.Millisecond)
	r.addStatistic(bucket, statistic{1, d})
	switch typ {
	case "get":
		r.getCount++
	case "set":
		r.setCount++
	default:
		r.missCount++
	}
}

func (r *result) addResult(src *result) {
	for b, s := range src.statBuckets {
		r.addStatistic(b, s)
	}
	r.getCount += src.getCount
	r.setCount += src.setCount
	r.missCount += src.missCount
}

func run(c client.Client, msg *client.Message, r *result) {
	expect := msg.Value
	start := time.Now()
	c.Run(msg)
	d := time.Now().Sub(start)
	resultType := msg.Name
	if resultType == "get" {
		if msg.Value == "" {
			resultType = "miss"
		} else if msg.Value != expect {
			panic(msg)
		}
	}
	r.addDuration(d, resultType)
}

func pipeline(c client.Client, msgs []*client.Message, r *result) {
	except := make([]string, len(msgs))
	for i, c := range msgs {
		if c.Name == "get" {
			except[i] = c.Value
		}
	}
	start := time.Now()
	c.PipeLineRun(msgs)
	d := time.Now().Sub(start)
	for i, msg := range msgs {
		resultType := msg.Name
		if resultType == "get" {
			if msg.Value == "" {
				resultType = "miss"
			} else if msg.Value != except[i] {
				fmt.Println(except[i])
				panic(msg.Value)
			}
		}
		r.addDuration(d, resultType)
	}
}

func operator(id, count int, ch chan *result) {
	//c := client.New(typ, server)
	//cmds := []*client.Client{}
	//valueStr := strings.Repeat("a", valueSize)
	//r := &result{0, 0, 0, []statistic{}}
	//for i := 0 ;i < count;i++ {
	//	var tmp int
	//	if keyspacelen > 0 {
	//		tmp = rand.Intn(keyspacelen)
	//	} else {
	//		tmp = id * count + i
	//	}
	//	key := fmt.Sprintf("%d",tmp)
	//	value := fmt.Sprintf("%s%d",valueStr,tmp)
	//
	//}
}

var (
	server, operation string
	typ               int
	valueSize         int
	threads           int
	keyspacelen       int
	pipelen           int
)

func init() {
	flag.IntVar(&typ, "type", client.RedisClient, "cache server type")
	flag.StringVar(&server, "h", "localhost", "ip address")
	flag.StringVar(&operation, "t", "set", "test set")

}

func main() {

}
