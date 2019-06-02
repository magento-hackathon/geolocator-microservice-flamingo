package infrastructure

import (
	"encoding/json"
	"errors"
	"fmt"
	appDomain "github.com/magento-hackathon/geolocator-microservice-flamingo/src/app/domain"
	"github.com/magento-hackathon/geolocator-microservice-flamingo/src/ipdata/infrastructure/dto"
	"net"
	"net/http"
)

// Constants
const (
	ProviderCode = "ipdata"
)

type (
	// IPDataAdapter concrete ipdata.co implementation
	IPDataAdapter struct {
		Config struct {
			activeflag bool
			apiKey     string
		}
	}
)

var _ appDomain.LocationProvider = new(IPDataAdapter)

// Inject dependencies
func (p *IPDataAdapter) Inject(
	cfg *struct {
	ActiveFlag bool   `inject:"config:providers.ipdata.active"`
	APIKey     string `inject:"config:providers.ipdata.apiKey"`
},
) {
	if cfg != nil {
		p.Config.activeflag = cfg.ActiveFlag
		p.Config.apiKey = cfg.APIKey
	}
}

// GetLocationByIP retrieves the result from ipdata.co
func (p *IPDataAdapter) GetLocationByIP(ipAddress net.IP) (*appDomain.LocationData, error) {
	if p.Config.activeflag == false {
		return nil, errors.New(appDomain.ProviderInactive)
	}

	requestURL := fmt.Sprintf("https://api.ipdata.co/%s?api-key=%s", ipAddress, p.Config.apiKey)

	response, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}

	jsonResult := &dto.IPDataResponse{}

	err = json.NewDecoder(response.Body).Decode(jsonResult)
	if err != nil {
		return nil, err
	}

	locationData := &appDomain.LocationData{
		ProviderCode:  ProviderCode,
		Longitude:     float32(jsonResult.Longitude),
		Latitude:      float32(jsonResult.Latitude),
		ContinentCode: jsonResult.ContinentCode,
		ContinentName: jsonResult.ContinentName,
		CountryCode:   jsonResult.CountryCode,
		CountryName:   jsonResult.CountryName,
		RegionCode:    jsonResult.RegionCode,
		RegionName:    jsonResult.Region,
		City:          jsonResult.City,
		Zip:           jsonResult.Postal,
	}

	return locationData, nil
}
