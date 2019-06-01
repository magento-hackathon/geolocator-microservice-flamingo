package ipstack

import (
	"flamingo.me/dingo"
	appDomain "github.com/magento-hackathon/geolocator-microservice-flamingo/src/app/domain"
	"github.com/magento-hackathon/geolocator-microservice-flamingo/src/ipstack/infrastructure"
)

type (
	// Module basic struct
	Module struct{}
)

// Configure Rating module
func (m *Module) Configure(injector *dingo.Injector) {
	injector.BindMulti(new(appDomain.LocationProvider)).To(infrastructure.IPStackProvider{})
}
