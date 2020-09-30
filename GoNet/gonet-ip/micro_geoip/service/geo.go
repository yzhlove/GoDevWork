package service

import (
	"context"
	"encoding/json"
	"github.com/oschwald/geoip2-golang"
	log "github.com/sirupsen/logrus"
	"micro_geoip/config"
	"micro_geoip/proto"
	"net"
)

// Geo Ip信息
type Geo struct {
	Country  string   `json:"country"`
	Province string   `json:"province"`
	City     string   `json:"city"`
	ISO      string   `json:"iso"`
	TimeZone string   `json:"time_zone"`
	Local    Location `json:"location"`
}

// Location 经纬度
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GeoService struct {
	reader *geoip2.Reader
}

func (s *GeoService) Init(cfg *config.Config) {
	if reader, err := geoip2.Open(cfg.Path); err != nil {
		log.Fatal(err)
	} else {
		s.reader = reader
		log.Info(config.SERVICE, " City Load Complete.")
	}
}

func (s *GeoService) Query(ctx context.Context, in *proto.Geo_IP) (*proto.String, error) {
	record, err := s.reader.City(net.ParseIP(in.Ip))
	if err != nil {
		return nil, err
	}
	res := &Geo{
		City:     record.City.Names[config.AREA],
		Country:  record.Country.Names[config.LANG],
		ISO:      record.Country.IsoCode,
		TimeZone: record.Location.TimeZone,
		Local:    Location{Latitude: record.Location.Latitude, Longitude: record.Location.Longitude},
	}
	if len(record.Subdivisions) > 0 {
		res.Province = record.Subdivisions[0].Names[config.LANG]
	}
	if result, err := json.Marshal(res); err != nil {
		return nil, err
	} else {
		return &proto.String{Msg: string(result)}, nil
	}
}
