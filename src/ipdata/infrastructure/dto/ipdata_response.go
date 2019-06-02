package dto

import "time"

type (
	// IPDataResponse response object from api
	IPDataResponse struct {
		IP            string  `json:"ip"`
		IsEu          bool    `json:"is_eu"`
		City          string  `json:"city"`
		Region        string  `json:"region"`
		RegionCode    string  `json:"region_code"`
		CountryName   string  `json:"country_name"`
		CountryCode   string  `json:"country_code"`
		ContinentName string  `json:"continent_name"`
		ContinentCode string  `json:"continent_code"`
		Latitude      float64 `json:"latitude"`
		Longitude     float64 `json:"longitude"`
		Asn           string  `json:"asn"`
		Organisation  string  `json:"organisation"`
		Postal        string  `json:"postal"`
		CallingCode   string  `json:"calling_code"`
		Flag          string  `json:"flag"`
		EmojiFlag     string  `json:"emoji_flag"`
		EmojiUnicode  string  `json:"emoji_unicode"`
		Carrier       struct {
			Name string `json:"name"`
			Mcc  string `json:"mcc"`
			Mnc  string `json:"mnc"`
		} `json:"carrier"`
		Languages []struct {
			Name   string `json:"name"`
			Native string `json:"native"`
		} `json:"languages"`
		Currency struct {
			Name   string `json:"name"`
			Code   string `json:"code"`
			Symbol string `json:"symbol"`
			Native string `json:"native"`
			Plural string `json:"plural"`
		} `json:"currency"`
		TimeZone struct {
			Name        string    `json:"name"`
			Abbr        string    `json:"abbr"`
			Offset      string    `json:"offset"`
			IsDst       bool      `json:"is_dst"`
			CurrentTime time.Time `json:"current_time"`
		} `json:"time_zone"`
		Threat struct {
			IsTor           bool `json:"is_tor"`
			IsProxy         bool `json:"is_proxy"`
			IsAnonymous     bool `json:"is_anonymous"`
			IsKnownAttacker bool `json:"is_known_attacker"`
			IsKnownAbuser   bool `json:"is_known_abuser"`
			IsThreat        bool `json:"is_threat"`
			IsBogon         bool `json:"is_bogon"`
		} `json:"threat"`
		Count string `json:"count"`
	}
)
