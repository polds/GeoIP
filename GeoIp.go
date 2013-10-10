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
	Ip          string  `json:"ip,omitempty"`
	CountryCode string  `json:"country_code,omitempty"`
	CountryName string  `json:"country_name,omitempty"`
	RegionCode  string  `json:"region_code,omitempty"`
	RegionName  string  `json:"region_name,omitempty"`
	City        string  `json:"city,omitempty"`
	Zipcode     string  `json:"zipcode,omitempty"`
	Latitude    float32 `json:"latitude,omitempty"`
	Longitude   float32 `json:"longitude,omitempty"`
	MetroCode   string  `json:"metro_code,omitempty"`
	Areacode    string  `json:"areacode,omitempty"`
}

// Fetch makes the call to freegeoip to search
// for the provided IP Address
func Fetch(ip string) (Ip, error) {
	resp, err := http.Get("http://freegeoip.net/json/" + ip)
	if err != nil {
		return Ip{}, err
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case 403:
		return Ip{}, errors.New("Freegeoip hourly limit reached")
	case 404:
		return Ip{}, errors.New("Invalid IP Address, or Not Found")
	default:
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Ip{}, err
	}
	var response Ip
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Ip{}, err
	}

	return response, nil
}
