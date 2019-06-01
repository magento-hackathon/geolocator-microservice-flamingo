package domain

import "net"

type (
	// LocationData structures response data from a locaton provider
	LocationData struct {
		Longitude float32
		Latitude float32
		ContinentCode string
		ContinentName string
		CountryCode string
		CountryName string
		RegionCode string
		RegionName string
		City string
		Zip string
	}

	// LocationProvider interface to fulfill for providers
	LocationProvider interface {
		GetLocationByIP(ipAdress net.IP) (*LocationData, error)
	}
)
