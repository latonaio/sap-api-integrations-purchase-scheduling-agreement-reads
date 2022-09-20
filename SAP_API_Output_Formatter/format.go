package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-purchase-scheduling-agreement-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			SchedulingAgreement:            data.SchedulingAgreement,
			CompanyCode:                    data.CompanyCode,
			PurchasingDocumentCategory:     data.PurchasingDocumentCategory,
			PurchasingDocumentType:         data.PurchasingDocumentType,
			CreationDate:                   data.CreationDate,
			Language:                       data.Language,
			PurchasingOrganization:         data.PurchasingOrganization,
			PurchasingGroup:                data.PurchasingGroup,
			DocumentCurrency:               data.DocumentCurrency,
			IncotermsClassification:        data.IncotermsClassification,
			PaymentTerms:                   data.PaymentTerms,
			NetPaymentDays:                 data.NetPaymentDays,
			TargetAmount:                   data.TargetAmount,
			ExchangeRate:                   data.ExchangeRate,
			PurchasingDocumentOrderDate:    data.PurchasingDocumentOrderDate,
			Supplier:                       data.Supplier,
			SupplierAddressID:              data.SupplierAddressID,
			ValidityStartDate:              data.ValidityStartDate,
			ValidityEndDate:                data.ValidityEndDate,
			PurchasingDocumentOrigin:       data.PurchasingDocumentOrigin,
			PurchasingDocumentDeletionCode: data.PurchasingDocumentDeletionCode,
			SupplierRespSalesPersonName:    data.SupplierRespSalesPersonName,
			SupplierPhoneNumber:            data.SupplierPhoneNumber,
			InvoicingParty:                 data.InvoicingParty,
			SchedulingAgreementStatus:      data.SchedulingAgreementStatus,
			ToHeaderPartner:                data.ToHeaderPartner.Deferred.URI,
			ToItem:                         data.ToItem.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) ([]Item, error) {
	pm := &responses.Item{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	item := make([]Item, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		item = append(item, Item{
			SchedulingAgreement:            data.SchedulingAgreement,
			SchedulingAgreementItem:        data.SchedulingAgreementItem,
			CompanyCode:                    data.CompanyCode,
			PurchasingDocumentCategory:     data.PurchasingDocumentCategory,
			PurchasingDocumentItemCategory: data.PurchasingDocumentItemCategory,
			PurchasingDocumentItemText:     data.PurchasingDocumentItemText,
			Material:                       data.Material,
			SupplierMaterialNumber:         data.SupplierMaterialNumber,
			MaterialGroup:                  data.MaterialGroup,
			Plant:                          data.Plant,
			ReferenceDeliveryAddressID:     data.ReferenceDeliveryAddressID,
			IncotermsClassification:        data.IncotermsClassification,
			OrderQuantityUnit:              data.OrderQuantityUnit,
			TargetQuantity:                 data.TargetQuantity,
			PurchaseRequisition:            data.PurchaseRequisition,
			PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
			SchedAgrmtAgreedCumQty:         data.SchedAgrmtAgreedCumQty,
			SchedAgrmtCumQtyReconcileDate:  data.SchedAgrmtCumQtyReconcileDate,
			AccountAssignmentCategory:      data.AccountAssignmentCategory,
			NetPriceAmount:                 data.NetPriceAmount,
			NetPriceQuantity:               data.NetPriceQuantity,
			OrderPriceUnit:                 data.OrderPriceUnit,
			ProductType:                    data.ProductType,
			MaterialType:                   data.MaterialType,
			StorageLocation:                data.StorageLocation,
			DocumentCurrency:               data.DocumentCurrency,
			PurchasingInfoRecord:           data.PurchasingInfoRecord,
			PurchasingDocumentDeletionCode: data.PurchasingDocumentDeletionCode,
			UnderdelivTolrtdLmtRatioInPct:  data.UnderdelivTolrtdLmtRatioInPct,
			OverdelivTolrtdLmtRatioInPct:   data.OverdelivTolrtdLmtRatioInPct,
			UnlimitedOverdeliveryIsAllowed: data.UnlimitedOverdeliveryIsAllowed,
			StockType:                      data.StockType,
			TaxCode:                        data.TaxCode,
			TaxCountry:                     data.TaxCountry,
			GoodsReceiptIsExpected:         data.GoodsReceiptIsExpected,
			GoodsReceiptIsNonValuated:      data.GoodsReceiptIsNonValuated,
			InvoiceIsExpected:              data.InvoiceIsExpected,
			InvoiceIsGoodsReceiptBased:     data.InvoiceIsGoodsReceiptBased,
			EvaldRcptSettlmtIsAllowed:      data.EvaldRcptSettlmtIsAllowed,
			ToItemScheduleLine:             data.ToItemScheduleLine.Deferred.URI,
			ToItemDeliveryAddress:          data.ToItemDeliveryAddress.Deferred.URI,
		})
	}

	return item, nil
}

func ConvertToToHeaderPartner(raw []byte, l *logger.Logger) ([]ToHeaderPartner, error) {
	pm := &responses.ToHeaderPartner{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToHeaderPartner. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toHeaderPartner := make([]ToHeaderPartner, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toHeaderPartner = append(toHeaderPartner, ToHeaderPartner{
			SchedulingAgreement:     data.SchedulingAgreement,
			SchedulingAgreementItem: data.SchedulingAgreementItem,
			PurchasingOrganization:  data.PurchasingOrganization,
			SupplierSubrange:        data.SupplierSubrange,
			Plant:                   data.Plant,
			PartnerFunction:         data.PartnerFunction,
			PartnerCounter:          data.PartnerCounter,
			Supplier:                data.Supplier,
			DefaultPartner:          data.DefaultPartner,
		})
	}

	return toHeaderPartner, nil
}

func ConvertToToItem(raw []byte, l *logger.Logger) ([]ToItem, error) {
	pm := &responses.ToItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItem := make([]ToItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItem = append(toItem, ToItem{
			SchedulingAgreement:            data.SchedulingAgreement,
			SchedulingAgreementItem:        data.SchedulingAgreementItem,
			CompanyCode:                    data.CompanyCode,
			PurchasingDocumentCategory:     data.PurchasingDocumentCategory,
			PurchasingDocumentItemCategory: data.PurchasingDocumentItemCategory,
			PurchasingDocumentItemText:     data.PurchasingDocumentItemText,
			Material:                       data.Material,
			SupplierMaterialNumber:         data.SupplierMaterialNumber,
			MaterialGroup:                  data.MaterialGroup,
			Plant:                          data.Plant,
			ReferenceDeliveryAddressID:     data.ReferenceDeliveryAddressID,
			IncotermsClassification:        data.IncotermsClassification,
			OrderQuantityUnit:              data.OrderQuantityUnit,
			TargetQuantity:                 data.TargetQuantity,
			PurchaseRequisition:            data.PurchaseRequisition,
			PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
			SchedAgrmtAgreedCumQty:         data.SchedAgrmtAgreedCumQty,
			SchedAgrmtCumQtyReconcileDate:  data.SchedAgrmtCumQtyReconcileDate,
			AccountAssignmentCategory:      data.AccountAssignmentCategory,
			NetPriceAmount:                 data.NetPriceAmount,
			NetPriceQuantity:               data.NetPriceQuantity,
			OrderPriceUnit:                 data.OrderPriceUnit,
			ProductType:                    data.ProductType,
			MaterialType:                   data.MaterialType,
			StorageLocation:                data.StorageLocation,
			DocumentCurrency:               data.DocumentCurrency,
			PurchasingInfoRecord:           data.PurchasingInfoRecord,
			PurchasingDocumentDeletionCode: data.PurchasingDocumentDeletionCode,
			UnderdelivTolrtdLmtRatioInPct:  data.UnderdelivTolrtdLmtRatioInPct,
			OverdelivTolrtdLmtRatioInPct:   data.OverdelivTolrtdLmtRatioInPct,
			UnlimitedOverdeliveryIsAllowed: data.UnlimitedOverdeliveryIsAllowed,
			StockType:                      data.StockType,
			TaxCode:                        data.TaxCode,
			TaxCountry:                     data.TaxCountry,
			GoodsReceiptIsExpected:         data.GoodsReceiptIsExpected,
			GoodsReceiptIsNonValuated:      data.GoodsReceiptIsNonValuated,
			InvoiceIsExpected:              data.InvoiceIsExpected,
			InvoiceIsGoodsReceiptBased:     data.InvoiceIsGoodsReceiptBased,
			EvaldRcptSettlmtIsAllowed:      data.EvaldRcptSettlmtIsAllowed,
			ToItemScheduleLine:             data.ToItemScheduleLine.Deferred.URI,
			ToItemDeliveryAddress:          data.ToItemDeliveryAddress.Deferred.URI,
		})
	}

	return toItem, nil
}

func ConvertToToItemScheduleLine(raw []byte, l *logger.Logger) ([]ToItemScheduleLine, error) {
	pm := &responses.ToItemScheduleLine{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemScheduleLine. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItemScheduleLine := make([]ToItemScheduleLine, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemScheduleLine = append(toItemScheduleLine, ToItemScheduleLine{
			SchedulingAgreement:       data.SchedulingAgreement,
			SchedulingAgreementItem:   data.SchedulingAgreementItem,
			ScheduleLine:              data.ScheduleLine,
			DelivDateCategory:         data.DelivDateCategory,
			ScheduleLineDeliveryDate:  data.ScheduleLineDeliveryDate,
			ScheduleLineDeliveryTime:  data.ScheduleLineDeliveryTime,
			OrderQuantityUnit:         data.OrderQuantityUnit,
			ScheduleLineOrderQuantity: data.ScheduleLineOrderQuantity,
			PurchaseRequisition:       data.PurchaseRequisition,
			PurchaseRequisitionItem:   data.PurchaseRequisitionItem,
			ScheduleLineIsFixed:       data.ScheduleLineIsFixed,
		})
	}

	return toItemScheduleLine, nil
}

func ConvertToToItemDeliveryAddress(raw []byte, l *logger.Logger) (*ToItemDeliveryAddress, error) {
	pm := &responses.ToItemDeliveryAddress{}

	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemDeliveryAddress. unmarshal error: %w", err)
	}

	return &ToItemDeliveryAddress{
		SchedulingAgreement:     pm.D.SchedulingAgreement,
		SchedulingAgreementItem: pm.D.SchedulingAgreementItem,
		DeliveryAddressID:       pm.D.DeliveryAddressID,
		AddressType:             pm.D.AddressType,
		StreetName:              pm.D.StreetName,
		PostalCode:              pm.D.PostalCode,
		CityName:                pm.D.CityName,
		MobileNumber:            pm.D.MobileNumber,
		Region:                  pm.D.Region,
		Country:                 pm.D.Country,
		EmailAddress:            pm.D.EmailAddress,
		Plant:                   pm.D.Plant,
		CorrespondenceLanguage:  pm.D.CorrespondenceLanguage,
		PhoneNumber:             pm.D.PhoneNumber,
		FaxNumber:               pm.D.FaxNumber,
	}, nil
}
