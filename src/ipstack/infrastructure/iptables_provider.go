package infrastructure

import (
	"encoding/json"
	"fmt"
	appDomain "github.com/magento-hackathon/geolocator-microservice-flamingo/src/app/domain"
	"github.com/magento-hackathon/geolocator-microservice-flamingo/src/ipstack/infrastructure/provider_response"
	"net"
	"net/http"
)

const (
	APIKEY = "87074d3a66828bfcbade49ca0edbf99b"
	APIURL = "http://api.ipstack.com/%s?access_key=%s"
)

type (
	// IPStackProvider concrete ipstack.com implementation
	IPStackProvider struct{}
)

var _ appDomain.LocationProvider = new(IPStackProvider)

// GetLocationByIP retrieves the result from ipstack.com
func (p *IPStackProvider) GetLocationByIP(ipAdress net.IP) (*appDomain.LocationData, error) {
	requestUrl := fmt.Sprintf(APIURL, ipAdress, APIKEY)

	response, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}

	jsonResult := &provider_response.IpstackResponse{}

	err = json.NewDecoder(response.Body).Decode(jsonResult)
	if err != nil {
		return nil, err
	}

	locationData := &appDomain.LocationData{
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
