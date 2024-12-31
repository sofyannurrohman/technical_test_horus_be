package main

import (
	"horus/auth"
	"horus/handler"
	"horus/helper"
	"horus/user"
	"horus/voucher"
	voucherclaim "horus/voucher_claim"
	"net/http"
	"strings"

	// voucherclaim "horus/voucher_claim"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/horus_sofyan_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	/* Repository */
	userRepository := user.NewRepository(db)
	voucherRepository := voucher.NewRepository(db)
	voucherClaimRepository := voucherclaim.NewRepository(db)

	/* Service */
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	voucherService := voucher.NewService(voucherRepository)
	voucherClaimService := voucherclaim.NewService(voucherClaimRepository)

	/* Handler */
	userHandler := handler.NewUserHandler(userService, authService)
	voucherHandler := handler.NewVoucherHandler(voucherService)
	voucherClaimHandler := handler.NewVoucherClaimHandler(voucherClaimService)

	/* Router */
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60, 
	}))
	router.Static("/images", "./images")

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.GET("/users/:id", userHandler.GetUserByID)
	api.POST("/sessions", userHandler.LoginUser)
	api.POST("/vouchers", voucherHandler.CreateVoucher)
	api.POST("/vouchers/:id/foto", voucherHandler.UploadFoto)
	api.DELETE("/vouchers/:id", voucherHandler.DeleteVoucher)
	api.GET("/vouchers", voucherHandler.GetAllVoucher)
	api.GET("/vouchers/:id", voucherHandler.GetVoucherByID)
	api.POST("/vouchers/:id/claim-vouchers",authMiddleware(authService,userService), voucherClaimHandler.CreateVoucherClaim)
	api.GET("/claim-vouchers",authMiddleware(authService,userService), voucherClaimHandler.GetVoucherClaimByUserID)
	api.DELETE("/claim-vouchers/:id", voucherClaimHandler.DeleteVoucherClaim)
	router.Run(":8000")
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userIDStr, ok := claim["user_id"].(string)
		if !ok {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			response := helper.APIResponse("Invalid User ID", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
}
