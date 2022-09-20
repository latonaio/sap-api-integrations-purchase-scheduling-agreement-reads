package responses

type ToItem struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			SchedulingAgreement            string      `json:"SchedulingAgreement"`
			SchedulingAgreementItem        string      `json:"SchedulingAgreementItem"`
			CompanyCode                    string      `json:"CompanyCode"`
			PurchasingDocumentCategory     string      `json:"PurchasingDocumentCategory"`
			PurchasingDocumentItemCategory string      `json:"PurchasingDocumentItemCategory"`
			PurchasingDocumentItemText     string      `json:"PurchasingDocumentItemText"`
			Material                       string      `json:"Material"`
			SupplierMaterialNumber         string      `json:"SupplierMaterialNumber"`
			MaterialGroup                  string      `json:"MaterialGroup"`
			Plant                          string      `json:"Plant"`
			ReferenceDeliveryAddressID     string      `json:"ReferenceDeliveryAddressID"`
			IncotermsClassification        string      `json:"IncotermsClassification"`
			OrderQuantityUnit              string      `json:"OrderQuantityUnit"`
			TargetQuantity                 string      `json:"TargetQuantity"`
			PurchaseRequisition            string      `json:"PurchaseRequisition"`
			PurchaseRequisitionItem        string      `json:"PurchaseRequisitionItem"`
			SchedAgrmtAgreedCumQty         string      `json:"SchedAgrmtAgreedCumQty"`
			SchedAgrmtCumQtyReconcileDate  string      `json:"SchedAgrmtCumQtyReconcileDate"`
			AccountAssignmentCategory      string      `json:"AccountAssignmentCategory"`
			NetPriceAmount                 string      `json:"NetPriceAmount"`
			NetPriceQuantity               string      `json:"NetPriceQuantity"`
			OrderPriceUnit                 string      `json:"OrderPriceUnit"`
			ProductType                    string      `json:"ProductType"`
			MaterialType                   string      `json:"MaterialType"`
			StorageLocation                string      `json:"StorageLocation"`
			DocumentCurrency               string      `json:"DocumentCurrency"`
			PurchasingInfoRecord           string      `json:"PurchasingInfoRecord"`
			PurchasingDocumentDeletionCode string      `json:"PurchasingDocumentDeletionCode"`
			UnderdelivTolrtdLmtRatioInPct  string      `json:"UnderdelivTolrtdLmtRatioInPct"`
			OverdelivTolrtdLmtRatioInPct   string      `json:"OverdelivTolrtdLmtRatioInPct"`
			UnlimitedOverdeliveryIsAllowed bool        `json:"UnlimitedOverdeliveryIsAllowed"`
			StockType                      string      `json:"StockType"`
			TaxCode                        string      `json:"TaxCode"`
			TaxCountry                     string      `json:"TaxCountry"`
			GoodsReceiptIsExpected         bool        `json:"GoodsReceiptIsExpected"`
			GoodsReceiptIsNonValuated      bool        `json:"GoodsReceiptIsNonValuated"`
			InvoiceIsExpected              bool        `json:"InvoiceIsExpected"`
			InvoiceIsGoodsReceiptBased     bool        `json:"InvoiceIsGoodsReceiptBased"`
			EvaldRcptSettlmtIsAllowed      bool        `json:"EvaldRcptSettlmtIsAllowed"`
			ToItemScheduleLine struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_SchAgrmtSchLine"`
			ToItemDeliveryAddress struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_SchedgAgrmtDeliveryAddress"`
		} `json:"results"`
	} `json:"d"`
}
