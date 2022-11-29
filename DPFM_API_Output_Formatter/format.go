package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "convert-to-dpfm-orders-from-sap-sales-order/DPFM_API_Input_Reader"
	"fmt"
	"strconv"
)

func ConvertToHeader(sdc *dpfm_api_input_reader.SDC) (*Header, error) {
	data := sdc.SalesOrder
	var err error

	var priceDetnExchangeRateFloat32 *float32
	if data.PriceDetnExchangeRate != nil {
		priceDetnExchangeRateFloat32, err = parseFolat32Ptr(*data.PriceDetnExchangeRate)
		if err != nil {
			return nil, err
		}
	} else {
		priceDetnExchangeRateFloat32 = nil
	}

	var accountingExchangeRateFloat32 *float32
	if data.AccountingExchangeRate != nil {
		accountingExchangeRateFloat32, err = parseFolat32Ptr(*data.AccountingExchangeRate)
		if err != nil {
			return nil, err
		}
	} else {
		accountingExchangeRateFloat32 = nil
	}

	header := Header{
		// OrderID:                         data.SalesOrder,
		OrderDate: *data.SalesOrderDate,
		// OrderType:                       data.OrderType,
		// Buyer:                           data.Buyer,
		// Seller:                          data.Seller,
		// CreationDate:                    data.CreationDate,
		// LastChangeDate:                  data.LastChangeDate,
		// ContractType:                    data.ContractType,
		// ValidityStartDate:               data.ValidityStartDate,
		// ValidityEndDate:                 data.ValidityEndDate,
		// InvoiceScheduleStartDate:        data.InvoiceScheduleStartDate,
		// InvoiceScheduleEndDate:          data.InvoiceScheduleEndDate,
		// TotalNetAmount:                  data.TotalNetAmount,
		// TotalTaxAmount:                  data.TotalTaxAmount,
		// TotalGrossAmount:                data.TotalGrossAmount,
		// OverallDeliveryStatus:           data.OverallDeliveryStatus,
		// TotalBlockStatus:                data.TotalBlockStatus,
		// OverallOrdReltdBillgStatus:      data.OverallOrdReltdBillgStatus,
		// OverallDocReferenceStatus:       data.OverallDocReferenceStatus,
		TransactionCurrency:   *data.TransactionCurrency,
		PricingDate:           *data.PricingDate,
		PriceDetnExchangeRate: priceDetnExchangeRateFloat32,
		RequestedDeliveryDate: *data.RequestedDeliveryDate,
		// HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		// HeaderBillingBlockReason:        data.HeaderBillingBlockReason,
		// DeliveryBlockReason:             data.DeliveryBlockReason,
		Incoterms:     data.IncotermsClassification,
		PaymentTerms:  *data.CustomerPaymentTerms,
		PaymentMethod: *data.PaymentMethod,
		// ReferenceDocument:        data.ReferenceDocument,
		// ReferenceDocumentItem:    data.ReferenceDocumentItem,
		// BPAccountAssignmentGroup: data.BPAccountAssignmentGroup,
		AccountingExchangeRate: accountingExchangeRateFloat32,
		BillingDocumentDate:    *data.BillingDocumentDate,
		// IsExportImportDelivery:   data.IsExportImportDelivery,
		// HeaderText:               data.HeaderText,
	}

	return &header, nil
}

func parseFolat32Ptr(s string) (*float32, error) {
	resFloat64, err := strconv.ParseFloat(s, 32)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return nil, err
	}
	resFloat32 := getFloat32Ptr(float32(resFloat64))

	return resFloat32, nil
}

func getFloat32Ptr(f float32) *float32 {
	return &f
}
