package voucher

type Service interface {
	CreateVoucher(input CreateVoucherInput) (Voucher, error)
	SaveVoucherFoto(ID int, fileLocation string) (Voucher, error)
	GetAllVoucher() ([]Voucher, error)
	GetVoucherByID(ID int) (Voucher, error)
	DeleteVoucher(ID int)error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateVoucher(input CreateVoucherInput) (Voucher, error) {
	voucher := Voucher{}
	voucher.Name = input.Name
	voucher.Category = input.Category
	voucher.Status = input.Status

	newVoucher, err := s.repository.Save(voucher)
	if err != nil {
		return newVoucher, err
	}
	return newVoucher, nil
}
func (s *service) SaveVoucherFoto(ID int, fileLocation string) (Voucher, error) {
	//getuser by id
	voucher, err := s.repository.FindByID(ID)
	if err != nil {
		return voucher, err
	}

	//update attribute avatar filename
	voucher.Foto = fileLocation
	//save update to db
	updatedVoucher, err := s.repository.Update(voucher)
	if err != nil {
		return updatedVoucher, err
	}
	return updatedVoucher, nil
}

func (s *service) GetAllVoucher() ([]Voucher, error) {
	vouchers, err := s.repository.FindAll()
	if err != nil {
		return vouchers, err
	}
	return vouchers, nil
}

func (s *service) GetVoucherByID(ID int) (Voucher, error) {
	voucher, err := s.repository.FindByID(ID)
	if err != nil {
		return voucher, err
	}
	return voucher, nil
}
func (s *service) DeleteVoucher(ID int) error {
	return s.repository.Delete(ID)
}

