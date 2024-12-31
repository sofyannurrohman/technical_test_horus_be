package voucher

type VoucherFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Status   bool   `json:"status"`
	Foto string `json:"foto"`
}

func FormatVoucher(voucher Voucher) VoucherFormatter {
	formatter := VoucherFormatter{
		ID:       voucher.ID,
		Name:     voucher.Name,
		Category: voucher.Category,
		Status:   voucher.Status,
		Foto: voucher.Foto,
	}
	return formatter
}
func FormatVouchers(vouchers []Voucher) []VoucherFormatter {

	vouchersFormatter := []VoucherFormatter{}

	for _, voucher := range vouchers {
		voucherFormatter := FormatVoucher(voucher)
		vouchersFormatter = append(vouchersFormatter, voucherFormatter)
	}
	return vouchersFormatter
}
