package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/morodik/convoy/internal/services"
)

type ConvoyHandler struct {
	convoyService *services.ConvoyService
}

func NewConvoyHandler(convoyService *services.ConvoyService) *ConvoyHandler {
	return &ConvoyHandler{convoyService: convoyService}
}

type CreateConvoyRequest struct {
	Title     string     `json:"title" binding:"required"`
	StartTime time.Time  `json:"start_time" binding:"required"`
	EndTime   *time.Time `json:"end_time"`
	IsPrivate bool       `json:"is_private"`
}

func (h *ConvoyHandler) CreateConvoy(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		c.Abort()
		return
	}

	// Логирование для отладки
	log.Printf("Тип userID: %T, значение: %v", userID, userID)

	id, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обработки ID пользователя"})
		c.Abort()
		return
	}

	var req CreateConvoyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных: " + err.Error()})
		c.Abort()
		return
	}

	convoy, err := h.convoyService.CreateConvoy(c.Request.Context(), id, req.Title, req.StartTime, req.EndTime, req.IsPrivate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, convoy)
}
