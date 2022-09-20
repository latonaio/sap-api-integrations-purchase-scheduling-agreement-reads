package responses

type ToItemDeliveryAddress struct {
	D struct {
		Metadata struct {
			ID   string `json:"id"`
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
		SchedulingAgreement     string `json:"SchedulingAgreement"`
		SchedulingAgreementItem string `json:"SchedulingAgreementItem"`
		DeliveryAddressID       string `json:"DeliveryAddressID"`
		AddressType             string `json:"AddressType"`
		StreetName              string `json:"StreetName"`
		PostalCode              string `json:"PostalCode"`
		CityName                string `json:"CityName"`
		MobileNumber            string `json:"MobileNumber"`
		Region                  string `json:"Region"`
		Country                 string `json:"Country"`
		EmailAddress            string `json:"EmailAddress"`
		Plant                   string `json:"Plant"`
		CorrespondenceLanguage  string `json:"CorrespondenceLanguage"`
		PhoneNumber             string `json:"PhoneNumber"`
		FaxNumber               string `json:"FaxNumber"`
	} `json:"d"`
}
