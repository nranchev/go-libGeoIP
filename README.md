## General Description ##

This is the Go implementation of the [Maxmind](http://www.maxmind.com/app/ip-location) GeoIP API. It is incomplete and work in progress the initial goal is support only two of the database types - the City Lite and Country Lite. The only supported method is loading the full db on startup into memory (memory cache).

### Supported Access Methods ###

* In Memory (Load(string))

### Supported Database Formats ###

* Country Edition    (dbType=1)
* City Edition REV 0 (dbType=6)
* City Edition REV 1 (dbType=2)

### Supported Lookups ###

* By IP Address (GetLocationByIP(string))
* By IP Number  (GetLocationByIPNum(uint32))

### Supported Responses ###

* CountryCode string (available in all databases)
* CountryName string (available in all databases)
* City string
* Region string
* PostalCode string
* Latitude float32
* Longitude float32

### To Do ###
* Implement better error handling (report the error on load and lookups)
* Better returns, country edition has only code and name (perhaps use interfaces)
* Add test cases and benchmarking
* Add support for more database formats

### Build ###

        go get github.com/zyxar/go-libGeoIP
        go install github.com/zyxar/go-libGeoIP
        go build -o example

### Example ###

./example DBFILE IPADDRESS (i.e. ./example GeoIP.dat 1.1.1.1)

### Usage ###

Please see example.go for a complete example of how to use this library.
