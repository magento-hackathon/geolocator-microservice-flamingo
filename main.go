package main

import (
	"github.com/magento-hackathon/geolocator-microservice-flamingo/src/app"
	"github.com/magento-hackathon/geolocator-microservice-flamingo/src/ipstack"

	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/cache"
	"flamingo.me/flamingo/v3/core/locale"
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
			new(locale.Module),
			new(app.Module),
			new(application),
			new(ipstack.Module),
		},
	)
}
