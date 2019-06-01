package infrastructure

import (
	appDomain "github.com/magento-hackathon/geolocator-microservice-flamingo/src/app/domain"
	"net"
)

type (
	// IPStackProvider concrete ipstack.com implementation
	IPStackProvider struct {}
)

var _ appDomain.LocationProvider = new(IPStackProvider)

// GetLocationByIP retrieves the result from ipstack.com
func (p *IPStackProvider) GetLocationByIP (ipAdress net.IP) (*appDomain.LocationData, error) {
	return nil, nil
}
