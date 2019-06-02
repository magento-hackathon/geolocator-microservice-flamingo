package main

import (
	"github.com/magento-hackathon/geolocator-microservice-flamingo/src/app"
	"github.com/magento-hackathon/geolocator-microservice-flamingo/src/ipdata"
	"github.com/magento-hackathon/geolocator-microservice-flamingo/src/ipinfo"
	"github.com/magento-hackathon/geolocator-microservice-flamingo/src/ipstack"
	"github.com/magento-hackathon/geolocator-microservice-flamingo/src/maxmind"

	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/cache"
)

type (
	application   struct{}
)

func (a *application) Configure(injector *dingo.Injector) {
	injector.Bind((*cache.Backend)(nil)).ToInstance(cache.NewInMemoryCache())
}

func main() {
	flamingo.App(
		[]dingo.Module{
			new(app.Module),
			new(application),
			new(ipstack.Module),
			new(ipdata.Module),
			new(maxmind.Module),
			new(ipinfo.Module),
		},
	)
}
