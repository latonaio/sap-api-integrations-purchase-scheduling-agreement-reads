package responses

type ToHeaderPartner struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			SchedulingAgreement     string `json:"SchedulingAgreement"`
			SchedulingAgreementItem string `json:"SchedulingAgreementItem"`
			PurchasingOrganization  string `json:"PurchasingOrganization"`
			SupplierSubrange        string `json:"SupplierSubrange"`
			Plant                   string `json:"Plant"`
			PartnerFunction         string `json:"PartnerFunction"`
			PartnerCounter          string `json:"PartnerCounter"`
			Supplier                string `json:"Supplier"`
			DefaultPartner          bool   `json:"DefaultPartner"`
		} `json:"results"`
	} `json:"d"`
}
