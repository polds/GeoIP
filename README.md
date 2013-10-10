# GeoIP

GeoIP is a [FreeGeoIP](https://github.com/fiorix/freegeoip) API abstraction library written in Go like the original package. This package utilizes the freely available RESTful API found on their [website](http://freegeoip.net). [API Documentation](http://godoc.org/github.com/polds/GeoIP) on go.pkgdoc.org.

## [Installation](https://github.com/polds/GeoIP#installation)

    go get github.com/polds/GeoIP

## [Usage](https://github.com/polds/GeoIP#usage)

	package main

	import (
		"fmt"
		"github.com/polds/GeoIP"
	)

	func main() {
		info, err := GeoIP.Fetch("127.0.0.1")
		if err != nil {
			panic(err) // Could be rate limiting or not found
		}

		fmt.Println(info.City)
		fmt.Println(info.CountryCode)
	}