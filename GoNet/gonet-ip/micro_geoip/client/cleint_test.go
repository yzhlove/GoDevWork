package client

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"micro_geoip/proto"
	"testing"
)

var (
	TestIp = "120.24.37.249"
)

func TestGeoIp(t *testing.T) {

	conn, err := grpc.Dial("127.0.0.1:4388", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	// client
	client := proto.NewGeoIpClient(conn)

	res, err := client.Query(context.Background(), &proto.Geo_IP{Ip: TestIp})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Query => ", res.Msg)

}
