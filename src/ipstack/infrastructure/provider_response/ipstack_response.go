package provider_response

type IpstackResponse struct {
	IP            string  `json:"ip"`
	Hostname      string  `json:"hostname"`
	Type          string  `json:"type"`
	ContinentCode string  `json:"continent_code"`
	ContinentName string  `json:"continent_name"`
	CountryCode   string  `json:"country_code"`
	CountryName   string  `json:"country_name"`
	RegionCode    string  `json:"region_code"`
	RegionName    string  `json:"region_name"`
	City          string  `json:"city"`
	Zip           string  `json:"zip"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Location      struct {
		GeonameID int    `json:"geoname_id"`
		Capital   string `json:"capital"`
		Languages []struct {
			Code   string `json:"code"`
			Name   string `json:"name"`
			Native string `json:"native"`
		} `json:"languages"`
		CountryFlag             string `json:"country_flag"`
		CountryFlagEmoji        string `json:"country_flag_emoji"`
		CountryFlagEmojiUnicode string `json:"country_flag_emoji_unicode"`
		CallingCode             string `json:"calling_code"`
		IsEu                    bool   `json:"is_eu"`
	} `json:"location"`
	TimeZone struct {
		ID               string `json:"id"`
		CurrentTime      string `json:"current_time"`
		GmtOffset        int    `json:"gmt_offset"`
		Code             string `json:"code"`
		IsDaylightSaving bool   `json:"is_daylight_saving"`
	} `json:"time_zone"`
	Currency struct {
		Code         string `json:"code"`
		Name         string `json:"name"`
		Plural       string `json:"plural"`
		Symbol       string `json:"symbol"`
		SymbolNative string `json:"symbol_native"`
	} `json:"currency"`
	Connection struct {
		Asn int    `json:"asn"`
		Isp string `json:"isp"`
	} `json:"connection"`
	Security struct {
		IsProxy     bool        `json:"is_proxy"`
		ProxyType   interface{} `json:"proxy_type"`
		IsCrawler   bool        `json:"is_crawler"`
		CrawlerName interface{} `json:"crawler_name"`
		CrawlerType interface{} `json:"crawler_type"`
		IsTor       bool        `json:"is_tor"`
		ThreatLevel string      `json:"threat_level"`
		ThreatTypes interface{} `json:"threat_types"`
	} `json:"security"`
}
