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

## [Tests](https://github.com/polds/GeoIP#tests)

[GoConvey](http://smartystreets.github.io/goconvey/) is used in this library for testing. Testing can be done two different ways:

    go test -v

or likewise you can run the GoConvey web UI to run the tests

    goconvey .


## [Notes](https://github.com/polds/GeoIP#notes)
 - As of February 19, 2014 this library is considered complete. No additional features will be added. Should a Go version change require any changes to the library it will be updated. Please be aware this is not an abandoned project but is used on a daily basis.
 - This library has been run through golint. 



 Any feature requets, comments, concerns are welcome: [@Peter_Olds](https://twitter.com/Peter_Olds)