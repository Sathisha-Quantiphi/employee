package voucher

type Voucher struct {
	ID          string
	Description string
	Type        string
	Status      string
	Invoice     string
	Vendor      string
	Amount      float64
}
