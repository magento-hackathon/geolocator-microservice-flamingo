package infrastructure

import (
	"errors"
	appDomain "github.com/magento-hackathon/geolocator-microservice-flamingo/src/app/domain"
	maxmindGeoIP "github.com/savaki/geoip2"
	"net"
)

// Constants
const (
	ProviderCode = "maxmind"
)

type (
	// MaxmindAdapter concrete maxmind.com implementation
	MaxmindAdapter struct {
		Config struct {
			activeflag bool
			userID     string
			licenseKey string
		}
	}
)

var _ appDomain.LocationProvider = new(MaxmindAdapter)

// Inject dependencies
func (p *MaxmindAdapter) Inject(
	cfg *struct {
	ActiveFlag bool   `inject:"config:providers.maxmind.active"`
	UserID     string `inject:"config:providers.maxmind.userID"`
	LicenseKey string `inject:"config:providers.maxmind.licenseKey"`
},
) {
	if cfg != nil {
		p.Config.activeflag = cfg.ActiveFlag
		p.Config.userID = cfg.UserID
		p.Config.licenseKey = cfg.LicenseKey
	}
}

// GetLocationByIP retrieves the result from maxmind.com
func (p *MaxmindAdapter) GetLocationByIP(ipAddress net.IP) (*appDomain.LocationData, error) {
	if p.Config.activeflag == false {
		return nil, errors.New(appDomain.ProviderInactive)
	}

	api := maxmindGeoIP.New(p.Config.userID, p.Config.licenseKey)

	maxmindResponse, err := api.City(nil, ipAddress.String())
	if err != nil {
		return nil, err
	}

	locationData := &appDomain.LocationData{
		ProviderCode:  ProviderCode,
		Longitude:     float32(maxmindResponse.Location.Longitude),
		Latitude:      float32(maxmindResponse.Location.Latitude),
		ContinentCode: maxmindResponse.Continent.Code,
		ContinentName: maxmindResponse.Continent.Names["en"],
		CountryCode:   maxmindResponse.Country.IsoCode,
		CountryName:   maxmindResponse.Country.Names["en"],
		RegionCode:    maxmindResponse.Subdivisions[0].IsoCode,
		RegionName:    maxmindResponse.Subdivisions[0].Names["en"],
		City:          maxmindResponse.City.Names["en"],
		Zip:           maxmindResponse.Postal.Code,
	}

	return locationData, nil
}
