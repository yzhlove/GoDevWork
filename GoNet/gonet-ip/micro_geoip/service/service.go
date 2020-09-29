package service

import (
	"context"
	"errors"
	"github.com/oschwald/maxminddb-golang"
	log "github.com/sirupsen/logrus"
	"micro_geoip/proto"
	"net"
)

const (
	SERVICE  = "[GEOIP]"
	Language = "en"
)

var (
	ErrorCannotQueryIp = errors.New("cannot query ip")
)

type Geo struct {
	Co  Country    `maxminddb:"country"`
	Ct  City       `maxminddb:"city"`
	Pov []Province `maxminddb:"subdivisions"`
}

// Country 国家
type Country struct {
	GeoNameId uint   `maxminddb:"geoname_id"`
	IsoCode   string `maxminddb:"iso_code"`
}

// City 城市
type City struct {
	Names map[string]string `maxminddb:"city"`
}

// Province 省份
type Province struct {
	Names map[string]string `maxminddb:"names"`
}

type GeoService struct {
	reader *maxminddb.Reader
}

func (s *GeoService) init() {

}

func (s *GeoService) load(file string) {

}

func (s *GeoService) query(ip net.IP) *Geo {
	geo := new(Geo)
	if err := s.reader.Lookup(ip, geo); err != nil {
		log.Error(err)
		return nil
	}
	return geo
}

func (s *GeoService) QueryCountry(ctx context.Context, in *proto.Geo_IP) (*proto.Geo_Res, error) {
	if res := s.query(net.ParseIP(in.Ip)); res != nil {
		return &proto.Geo_Res{Result: res.Co.IsoCode}, nil
	}
	return nil, ErrorCannotQueryIp
}

func (s *GeoService) QueryProvince(ctx context.Context, in *proto.Geo_IP) (*proto.Geo_Res, error) {
	if res := s.query(net.ParseIP(in.Ip)); res != nil {
		if len(res.Pov) > 0 {
			return &proto.Geo_Res{Result: res.Pov[0].Names[Language]}, nil
		}
	}
	return nil, ErrorCannotQueryIp
}

func (s *GeoService) QueryCity(ctx context.Context, in *proto.Geo_IP) (*proto.Geo_Res, error) {
	if res := s.query(net.ParseIP(in.Ip)); res != nil {
		return &proto.Geo_Res{Result: res.Ct.Names[Language]}, nil
	}
	return nil, ErrorCannotQueryIp
}
