package domain

import "net"

// Error constants
const (
	ProviderInactive = "provider inactive"
)

type (
	// LocationData structures response data from a locaton provider
	LocationData struct {
		Longitude     float32
		Latitude      float32
		ContinentCode string
		ContinentName string
		CountryCode   string
		CountryName   string
		RegionCode    string
		RegionName    string
		City          string
		Zip           string
		ErrorMessage  string
	}

	// LocationProvider interface to fulfill for providers
	LocationProvider interface {
		GetLocationByIP(ipAddress net.IP) (*LocationData, error)
	}
)
