package handler

import (
	"horus/helper"
	"horus/user"
	voucherclaim "horus/voucher_claim"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type voucherClaimHandler struct {
	service voucherclaim.Service
}

func NewVoucherClaimHandler(service voucherclaim.Service) *voucherClaimHandler {
	return &voucherClaimHandler{service}
}

// CreateVoucherClaim handles the creation of a new voucher claim
func (h *voucherClaimHandler) CreateVoucherClaim(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(user.User)
	IDParam := c.Param("id")
	userID := currentUser.ID
	voucherID, _ := strconv.Atoi(IDParam)
	// Call the service to save the voucher claim
	newVoucherClaim, err := h.service.CreateVoucherClaim(userID, voucherID)
	if err != nil {
		response := helper.APIResponse("Failed to create voucher claim", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Return success response
	response := helper.APIResponse("Voucher claim created successfully", http.StatusOK, "success", newVoucherClaim)
	c.JSON(http.StatusOK, response)
}

// GetVoucherClaimByID handles fetching a voucher claim by ID
func (h *voucherClaimHandler) GetVoucherClaimByUserID(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	// Fetch the voucher claim
	voucherClaim, err := h.service.GetVoucherClaimByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Voucher claim not found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	// Return success response
	response := helper.APIResponse("Voucher claim's retrieved successfully", http.StatusOK, "success", voucherClaim)
	c.JSON(http.StatusOK, response)
}

// GetVoucherClaimByID handles fetching a voucher claim by ID
func (h *voucherClaimHandler) DeleteVoucherClaim(c *gin.Context) {
	idParam := c.Param("id")
	id,_ := strconv.Atoi(idParam)

	// Fetch the voucher claim
	voucherClaim, err := h.service.GetVoucherClaimByID(id)
	if err != nil {
		response := helper.APIResponse("Voucher claim not found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	err = h.service.DeleteVoucherClaim(voucherClaim.ID)
	if err != nil {
		response := helper.APIResponse("Failed delete voucher", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	response := helper.APIResponse("Voucher claim successfully deleted", http.StatusOK, "success", voucherClaim)
	c.JSON(http.StatusOK, response)
}
