package responses

type ToItemScheduleLine struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
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
		} `json:"results"`
	} `json:"d"`
}
