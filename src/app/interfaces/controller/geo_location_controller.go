package controller

import (
	"context"
	"flamingo.me/flamingo/v3/framework/web"
	"github.com/magento-hackathon/geolocator-microservice-flamingo/src/app/domain"
	"net"
	"net/http"
)

type (
	// GeoLocationController is the app main controller
	GeoLocationController struct {
		responder *web.Responder
	}
)

// GetGeoLocation returns a geolocation for a provided ipaddress param
func (c *GeoLocationController) GetGeoLocation(ctx context.Context, r *web.Request) web.Result {
	ipAddress, err := r.Query1("ipaddress")
	if err != nil {
		return c.responder.Data(nil).Status(http.StatusUnprocessableEntity)
	}

	validIP := net.ParseIP(ipAddress)
	if validIP == nil {
		return c.responder.Data(nil).Status(http.StatusUnprocessableEntity)
	}

	res := domain.GeoLocationResult{
		Latitude: 1,
		Longitude: 2,
	}

	return c.responder.Data(res).Status(http.StatusOK)
}
