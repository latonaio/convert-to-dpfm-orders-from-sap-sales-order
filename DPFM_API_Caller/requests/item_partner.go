package requests

type ItemPartner struct {
	SalesOrder      string `json:"SalesOrder"`
	SalesOrderItem  string `json:"SalesOrderItem"`
	PartnerFunction string `json:"PartnerFunction"`
	Customer        string `json:"Customer"`
	Supplier        string `json:"Supplier"`
	Personnel       string `json:"Personnel"`
	ContactPerson   string `json:"ContactPerson"`
}
