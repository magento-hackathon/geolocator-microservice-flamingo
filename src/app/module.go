package app

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
	"github.com/magento-hackathon/geolocator-microservice-flamingo/src/app/interfaces/controller"
)

type (
	// Module basic struct
	Module struct{}

	routes struct {
		geolocationController *controller.GeoLocationController
	}
)

// Inject dependencies
func (r *routes) Inject(
	geoController *controller.GeoLocationController,
) {
	r.geolocationController = geoController
}

// Configure Rating module
func (m *Module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(routes))
}

// Routes served by this module
func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.Route("/geolocation/:ipaddress", "getgeolocation")
	registry.HandleGet("getgeolocation", r.geolocationController.GetGeoLocation)
}
