package handler

import (
	"fmt"
	"horus/helper"
	"horus/voucher"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type voucherHandler struct {
	voucherService voucher.Service
}

func NewVoucherHandler(voucherService voucher.Service) *voucherHandler {
	return &voucherHandler{voucherService}
}

func (h *voucherHandler) CreateVoucher(c *gin.Context) {
	var input voucher.CreateVoucherInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to create voucher", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newVoucher, err := h.voucherService.CreateVoucher(input)
	if err != nil {

		response := helper.APIResponse("Failed to create voucher", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create voucher", http.StatusOK, "success", voucher.FormatVoucher(newVoucher))
	c.JSON(http.StatusOK, response)
}

func (h *voucherHandler) UploadFoto(c *gin.Context) {
	voucherID := c.Param("id")
	if voucherID == "" {
		response := helper.APIResponse("Missing voucher ID", http.StatusBadRequest, "error", gin.H{"is_uploaded": false})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	file, err := c.FormFile("foto")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed upload foto", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	ID, _ := strconv.Atoi(voucherID)
	path := fmt.Sprintf("images/%d-%s", ID, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed upload foto image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.voucherService.SaveVoucherFoto(ID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed upload foto image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("foto successfully update", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

	//catch input user form body
	//save gambar di folder "images/"
	//service memanggil repo
	//JWT untuk memperoleh id (hrdcode = ID 1)
	//repo get user id = 1
	//repo update data user simpan lokasi file
}

func (h *voucherHandler) GetAllVoucher(c *gin.Context) {

	vouchers, err := h.voucherService.GetAllVoucher()
	if err != nil {

		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List Vouchers", http.StatusOK, "success", voucher.FormatVouchers(vouchers))
	c.JSON(http.StatusOK, response)
}
func (h *voucherHandler) GetVoucherByID(c *gin.Context) {
	idParam := c.Param("id")
	ID, _ := strconv.Atoi(idParam)
	result, err := h.voucherService.GetVoucherByID(ID)
	if err != nil {

		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Voucher by id", http.StatusOK, "success", voucher.FormatVoucher(result))
	c.JSON(http.StatusOK, response)
}

// DeleteVoucher handles the deletion of a voucher by its ID
func (h *voucherHandler) DeleteVoucher(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam) // Convert ID from string to int
	if err != nil {
		response := helper.APIResponse("Invalid voucher ID", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.voucherService.DeleteVoucher(id)
	if err != nil {
		response := helper.APIResponse("Failed to delete voucher", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.APIResponse("Voucher deleted successfully", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}