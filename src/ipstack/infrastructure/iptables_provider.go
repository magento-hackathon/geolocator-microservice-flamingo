package infrastructure

import (
	"encoding/json"
	"errors"
	"fmt"
	appDomain "github.com/magento-hackathon/geolocator-microservice-flamingo/src/app/domain"
	"github.com/magento-hackathon/geolocator-microservice-flamingo/src/ipstack/infrastructure/dto"
	"net"
	"net/http"
)

// Constants
const (
	ProviderCode = "ipstack.com"
)

type (
	// IPStackProvider concrete ipstack.com implementation
	IPStackProvider struct {
		Config struct {
			activeflag bool
			apiKey     string
			apiURL     string
		}
	}
)

var _ appDomain.LocationProvider = new(IPStackProvider)

// Inject dependencies
func (p *IPStackProvider) Inject(
	cfg *struct {
	ActiveFlag bool   `inject:"config:providers.ipstack.active"`
	APIKey     string `inject:"config:providers.ipstack.apiKey"`
	APIUrl     string `inject:"config:providers.ipstack.apiUrl"`
},
) {
	if cfg != nil {
		p.Config.activeflag = cfg.ActiveFlag
		p.Config.apiKey = cfg.APIKey
		p.Config.apiURL = cfg.APIUrl
	}
}

// GetLocationByIP retrieves the result from ipstack.com
func (p *IPStackProvider) GetLocationByIP(ipAddress net.IP) (*appDomain.LocationData, error) {
	if p.Config.activeflag == false {
		return nil, errors.New(appDomain.ProviderInactive)
	}

	requestURL := fmt.Sprintf(p.Config.apiURL, ipAddress, p.Config.apiKey)

	response, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}

	jsonResult := &dto.IpstackResponse{}

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
		RegionName:    jsonResult.RegionName,
		City:          jsonResult.City,
		Zip:           jsonResult.Zip,
	}

	return locationData, nil
}
