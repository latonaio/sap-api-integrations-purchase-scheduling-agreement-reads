package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			SchedulingAgreement            string      `json:"SchedulingAgreement"`
			CompanyCode                    string      `json:"CompanyCode"`
			PurchasingDocumentCategory     string      `json:"PurchasingDocumentCategory"`
			PurchasingDocumentType         string      `json:"PurchasingDocumentType"`
			CreationDate                   string      `json:"CreationDate"`
			Language                       string      `json:"Language"`
			PurchasingOrganization         string      `json:"PurchasingOrganization"`
			PurchasingGroup                string      `json:"PurchasingGroup"`
			DocumentCurrency               string      `json:"DocumentCurrency"`
			IncotermsClassification        string      `json:"IncotermsClassification"`
			PaymentTerms                   string      `json:"PaymentTerms"`
			NetPaymentDays                 string      `json:"NetPaymentDays"`
			TargetAmount                   string      `json:"TargetAmount"`
			ExchangeRate                   string      `json:"ExchangeRate"`
			PurchasingDocumentOrderDate    string      `json:"PurchasingDocumentOrderDate"`
			Supplier                       string      `json:"Supplier"`
			SupplierAddressID              string      `json:"SupplierAddressID"`
			ValidityStartDate              string      `json:"ValidityStartDate"`
			ValidityEndDate                string      `json:"ValidityEndDate"`
			PurchasingDocumentOrigin       string      `json:"PurchasingDocumentOrigin"`
			PurchasingDocumentDeletionCode string      `json:"PurchasingDocumentDeletionCode"`
			SupplierRespSalesPersonName    string      `json:"SupplierRespSalesPersonName"`
			SupplierPhoneNumber            string      `json:"SupplierPhoneNumber"`
			InvoicingParty                 string      `json:"InvoicingParty"`
			SchedulingAgreementStatus      string      `json:"SchedulingAgreementStatus"`
			ToHeaderPartner              struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_SchAgrmtPartner"`
			ToItem struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_SchedgAgrmtItm"`
		} `json:"results"`
	} `json:"d"`
}
