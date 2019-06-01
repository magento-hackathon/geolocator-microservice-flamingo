package main

import (
	"flamingo.me/flamingo/v3/core/gotemplate"
	"github.com/magento-hackathon/geolocator-microservice-flamingo/src/app"
	"net/http"

	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/cache"
	"flamingo.me/flamingo/v3/core/locale"
	"flamingo.me/flamingo/v3/framework/web"
)

type (
	application   struct{}
	defaultRoutes struct{}
)

func (a *application) Configure(injector *dingo.Injector) {
	injector.Bind((*cache.Backend)(nil)).ToInstance(cache.NewInMemoryCache())
	web.BindRoutes(injector, &defaultRoutes{})
}

// Routes
func (a *defaultRoutes) Routes(registry *web.RouterRegistry) {
	registry.Route("/static/*n", "_static")
	registry.HandleGet(
		"_static",
		web.WrapHTTPHandler(http.StripPrefix("/static/", http.FileServer(http.Dir("static")))),
	)
}

func main() {
	flamingo.App(
		[]dingo.Module{
			new(locale.Module),
			new(gotemplate.Module),
			//new(opencensus.Module),
			new(app.Module),
			new(application),
		},
	)
}
