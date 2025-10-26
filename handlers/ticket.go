package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Kartik30R/Tiket.git/models"
	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	repo models.TicketRepository
}


func (h *TicketHandler) GetOne(ctx *gin.Context) {
	idParam := ctx.Param("ticketId")
	ticketID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "invalid ticket ID"})
		return
	}

	c, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	ticket, err := h.repo.GetOne(c, uint(ticketID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": ticket})
}


func (h *TicketHandler) GetMany(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()


	tickets, err := h.repo.GetMany(c)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": tickets})
}

func (h *TicketHandler) CreateOne(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	ticket := &models.Ticket{}
 
	if err := ctx.ShouldBindJSON(ticket); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ticket, err := h.repo.CreateOne(c,ticket)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Ticket created", "data": ticket})
}


func (h *TicketHandler) ValidateOne(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	validateBody := &models.ValidateTicket{}
	if err := ctx.ShouldBindJSON(validateBody); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	updateData := map[string]any{"entered": true}

	ticket, err := h.repo.UpdateOne(c, validateBody.TicketId, updateData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Welcome to the show!", "data": ticket})
}


func NewTicketHandler(router gin.IRouter, repo models.TicketRepository) *TicketHandler{
	handler:= &TicketHandler{
		repo: repo,
	}

	router.GET("/", handler.GetMany)
	router.POST("/", handler.CreateOne)
	router.GET("/:ticketId", handler.GetOne)
  	router.POST("/validate", handler.ValidateOne)

	return handler
}

