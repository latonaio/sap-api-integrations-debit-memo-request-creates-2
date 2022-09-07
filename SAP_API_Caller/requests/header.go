package requests

type Header struct {
	DebitMemoRequest              string  `json:"DebitMemoRequest"`
	DebitMemoRequestType          string  `json:"DebitMemoRequestType"`
	SalesOrganization             string  `json:"SalesOrganization"`
	DistributionChannel           string  `json:"DistributionChannel"`
	OrganizationDivision          string  `json:"OrganizationDivision"`
	SalesGroup                    *string `json:"SalesGroup"`
	SalesOffice                   *string `json:"SalesOffice"`
	SalesDistrict                 *string `json:"SalesDistrict"`
	SoldToParty                   *string `json:"SoldToParty"`
	CreationDate                  *string `json:"CreationDate"`
	LastChangeDate                *string `json:"LastChangeDate"`
	LastChangeDateTime            *string `json:"LastChangeDateTime"`
	PurchaseOrderByCustomer       *string `json:"PurchaseOrderByCustomer"`
	CustomerPurchaseOrderType     *string `json:"CustomerPurchaseOrderType"`
	CustomerPurchaseOrderDate     *string `json:"CustomerPurchaseOrderDate"`
	DebitMemoRequestDate          *string `json:"DebitMemoRequestDate"`
	TotalNetAmount                *string `json:"TotalNetAmount"`
	TransactionCurrency           *string `json:"TransactionCurrency"`
	SDDocumentReason              *string `json:"SDDocumentReason"`
	PricingDate                   *string `json:"PricingDate"`
	CustomerTaxClassification1    *string `json:"CustomerTaxClassification1"`
	HeaderBillingBlockReason      *string `json:"HeaderBillingBlockReason"`
	IncotermsClassification       *string `json:"IncotermsClassification"`
	CustomerPaymentTerms          *string `json:"CustomerPaymentTerms"`
	PaymentMethod                 *string `json:"PaymentMethod"`
	BillingDocumentDate           *string `json:"BillingDocumentDate"`
	ReferenceSDDocument           *string `json:"ReferenceSDDocument"`
	ReferenceSDDocumentCategory   *string `json:"ReferenceSDDocumentCategory"`
	OverallSDProcessStatus        *string `json:"OverallSDProcessStatus"`
	TotalCreditCheckStatus        *string `json:"TotalCreditCheckStatus"`
	OverallSDDocumentRejectionSts *string `json:"OverallSDDocumentRejectionSts"`
	OverallOrdReltdBillgStatus    *string `json:"OverallOrdReltdBillgStatus"`
}
