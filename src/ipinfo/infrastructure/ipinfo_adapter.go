package infrastructure

import (
	"encoding/json"
	"errors"
	"fmt"
	appDomain "github.com/magento-hackathon/geolocator-microservice-flamingo/src/app/domain"
	"github.com/magento-hackathon/geolocator-microservice-flamingo/src/ipinfo/infrastructure/dto"
	"github.com/pariz/gountries"
	"net"
	"net/http"
	"strconv"
	"strings"
)

// Constants
const (
	ProviderCode = "ipinfo"
)

type (
	// IPInfoAdapter concrete ipinfo.io implementation
	IPInfoAdapter struct {
		Config struct {
			activeflag bool
		}
	}
)

var _ appDomain.LocationProvider = new(IPInfoAdapter)

// Inject dependencies
func (p *IPInfoAdapter) Inject(
	cfg *struct {
	ActiveFlag bool `inject:"config:providers.ipinfo.active"`
},
) {
	if cfg != nil {
		p.Config.activeflag = cfg.ActiveFlag
	}
}

// GetLocationByIP retrieves the result from ipinfo.io
func (p *IPInfoAdapter) GetLocationByIP(ipAddress net.IP) (*appDomain.LocationData, error) {
	if p.Config.activeflag == false {
		return nil, errors.New(appDomain.ProviderInactive)
	}

	requestURL := fmt.Sprintf("https://ipinfo.io/%s/json", ipAddress)

	response, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}

	jsonResult := &dto.IPInfoResponse{}

	err = json.NewDecoder(response.Body).Decode(jsonResult)
	if err != nil {
		return nil, err
	}

	coordsParts := strings.Split(jsonResult.Loc, ",")

	longitude, err := strconv.ParseFloat(coordsParts[0], 32)
	if err != nil {
		longitude = 0
	}

	latitude, err := strconv.ParseFloat(coordsParts[1], 32)
	if err != nil {
		latitude = 0
	}

	countryName := ""

	gountriesQuery := gountries.New()
	countryResult, err := gountriesQuery.FindCountryByAlpha(jsonResult.Country)
	if err == nil {
		countryName = countryResult.Name.Common
	}

	locationData := &appDomain.LocationData{
		ProviderCode:  ProviderCode,
		Longitude:     float32(latitude),
		Latitude:      float32(longitude),
		ContinentCode: "",
		ContinentName: "",
		CountryCode:   jsonResult.Country,
		CountryName:   countryName,
		RegionCode:    "",
		RegionName:    jsonResult.Region,
		City:          jsonResult.City,
		Zip:           jsonResult.Postal,
	}

	return locationData, nil
}
