package handlers

import (
	"context"
 	"net/http"
	"time"

	"github.com/Kartik30R/Tiket.git/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type AuthHandler struct {
	service models.AuthService
}

// Login endpoint
func (h *AuthHandler) Login(c *gin.Context) {
	creds := &models.AuthCredentials{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := c.ShouldBindJSON(creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	if err := validate.Struct(creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	token, user, err := h.service.Login(ctx, creds)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Successfully logged in",
		"data": gin.H{
			"token": token,
			"user":  user,
		},
	})
}

// Register endpoint
func (h *AuthHandler) Register(c *gin.Context) {
	creds := &models.AuthCredentials{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := c.ShouldBindJSON(creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	if err := validate.Struct(creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "please, provide a valid name, email and password",
		})
		return
	}

	token, user, err := h.service.Register(ctx, creds)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Successfully registered",
		"data": gin.H{
			"token": token,
			"user":  user,
		},
	})
}

// Route registration
func NewAuthHandler(router gin.IRouter, service models.AuthService) {
	handler := &AuthHandler{
		service: service,
	}

	router.POST("/login", handler.Login)
	router.POST("/register", handler.Register)
}
