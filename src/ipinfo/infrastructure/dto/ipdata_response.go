package dto

type (
	// IPInfoResponse response object from api
	IPInfoResponse struct {
		IP       string `json:"ip"`
		Hostname string `json:"hostname"`
		City     string `json:"city"`
		Region   string `json:"region"`
		Country  string `json:"country"`
		Loc      string `json:"loc"`
		Postal   string `json:"postal"`
		Org      string `json:"org"`
	}
)
