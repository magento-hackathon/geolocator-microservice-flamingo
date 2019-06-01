package controller

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"flamingo.me/flamingo/v3/framework/web"
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
				responder:  &web.Responder{},
			},
			args: args{
				ctx: context.Background(),
				r: mockRequest("http://www.mock.url"),
			},
			want: &web.DataResponse{
				Response: web.Response{
					Status: http.StatusUnprocessableEntity,
					Body: nil,
					Header: http.Header{},
					CacheDirective: nil,
				},
				Data: nil,
			},
		},
		{
			name: "invalid ip param",
			fields: fields{
				responder:  &web.Responder{},
			},
			args: args{
				ctx: context.Background(),
				r: mockRequest(
					"http://www.mock.url?ipaddress=bjkhlwarekjlbhjbhk",
				),
			},
			want: &web.DataResponse{
				Response: web.Response{
					Status: http.StatusUnprocessableEntity,
					Body: nil,
					Header: http.Header{},
					CacheDirective: nil,
				},
				Data: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &GeoLocationController{
				responder: tt.fields.responder,
			}

			result := c.GetGeoLocation(tt.args.ctx, tt.args.r)

			if got := result; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeoLocationController.GetGeoLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func mockRequest(reqURL string) *web.Request {
	res := &web.Request{}

	parsedURL, _ := url.Parse(reqURL)

	res.Request().URL = parsedURL

	return res
}
