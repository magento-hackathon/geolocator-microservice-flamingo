package controller

import (
	"context"
	"github.com/magento-hackathon/geolocator-microservice-flamingo/src/app/domain"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"flamingo.me/flamingo/v3/framework/web"
)

type (
	mockLocationProvider struct{}
)

func TestGeoLocationController_GetGeoLocation(t *testing.T) {
	type fields struct {
		responder *web.Responder
	}
	type args struct {
		ctx context.Context
		r   *web.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   web.Result
	}{
		{
			name: "empty ip param",
			fields: fields{
				responder: &web.Responder{},
			},
			args: args{
				ctx: context.Background(),
				r: mockRequest(
					"http://www.mock.url/geolocation",
					web.RequestParams{"ipaddress": ""},
				),
			},
			want: &web.DataResponse{
				Response: web.Response{
					Status:         http.StatusBadRequest,
					Body:           nil,
					Header:         http.Header{},
					CacheDirective: nil,
				},
				Data: nil,
			},
		},
		{
			name: "invalid ip param",
			fields: fields{
				responder: &web.Responder{},
			},
			args: args{
				ctx: context.Background(),
				r: mockRequest(
					"http://www.mock.url/geolocation/notvalid",
					web.RequestParams{"ipaddress": "notvalid"},
				),
			},
			want: &web.DataResponse{
				Response: web.Response{
					Status:         http.StatusUnprocessableEntity,
					Body:           nil,
					Header:         http.Header{},
					CacheDirective: nil,
				},
				Data: nil,
			},
		},
		{
			name: "valid ip param",
			fields: fields{
				responder: &web.Responder{},
			},
			args: args{
				ctx: context.Background(),
				r: mockRequest(
					"http://www.mock.url/geolocation/91.41.212.114",
					web.RequestParams{"ipaddress": "91.41.212.114"},
				),
			},
			want: &web.DataResponse{
				Response: web.Response{
					Status:         http.StatusOK,
					Body:           nil,
					Header:         http.Header{"Content-Type": []string{"application/json"}},
					CacheDirective: nil,
				},
				Data: append([]*domain.LocationData{}, dummyLocationData()),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &GeoLocationController{
				responder:         tt.fields.responder,
				LocationProviders: getMockLocationProviders(),
			}

			result := c.GetGeoLocation(tt.args.ctx, tt.args.r)

			if got := result; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeoLocationController.GetGeoLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func mockRequest(reqURL string, params web.RequestParams) *web.Request {
	res := &web.Request{}

	parsedURL, _ := url.Parse(reqURL)

	res.Request().URL = parsedURL
	res.Params = params

	return res
}

func getMockLocationProviders() []domain.LocationProvider {
	var res []domain.LocationProvider
	res = append(res, &mockLocationProvider{})

	return res
}

func (m *mockLocationProvider) GetLocationByIP(ipAddress net.IP) (*domain.LocationData, error) {
	dummyData := dummyLocationData()

	return dummyData, nil
}

func dummyLocationData() *domain.LocationData {
	return &domain.LocationData{
		ProviderCode:  "dummyProviderCode",
		Longitude:     10,
		Latitude:      20,
		ContinentCode: "dummyContinent",
		ContinentName: "dummyContinentName",
		CountryCode:   "dummyCountryCode",
		CountryName:   "dummyCountryName",
		RegionCode:    "dummyRegionCode",
		RegionName:    "dummyRegionName",
		City:          "dummyCity",
		Zip:           "dummyZip",
		ErrorMessage:  "",
	}
}
