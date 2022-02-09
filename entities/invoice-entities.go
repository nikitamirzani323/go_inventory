package entities

type Model_invoicehome struct {
	Idinvoice string `json:"invoice_id"`
	Company   string `json:"invoice_company"`
	Date      string `json:"invoice_date"`
	Name      string `json:"invoice_name"`
	Winlose   int    `json:"invoice_winlose"`
	Status    string `json:"invoice_status"`
	Statuscss string `json:"invoice_statuscss"`
}

type Controller_invoicehome struct {
	Master string `json:"master" validate:"required"`
}
type Controller_invoicesave struct {
	Sdata   string `json:"sdata" validate:"required"`
	Master  string `json:"master" validate:"required"`
	Periode string `json:"periode" validate:"required"`
}
type Controller_invoicesavestatus struct {
	Master  string `json:"master" validate:"required"`
	Invoice string `json:"invoice" validate:"required"`
	Tipe    string `json:"tipe" validate:"required"`
}
