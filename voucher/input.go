package voucher

type CreateVoucherInput struct{
	Name string `json:"name" binding:"required"`
	Category string `json:"category" binding:"required"`
	Status bool `json:"status" binding:"boolean"`
}
