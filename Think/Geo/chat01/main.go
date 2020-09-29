package main

//geo ip 使用

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"net"
)

func main() {
	db, err := geoip2.Open("./GeoIP2-City.mmdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	ip := net.ParseIP("120.24.37.249")
	record, err := db.City(ip)
	fmt.Printf("Portuguese (BR) city name: %v\n", record.City.Names["zh-CN"])
	if len(record.Subdivisions) > 0 {
		fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
	}
	fmt.Printf("Russian country name: %v\n", record.Country.Names["en"])
	fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
	fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
}
