package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Kartik30R/Tiket.git/models"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	repo models.EventRepository
}

func NewEventHandler(router gin.IRouter, repo models.EventRepository) *EventHandler {
	handler := &EventHandler{
		repo: repo,
	}

	router.GET("/", handler.GetMany)
	router.POST("/", handler.CreateOne)
	router.GET("/:eventId", handler.GetOne)
	router.PUT("/:eventId", handler.UpdateOne)
	router.DELETE("/:eventId", handler.DeleteOne)

	return handler
}

func (h *EventHandler) GetMany(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	events, err := h.repo.GetMany(c)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": events})
}

func (h *EventHandler) GetOne(ctx *gin.Context) {
	idParam := ctx.Param("eventId")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "invalid event ID"})
		return
	}

	c, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	event, err := h.repo.GetOne(c, uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": event})
}

func (h *EventHandler) CreateOne(ctx *gin.Context) {
	var event models.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	newEvent, err := h.repo.CreateOne(c, &event)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Event created", "data": newEvent})
}

func (h *EventHandler) UpdateOne(ctx *gin.Context) {
	idParam := ctx.Param("eventId")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "invalid event ID"})
		return
	}

	var updateData map[string]interface{}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	updatedEvent, err := h.repo.UpdateOne(c, uint(id), updateData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Event updated", "data": updatedEvent})
}

func (h *EventHandler) DeleteOne(ctx *gin.Context) {
	idParam := ctx.Param("eventId")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "invalid event ID"})
		return
	}

	c, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	if err := h.repo.DeleteOne(c, uint(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "event deleted successfully"})
}
