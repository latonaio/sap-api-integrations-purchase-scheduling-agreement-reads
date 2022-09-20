package sap_api_output_formatter

type PurchaseSchedulingAgreement struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	SchedulingAgreement string `json:"purchase_scheduling_agreement"`
	Deleted       string `json:"deleted"`
}

type Header struct {
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
    ToHeaderPartner                string      `json:"to_SchAgrmtPartner"`
    ToItem                         string      `json:"to_SchedgAgrmtItm"`
}

type Item struct {
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
	ToItemScheduleLine             string      `json:"to_SchAgrmtSchLine"`
	ToItemDeliveryAddress          string      `json:"to_SchedgAgrmtDeliveryAddress"`
}

type ToHeaderPartner struct {
	SchedulingAgreement     string `json:"SchedulingAgreement"`
	SchedulingAgreementItem string `json:"SchedulingAgreementItem"`
	PurchasingOrganization  string `json:"PurchasingOrganization"`
	SupplierSubrange        string `json:"SupplierSubrange"`
	Plant                   string `json:"Plant"`
	PartnerFunction         string `json:"PartnerFunction"`
	PartnerCounter          string `json:"PartnerCounter"`
	Supplier                string `json:"Supplier"`
	DefaultPartner          bool   `json:"DefaultPartner"`
}

type ToItem struct {
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
	ToItemScheduleLine             string      `json:"to_SchAgrmtSchLine"`
	ToItemDeliveryAddress          string      `json:"to_SchedgAgrmtDeliveryAddress"`
}

type ToItemScheduleLine struct {
	SchedulingAgreement          string `json:"SchedulingAgreement"`
	SchedulingAgreementItem      string `json:"SchedulingAgreementItem"`
	ScheduleLine                 string `json:"ScheduleLine"`
	DelivDateCategory            string `json:"DelivDateCategory"`
	ScheduleLineDeliveryDate     string `json:"ScheduleLineDeliveryDate"`
	ScheduleLineDeliveryTime     string `json:"ScheduleLineDeliveryTime"`
	OrderQuantityUnit            string `json:"OrderQuantityUnit"`
	ScheduleLineOrderQuantity    string `json:"ScheduleLineOrderQuantity"`
	PurchaseRequisition          string `json:"PurchaseRequisition"`
	PurchaseRequisitionItem      string `json:"PurchaseRequisitionItem"`
	ScheduleLineIsFixed          bool   `json:"ScheduleLineIsFixed"`
}

type ToItemDeliveryAddress struct {
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
}
