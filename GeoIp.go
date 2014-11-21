// Package GeoIP is a [freegeoip.net](https://github.com/fiorix/freegeoip)
// api abstraction layer also written in Go like the original package.
package GeoIP

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Ip struct {
	IP          string  `json:"ip,omitempty"`
	CountryCode string  `json:"country_code,omitempty"`
	CountryName string  `json:"country_name,omitempty"`
	RegionCode  string  `json:"region_code,omitempty"`
	RegionName  string  `json:"region_name,omitempty"`
	City        string  `json:"city,omitempty"`
	Zipcode     string  `json:"zipcode,omitempty"`
	TimeZone    string  `json:"time_zone,omitempty"`
	Latitude    float32 `json:"latitude,omitempty"`
	Longitude   float32 `json:"longitude,omitempty"`
	MetroCode   int     `json:"metro_code,omitempty"`
}

// Fetch makes the call to freegeoip to search
// for the provided ip Address
func Fetch(query string) (Ip, error) {
	var response Ip

	resp, err := http.Get("http://freegeoip.net/json/" + query)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 403:
		return response, errors.New("freegeoip hourly limit reached")
	case 404:
		return response, errors.New("invalid ip Address, or not found")
	default:
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(body, &response)
	return response, err
}
